package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveLandingPage)
	http.HandleFunc("/post/", serveDetailedPage)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	log.Println("Listening on port 80")

	http.ListenAndServe(":80", nil)
}
