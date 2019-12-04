package main

import (
	"fmt"
	"github.com/Hiroshii8/GoTrans/Server"
)

func main() {
	fmt.Print("start server")
	srv := Server.NewServer(":8080")
	if err := srv.Start(); err != nil {
		panic(err)
	}
}
