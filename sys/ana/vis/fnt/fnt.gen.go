package fnt

import (
	"golang.org/x/image/font"
	"sync"
)

type (
	Fnt struct {
		Face font.Face
		Mu   sync.Mutex
	}
	FntScp struct {
		Idx uint32
		Arr []*Fnt
	}
)
