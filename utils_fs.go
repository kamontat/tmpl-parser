package main

import (
	"os"
	"path/filepath"
	"strings"
)

func IsTmpl(f string) bool {
	return filepath.Ext(f) == TEMPLATE_EXTENSION
}

func CreateOutFile(s string) (*os.File, error) {
	return os.OpenFile(
		s,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0666,
	)
}

func JoinPath(base, f string) string {
	return filepath.Join(base, f)
}

func GetOutFile(s string) string {
	var file = filepath.Base(s)
	return strings.Replace(file, TEMPLATE_EXTENSION, "", 1)
}

func GetInFiles(in string) (ins []string) {
	var inDirs, err = os.ReadDir(in)
	if err != nil {
		return
	}

	for _, inDir := range inDirs {
		var abs = filepath.Join(in, inDir.Name())
		if inDir.IsDir() {
			ins = append(ins, GetInFiles(abs)...)
		} else if IsTmpl(abs) {
			ins = append(ins, abs)
		}
	}

	return
}
