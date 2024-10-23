package ana

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"sys"
	"sys/ana/cfg"
	"sys/bsc/bol"
	"sys/bsc/flt"
	bscint "sys/bsc/int"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/err"
	"sys/lng/jsn"
	"sys/trc"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const (
	No  = uint32(0)
	Yes = uint32(1)
)

var (
	PrvOan *Oan
)

type (
	Oan struct {
		Cfg           *cfg.Cfg
		Ticr          *Ticr
		Instrs        Instrs
		instrsMu      sync.RWMutex
		subs          Instrs
		subsMu        sync.RWMutex
		rxExitC       chan bool
		rxExitedC     chan bool
		acnt          *Acnt
		opning        uint32
		filing        uint32
		trding        uint32 // in trd: 0/1
		trdMu         sync.Mutex
		opningPending []string
	}
	Instrs map[str.Str]*Instr
	TicRow struct {
		Time time.Time
		Tick string
	}
)

func NewOan(cfg *cfg.Cfg, ticr *Ticr) (r *Oan) {
	r = &Oan{}
	r.Cfg = cfg
	r.Ticr = ticr
	r.Instrs = make(Instrs)
	if cfg.Test {
		// for portfolio hst/rlt pairity testing
		r.acnt = &Acnt{}
		r.acnt.Balance = cfg.Hst.BalUsd
	}
	PrvOan = r // pkg glbl set
	return r
}
func (x *Oan) Cls() {
	x.subsMu.Lock()
	if len(x.subs) > 0 && !x.Cfg.Test {
		x.rxExitC <- true // request current http stream to exit
		<-x.rxExitedC     // wait for exit completion
	}
	x.rxExitC = nil // initialized in Sub
	x.rxExitedC = nil
	x.subs = nil
	x.subsMu.Unlock()
}

// Instr loads instrument details from an Oanda server.
func (x *Oan) Instr(name str.Str) (r *Instr) {
	// sys.Log("Oan.Instr")
	x.instrsMu.RLock()
	r, ok := x.Instrs[name]
	if ok {
		x.instrsMu.RUnlock()
		return r
	}
	x.instrsMu.RUnlock()
	x.instrsMu.Lock()
	defer x.instrsMu.Unlock()
	r, ok = x.Instrs[name]
	if ok {
		return r
	}
	r = &Instr{}
	r.Name = name
	if x.DskLoadInstrDetail(r) { // try load from local disk
		sys.Log("Oan.Instr", r)
		r.Prv = x
		x.Instrs[name] = r
		return r
	}
	trcr := trc.New("Oan.Instr: Network")
	defer trcr.End()
	client := &http.Client{ // load from network
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
	u, er := url.Parse(x.Cfg.BaseURL.Unquo())
	if er != nil {
		err.Panic(er)
	}
	u.Path = fmt.Sprintf("/v3/accounts/%v/instruments", x.Cfg.AccountID.Unquo())
	q := u.Query()

	q.Set("instruments", name.Upper().Unquo()) // Oanda requires uppercase instruments
	u.RawQuery = q.Encode()
	req, er := http.NewRequest("GET", u.String(), nil)
	if er != nil {
		err.Panic(er)
	}
	req.Header.Set("Authorization", "Bearer "+x.Cfg.Token.Unquo())
	resp, er := client.Do(req)
	if er != nil {
		err.Panic(er)
	}
	b, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		err.Panic(er)
	}
	// sys.Log("string(b)", string(b))
	r = &Instr{}
	r.OanJsnRed(string(b))
	r.Prv = x
	x.Instrs[name] = r
	r.CalcMktWeeks(Cfg.MktHr)
	return r
}

