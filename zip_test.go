package zip

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	src, dest := "./zip.go", "./file.zip"
	defer os.Remove(dest)
	err := File(src, dest)
	if err != nil {
		t.Error(err)
	}
}

func TestDir(t *testing.T) {
	src, dest, destIncludeSrc := "./test", "./dir.zip", "./includeSrc.zip"
	// src, _ = filepath.Abs(src)
	defer os.Remove(dest)
	defer os.Remove(destIncludeSrc)
	err := Dir(src, dest, false)
	if err != nil {
		t.Error(err)
	}
	err = Dir(src, destIncludeSrc, true)
	if err != nil {
		t.Error(err)
	}
}

func TestUnzip(t *testing.T) {
	src, zipFile, unzipDir := "./test", "./toUzip.zip", "./unzip"
	defer os.Remove(zipFile)
	defer os.RemoveAll(unzipDir)
	err := Dir(src, zipFile, true)
	if err != nil {
		t.Error(err)
	}
	err = Unzip(zipFile, unzipDir)
	if err != nil {
		t.Error(err)
	}
}
