package dsk

import (
	"encoding/binary"
	"path/filepath"
	"sys/err"
	"sys/fs"

	"github.com/dgraph-io/badger"
)

const ( // declare explicit value to avoid maintenance bugs
	InstrStm    Prefix = 1
	InstrDetail Prefix = 2
	Ml          Prefix = 4
)

type (
	Prefix uint16
	// Dskr is disk storage.
	Dskr struct {
		db *badger.DB
	}
)

// New creates a Dsk.
func New(dir string, name ...string) (r *Dskr) {
	r = &Dskr{}
	opt := badger.DefaultOptions
	if len(name) == 0 {
		opt.Dir = dir
	} else {
		opt.Dir = filepath.Join(dir, name[0])
	}
	opt.ValueDir = opt.Dir
	fs.EnsureDir(opt.Dir)
	var er error
	r.db, er = badger.Open(opt)
	if er != nil {
		err.Panic(er)
	}
	return r
}

// Close closes a Dsk.
func (x *Dskr) Cls() {
	if x.db != nil {
		er := x.db.Close()
		if er != nil {
			err.Panic(er)
		}
	}
}

// Sav saves bytes to disk.
func (x *Dskr) Sav(key []byte, val []byte) {
	er := x.db.Update(func(tx *badger.Txn) error {
		return tx.Set(key, val)
	})
	if er != nil {
		err.Panic(er)
	}
}

// Load loads bytes from disk.
func (x *Dskr) Load(key []byte, retNil ...bool) (r []byte) {
	//x.mu.Lock()
	er := x.db.View(func(tx *badger.Txn) (er error) {
		item, err := tx.Get(key)
		if err != nil {
			return err
		}
		r, er = item.Value()
		return er
	})
	//x.mu.Unlock()
	if er != nil && (len(retNil) == 0 || !retNil[0]) {
		err.Panic(er)
	}
	return r
}

// Loads loads keys and values with the specified key prefix.
func (x *Dskr) Loads(keyPrefix ...byte) (ks, vs [][]byte) {
	x.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		for it.Seek(keyPrefix); it.ValidForPrefix(keyPrefix); it.Next() {
			item := it.Item()
			v, err := item.Value()
			if err != nil {
				return err
			}
			ks = append(ks, item.Key())
			vs = append(vs, v)
		}
		return nil
	})
	return ks, vs
}

// Del deletes bytes from disk.
func (x *Dskr) Del(key []byte) {
	//x.mu.Lock()
	er := x.db.Update(func(tx *badger.Txn) error {
		return tx.Delete(key)
	})
	//x.mu.Unlock()
	if er != nil {
		err.Panic(er)
	}
}

// Upd runs a disk transaction.
func (x *Dskr) Upd(f func(tx *badger.Txn) error) {
	//x.mu.Lock()
	er := x.db.Update(f)
	//x.mu.Unlock()
	if er != nil {
		err.Panic(er)
	}
}

func (x Prefix) Key(elmKey []byte) []byte {
	pb := make([]byte, 4) // array length
	binary.LittleEndian.PutUint32(pb, uint32(x))
	return append(pb, elmKey...)
}

func (x *Dskr) SavInstrStm(key []byte, val []byte) { x.Sav(InstrStm.Key(key), val) }
func (x *Dskr) LoadInstrStm(key []byte) []byte     { return x.Load(InstrStm.Key(key), true) }
func (x *Dskr) DelInstrStm(key []byte)             { x.Del(InstrStm.Key(key)) }

func (x *Dskr) SavInstrDetail(key []byte, val []byte) { x.Sav(InstrDetail.Key(key), val) }
func (x *Dskr) LoadInstrDetail(key []byte) []byte     { return x.Load(InstrDetail.Key(key), true) }
func (x *Dskr) DelInstrDetail(key []byte)             { x.Del(InstrDetail.Key(key)) }

func (x *Dskr) SavMl(key []byte, val []byte) { x.Sav(Ml.Key(key), val) }
func (x *Dskr) LoadMl(key []byte) []byte     { return x.Load(Ml.Key(key), true) }
func (x *Dskr) DelMl(key []byte)             { x.Del(Ml.Key(key)) }
