package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	conf := getConfig()
	clearTempDirectory()
	saveDetailedPages(conf)
	saveLandingPage(conf)

	http.Handle("/", http.FileServer(http.Dir(path.Join(os.TempDir(), "parellagram"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir(conf.Resources.Styles))))

	log.Println("Listening on port 80")
	http.ListenAndServe(":80", nil)
}
