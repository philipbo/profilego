package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
)

var visitors int64

var colorRx = regexp.MustCompile(`\w*$`)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

var colorRxPool = sync.Pool{
	New: func() interface{} { return regexp.MustCompile(`\w*$`) },
}

func handleHi(w http.ResponseWriter, r *http.Request) {

	if !colorRxPool.Get().(*regexp.Regexp).MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}

	num := atomic.AddInt64(&visitors, 1)

	buf := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buf)
	buf.Reset()

	buf.WriteString("<h1 style='color: ")
	buf.WriteString(r.FormValue("color"))
	buf.WriteString("'>Welcome!</h1>You are visitor number ")
	b := strconv.AppendInt(buf.Bytes(), int64(num), 10)
	b = append(b, '!')
	w.Write(b)

	fmt.Fprintf(w, "<html><h1 stype='color: %s'>Welcome!</h1>You are visitor number %d!", r.FormValue("color"), num)
}

func main() {
	log.Println("Starting on port 3000")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