func (x *Oan) OpnStart() {
	atomic.StoreUint32(&x.opning, Yes) // TOGGLE RLT
}
func (x *Oan) OpnEnd() {
	atomic.StoreUint32(&x.filing, Yes) // BLOCK POTENTIAL RLT TRDS WHILE FILLING
	for _, i := range x.Instrs {       // instrs instantiated in App.Run
		x.LoadHst(i) // ensure hst is loaded; cfg must have proper rng
		// sys.Logf("Oan.OpnEnd %v RltInrvlMax:%v", i.Name, i.RltInrvlMax)
		start := tme.Now() - i.RltInrvlMax
		// sys.Log("start", start)
		mktWeekIdx := i.MktWeeks.SrchIdx(start)
		if mktWeekIdx == i.MktWeeks.Cnt() {
			start = MktPrvWeekMax(tme.Now()) - i.RltInrvlMax
		}
		// sys.Log("start", start)
		mktWeekIdx = i.MktWeeks.SrchIdx(start)
		if mktWeekIdx == i.MktWeeks.Cnt() {
			sys.Logf("Oan.OpnEnd UNABLE TO FIND START WITHIN HST %v %v", i.Name, start)
			continue
		}
		tmeIdx := i.HstStm.Tmes.SrchIdxEql(start)
		// sys.Logf("tmeIdx:%v %v", tmeIdx, i.HstStm.Tmes.At(tmeIdx))
		// RltStm WILL BE 1S LESS; SO FST RLT TME MAY RAISE EVENT
		// RLT WILL SEND THESE VALUES THROUGH RLT GRAPH
		if tmeIdx < i.HstStm.Tmes.Cnt() {
			sys.Logf(" *** fst:%v lst%v", i.HstStm.Tmes.At(tmeIdx), i.HstStm.Tmes.Lst())
			for n := tmeIdx; n < i.HstStm.Cnt(); n++ {
				x.Ticr.RxMu.Lock() // AGGREGATE WITHIN 1 SECOND TO SIMPLIFY STGY PKT ORDERING AND OVERALL CALCULATION; AS LONG AS tme.Tme is 32-bit
				i.RltStm.PushTic(i.HstStm.Tic(n))
				x.Ticr.RxPktC <- 0
				x.Ticr.RxMu.Unlock()
			}
		}
	}
	atomic.StoreUint32(&x.opning, No)
	go func() { // wait for hst graph processing to complete before allowing rlt trds

		<-time.NewTimer(time.Second * 20).C
		atomic.StoreUint32(&x.filing, No)
	}()
}
func (x *Oan) Load(i *Instr, rng tme.Rng) (r *Stm) { // https://godoc.org/cloud.google.com/go/bigquery
	trc := trc.New("Oan.Load: Network")
	defer trc.End()
	ctx := context.Background()
	client, er := bigquery.NewClient(ctx, x.Cfg.BqProject.Unquo())
	if er != nil {
		err.Panic(er)
	}
	var where string
	if rng.Min != tme.Min || rng.Max != tme.Max {
		var b strings.Builder
		b.WriteString("WHERE ")
		if rng.Min != tme.Min && rng.Max != tme.Max {
			b.WriteString(fmt.Sprintf("Time >= '%v' AND Time < '%v'", rng.Min.Time().Format(time.RFC3339), rng.Max.Time().Format(time.RFC3339)))
		} else if rng.Min != tme.Min {
			b.WriteString(fmt.Sprintf("Time >= '%v'", rng.Min.Time().Format(time.RFC3339)))
		} else {
			b.WriteString(fmt.Sprintf("Time < '%v'", rng.Max.Time().Format(time.RFC3339)))
		}
		where = b.String()
	}
	var limit string
	if x.Cfg.Test {
		limit = "LIMIT 1000"
	}
	// ORDER BY Time
	queryStr := fmt.Sprintf(`SELECT Time, Tick
		FROM %v.%v
		%v

		%v`,
		x.Cfg.Oan.BqDataset.Unquo(), i.Name.Unquo(),
		where, limit,
	)
	sys.Log("Oan.Load: LOADING...")
	sys.Log(queryStr)
	q := client.Query(queryStr)
	it, er := q.Read(ctx)
	if er != nil {
		err.Panic(er)
	}
	j := &jsn.Jsnr{}
	r = NewStm()
	var prsd uint64
	var rows []TicRow
	for { // receive non-sorted
		var row TicRow
		er = it.Next(&row)
		if er != nil {
			if er == iterator.Done {
				break
			}
			err.Panic(er)
		}
		rows = append(rows, row)
	}
	// isSorted := true
	// for n := 1; n < len(rows); n++ {
	// 	if !rows[n-1].Time.Before(rows[n].Time) && !rows[n-1].Time.Equal(rows[n].Time) {
	// 		isSorted = false
	// 		break
	// 	}
	// }
	// sys.Logf("BEFORE isSorted:%v", isSorted)
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Time.Before(rows[j].Time)
	})

	// isSorted = true
	// for n := 1; n < len(rows); n++ {
	// 	// sys.Log("AFTER", rows[n].Time)
	// 	if !rows[n-1].Time.Before(rows[n].Time) && !rows[n-1].Time.Equal(rows[n].Time) {
	// 		isSorted = false
	// 		break
	// 	}
	// }
	// sys.Logf("AFTER isSorted:%v", isSorted)

	for n := 0; n < len(rows); n++ {
		j.Reset(rows[n].Tick)
		if j.Str("type") == "PRICE" && j.Bol("tradeable") {
			r.OanJsnRed(j)
			prsd++
			if Cfg.Oan.GapFil {
				r.GapFil() // fill any data gaps
			}
		}
	}
	// var row TicRow
	// var prsd uint64
	// for {
	// 	er = it.Next(&row)
	// 	if er != nil {
	// 		if er == iterator.Done {
	// 			break
	// 		}
	// 		err.Panic(er)
	// 	}
	// 	j.Reset(row.Tick)
	// 	if j.Str("type") == "PRICE" && j.Bol("tradeable") {
	// 		r.OanJsnRed(j)
	// 		prsd++
	// 		if Cfg.Oan.GapFil {
	// 			r.GapFil() // fill any data gaps
	// 		}
	// 	}
	// }
	if Cfg.Oan.GapFil {
		r.GapFilLst()
	}
	if prsd > 0 {
		sys.Logf("Oan.Load %v (prsd:%v fst:%v lst:%v)\n", i.Name, prsd, r.Tmes.Fst(), r.Tmes.Lst())
	}
	return r
}

