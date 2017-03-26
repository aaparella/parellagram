package main

import (
	"fmt"
	"net/http"
	"time"
)

var launchTime time.Time

func init() {
	launchTime = time.Now()
}

func Uptime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This server has been running since %v (%v)", launchTime, time.Since(launchTime))
}
