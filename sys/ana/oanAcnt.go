package ana

import (
	"sys/bsc/bol"
	"sys/bsc/flt"
	"sys/bsc/str"
	"sys/bsc/tme"
	"sys/bsc/unt"
	"sys/lng/jsn"
)

const (
	prefix = "account"
)

type (
	// Acnt is an Oanda account.
	Acnt struct {
		Created                     tme.Tme
		Currency                    str.Str
		Alias                       str.Str
		MarginRate                  flt.Flt
		HedgingEnabled              bol.Bol
		LastTransactionID           unt.Unt
		Balance                     flt.Flt
		OpnTrdCnt                   unt.Unt
		OpnPosCnt                   unt.Unt
		PndOrdCnt                   unt.Unt
		Pl                          flt.Flt
		ResettablePL                flt.Flt
		Financing                   flt.Flt
		Commission                  flt.Flt
		UnrealizedPL                flt.Flt
		NAV                         flt.Flt
		MarginAvailable             flt.Flt
		PositionValue               flt.Flt
		MarginCloseoutUnrealizedPL  flt.Flt
		MarginCloseoutNAV           flt.Flt
		MarginCloseoutMarginUsed    flt.Flt
		MarginCloseoutPositionValue flt.Flt
		MarginCloseoutPercent       flt.Flt
		WithdrawalLimit             flt.Flt
		MarginCallMarginUsed        flt.Flt
		MarginCallPercent           flt.Flt
	}
)

func (x *Acnt) JsnRed(txt string) {
	var j jsn.Jsnr
	j.Reset(txt)
	x.Created = j.StrTme(prefix, "createdTime")
	x.Currency = j.Str(prefix, "currency")
	x.Alias = j.Str(prefix, "alias")
	x.MarginRate = j.StrFlt(prefix, "marginRate")
	x.HedgingEnabled = j.Bol(prefix, "hedgingEnabled")
	x.LastTransactionID = j.StrUnt(prefix, "lastTransactionID")
	x.Balance = j.StrFlt(prefix, "balance")
	x.OpnTrdCnt = j.Unt(prefix, "openTradeCount")
	x.OpnPosCnt = j.Unt(prefix, "openPositionCount")
	x.PndOrdCnt = j.Unt(prefix, "pendingOrderCount")
	x.Pl = j.StrFlt(prefix, "pl")
	x.ResettablePL = j.StrFlt(prefix, "resettablePL")
	x.Financing = j.StrFlt(prefix, "financing")
	x.Commission = j.StrFlt(prefix, "commission")
	x.UnrealizedPL = j.StrFlt(prefix, "unrealizedPL")
	x.NAV = j.StrFlt(prefix, "NAV")
	x.MarginAvailable = j.StrFlt(prefix, "marginAvailable")
	x.PositionValue = j.StrFlt(prefix, "positionValue")
	x.MarginCloseoutUnrealizedPL = j.StrFlt(prefix, "marginCloseoutUnrealizedPL")
	x.MarginCloseoutNAV = j.StrFlt(prefix, "marginCloseoutNAV")
	x.MarginCloseoutMarginUsed = j.StrFlt(prefix, "marginCloseoutMarginUsed")
	x.MarginCloseoutPositionValue = j.StrFlt(prefix, "marginCloseoutPositionValue")
	x.MarginCloseoutPercent = j.StrFlt(prefix, "marginCloseoutPercent")
	x.WithdrawalLimit = j.StrFlt(prefix, "withdrawalLimit")
	x.MarginCallMarginUsed = j.StrFlt(prefix, "marginCallMarginUsed")
	x.MarginCallPercent = j.StrFlt(prefix, "marginCallPercent")
}
