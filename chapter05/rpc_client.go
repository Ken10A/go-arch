package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	serverAddress := "127.0.0.1"

	c, e := rpc.DialHTTP("tcp", serverAddress+":1234")
	if e != nil {
		log.Fatal("dialing: ", e)
	}

	args := Args{3, 4}
	var reply int
	c.Call("MultiplyService.Do", args, &reply)
	log.Printf("%d * %d = %d", args.A, args.B, reply)
}
