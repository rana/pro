package fs

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sys"
	"sys/err"

	"github.com/sbinet/npyio"
)

const (
	// WriteDirMode is a directory write mode.
	WriteDirMode os.FileMode = 0766
	// WriteFileMode is a file write mode.
	WriteFileMode os.FileMode = 0666
)

// Exists returns true whether a file cortem path exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// EnsureDir creates a directory path if it doesn't exist.
func EnsureDir(path string) {
	if !Exists(path) {
		EnsureDir(filepath.Dir(path))
		er := os.Mkdir(path, WriteDirMode)
		if er != nil {
			err.Panic(er)
		}
	}
}

// WriteFile writes a file.
func WriteFile(filename string, content []byte) {
	er := ioutil.WriteFile(filename, content, WriteFileMode)
	if er != nil {
		err.Panic(er)
	}
}

// LoadText loads a text file.
func LoadText(filename string) string {
	content, er := ioutil.ReadFile(filename)
	if er != nil {
		err.Panic(er)
	}
	return string(content)
}

// WorkingDir returns the working directory.
func WorkingDir(file ...string) (r string) {
	r, er := os.Getwd()
	if er != nil {
		err.Panic(er)
	}
	if len(file) != 0 {
		r = filepath.Join(r, file[0])
	}
	return r
}

// Name return the name after the last separator.
func Name(pth string) string {
	idx := strings.LastIndex(pth, string(filepath.Separator))
	if idx < 0 {
		return pth
	}
	return pth[idx+1:]
}

func Fles(path string, pattern ...string) (r []string) {
	var dirs []string
	dirs = append(dirs, path)
	for len(dirs) > 0 {
		curDir := dirs[len(dirs)-1]
		dirs = dirs[:len(dirs)-1]
		fles, er := ioutil.ReadDir(curDir)
		if er != nil {
			err.Panic(er)
		}
		for _, fle := range fles {
			curPath := filepath.Join(curDir, fle.Name())
			if fle.IsDir() {
				dirs = append(dirs, curPath)
			} else {
				if len(pattern) > 0 {
					matched, er := regexp.MatchString(pattern[0], fle.Name())
					if er != nil {
						err.Panic(er)
					}
					if matched {
						r = append(r, curPath)
					}
				} else {
					r = append(r, curPath)
				}
			}
		}
	}
	return r
}

func Dirs(paths ...string) (r []string) {
	for _, path := range paths {
		fles, er := ioutil.ReadDir(path)
		if er != nil {
			err.Panic(er)
		}
		for _, fle := range fles {
			if fle.IsDir() {
				r = append(r, Dirs(filepath.Join(path, fle.Name()))...)
			}
		}
	}
	r = append(r, paths...)
	return r
}

func EmptyDirs(dirs ...string) (r []string) {
	for len(dirs) > 0 {
		dir := dirs[len(dirs)-1]
		dirs = dirs[:len(dirs)-1]
		fles, er := ioutil.ReadDir(dir)
		if er != nil {
			err.Panic(er)
		}
		if len(fles) == 0 {
			er := os.Remove(dir)
			if er != nil {
				err.Panic(er)
			}
			r = append(r, dir)
			continue
		}
		for _, fle := range fles { // delete inr
			if fle.IsDir() {
				r = append(r, EmptyDirs(filepath.Join(dir, fle.Name()))...)
			}
		}
		fles, er = ioutil.ReadDir(dir) // recheck current if all inr deleted
		if er != nil {
			err.Panic(er)
		}
		if len(fles) == 0 {
			er := os.Remove(dir)
			if er != nil {
				err.Panic(er)
			}
			r = append(r, dir)
		}
	}
	return r
}

func Del(paths ...string) {
	for _, path := range paths {
		er := os.Remove(path)
		if er != nil {
			err.Panic(er)
		}
	}
}

func Clean(path string, pattern ...string) (r []string) {
	paths := Fles(path, pattern...)
	r = append(r, paths...)
	Del(paths...)
	paths = EmptyDirs(path)
	r = append(r, paths...)
	sort.Slice(r, func(i, j int) bool { return r[i] > r[j] })
	return r
}

func Npz(filepath string, ks []string, vs ...interface{}) {
	sys.Log("Npz: saving", filepath)
	if len(ks) != len(vs) {
		err.Panicf("unequal lengths (ks:%v vs:%v)", len(ks), len(vs))
	}
	if len(ks) == 0 {
		return
	}
	fNpz, er := os.Create(filepath)
	if er != nil {
		err.Panic(er)
	}
	defer fNpz.Close()
	zw := zip.NewWriter(fNpz)
	defer zw.Close()
	for n := 0; n < len(ks); n++ {
		zEntry, er := zw.Create(ks[n])
		if er != nil {
			err.Panic(er)
		}
		er = npyio.Write(zEntry, vs[n])
		if er != nil {
			err.Panic(er)
		}
	}
}
