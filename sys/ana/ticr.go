package ana

import (
	"sync"
	"sys"
	"sys/ana/cfg"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/err"
	"sys/run"
	"sys/trc"
)

const (
	NextTier = 0
)

type (
	Ticr struct {
		RxPktC        chan byte
		RxHeartC      chan tme.Tme
		RxTicsExitC   chan bool
		RxTicsExitedC chan bool
		RxInstrs      map[str.Str]*Instr
		// RxPktTme      tme.Tme
		RxMu sync.Mutex

		TxTicsC       chan InstrTic // sent by Oan; monitored by rlt
		TxTicsExitedC chan bool

		Cfg *cfg.Cfg
	}
	Tx interface {
		Ret() []sys.Act
		Tier() int
		DecTier()
	}
)

func NewTicr(cfg *cfg.Cfg) (r *Ticr) {
	r = &Ticr{}
	r.Cfg = cfg
	r.OpnTx()
	return r
}
func (x *Ticr) Cls() {
	x.ClsRx()
	x.ClsTx()
}
func (x *Ticr) OpnTx() {
	x.Cls()
	x.TxTicsC = make(chan InstrTic, 4096)
	x.TxTicsExitedC = make(chan bool)
	go x.tx()
}
func (x *Ticr) ClsTx() {
	if x.TxTicsC != nil {
		close(x.TxTicsC)
		for len(x.TxTicsC) > 0 {
			<-x.TxTicsC
		}
		<-x.TxTicsExitedC
		close(x.TxTicsExitedC)
		x.TxTicsC = nil
		x.TxTicsExitedC = nil
	}
}

func (x *Ticr) OpnRx(instrs map[str.Str]*Instr) {
	x.RxPktC = make(chan byte, 2048)
	x.RxHeartC = make(chan tme.Tme, 2048)
	x.RxTicsExitC = make(chan bool)
	x.RxTicsExitedC = make(chan bool)
	x.RxInstrs = instrs
	// x.RxPktTme = fstPktTme // - tme.S1 // init pkt tme to 1s before
	// x.RxPktC <- 0          // signal fst pkt received
	go x.rx()
}
func (x *Ticr) ClsRx() {
	if x.RxPktC != nil {
		x.RxTicsExitC <- true
		<-x.RxTicsExitedC
		x.RxPktC = nil
		x.RxHeartC = nil
		x.RxTicsExitC = nil
		x.RxTicsExitedC = nil
		x.RxInstrs = nil
		// x.RxPktTme = 0
	}
}

