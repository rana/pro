package fs_test

import (
	"os"
	"path/filepath"
	"sys/fs"
	"sys/tst"
	"testing"
)

func TestEnsureDir(t *testing.T) {
	path := filepath.Join(os.TempDir(), "a", "b")
	fs.EnsureDir(path)
	tst.True(t, fs.Exists(path))
	os.RemoveAll(filepath.Join(os.TempDir(), "a"))
}
