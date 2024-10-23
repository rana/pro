package app

import "sys/fs"

type (
	Fle struct {
		Name string
		Pth  string
		Txt  string
	}
)

func (x *Fle) GetName() string { // Itm interface
	return x.Name
}

func (x *Fle) Load() {
	x.Txt = fs.LoadText(x.Pth)
}
func (x *Fle) Save() {
	fs.WriteFile(x.Pth, []byte(x.Txt))
}