func (x *Ticr) rx() {
	// tic time must be based on pkt time (not local machine time)
	// to be insync with the actual data
	// heartbeat will e used to replay historical data
	// through realtime for testing parity between hst and rlt
	// therefore use tic pkt time and not local machine time
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		x.RxTicsExitedC <- true
	}()
	var durr *trc.Durr
	for {
		select {
		case <-x.RxTicsExitC:
			return
		case heartTme := <-x.RxHeartC: // RX HEARTBEAT; SEND GAP-FILLED VALUES ONLY
			x.RxMu.Lock()
			if x.Cfg.Trc.IsTicrRx() {
				sys.Logf("Ticr.rx ))) RX HEARTBEAT lcl:%v heartTme:%v", tme.Now(), heartTme)
				durr = trc.NewDurr("Ticr.rx")
			}
			for _, i := range x.RxInstrs {
				if i.RltStm.Cnt() > 0 {
					i.RltStm.GapFilTo(heartTme) // attempt to gap fill any not filled by ticker; result will be gap fil burst
					for i.RltStm.RxIdx < i.RltStm.Cnt() {
						pkt := TmeIdx{Tme: i.RltStm.Tmes.At(i.RltStm.RxIdx), Idx: i.RltStm.RxIdx}
						// sys.Log("Ticr.rx ))) RX HEARTBEAT SEND", pkt, "LST RX TME", i.RltStm.RxTme, "LST IDX", i.RltStm.RxIdx)
						x.TxTicsC <- InstrTic{I: i, Pkt: pkt}
						i.RltStm.RxIdx++
					}
					i.RltStm.RxTme = i.RltStm.Tmes.Lst() // 1s minimum resolution for 32-bit tme.Tme
				}
			}
			if x.Cfg.Trc.IsTicrRx() {
				sys.Logf("Ticr.rx ((( RX HEARTBEAT lcl:%v %v", tme.Now(), durr.Dur())
			}
			x.RxMu.Unlock()

		case <-x.RxPktC: // RX PKT
			x.RxMu.Lock()
			if x.Cfg.Trc.IsTicrRx() {
				sys.Logf("Ticr.rx ))) RX PKT lcl:%v", tme.Now())
				durr = trc.NewDurr("Ticr.rx")
			}
			for _, i := range x.RxInstrs {
				if i.RltStm.Cnt() > 0 {
					if i.RltStm.RxIdx != i.RltStm.Cnt() { // more tics available to send?
						i.RltStm.GapFil() // attempt to gap fill any not filled by ticker; result will be gap fil burst
						for i.RltStm.RxIdx < i.RltStm.Cnt() {
							pkt := TmeIdx{Tme: i.RltStm.Tmes.At(i.RltStm.RxIdx), Idx: i.RltStm.RxIdx}
							// sys.Log("Ticr.rx ))) RX PKT SEND", pkt, "LST RX TME", i.RltStm.RxTme, "LST IDX", i.RltStm.RxIdx)
							x.TxTicsC <- InstrTic{I: i, Pkt: pkt}
							i.RltStm.RxIdx++
						}
						i.RltStm.RxTme = i.RltStm.Tmes.Lst() // 1s minimum resolution for 32-bit tme.Tme
					}
				}
			}
			if x.Cfg.Trc.IsTicrRx() {
				sys.Logf("Ticr.rx ((( RX PKT lcl:%v %v", tme.Now(), durr.Dur())
			}
			x.RxMu.Unlock()
		}
	}
}

// Ensure each tick is processed in chronological order through out the entire subscription graph.
// Process each tick through entire subscription graph before processing next tick.
// Within graph, process single graph tier at a time.
func (x *Ticr) tx() {
	// send heartbeat so that inrvls complete without waiting for network pkts
	// inrvls between network pkts may be 10s or more
	// ability to process the rlt graph within a second is important
	defer func() {
		if v := recover(); v != nil {
			switch t := v.(type) {
			case *err.Err:
				sys.Log(t.Full())
			default:
				sys.Log(err.New(v).Full())
			}
		}
		x.TxTicsExitedC <- true
	}()
	var durr *trc.Durr
	txsA, txsB := run.NewActs(), run.NewActs()
	for {
		select {
		case tic, open := <-x.TxTicsC: // tic may be Heartbeat or tradeable
			if !open {
				return
			}
			if x.Cfg.Trc.IsTicrTx() {
				sys.Log("Ticr.tx >>> TX", tic.I.Name, tic.Pkt, "x.TxTicsC", len(x.TxTicsC))
				durr = trc.NewDurr("Ticr.tx")
			}
			tic.I.RltSubsMu.Lock() // gather root graph tier
			for _, rx := range tic.I.RltSubs {
				txsA.Push(&TmeIdxTx{Pkt: tic.Pkt, Rx: rx})
			}
			tic.I.RltSubsMu.Unlock()
			for txsA.Cnt() > 0 {
				for n := txsA.Cnt() - 1; n > -1; n-- { // push later tier items back
					tx := txsA.Elm(n).(Tx)
					if tx.Tier() != NextTier {
						tx.DecTier()
						txsB.Push(tx.(sys.Act)) // push item to later tier
						txsA.Del(n)
					}
				}
				sys.Run().Pll(*txsA...)     // process single graph tier in parallel
				for _, act := range *txsA { // prepare next graph tier
					txsB.Push(act.(Tx).Ret()...)
				}
				txsA.Clr()              // clear previous graph tier
				txsA, txsB = txsB, txsA // swap to next graph tier
			}
			txsA.Clr()
			txsB.Clr()

			if x.Cfg.Trc.IsTicrTx() {
				sys.Log("Ticr.tx <<< TX", tic.I.Name, tic.Pkt, durr.Dur())
			}
		}
	}
}
