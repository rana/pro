package cld

import (
	"context"
	"sync"
	"sys/err"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

type ( // https://godoc.org/cloud.google.com/go/bigquery
	Cldr struct {
		client      *bigquery.Client
		dataset     *bigquery.Dataset
		projectName string
		datasetName string
		tables      map[string]*bigquery.Table
		mu          sync.Mutex
	}
	Tblr interface {
		CldName() string
		CldSchema() bigquery.Schema
	}
)

func New(project, dataset string) (r *Cldr) {
	r = &Cldr{}
	r.projectName = project
	r.datasetName = dataset
	return r
}
func (x *Cldr) Cfg() (projectName, datasetName string) { return x.projectName, x.datasetName }
func (x *Cldr) Mu() *sync.Mutex                        { return &x.mu }
func (x *Cldr) init() {
	if x.client == nil {
		var er error
		var opts []option.ClientOption
		// flepath := filepath.Join(c.Wd, "pro-cld.json")
		// // fmt.Println("loading...", flepath)
		// opts = append(opts, option.WithServiceAccountFile(flepath))
		ctx := context.Background()
		x.client, er = bigquery.NewClient(ctx, x.projectName, opts...)
		if er != nil {
			err.Panic(er)
		}
		x.dataset = x.client.Dataset(x.datasetName)
		x.tables = make(map[string]*bigquery.Table)
	}
}
func (x *Cldr) GetTable(v interface{}) (r *bigquery.Table) {
	x.init()
	var ok bool
	tblr, ok := v.(Tblr)
	if !ok {
		err.Panicf("cld: does not implement Tblr interface")
	}
	r, ok = x.tables[tblr.CldName()]
	if !ok {
		r = x.dataset.Table(tblr.CldName())
		x.tables[tblr.CldName()] = r
		r.Create(context.Background(), &bigquery.TableMetadata{Schema: tblr.CldSchema()})
	}

	return r
}
func (x *Cldr) Query(txt string) *bigquery.RowIterator {
	x.init()
	q := x.client.Query(txt)
	ctx := context.Background()
	itr, er := q.Read(ctx)
	if er != nil {
		err.Panic(er)
	}
	return itr
}
