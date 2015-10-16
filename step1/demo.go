package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var visitors int

func handleHi(w http.ResponseWriter, r *http.Request) {
	if match, _ := regexp.MatchString(`^\w*$`, r.FormValue("color")); !match {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}

	visitors++
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("<h1 style='color:%s'>Welcome!</h1>You are visitor number %d!", r.FormValue("color"), visitors)))
}

func main() {
	log.Println("Starting on port 3000")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
