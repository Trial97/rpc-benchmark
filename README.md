# The Bench result 
```
go test -run=:"^$" -benchmem -v -bench=. -benchtime=30s                                                
goos: linux
goarch: amd64
pkg: rpc-bench
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkRPCCallWithContext
2021/04/02 13:03:58 Test RPC server listening on 127.0.0.1:43335
2021/04/02 13:03:58 Test RPC server listening on 127.0.0.1:42405
2021/04/02 13:03:58 Test RPC server listening on 127.0.0.1:45319
2021/04/02 13:03:58 Test RPC server listening on 127.0.0.1:39271
BenchmarkRPCCallWithContext-8   	  911137	     39466 ns/op	     416 B/op	      11 allocs/op
BenchmarkRPCCall
2021/04/02 13:04:34 Test RPC server listening on 127.0.0.1:41101
2021/04/02 13:04:34 Test RPC server listening on 127.0.0.1:38679
2021/04/02 13:04:34 Test RPC server listening on 127.0.0.1:39067
2021/04/02 13:04:35 Test RPC server listening on 127.0.0.1:40379
BenchmarkRPCCall-8              	  980272	     36727 ns/op	     312 B/op	       9 allocs/op
PASS
ok  	rpc-bench	72.742s
```