# pro

A sophisticated algorithmic trading system, compiler, and domain-specific language written in Go.

Here's a summary of the key components:

1. Core System Architecture:
- `sys.go`: Contains the main `Sys` struct with interfaces for core functionalities:
  - `IRunr`: Runner interface for sequential/parallel execution
  - `ICldr`: Cloud interface for BigQuery operations
  - `IDskr`: Storage interface for instrument data
  - `IActr`: Actor interface for running different types of operations
  - `ILrnr`: Machine learning interface for model fitting/prediction

1. Major Components:
- `/ana`: Analytics package for historical and real-time analysis
- `/app`: UI application components 
- `/bsc`: Basic data types and operations
- `/cld`: Cloud services integration
- `/dsk`: Disk storage operations
- `/lng`: Language processing (includes JSON and custom DSL parser)
- `/run`: Runtime execution engine
- `/vis`: Visualization package for charts and graphics
- `/tpl`: Templates for code generation
- `/tst`: Testing utilities

1. Key Features:
- Custom DSL (Domain Specific Language) for trading strategies
- Real-time data processing capabilities
- Integration with BigQuery for data storage/retrieval
- Machine learning model integration
- Visualization tools for market data
- Comprehensive testing framework
- Both historical and real-time analysis capabilities

1. Project Scale:
- ~288,552 lines of code
- 64 directories
- 503 files
- Extensive use of code generation (many `.gen.go` files)

A production-grade trading system with capabilities for:
- Strategy development and backtesting
- Real-time market data processing
- Machine learning integration
- Data visualization
- Cloud integration
- High-performance execution

The architecture emphasizes modularity through interfaces and is designed for high-performance trading operations with both historical analysis and real-time execution capabilities.

# Project File Tree

64 directories, 503 files