func (x *Oan) LoadHst(i *Instr) {
	i.HstMu.Lock()
	defer i.HstMu.Unlock()
	if i.HstStm != nil {
		return
	}
	if x.DskLoadInstrStm(i) { // try load from local disk
		if i.HstStm != nil {
			return
		}
	}
	// sys.Logf("Oan.LoadHst x.Cfg.Hst.Rng:%v", x.Cfg.Hst.Rng)
	i.HstStm = x.Load(i, x.Cfg.Hst.Rng)
	i.CalcStats()
	sys.Logf("Oan.LoadHst: LOADED %v \n", i)
	// for _, t := range *i.HstStm.Tmes {
	// 	sys.Log(t)
	// }
	x.DskSavInstrDetail(i) // save to local disk
	x.DskSavInstrStm(i)
}
func (x *Oan) DskSavInstrStm(i *Instr) {
	if i != nil && i.HstStm != nil && sys.HasDsk() {
		key := &bytes.Buffer{}
		i.Name.BytWrt(key)
		val := &bytes.Buffer{}
		i.HstStm.BytWrt(val)
		sys.Dsk().SavInstrStm(key.Bytes(), val.Bytes())
	}
}
func (x *Oan) DskLoadInstrStm(i *Instr) bool { // use separate method for testing
	if !sys.HasDsk() {
		return false
	}
	trc := trc.New("Oan.DskLoadInstrStm:", i.Name.Unquo())
	defer trc.End()
	key := &bytes.Buffer{}
	i.Name.BytWrt(key)
	val := sys.Dsk().LoadInstrStm(key.Bytes())
	if val == nil {
		return false
	}
	i.HstStm = &Stm{}
	i.HstStm.BytRed(val)
	return true
}
func (x *Oan) DskSavInstrDetail(i *Instr) {
	if i != nil && i.HstStm != nil && sys.HasDsk() {
		key := &bytes.Buffer{}
		i.Name.BytWrt(key)
		val := &bytes.Buffer{}
		i.BytWrt(val)
		sys.Dsk().SavInstrDetail(key.Bytes(), val.Bytes())
		i.CalcMktWeeks(Cfg.MktHr)
	}
}
func (x *Oan) DskLoadInstrDetail(i *Instr) bool {
	if !sys.HasDsk() {
		return false
	}
	key := &bytes.Buffer{}
	i.Name.BytWrt(key)
	val := sys.Dsk().LoadInstrDetail(key.Bytes())
	if val == nil {
		return false
	}
	i.BytRed(val)
	i.CalcMktWeeks(Cfg.MktHr)
	return true
}
func (x *Oan) Sub(i *Instr) {
	// expect Sub to be called once per instr, sub counts done by Instr
	x.subsMu.Lock()
	if x.subs == nil {
		x.subs = make(map[str.Str]*Instr)
		x.rxExitC = make(chan bool, 1)   // init run chan
		x.rxExitedC = make(chan bool, 1) // init exit chan
	}
	i.RltStm = NewStm()
	i.RltSubs = make(map[uint64]TmeIdxRx)
	x.subs[i.Name] = i // track for restartRx()
	if !x.Cfg.Test {
		x.rxRestart()
	}
	x.subsMu.Unlock()
}

