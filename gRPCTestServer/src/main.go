package main

import (
	"fmt"
	"net/http"

	pb "./protocol"
	"./service"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	/*listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}*/
	port := "19003"
	server := grpc.NewServer()
	catService := &service.MyCatService{}
	uploadService := &service.ServerGRPC{}
	// 実行したい実処理をseverに登録する
	//pb.RegisterCatServer(server, catService)
	pb.RegisterCatServer(server, uploadService)

	wrapedGrpc := grpcweb.WrapServer(server)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		print(req)
		wrapedGrpc.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(handler),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}

	//server.Serve(listenPort)
}