```sh
> pwd
/home/rana/prj/pro
> tree
.
├── LICENSE
├── README.md
└── sys
    ├── act.gen.go
    ├── ana
    │   ├── ana.go
    │   ├── cfg
    │   │   ├── cfg.go
    │   │   └── cfg_test.go
    │   ├── hst
    │   │   ├── cndBse_test.go
    │   │   ├── cnd.gen.go
    │   │   ├── cnd.gen_test.go
    │   │   ├── cnds.gen.go
    │   │   ├── cnd_test.go
    │   │   ├── ftr.go
    │   │   ├── ftrs.npz
    │   │   ├── inrvl.gen.go
    │   │   ├── inrvl.gen_test.go
    │   │   ├── inrvls.gen.go
    │   │   ├── instr.gen.go
    │   │   ├── instr.gen_test.go
    │   │   ├── instrs.gen.go
    │   │   ├── mdl
    │   │   │   └── lab.py
    │   │   ├── port.go
    │   │   ├── port_test.go
    │   │   ├── prcp.go
    │   │   ├── prv.gen.go
    │   │   ├── prv.gen_test.go
    │   │   ├── prvs.gen.go
    │   │   ├── side.gen.go
    │   │   ├── side.gen_test.go
    │   │   ├── sides.gen.go
    │   │   ├── srch
    │   │   │   ├── perm.go
    │   │   │   ├── perm_test.go
    │   │   │   ├── srch0.go
    │   │   │   ├── srchFam.go
    │   │   │   ├── srch.go
    │   │   │   ├── srchMl0.go
    │   │   │   ├── srchPerm0.go
    │   │   │   ├── srchPerm1.go
    │   │   │   └── srch_test.go
    │   │   ├── stgy.gen.go
    │   │   ├── stgy.gen_test.go
    │   │   ├── stgy.go
    │   │   ├── stgyRlng.go
    │   │   ├── stgys.gen.go
    │   │   ├── stgy_test.go
    │   │   ├── stmBse_test.go
    │   │   ├── stm.gen.go
    │   │   ├── stm.gen_test.go
    │   │   ├── stm.go
    │   │   └── stms.gen.go
    │   ├── instr.gen.go
    │   ├── instr.gen_test.go
    │   ├── instr.go
    │   ├── json
    │   │   ├── acnt.json
    │   │   ├── hst.json
    │   │   ├── instr.json
    │   │   ├── ordCancel.json
    │   │   ├── ordFill.json
    │   │   ├── stream_heartbeat.json
    │   │   ├── stream_price.json
    │   │   └── tics.json
    │   ├── ml
    │   │   └── lrnr.go
    │   ├── oanAcnt.go
    │   ├── oan.go
    │   ├── oan_test.go
    │   ├── ord.go
    │   ├── port.gen.go
    │   ├── prfmDlt.gen.go
    │   ├── prfm.gen.go
    │   ├── prfm.gen_test.go
    │   ├── prfm.go
    │   ├── prfms.gen.go
    │   ├── prv.gen.go
    │   ├── pth.gen.go
    │   ├── rlt
    │   │   ├── cnd.gen.go
    │   │   ├── cnd.gen_test.go
    │   │   ├── cnds.gen.go
    │   │   ├── cnd_test.go
    │   │   ├── inrvl.gen.go
    │   │   ├── inrvl.gen_test.go
    │   │   ├── inrvls.gen.go
    │   │   ├── instr.gen.go
    │   │   ├── instr.gen_test.go
    │   │   ├── instrs.gen.go
    │   │   ├── port.go
    │   │   ├── port_test.go
    │   │   ├── prv.gen.go
    │   │   ├── prv.gen_test.go
    │   │   ├── prvs.gen.go
    │   │   ├── rlt.go
    │   │   ├── side.gen.go
    │   │   ├── side.gen_test.go
    │   │   ├── sides.gen.go
    │   │   ├── stgy.gen.go
    │   │   ├── stgy.gen_test.go
    │   │   ├── stgy.go
    │   │   ├── stgys.gen.go
    │   │   ├── stgy_test.go
    │   │   ├── stm.gen.go
    │   │   ├── stm.gen_test.go
    │   │   ├── stm.go
    │   │   └── stms.gen.go
    │   ├── stm.gen.go
    │   ├── stm.go
    │   ├── stm_test.go
    │   ├── tic.gen.go
    │   ├── ticr.go
    │   ├── tics
    │   │   ├── pro-cld.json
    │   │   ├── tics.go
    │   │   └── tics.sh
    │   ├── tmeFlt.gen.go
    │   ├── tmeFlts.gen.go
    │   ├── tmeIdx.gen.go
    │   ├── tmeIdxs.gen.go
    │   ├── trd.gen.go
    │   ├── trd.go
    │   ├── trdRsnCls.gen.go
    │   ├── trdRsnCls.go
    │   ├── trdRsnOpn.gen.go
    │   ├── trds.gen.go
    │   ├── trd_test.go
    │   └── vis
    │       ├── clr
    │       │   ├── clr.gen.go
    │       │   └── clr.go
    │       ├── fnt
    │       │   ├── fnt.gen.go
    │       │   ├── fnt.go
    │       │   └── roboto
    │       │       └── roboto.go
    │       ├── len.go
    │       ├── mtrx.go
    │       ├── pen
    │       │   ├── pen.gen.go
    │       │   └── pens.gen.go
    │       ├── plt
    │       │   ├── cndStk.go
    │       │   ├── dpth.gen.go
    │       │   ├── dpth.go
    │       │   ├── fltAxisX.go
    │       │   ├── fltAxisY.gen.go
    │       │   ├── fltAxisY.go
    │       │   ├── flts.go
    │       │   ├── fltsSctrDist.gen.go
    │       │   ├── fltsSctrDist.go
    │       │   ├── fltsSctr.gen.go
    │       │   ├── fltsSctr.go
    │       │   ├── hrzBnd.go
    │       │   ├── hrz.gen.go
    │       │   ├── hrz.go
    │       │   ├── hrzLn.go
    │       │   ├── pltBse.go
    │       │   ├── plt.gen.go
    │       │   ├── plt.go
    │       │   ├── plts.gen.go
    │       │   ├── plt_test.go
    │       │   ├── prcp.go
    │       │   ├── prcpSplt.go
    │       │   ├── serTrd.go
    │       │   ├── stgy.go
    │       │   ├── stmBnd.go
    │       │   ├── stm.gen.go
    │       │   ├── stm.go
    │       │   ├── stmSplt.go
    │       │   ├── stmStk.go
    │       │   ├── stmStkTrd.go
    │       │   ├── tmeAxisX.gen.go
    │       │   ├── tmeAxisX.go
    │       │   ├── vrtBnd.go
    │       │   ├── vrt.gen.go
    │       │   ├── vrt.go
    │       │   └── vrtLn.go
    │       ├── pnt.go
    │       ├── pnt_test.go
    │       ├── pos.go
    │       ├── pxl.go
    │       ├── rct.go
    │       ├── shp.go
    │       ├── siz.go
    │       ├── stk.go
    │       ├── stk_test.go
    │       └── vis.go
    ├── app
    │   ├── alignment.go
    │   ├── app.go
    │   ├── borderBse.go
    │   ├── box.go
    │   ├── ctrl.go
    │   ├── event.go
    │   ├── fle.go
    │   ├── HList.go
    │   ├── item.go
    │   ├── keybinding.go
    │   ├── label.go
    │   ├── length.go
    │   ├── list.go
    │   ├── logger.go
    │   ├── painter.go
    │   ├── proBox.go
    │   ├── prtBorder.go
    │   ├── prtTitle.go
    │   ├── readBox.go
    │   ├── sizePolicy.go
    │   ├── spacer.go
    │   ├── surface.go
    │   ├── tab.go
    │   ├── text.go
    │   ├── theme.go
    │   ├── uiBse.go
    │   ├── ui.go
    │   └── VList.go
    ├── bsc
    │   ├── bnd
    │   │   ├── bnd.gen.go
    │   │   └── bnd.gen_test.go
    │   ├── bnds
    │   │   └── bnds.gen.go
    │   ├── bol
    │   │   ├── bol.gen.go
    │   │   └── bol.gen_test.go
    │   ├── bols
    │   │   └── bols.gen.go
    │   ├── flt
    │   │   ├── flt.gen.go
    │   │   ├── flt.gen_test.go
    │   │   ├── rng.gen.go
    │   │   └── rng.gen_test.go
    │   ├── flts
    │   │   └── flts.gen.go
    │   ├── int
    │   │   ├── int.gen.go
    │   │   └── int.gen_test.go
    │   ├── ints
    │   │   └── ints.gen.go
    │   ├── str
    │   │   ├── str.gen.go
    │   │   └── str.gen_test.go
    │   ├── strs
    │   │   └── strs.gen.go
    │   ├── tme
    │   │   ├── rng.gen.go
    │   │   ├── rng.gen_test.go
    │   │   ├── rngs.gen.go
    │   │   ├── tme.gen.go
    │   │   └── tme.gen_test.go
    │   ├── tmes
    │   │   └── tmes.gen.go
    │   ├── unt
    │   │   ├── unt.gen.go
    │   │   └── unt.gen_test.go
    │   └── unts
    │       └── unts.gen.go
    ├── cld
    │   └── cldr.go
    ├── cmd
    │   ├── cmd.go
    │   ├── ml.jl
    │   ├── ml-linear.jl
    │   ├── pro-cld.json
    │   ├── scripts
    │   │   ├── _0.pro
    │   │   ├── _1.pro
    │   │   ├── _2.pro
    │   │   ├── _3.pro
    │   │   ├── _4.pro
    │   │   ├── blngrBtmCrs.pro
    │   │   ├── blngrTopCrs.pro
    │   │   ├── clr.pro
    │   │   ├── cndStm2.pro
    │   │   ├── cndStm.pro
    │   │   ├── c.pro
    │   │   ├── d.pro
    │   │   ├── _long0.pro
    │   │   ├── _long3.pro
    │   │   ├── max.pro
    │   │   ├── pll.pro
    │   │   ├── plt.pro
    │   │   ├── __.pro
    │   │   ├── _shrt.pro
    │   │   ├── singlPass.pro
    │   │   └── tmp.pro
    │   └── sys.cfg
    ├── dsk
    │   └── dskr.go
    ├── err
    │   └── err.go
    ├── ext.gen.go
    ├── fs
    │   ├── fs.go
    │   └── fs_test.go
    ├── idn.gen.go
    ├── ifc.gen.go
    ├── k
    │   └── k.go
    ├── ks
    │   └── ks.go
    ├── lng
    │   ├── jsn
    │   │   ├── jsnr.gen.go
    │   │   ├── jsnr.gen_test.go
    │   │   ├── jsnr_test.go
    │   │   └── trm
    │   │       ├── prs
    │   │       │   ├── prs.gen.go
    │   │       │   └── prs.gen_test.go
    │   │       ├── trmr.gen.go
    │   │       └── trmr.gen_test.go
    │   ├── pro
    │   │   ├── act
    │   │   │   ├── actr.gen.go
    │   │   │   ├── actr.gen_test.go
    │   │   │   └── scp.gen.go
    │   │   ├── cfg
    │   │   │   ├── cfgr.gen.go
    │   │   │   └── cfgr.gen_test.go
    │   │   ├── script_test.go
    │   │   ├── trm
    │   │   │   ├── prs
    │   │   │   │   ├── prs.gen.go
    │   │   │   │   └── prs.gen_test.go
    │   │   │   ├── trmr.gen.go
    │   │   │   └── trmr.gen_test.go
    │   │   └── xpr
    │   │       ├── knd
    │   │       │   └── knd.gen.go
    │   │       ├── scp.gen.go
    │   │       ├── xprr.gen.go
    │   │       └── xprr.gen_test.go
    │   └── scn
    │       ├── scn.gen.go
    │       ├── scnr.gen.go
    │       └── scnr.gen_test.go
    ├── log
    │   └── log
    │       └── log.gen.go
    ├── mu.gen.go
    ├── run
    │   ├── act.go
    │   ├── run.go
    │   ├── runr.go
    │   ├── runr_test.go
    │   ├── runs.go
    │   └── tkr.go
    ├── sql
    │   └── rltOpn.sql
    ├── srt
    │   └── srt.go
    ├── sys.go
    ├── sys_test.go
    ├── todo
    ├── tpl
    │   ├── act.go
    │   ├── actr.go
    │   ├── ana.go
    │   ├── anaInstr.go
    │   ├── anaPort.go
    │   ├── anaPrfmDlt.go
    │   ├── anaPrfm.go
    │   ├── anaPrv.go
    │   ├── anaPth.go
    │   ├── anaStm.go
    │   ├── anaTic.go
    │   ├── anaTmeFlt.go
    │   ├── anaTmeflts.go
    │   ├── anaTmeIdx.go
    │   ├── anaTrd.go
    │   ├── anaTrdRsnCls.go
    │   ├── anaTrdRsnOpn.go
    │   ├── atr
    │   │   └── atr.go
    │   ├── block.go
    │   ├── bnd.go
    │   ├── bol.go
    │   ├── bsc.go
    │   ├── cfg.go
    │   ├── cfgr.go
    │   ├── clr.go
    │   ├── cmd
    │   │   └── cmd.go
    │   ├── cmnts.go
    │   ├── cnst.go
    │   ├── ext.go
    │   ├── fld.go
    │   ├── fle.go
    │   ├── flt.go
    │   ├── fltRng.go
    │   ├── fn.go
    │   ├── fnt.go
    │   ├── hstCnd.go
    │   ├── hstFtr.go
    │   ├── hst.go
    │   ├── hstInrvl.go
    │   ├── hstInstr.go
    │   ├── hstPort.go
    │   ├── hstPrcp.go
    │   ├── hstPrcpSplt.go
    │   ├── hstPrfm.go
    │   ├── hstPrv.go
    │   ├── hstSide.go
    │   ├── hstSplt.go
    │   ├── hstStgy.go
    │   ├── hstStm.go
    │   ├── hstStmSplt.go
    │   ├── import.go
    │   ├── int.go
    │   ├── jsn.go
    │   ├── jsnPrs.go
    │   ├── jsnr.go
    │   ├── jsnTrm.go
    │   ├── jsnTrmr.go
    │   ├── knd.go
    │   ├── lbl.go
    │   ├── lines.go
    │   ├── lng.go
    │   ├── log.go
    │   ├── memSig.go
    │   ├── mod
    │   │   └── mod.go
    │   ├── node.go
    │   ├── pen.go
    │   ├── pkgFn.go
    │   ├── pkg.go
    │   ├── pltDpth.go
    │   ├── pltFltAxisY.go
    │   ├── pltFlts.go
    │   ├── pltFltsSctrDist.go
    │   ├── pltFltsSctr.go
    │   ├── plt.go
    │   ├── pltHrz.go
    │   ├── pltPrcp.go
    │   ├── pltPrcpSplt.go
    │   ├── pltStgy.go
    │   ├── pltStm.go
    │   ├── pltStmSplt.go
    │   ├── pltTmeAxisX.go
    │   ├── pltVrt.go
    │   ├── prm.go
    │   ├── pro.go
    │   ├── prs.go
    │   ├── prtActFn.go
    │   ├── prtActTyp.go
    │   ├── prtAri.go
    │   ├── prtArrAgg.go
    │   ├── prtArrBytWrt.go
    │   ├── prtArrCld.go
    │   ├── prtArrCnt.go
    │   ├── prtArrFld.go
    │   ├── prtArrFldGrp.go
    │   ├── prtArrFldSel.go
    │   ├── prtArrFldSrt.go
    │   ├── prtArr.go
    │   ├── prtArrIdn.go
    │   ├── prtArrInr.go
    │   ├── prtArrRel.go
    │   ├── prtArrRng.go
    │   ├── prtArrScl.go
    │   ├── prtArrSel.go
    │   ├── prtArrSer.go
    │   ├── prtArrSrt.go
    │   ├── prtArrStrWrt.go
    │   ├── prtArrUna.go
    │   ├── prtBytes.go
    │   ├── prtElmArr.go
    │   ├── prtEnumFlg.go
    │   ├── prtEnum.go
    │   ├── prt.go
    │   ├── prtIdn.go
    │   ├── prtIfc.go
    │   ├── prtLog.go
    │   ├── prtPkt.go
    │   ├── prtPltArng.go
    │   ├── prtPltArngNew.go
    │   ├── prtPlt.go
    │   ├── prtPrm.go
    │   ├── prtPrntIfc.go
    │   ├── prtPth.go
    │   ├── prtRel.go
    │   ├── prtRng.go
    │   ├── prtSel.go
    │   ├── prtSgn.go
    │   ├── prtString.go
    │   ├── prtStructBytRed.go
    │   ├── prtStructBytWrt.go
    │   ├── prtStructCld.go
    │   ├── prtStructCpy.go
    │   ├── prtStructFldSet.go
    │   ├── prtStructGrp.go
    │   ├── prtStructIdn.go
    │   ├── prtStructKeyWrt.go
    │   ├── prtStructRel.go
    │   ├── prtStructStrWrt.go
    │   ├── prtWve.go
    │   ├── pth.go
    │   ├── rltCnd.go
    │   ├── rlt.go
    │   ├── rltInrvl.go
    │   ├── rltInstr.go
    │   ├── rltPort.go
    │   ├── rltPrfm.go
    │   ├── rltPrv.go
    │   ├── rltSide.go
    │   ├── rltStgy.go
    │   ├── rltStm.go
    │   ├── scn.go
    │   ├── scnr.go
    │   ├── scpAct.go
    │   ├── scpXpr.go
    │   ├── str.go
    │   ├── sysAct.go
    │   ├── sys.go
    │   ├── sysIdn.go
    │   ├── sysIfc.go
    │   ├── sysMu.go
    │   ├── test.go
    │   ├── tme.go
    │   ├── tmeRng.go
    │   ├── tpl.go
    │   ├── trc.go
    │   ├── trcOpt.go
    │   ├── trdsStm.go
    │   ├── trdStmSeg.go
    │   ├── trm.go
    │   ├── trmr.go
    │   ├── tst.go
    │   ├── typFn.go
    │   ├── typ.go
    │   ├── unt.go
    │   ├── var.go
    │   ├── vis.go
    │   ├── xpr.go
    │   └── xprr.go
    ├── trc
    │   ├── durr.go
    │   ├── opt.gen.go
    │   └── trcr.go
    └── tst
        ├── bool.go
        ├── bool_test.go
        ├── cndMnr.go
        ├── inrvlMnr.go
        ├── instrMnr.go
        ├── int.go
        ├── mnr.go
        ├── nil.go
        ├── nil_test.go
        ├── panic.go
        ├── panic_test.go
        ├── sideMnr.go
        ├── stgyMnr.go
        ├── stmMnr.go
        ├── strings.go
        ├── strings_test.go
        ├── tics.go
        ├── tst.gen.go
        ├── tst.go
        ├── type.go
        └── type_test.go
```
