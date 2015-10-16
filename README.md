$ go test -v

#### Race Detector
 
$ go test -v -race

#### Run the benchmarks:

$ go test -v -run=^$ -bench=.


#### CPU Profiling

$ go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -cpuprofile=prof.cpu

会生成prof.cpu和xxx.test两个文件

$ go tool pprof xxx.test prof.cpu

(pprof) top

(pprof) top –cum  10 可以指定数字，按排名

(pprof) list handleHi

#### Memory profiling

$ go test -v -run=^$ -bench=^BenchmarkHi$ -benchtime=2s -memprofile=prof.mem

$ go tool pprof -alloc_space xxx.test prof.mem

#### install benchcmp  

需要自行搭梯子

go get golang.org/x/tools/cmd/benchcmp

$ go test -bench=. -memprofile=prof.mem | tee mem.0
$ go test -bench=. -memprofile=prof.mem | tee mem.1
$ go test -bench=. -memprofile=prof.mem | tee mem.2
$ go test -bench=. -memprofile=prof.mem | tee mem.3

$ benchcmp step0/mem.0 step3/mem.3


并行分析 step6

$ go test -bench=Parallel -blockprofile=prof.block