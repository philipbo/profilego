package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync/atomic"
)

var visitors int64

var colorRx = regexp.MustCompile(`\w*$`)

func handleHi(w http.ResponseWriter, r *http.Request) {

	if !colorRx.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}

	num := atomic.AddInt64(&visitors, 1)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("<h1 style='color:%s'>Welcome!</h1>You are visitor number %d!", r.FormValue("color"), num)))
}

func main() {
	log.Println("Starting on port 3000")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
