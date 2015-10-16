package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync/atomic"
)

var visitors int64

var rxOptionalID = regexp.MustCompile(`^\d*$`)

func handleHi(w http.ResponseWriter, r *http.Request) {

	if !rxOptionalID.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}

    num := atomic.AddInt64(&visitors, 1)

    fmt.Fprintf(w, "<html><h1 stype='color: \"%s\"'>Welcome!</h1>You are visitor number %d!", r.FormValue("color"), num)
}

func main() {
	log.Println("Starting on port 3000")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
