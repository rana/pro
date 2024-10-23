package ana

import (
	"strings"
	"sys/bsc/str"
)

type (
	Pth interface {
		Name() str.Str
		PrmWrt(b *strings.Builder)
		Prm() string
		StrWrt(b *strings.Builder)
		String() string
	}
)
