package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sys"
	"sys/err"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

const (
	datasetName = "oan_cor_tic"
)

type (
	app struct {
		names      []string
		instrs     map[string]*instr
		goo        goo
		oan        oan
		ticksC     chan tick
		rxExitC    chan bool
		rxExitedC  chan bool
		txExitC    chan bool
		txExitedC  chan bool
		retryDelay time.Duration
		ln         string
		row        Row
	}
	instr struct {
		name string
		tbl  *bigquery.Table
		ldr  *bigquery.Uploader
	}
	tick struct {
		instr string
		row   Row
	}
	Row struct {
		Time time.Time
		Tick string
	}
	goo struct {
		projectID string
		ctx       context.Context
		client    *bigquery.Client
	}
	oan struct {
		accountID     string
		token         string
		baseStreamURL string
		client        *http.Client
	}
)

var (
	del = flag.Bool("del", false, "deletes all tables and exits")
	fle = flag.Bool("fle", false, "loads a google account file")
)

func main() {
	fmt.Println("--- start", time.Now())
	defer fmt.Println("--- exit", time.Now())
	app := newApp()
	app.open()
	if *del {
		app.clrStore()
	} else {
		app.initStore()
		app.run()
	}
}

func newApp() (r *app) {
	r = &app{}
	r.names = []string{"aud_usd", "eur_usd", "nzd_usd", "gbp_usd"}
	r.instrs = make(map[string]*instr)
	r.goo.projectID = "pro-cld"
	r.oan.accountID = "101-001-4093945-001"
	r.oan.token = "d3e81b3ca79ebd8cf471e31b7d3a4f74-d5cee5a343849dcbb45e3fd6f21870f0"
	r.oan.baseStreamURL = "https://stream-fxpractice.oanda.com"
	r.rxExitC = make(chan bool)
	r.rxExitedC = make(chan bool)
	r.txExitC = make(chan bool)
	r.txExitedC = make(chan bool)
	r.retryDelay = time.Second * 15
	return r
}
func (x *app) open() {
	fmt.Println(">>> HERE")
	flag.Parse()
	var err error
	var opts []option.ClientOption
	if *fle {
		curDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		flepath := filepath.Join(curDir, "pro-cld.json")
		fmt.Println("loading...", flepath)
		opts = append(opts, option.WithServiceAccountFile(flepath))
	}
	x.goo.ctx = context.Background()
	x.goo.client, err = bigquery.NewClient(x.goo.ctx, x.goo.projectID, opts...)
	if err != nil {
		panic(err)
	}
	x.oan.client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
			MaxIdleConnsPerHost: 2,
		},
	}
	x.ticksC = make(chan tick, 2096)
}
func (x *app) initStore() {
	ds := x.goo.client.Dataset(datasetName)
	err := ds.Create(x.goo.ctx, &bigquery.DatasetMetadata{}) // err returned if already exists
	if err == nil {
		fmt.Printf("dataset: %v: created \n", datasetName)
	} else {
		fmt.Printf("dataset: %v: not created, already exists (%v) \n", datasetName, err)
	}
	for _, name := range x.names {
		tbl := ds.Table(name)
		mta := &bigquery.TableMetadata{}
		mta.Name = name
		mta.Schema = append(mta.Schema, &bigquery.FieldSchema{
			Name:        "Time",
			Description: "Time sent from the Oanda server",
			Type:        bigquery.TimestampFieldType,
			Required:    true,
		})
		mta.Schema = append(mta.Schema, &bigquery.FieldSchema{
			Name:        "Tick",
			Description: "Price tick in JSON format",
			Type:        bigquery.StringFieldType,
			Required:    true,
		})
		err = tbl.Create(x.goo.ctx, mta)
		if err == nil {
			fmt.Printf("table: %v: created \n", name)
		} else {
			fmt.Printf("table: %v: not created, already exists (%v) \n", name, err)
		}
		// use upper case key; Oanda tick data recieved as upper case
		// storage uses lowercase
		x.instrs[strings.ToUpper(name)] = &instr{
			name: name,
			tbl:  tbl,
			ldr:  tbl.Uploader(),
		}
	}
}
func (x *app) clrStore() {
	fmt.Println("deleting...")
	ds := x.goo.client.Dataset(datasetName)
	itr := ds.Tables(x.goo.ctx)
	for {
		tbl, err := itr.Next()
		if err != nil {
			break
		}
		err = tbl.Delete(x.goo.ctx)
		if err == nil {
			fmt.Printf("table: %v: deleted \n", tbl.FullyQualifiedName())
		} else {
			fmt.Printf("table: %v: not deleted (%v) \n", tbl.FullyQualifiedName(), err)
		}
	}
	err := ds.Delete(x.goo.ctx)
	if err == nil {
		fmt.Printf("dataset: %v: deleted \n", datasetName)
	} else {
		fmt.Printf("dataset: %v: not deleted (%v) \n", datasetName, err)
	}
}
func (x *app) run() {
	fmt.Println("running... ", time.Now())
	go x.rx()
	go x.tx()
	go x.key()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait() // wait indefintely
}
func (x *app) rx() {
	fmt.Println("rx starting...", time.Now())
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		sys.Log("rx restarting in... ", x.retryDelay, time.Now())
		<-time.NewTimer(x.retryDelay).C
		go x.rx()
	}()
	u, err := url.Parse(x.oan.baseStreamURL)
	if err != nil {
		panic(err)
	}
	u.Path = fmt.Sprintf("/v3/accounts/%v/pricing/stream", x.oan.accountID)
	var b bytes.Buffer
	n := 0
	for _, name := range x.names {
		b.WriteString(strings.ToUpper(name))
		if n != len(x.names)-1 {
			b.WriteString(",")
		}
		n++
	}
	q := u.Query()
	q.Set("instruments", b.String())
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+x.oan.token)
	resp, err := x.oan.client.Do(req)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(resp.Body)
