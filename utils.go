package main

import (
	"os"
	"path"
)

func clearTempDirectory(conf Config) {
	p := path.Join(os.TempDir(), "parellagram")
	clearDirContents(p)
	os.Mkdir(path.Join(p, conf.Resources.Posts), os.ModeTemporary)
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

func clearDirContents(dirname string) error {
	if err := os.RemoveAll(dirname); err != nil {
		return err
	}
	os.Mkdir(dirname, os.ModeDir|os.ModeTemporary)
	return nil
}
