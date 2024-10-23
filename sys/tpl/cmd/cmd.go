package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sys/err"
	"sys/fs"
	"sys/tpl"
)

func main() {
	fmt.Println("- tpl")
	defer fmt.Println("- tpl")

	curDir, er := os.Getwd()
	if er != nil {
		err.Panic(er)
	}
	// curDir: /home/rana/go/src/sys/tpl/cmd
	// srcDir: /home/rana/go/src
	srcDir := filepath.Dir(filepath.Dir(filepath.Dir(curDir)))
	if *tpl.Clean {
		paths := fs.Clean(filepath.Join(srcDir, "sys"), ".gen.go$|.gen_test.go$|debug*|.log$|^cmd$")
		for _, path := range paths {
			fmt.Println("  - Del", path)
		}
	}
	if *tpl.Wrt {
		sys := tpl.NewSys()
		sys.Init()
		sys.WriteToDisk(srcDir)
	}
}
