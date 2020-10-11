package main

import (
	"fmt"

	"github.com/Hiroshii8/GoTrans/server"
)

func main() {
	fmt.Print("start server")

	grpcSrv := server.InitGrpc(":8081")
	grpcSrv.Start()

	srv := server.NewServer(":8080")
	if err := srv.Start(); err != nil {
		panic(err)
	}
}
