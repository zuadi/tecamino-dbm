package main

import (
	"fmt"

	"github.com/zuadi/tecamino-dbm.git/server"
)

func main() {
	fmt.Println("start")
	server := server.NewServer()
	panic(server.Serve(8100))
}