func (x *Oan) Unsub(i *Instr) {
	// expect Unsub to be called once per instr, sub counts done by Instr
	x.subsMu.Lock()
	delete(x.subs, i.Name) // track for restartRx()
	if !x.Cfg.Test {
		if len(x.subs) == 0 {
			x.rxExitC <- true // request current http stream to exit
			<-x.rxExitedC     // wait for exit completion
			x.rxExitC = nil
			x.rxExitedC = nil
			x.subs = nil // ensure rx vars are cleared out
		} else {
			x.rxRestart()
		}
	}
	x.subsMu.Unlock()
}

func (x *Oan) rxRestart() { // guarded by x.subsMu
	subs := make(map[str.Str]*Instr) // copy subscribed map to ensure no changes occur in rx loop
	for _, i := range x.subs {
		subs[i.Name] = i
	}
	if len(subs) > 1 {
		x.rxExitC <- true // request current http stream to exit
		<-x.rxExitedC     // wait for exit completion
	}
	go x.rx(subs) // subscribe to network stream
}

func (x *Oan) rx(subs map[str.Str]*Instr) {
	sys.Log("Oan.rx", "START", len(subs))
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		x.rxExitedC <- true
		if x.Ticr != nil {
			x.Ticr.ClsRx()
		}
	}()
	x.Ticr.OpnRx(subs)
	u, er := url.Parse(x.Cfg.BaseStreamURL.Unquo())
	if er != nil {
		err.Panic(er)
	}
	u.Path = fmt.Sprintf("/v3/accounts/%v/pricing/stream", x.Cfg.AccountID.Unquo())
	var buf bytes.Buffer
	n := 0
	for _, i := range subs {
		// sys.Log("Oan.rx", i)
		if n != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(i.Name.Upper().Unquo())
		n++
	}
	q := u.Query()
	q.Set("instruments", buf.String())
	u.RawQuery = q.Encode()
	req, er := http.NewRequest("GET", u.String(), nil)
	if er != nil {
		err.Panic(er)
	}
	req.Header.Set("Authorization", "Bearer "+x.Cfg.Token.Unquo())
	client := &http.Client{
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
	sys.Log("Oan.rx", "req", req)
	res, er := client.Do(req)
	if er != nil {
		err.Panic(er)
	}

	j := &jsn.Jsnr{}
	rdr := bufio.NewReader(res.Body)
	for {
		select {
		case <-x.rxExitC:
			return
		default:
			ln, er := rdr.ReadBytes('\n')
			if er != nil {
				err.Panic(er)
			}
			if atomic.LoadUint32(&x.opning) == Yes { // OPN PENDING
				x.opningPending = append(x.opningPending, string(ln))
				continue
			}
			if len(x.opningPending) != 0 {
				// atomic.StoreUint32(&x.filing, Yes)          // BLOCK POTENTIAL RLT TRDS WHILE FILLING
				for _, pendingLn := range x.opningPending { // SEND PENDING VALS
					j.Reset(pendingLn)
					if j.Str("type") == "PRICE" && j.Bol("tradeable") {
						i := subs[j.Str("instrument").Lower()]
						x.Ticr.RxMu.Lock()
						i.RltStm.OanJsnRed(j)
						sys.Log("Oan.rx", i.Name, "lstPktTme", i.RltStm.Tmes.Lst())
						x.Ticr.RxPktC <- 0 // signal pkt received
						x.Ticr.RxMu.Unlock()
					}
				}
				x.opningPending = nil
				// atomic.StoreUint32(&x.filing, No)
			}

			// sys.Logf(string(ln))
			j.Reset(string(ln))
			if j.Str("type") == "PRICE" && j.Bol("tradeable") { // tradeable tic
				i := subs[j.Str("instrument").Lower()]
				x.Ticr.RxMu.Lock()    // AGGREGATE WITHIN 1 SECOND TO SIMPLIFY STGY PKT ORDERING AND OVERALL CALCULATION; AS LONG AS tme.Tme is 32-bit
				i.RltStm.OanJsnRed(j) // RltStm.OanJsnRed AGGREGATES WITHIN 1S; WILL BE READ BY TICR
				sys.Log("Oan.rx", i.Name, "lstPktTme", i.RltStm.Tmes.Lst())
				x.Ticr.RxPktC <- 0 // signal pkt received
				x.Ticr.RxMu.Unlock()
			} else if j.Str("type") == "HEARTBEAT" { // HEARTBEAT
				x.Ticr.RxMu.Lock()
				x.Ticr.RxHeartC <- j.StrTme("time")
				x.Ticr.RxMu.Unlock()
			}
		}
	}
}