Loop:
	for {
		select {
		case <-x.rxExitC:
			break Loop
		default:
			lnbytes, err := reader.ReadBytes('\n')
			if err != nil {
				panic(err)
			}
			x.ln = string(lnbytes)
			fmt.Println("ln", x.ln)
			if strings.Index(x.ln, "PRICE") > -1 { //&& lex(x.ln, "status") == "tradeable" {
				t, err := time.Parse(time.RFC3339Nano, lex(x.ln, "time"))
				if err != nil {
					panic(err)
				}
				x.ticksC <- tick{
					instr: lex(x.ln, "instrument"),
					row: Row{
						Time: t,
						Tick: x.ln,
					},
				}
			}
		}
	}
}
func lex(ln, key string) string {
	idx := strings.Index(ln, key)
	idx += strings.Index(ln[idx:], ":")
	idx += strings.Index(ln[idx:], "\"")
	idx++
	lim := idx + strings.Index(ln[idx:], "\"")
	//fmt.Println("lex", key, ln[idx+1:lim])
	return ln[idx:lim]
}
func (x *app) tx() {
	fmt.Println("tx starting...", time.Now())
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		sys.Log("tx restarting in... ", x.retryDelay, time.Now())
		<-time.NewTimer(x.retryDelay).C
		go x.tx()
	}()
Loop:
	for {
		select {
		case <-x.txExitC:
			break Loop
		case tick := <-x.ticksC:
			i, ok := x.instrs[tick.instr]
			if !ok {
				panic(fmt.Sprintf("x.instrs missing %v", tick.instr))
			}
			// fmt.Println("tx", tick.row)
			x.row = tick.row
			err := i.ldr.Put(x.goo.ctx, tick.row)
			if err != nil {
				fmt.Printf("tx: err: %v (%v)", i.name, err)
			}
		}
	}
}
func (x *app) key() {
	fmt.Println("key starting...", time.Now())
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		sys.Log("key restarting in... ", x.retryDelay, time.Now())
		<-time.NewTimer(x.retryDelay).C
		go x.key()
	}()
	scnr := bufio.NewScanner(os.Stdin)
	for scnr.Scan() { // print last line recieved to allow manual monitoring
		// fmt.Println("--- last rx", x.ln[:len(x.ln)-1])
		// fmt.Println("--- last tx", x.row)
		if err := scnr.Err(); err != nil {
			panic(err)
		}
	}
}
