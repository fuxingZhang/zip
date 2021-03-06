package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// File Compress a file into a zip file
func File(src string, dst string) error {
	zipfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	file, err := os.Open(src)
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	w, err := archive.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, file)
	return err
}

// Another implementation for File
func file(src string, dst string) error {
	zipfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	file, err := os.Open(src)
	if err != nil {
		return err
	}

	w, err := archive.Create(filepath.Base(src))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, file)
	return err
}

// Dir Compress a directory into a zip file
func Dir(src string, dst string, includeSrc bool) error {
	zipfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var name string
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if includeSrc {
			name, err = filepath.Rel(filepath.Dir(src), path)
			if err != nil {
				return err
			}
		} else {
			if path == src {
				return nil
			}
			name, err = filepath.Rel(src, path)
			if err != nil {
				return err
			}
		}

		header.Name = filepath.ToSlash(name)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

// Another implementation for dir
func dir(src string, dst string, includeSrc bool) error {
	zipfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		var name string
		if includeSrc {
			name, err = filepath.Rel(filepath.Dir(src), path)
			if err != nil {
				return err
			}
		} else {
			if path == src {
				return nil
			}
			name, err = filepath.Rel(src, path)
			if err != nil {
				return err
			}
		}

		name = filepath.ToSlash(name)
		if info.IsDir() {
			name += "/"
		}

		writer, err := archive.Create(name)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

// Unzip Unzip zip file
func Unzip(zipFile string, dstDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(dstDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())

			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}

			// atime := time.Now()
			// _ = os.Chtimes(fpath, atime, f.Modified)
		}
	}
	return nil
}