func (x *Oan) AcntRefresh() flt.Flt {
	client := &http.Client{ // load from network
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
	u, er := url.Parse(x.Cfg.BaseURL.Unquo())
	if er != nil {
		err.Panic(er)
	}
	u.Path = fmt.Sprintf("/v3/accounts/%v/summary", x.Cfg.AccountID.Unquo())
	req, er := http.NewRequest("GET", u.String(), nil)
	if er != nil {
		err.Panic(er)
	}
	req.Header.Set("Authorization", "Bearer "+x.Cfg.Token.Unquo())
	resp, er := client.Do(req)
	if er != nil {
		err.Panic(er)
	}
	b, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		err.Panic(er)
	}
	x.acnt = &Acnt{}
	x.acnt.JsnRed(string(b))
	if Cfg.Trc.IsPrv() {
		sys.Logf("ana.Oan.AcntRefresh %p acnt %v", x, x.acnt)
	}
	return x.acnt.Balance
}

func (x *Oan) MayTrd() bol.Bol { return atomic.LoadUint32(&x.trding) == No }

// ONE TRD AT A TIME ACROSS ENTIRE ACCOUNT
func (x *Oan) OpnTrd(t *Trd, i *Instr) (bol.Bol, TrdRsnOpn) {
	if atomic.LoadUint32(&x.trding) == Yes {
		return false, InTrd // already in trd
	}
	x.trdMu.Lock()
	defer x.trdMu.Unlock()
	if Cfg.Trc.IsPrv() {
		sys.Logf("ana.Oan.OpnTrd %p START %v", x, t)
	}
	if atomic.LoadUint32(&x.trding) == Yes {
		return false, InTrd // already in trd
	}
	if atomic.LoadUint32(&x.filing) == Yes {
		return false, FilHstGap // IMPORTANT: FILLING RLT HST GAP; DO NOT OPN TRDS WITH HST GAP FILLED DATA
	}
	// CALCULATE OPEN QTY
	x.CalcOpn(t, i)
	if t.Units <= 0 { // set by CalcOpn
		return false, NoCapital
	}
	if Cfg.Test { // FOR RLT TESTING ONLY
		atomic.StoreUint32(&x.trding, Yes)
		return true, NoTrdRsnOpn
	}
	ordReq := OrdReq{ // TODO: SWITCH TO STOP LOSS ORDER
		Instrument:  i.Name.Upper(),
		Type:        "MARKET",
		TimeInForce: "FOK", // FILL OR KILL
		Units:       bscint.Int(t.Units),
	}
	if !t.IsLong {
		ordReq.Units *= -1 // a negative number of units results in a short Order.
	}
	// Ask: the price for me to buy at
	// Bid: the price for me to sell at
	if Cfg.Oan.PriceBnd { // REQUIRE NO PRICE SLIPPAGE
		if t.IsLong {
			ordReq.PriceBound = t.OpnAsk
		} else {
			ordReq.PriceBound = t.OpnBid
		}
	}

	ordReqBuf := &strings.Builder{}
	ordReq.JsnWrt(i, ordReqBuf)
	u, er := url.Parse(x.Cfg.BaseURL.Unquo())
	if er != nil {
		sys.Log(er)
		return false, PrvErr
	}
	u.Path = fmt.Sprintf("/v3/accounts/%v/orders", x.Cfg.AccountID.Unquo())
	if Cfg.Trc.IsPrv() {
		sys.Logf("ana.Oan.OpnTrd %p    u.Path %v", x, u.Path)
		sys.Logf("ana.Oan.OpnTrd %p ordReqBuf %v", x, ordReqBuf.String())
	}
	t.OpnReq = str.Str(ordReqBuf.String())
	// req, er := http.NewRequest("POST", u.String(), strings.NewReader(ordReqBuf.String()))
	req, er := http.NewRequest("POST", u.String(), bytes.NewBufferString(ordReqBuf.String()))
	if er != nil {
		sys.Log(er)
		return false, PrvErr
	}
	// sys.Log("Oan.OpnTrd", "x.Cfg.Token", x.Cfg.Token.Unquo())
	req.Header.Set("Authorization", "Bearer "+x.Cfg.Token.Unquo())
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
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
	res, er := client.Do(req)
	if er != nil {
		sys.Log(er)
		return false, PrvReject
	}
	// if Cfg.Trc.IsPrv() {
	// 	sys.Logf("ana.Oan.OpnTrd %p res %v", x, res)
	// }
	defer res.Body.Close()

	var body string
	bodyBytes, _ := ioutil.ReadAll(res.Body)
	body = string(bodyBytes)
	t.OpnRes = str.Str(body)
	if res.StatusCode != http.StatusCreated {
		sys.Logf("http err: %v %v \n", res.Status, body)
		return false, PrvReject
	}
	if strings.Index(body, "ORDER_FILL") >= 0 {
		if Cfg.Trc.IsPrv() {
			sys.Logf("ana.Oan.OpnTrd %p OPEN: SUCCESS", x)
		}
		atomic.StoreUint32(&x.trding, Yes)
		x.AcntRefresh()
		return true, NoTrdRsnOpn
	}
	if Cfg.Trc.IsPrv() {
		sys.Logf("ana.Oan.OpnTrd %p OPEN: FAIL", x)
		sys.Logf("ana.Oan.OpnTrd %p body %v", x, body)
	}
	return false, PrvReject
}

