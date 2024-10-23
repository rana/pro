WITH 
eur_usd_long AS (
  SELECT 
    Instr, Name, PnlPct, ScsPct, PipPerDay, UsdPerDay, ScsPerDay, OpnPerDay, PnlUsd, IsLong, PrfLim, LosLim, DurLim, PipAvg, PipMdn, PipMin, PipMax, PipSum, DurAvg, DurMdn, DurMin, DurMax, DayCnt, TrdCnt, BalFstUsd, BalLstUsd, CstTotUsd, CstSpdUsd, CstComUsd, MrgnRtio, TrdPct, Pth
  FROM 
    `pro-cld.oan_cor_tic.StgyPrfm`
  WHERE instr = 'eur_usd' AND islong = true
  ORDER BY PipPerDay DESC LIMIT 5
),
eur_usd_shrt AS (
  SELECT 
    Instr, Name, PnlPct, ScsPct, PipPerDay, UsdPerDay, ScsPerDay, OpnPerDay, PnlUsd, IsLong, PrfLim, LosLim, DurLim, PipAvg, PipMdn, PipMin, PipMax, PipSum, DurAvg, DurMdn, DurMin, DurMax, DayCnt, TrdCnt, BalFstUsd, BalLstUsd, CstTotUsd, CstSpdUsd, CstComUsd, MrgnRtio, TrdPct, Pth
  FROM 
    `pro-cld.oan_cor_tic.StgyPrfm`
  WHERE instr = 'eur_usd' AND islong = false
  ORDER BY PipPerDay DESC LIMIT 5
),
aud_usd_long AS (
  SELECT 
    Instr, Name, PnlPct, ScsPct, PipPerDay, UsdPerDay, ScsPerDay, OpnPerDay, PnlUsd, IsLong, PrfLim, LosLim, DurLim, PipAvg, PipMdn, PipMin, PipMax, PipSum, DurAvg, DurMdn, DurMin, DurMax, DayCnt, TrdCnt, BalFstUsd, BalLstUsd, CstTotUsd, CstSpdUsd, CstComUsd, MrgnRtio, TrdPct, Pth
  FROM 
    `pro-cld.oan_cor_tic.StgyPrfm`
  WHERE instr = 'aud_usd' AND islong = true
  ORDER BY PipPerDay DESC LIMIT 5
),
aud_usd_shrt AS (
  SELECT 
    Instr, Name, PnlPct, ScsPct, PipPerDay, UsdPerDay, ScsPerDay, OpnPerDay, PnlUsd, IsLong, PrfLim, LosLim, DurLim, PipAvg, PipMdn, PipMin, PipMax, PipSum, DurAvg, DurMdn, DurMin, DurMax, DayCnt, TrdCnt, BalFstUsd, BalLstUsd, CstTotUsd, CstSpdUsd, CstComUsd, MrgnRtio, TrdPct, Pth
  FROM 
    `pro-cld.oan_cor_tic.StgyPrfm`
  WHERE instr = 'aud_usd' AND islong = false
  ORDER BY PipPerDay DESC LIMIT 5
),
nzd_usd_long AS (
  SELECT 
    Instr, Name, PnlPct, ScsPct, PipPerDay, UsdPerDay, ScsPerDay, OpnPerDay, PnlUsd, IsLong, PrfLim, LosLim, DurLim, PipAvg, PipMdn, PipMin, PipMax, PipSum, DurAvg, DurMdn, DurMin, DurMax, DayCnt, TrdCnt, BalFstUsd, BalLstUsd, CstTotUsd, CstSpdUsd, CstComUsd, MrgnRtio, TrdPct, Pth
  FROM 
    `pro-cld.oan_cor_tic.StgyPrfm`
  WHERE instr = 'nzd_usd' AND islong = true
  ORDER BY PipPerDay DESC LIMIT 5
),
nzd_usd_shrt AS (
  SELECT 
    Instr, Name, PnlPct, ScsPct, PipPerDay, UsdPerDay, ScsPerDay, OpnPerDay, PnlUsd, IsLong, PrfLim, LosLim, DurLim, PipAvg, PipMdn, PipMin, PipMax, PipSum, DurAvg, DurMdn, DurMin, DurMax, DayCnt, TrdCnt, BalFstUsd, BalLstUsd, CstTotUsd, CstSpdUsd, CstComUsd, MrgnRtio, TrdPct, Pth
  FROM 
    `pro-cld.oan_cor_tic.StgyPrfm`
  WHERE instr = 'nzd_usd' AND islong = false
  ORDER BY PipPerDay DESC LIMIT 5
),
gbp_usd_long AS (
  SELECT 
    Instr, Name, PnlPct, ScsPct, PipPerDay, UsdPerDay, ScsPerDay, OpnPerDay, PnlUsd, IsLong, PrfLim, LosLim, DurLim, PipAvg, PipMdn, PipMin, PipMax, PipSum, DurAvg, DurMdn, DurMin, DurMax, DayCnt, TrdCnt, BalFstUsd, BalLstUsd, CstTotUsd, CstSpdUsd, CstComUsd, MrgnRtio, TrdPct, Pth
  FROM 
    `pro-cld.oan_cor_tic.StgyPrfm`
  WHERE instr = 'gbp_usd' AND islong = true
  ORDER BY PipPerDay DESC LIMIT 5
),
gbp_usd_shrt AS (
  SELECT 
    Instr, Name, PnlPct, ScsPct, PipPerDay, UsdPerDay, ScsPerDay, OpnPerDay, PnlUsd, IsLong, PrfLim, LosLim, DurLim, PipAvg, PipMdn, PipMin, PipMax, PipSum, DurAvg, DurMdn, DurMin, DurMax, DayCnt, TrdCnt, BalFstUsd, BalLstUsd, CstTotUsd, CstSpdUsd, CstComUsd, MrgnRtio, TrdPct, Pth
  FROM 
    `pro-cld.oan_cor_tic.StgyPrfm`
  WHERE instr = 'gbp_usd' AND islong = false
  ORDER BY PipPerDay DESC LIMIT 5
)

SELECT * FROM eur_usd_long
UNION ALL
SELECT * FROM eur_usd_shrt
UNION ALL
SELECT * FROM aud_usd_long
UNION ALL
SELECT * FROM aud_usd_shrt
UNION ALL
SELECT * FROM nzd_usd_long
UNION ALL
SELECT * FROM nzd_usd_shrt
UNION ALL
SELECT * FROM gbp_usd_long
UNION ALL
SELECT * FROM gbp_usd_shrt
ORDER BY PipPerDay DESC

