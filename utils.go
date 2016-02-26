package main

import (
	"log"
	"os"
	"path"
)

func clearTempDirectory() {
	if err := os.RemoveAll(path.Join(os.TempDir(), "parellagram")); err != nil {
		log.Fatal("Error writing to ", os.TempDir(), " : ", err)
	}
	os.Mkdir(path.Join(os.TempDir(), "parellagram"), os.ModeDir|os.ModeTemporary)
	os.Mkdir(path.Join(os.TempDir(), "parellagram", "posts"), os.ModeDir|os.ModeTemporary)
}

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
