package main

import (
	"fmt"

	"github.com/Hiroshii8/GoTrans/server"
)

func main() {
	fmt.Println("start server")

	// starting grpc via go routine
	go func() error {
		grpcSrv := server.InitGrpc(":8081")
		fmt.Println("starting grpc")
		return grpcSrv.Start()
	}()

	srv := server.NewServer(":8080")
	if err := srv.Start(); err != nil {
		panic(err)
	}
}
