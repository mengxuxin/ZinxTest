package main

import (
	"Zinx/src/zinx/znet"
	"fmt"
)

func main() {
	fmt.Println("start")
	s := znet.NewServer("[zinx v0.2]")
	s.Server()
}
