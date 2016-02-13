package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveLandingPage)
	http.HandleFunc("/post/", serveDetailedPage)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))

	log.Println("Listening on port 8080")

	http.ListenAndServe(":8080", nil)
}
