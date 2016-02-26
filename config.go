package main

import (
	"log"

	"code.google.com/p/gcfg"
)

type Config struct {
	Website struct {
		Title string
		Port  int
	}
	Resources struct {
		Posts  string
		Styles string
	}
}

func getConfig() Config {
	var conf Config
	if err := gcfg.ReadFileInto(&conf, "config.ini"); err != nil {
		log.Fatal("Error parsing configuration : ", err)
	}
	return conf
}