func (x *Oan) ClsTrd(t *Trd, i *Instr) (ok bol.Bol) {
	if atomic.LoadUint32(&x.trding) == No {
		return false
	}
	x.trdMu.Lock()
	defer x.trdMu.Unlock()
	if Cfg.Trc.IsPrv() {
		sys.Logf("ana.Oan.ClsTrd %p START %v", x, t)
	}
	if atomic.LoadUint32(&x.trding) == No {
		return false
	}
	if Cfg.Test { // FOR RLT TESTING ONLY
		atomic.StoreUint32(&x.trding, No)
		return true
	}
	posClsReqBuf := &strings.Builder{}
	posClsReq := PosClsReq{IsBuy: t.IsLong}
	posClsReq.JsnWrt(posClsReqBuf)
	t.ClsReq = str.Str(posClsReqBuf.String())
	u, er := url.Parse(x.Cfg.BaseURL.Unquo())
	if er != nil {
		err.Panic(er)
	}
	u.Path = fmt.Sprintf("/v3/accounts/%v/positions/%v/close", x.Cfg.AccountID.Unquo(), i.Name.Upper().Unquo())
	req, er := http.NewRequest("PUT", u.String(), bytes.NewBufferString((posClsReqBuf.String())))
	if er != nil {
		err.Panic(er)
	}
	req.Header.Set("Authorization", "Bearer "+x.Cfg.Token.Unquo())
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
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

	// sys.Log("Oan.ClsTrd: req", req)
	res, er := client.Do(req)
	if er != nil {
		err.Panic(er)
	}
	if Cfg.Trc.IsPrv() {
		sys.Logf("ana.Oan.ClsTrd %p res %v", x, res)
	}
	defer res.Body.Close()
	resBdy, er := ioutil.ReadAll(res.Body)
	if er != nil {
		err.Panic(er)
	}
	t.ClsRes = str.Str(string(resBdy))
	if res.StatusCode != http.StatusOK {
		sys.Log("ClsTrd", i.Name, "fail", "HTTP error", res.Status, string(resBdy))
		return false
	}
	atomic.StoreUint32(&x.trding, No)
	// x.CalcCls(t, i) // called in stgy
	x.AcntRefresh()
	t.ClsBalUsdAct = x.acnt.Balance
	return true
}

