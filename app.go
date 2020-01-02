package main

import (
	"fmt"
	Server "github.com/Hiroshii8/GoTrans/server"
)

func main() {
	fmt.Print("start server")

	grpcSrv := Server.InitGrpc(":8081")
	grpcSrv.Start()

	srv := Server.NewServer(":8080")
	if err := srv.Start(); err != nil {
		panic(err)
	}
}
