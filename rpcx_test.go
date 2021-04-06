package main

import (
	"context"
	"log"
	"net"
	"rpc-bench/airc"
	"testing"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
)

func startRPCxServer() {
	s := server.NewServer()
	s.Register(new(airc.Arith), "")

	var l net.Listener
	l, serverAddr = listenTCP()
	log.Println("Test RPC server listening on", serverAddr)
	go s.ServeListener("tcp", l)
}

func BenchmarkRPCXCall(b *testing.B) {
	startRPCxServer()

	c := client.NewClient(client.DefaultOption)
	err := c.Connect("tcp", serverAddr)
	if err != nil {
		b.Fatal("dialing", err)
	}
	defer c.Close()

	ctx := context.Background()
	args := &airc.Args{7, 8}
	reply := new(airc.Reply)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = c.Call(ctx, "Arith", "Add", args, reply)
		if err != nil {
			b.Errorf("Add: expected no error but got string %q", err.Error())
		}
		if reply.C != args.A+args.B {
			b.Errorf("Add: expected %d got %d", reply.C, args.A+args.B)
		}
	}
}
