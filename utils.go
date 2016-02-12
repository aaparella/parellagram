package main

import (
	"os"
	"path"
)

func getDirContents(dirpath string) ([]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dir, err := os.Open(path.Join(pwd, dirpath))
	if err != nil {
		return nil, err
	}
	return dir.Readdirnames(0)
}
