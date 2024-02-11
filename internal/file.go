package internal

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	data, err := io.ReadAll(xmlFile)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func Find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}
