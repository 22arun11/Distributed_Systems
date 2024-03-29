package main

import (
    "log"
    "net/rpc"
)

type Args struct {
    A, B int
}

func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    args := Args{7,8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    log.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}