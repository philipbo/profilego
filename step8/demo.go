package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var visitors struct {
	sync.Mutex
	n int
}

func nextVisitorNum() int {
	visitors.Lock()
	defer visitors.Unlock()
	visitors.n++
	return visitors.n
}

func nextVisitorNumNoDefer() int {
	visitors.Lock()
	visitors.n++
	visitors.Unlock()
	return visitors.n
}

var atomicVisitors struct {
	n int64
}

func nextVisitorNumAtomic() int {
	return int(atomic.AddInt64(&atomicVisitors.n, 1))
}

func main() {
	fmt.Println("visitors num: ", nextVisitorNum())
	fmt.Println("atomic visitors num: ", nextVisitorNumAtomic())

}
