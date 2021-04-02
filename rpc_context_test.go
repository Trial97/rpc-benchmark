package main

import (
	"context"
	"log"
	"net"
	"rpc-bench/airc"
	"testing"

	"github.com/keegancsmith/rpc"
)

func startRPCServerWithContext() {
	rpc.Register(new(airc.Arith))

	var l net.Listener
	l, serverAddr = listenTCP()
	log.Println("Test RPC server listening on", serverAddr)
	go rpc.Accept(l)
}

func BenchmarkRPCCallWithContext(b *testing.B) {
	startRPCServerWithContext()
	client, err := rpc.Dial("tcp", serverAddr)
	if err != nil {
		b.Fatal("dialing", err)
	}
	defer client.Close()

	// Synchronous calls
	ctx := context.Background()
	args := &airc.Args{7, 8}
	reply := new(airc.Reply)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = client.Call(ctx, "Arith.Add", args, reply)
		if err != nil {
			b.Errorf("Add: expected no error but got string %q", err.Error())
		}
		if reply.C != args.A+args.B {
			b.Errorf("Add: expected %d got %d", reply.C, args.A+args.B)
		}
	}
}
