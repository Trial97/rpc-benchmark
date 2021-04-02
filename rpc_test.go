package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc-bench/air"
	"testing"
)

func startRPCServer() {
	rpc.Register(new(air.Arith))

	var l net.Listener
	l, serverAddr = listenTCP()
	log.Println("Test RPC server listening on", serverAddr)
	go rpc.Accept(l)
}

func BenchmarkRPCCall(b *testing.B) {
	startRPCServer()
	client, err := rpc.Dial("tcp", serverAddr)
	if err != nil {
		b.Fatal("dialing", err)
	}
	defer client.Close()

	// Synchronous calls
	args := &air.Args{7, 8}
	reply := new(air.Reply)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = client.Call("Arith.Add", args, reply)
		if err != nil {
			b.Errorf("Add: expected no error but got string %q", err.Error())
		}
		if reply.C != args.A+args.B {
			b.Errorf("Add: expected %d got %d", reply.C, args.A+args.B)
		}
	}
}
