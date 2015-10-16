$ go test -v

#### Race Detector(竞态分析）
 
$ go test -v -race

#### Run the benchmarks:

$ go test -v -run=^$ -bench=.


#### CPU Profiling

$ go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -cpuprofile=prof.cpu

会生成prof.cpu和xxx.test两个文件

$ go tool pprof xxx.test prof.cpu

(pprof) top

(pprof) top –cum

(pprof) list handleHi