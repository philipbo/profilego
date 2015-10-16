package main

import "testing"

func TestNextVisitorNum(t *testing.T) {
	for i := 0; i < 2; i++ {
		a := i + 1
		n := nextVisitorNum()
		if n != a {
			t.Errorf("n: %d, a: %d, i: %d", n, a, i)
		}
	}
}

//并行
func BenchmarkVisitCount(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			nextVisitorNum()
		}
	})
}

func BenchmarkVisitCountNoDefer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			nextVisitorNumNoDefer()
		}
	})
}

func BenchmarkVisitCountAtomic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			nextVisitorNumAtomic()
		}
	})
}

//非并行
func BenchmarkVisitCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nextVisitorNum()
	}
}
func BenchmarkVisitCountNoDefer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nextVisitorNumNoDefer()
	}
}

func BenchmarkVisitCountAtomic1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nextVisitorNumAtomic()
	}
}
