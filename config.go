package main

import (
	"log"

	gcfg "gopkg.in/gcfg.v1"
)

// Config stores configuration of parellagram website.
type Config struct {
	Website struct {
		Title string
		Port  int
	}
	Resources struct {
		Posts  string
		Styles string
		Images string
	}
	Links struct {
		Twitter   string
		Email     string
		Github    string
		Bitbucket string
		Gitlab    string
	}
}

func getConfig() Config {
	var conf Config
	if err := gcfg.ReadFileInto(&conf, "config.ini"); err != nil {
		log.Fatal("Error parsing configuration : ", err)
	}
	return conf
}
