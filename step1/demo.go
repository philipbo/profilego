package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync/atomic"
)

var visitors int64

func handleHi(w http.ResponseWriter, r *http.Request) {
	if match, _ := regexp.MatchString(`^\w*$`, r.FormValue("color")); !match {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}

	//visitors++  //此处存在竟争
	//fix
	//1.使用channel
	//2.使用Mutex
	//3.使用atomic

	num := atomic.AddInt64(&visitors, 1)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("<h1 style='color:%s'>Welcome!</h1>You are visitor number %d!", r.FormValue("color"), num)))
}

func main() {
	log.Println("Starting on port 3000")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