func (x *Oan) CloneInstrs(vs map[str.Str]*Instr) { // for testing
	for _, v := range vs {
		i := v.Cpy()
		x.Instrs[v.Name] = i
		i.HstStm = v.HstStm     // shared instance
		i.MktWeeks = v.MktWeeks // shared instance
		i.MktDays = v.MktDays   // shared instance
		i.Prv = x
	}
}

func (x *Oan) Ifc() interface{} { return x }

func (x *Oan) CalcOpn(t *Trd, i *Instr) {
	// SEPARATE METHOD ALLOWS:
	// 	* HST PLL NOT TO BE LOCKED BY OPNTRD
	//	* RLT TO USE SAME METHOD AS HST

	// sys.Log("x.acnt == nil", x.acnt == nil)
	t.Instr = i.Name
	t.MrgnRtio = i.MrgnRtio
	t.OpnBalUsd = x.acnt.Balance
	opnBalUsedUsd := t.OpnBalUsd * t.TrdPct // trdPct is appetite for risk
	opnBalMrgnUsd := opnBalUsedUsd * t.MrgnRtio
	if t.IsLong { // Ask: the price for me to buy at
		t.Units = flt.Flt(bscint.Int(opnBalMrgnUsd / t.OpnAsk))
	} else { // Bid: the price for me to sell at
		t.Units = flt.Flt(bscint.Int(opnBalMrgnUsd / t.OpnBid))
	}
}

func (x *Oan) CalcCls(t *Trd, i *Instr) {
	opnBalUsedUsd := t.OpnBalUsd * t.TrdPct
	opnBalUnusedUsd := t.OpnBalUsd - opnBalUsedUsd
	opnBalMrgnUsd := opnBalUsedUsd * t.MrgnRtio
	var clsBalMrgnUsd flt.Flt
	if t.IsLong { // Ask: the price for me to buy at
		clsBalMrgnUsd = t.ClsBid * t.Units
		t.PnlGrsUsd = clsBalMrgnUsd - opnBalMrgnUsd
		t.Pip = t.ClsBid.Sub(t.OpnAsk).Div(i.Pip).Trnc(2)
	} else { // Bid: the price for me to sell at
		clsBalMrgnUsd = t.ClsAsk * t.Units
		t.PnlGrsUsd = opnBalMrgnUsd - clsBalMrgnUsd
		t.Pip = t.OpnBid.Sub(t.ClsAsk).Div(i.Pip).Trnc(2)
	}
	clsBalUsedUsd := clsBalMrgnUsd / t.MrgnRtio
	t.CstComUsd = (Cfg.Oan.ComPerUnitUsd * 2) * t.Units     // commission cost for buy and sell
	t.CstOpnSpdUsd = ((t.OpnAsk - t.OpnBid) * .5) * t.Units // hlf spd opn cst
	t.CstClsSpdUsd = ((t.ClsAsk - t.ClsBid) * .5) * t.Units // hlf spd cls cst
	t.ClsBalUsd = opnBalUnusedUsd + clsBalUsedUsd + t.PnlGrsUsd - t.CstComUsd
	t.PnlUsd = t.ClsBalUsd.Sub(t.OpnBalUsd).Trnc(2)
	t.PnlPct = t.PnlUsd.Div(t.OpnBalUsd).Mul(100).Trnc(3)
	// TODO: RLT: UPDATE WITH ACTUAL ACCOUNT BALANCE?
	// if compound {
	// 	x.BalLstUsd = t.ClsBalUsd // update portfolio balance
	// }
}
