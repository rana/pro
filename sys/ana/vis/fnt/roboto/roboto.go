package roboto

import (
	"sync"
	"sys/ana/vis/fnt"
	"github.com/golang/freetype/truetype"
)


var (
	medium *truetype.Font
	mediumFaces map[float32]*fnt.Fnt
	mediumMu sync.Mutex
)

func init() {
	mediumFaces = make(map[float32]*fnt.Fnt)
}

func Medium(size float32) (r *fnt.Fnt) {
	var ok bool
	mediumMu.Lock()
	r, ok = mediumFaces[size]
	if !ok {
		r = &fnt.Fnt{
			Face:truetype.NewFace(MediumFont(), &truetype.Options{Size: float64(size)}),
		}
		mediumFaces[size] = r
	}
	mediumMu.Unlock()
	return r
}

func MediumFont() *truetype.Font {
	if medium != nil {
		return medium
	}
	var er error
	medium, er = truetype.Parse(MediumBytes())
	if er != nil {
		panic(er)
	}
	return medium
}

// Roboto-Medium .ttf
func MediumBytes() []byte {
}