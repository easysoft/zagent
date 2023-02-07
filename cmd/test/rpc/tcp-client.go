package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smallnest/rpcx/client"
)

func main() {
	// #1
	url := fmt.Sprintf("tcp@127.0.0.1:%d", RpcPort)
	d, err := client.NewPeer2PeerDiscovery(url, "")
	// #2
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	// #3
	args := &Args{
		A: 10,
		B: 20,
	}

	// #4
	reply := &Reply{}

	// #5
	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
