go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

pprof -http ":8081" cpu.prof 