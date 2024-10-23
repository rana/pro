package ana

import "sys/bsc/str"

func (x TrdRsnCls) Str() str.Str { return str.Str(trdRsnClsNames[x]) }
