# The Bench result 
```
go test -run=:"^$" -benchmem -v -bench=. -benchtime=30s                                                
goos: linux
goarch: amd64
pkg: rpc-bench
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkRPCCallWithContext
2021/04/06 09:00:17 Test RPC server listening on 127.0.0.1:38549
2021/04/06 09:00:17 Test RPC server listening on 127.0.0.1:33137
2021/04/06 09:00:17 Test RPC server listening on 127.0.0.1:40217
2021/04/06 09:00:17 Test RPC server listening on 127.0.0.1:41137
BenchmarkRPCCallWithContext-8   	  867255	     40560 ns/op	     416 B/op	      11 allocs/op
BenchmarkRPCCall
2021/04/06 09:00:52 Test RPC server listening on 127.0.0.1:36377
2021/04/06 09:00:52 Test RPC server listening on 127.0.0.1:43979
2021/04/06 09:00:52 Test RPC server listening on 127.0.0.1:40967
2021/04/06 09:00:53 Test RPC server listening on 127.0.0.1:32781
BenchmarkRPCCall-8              	  915488	     38243 ns/op	     312 B/op	       9 allocs/op
BenchmarkRPCXCall
2021/04/06 09:01:28 Test RPC server listening on 127.0.0.1:37391
2021/04/06 09:01:28 server.go:182: INFO : server pid:31734
2021/04/06 09:01:28 server.go:406: INFO : client has closed this connection: 127.0.0.1:36620
2021/04/06 09:01:28 Test RPC server listening on 127.0.0.1:35987
2021/04/06 09:01:28 server.go:182: INFO : server pid:31734
2021/04/06 09:01:28 server.go:406: INFO : client has closed this connection: 127.0.0.1:56142
2021/04/06 09:01:28 Test RPC server listening on 127.0.0.1:37889
2021/04/06 09:01:28 server.go:182: INFO : server pid:31734
2021/04/06 09:01:28 server.go:406: INFO : client has closed this connection: 127.0.0.1:57122
2021/04/06 09:01:28 Test RPC server listening on 127.0.0.1:41143
2021/04/06 09:01:28 server.go:182: INFO : server pid:31734
BenchmarkRPCXCall-8             	  799888	     45368 ns/op	    2068 B/op	      39 allocs/op
2021/04/06 09:02:05 server.go:406: INFO : client has closed this connection: 127.0.0.1:60930
PASS
ok  	rpc-bench	107.767s
```