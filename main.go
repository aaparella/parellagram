package main

import (
	"log"
	"net/http"
)

func main() {
	conf := getConfig()
	saveDetailedPages(conf)
	saveLandingPage(conf)

	log.Println("Listening on port 80")

	http.Handle("/", http.FileServer(http.Dir(conf.Artifacts.Path)))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir(conf.Resources.Styles))))

	log.Println("Listening on port 80")
	http.ListenAndServe(":80", nil)
}
