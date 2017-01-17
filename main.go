package main

import (
	"fmt"
	"crypto/tls"
	"google.golang.org/grpc"
	"net"
	"log"
	"net/http"
	"strings"
	"stars-app/variables"
	"stars-app/utils"
	"stars-app/services"
)


func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main(){

	//Initializing the utilities and services
	utils.Init();
	services.Init();
	conn, err1 := net.Listen("tcp", fmt.Sprintf(":%d", 8587))
	if err1 != nil {
		panic(err1)
	}

	srv := &http.Server{
		Addr:    variables.Addr,
		Handler: grpcHandlerFunc(variables.GrpcServer, variables.Mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*variables.StarsAppKeyPair},
			NextProtos:   []string{"h2"},
		},
	}
	fmt.Printf("grpc on port: %d\n", 8587)
	err1 = srv.Serve(tls.NewListener(conn, srv.TLSConfig))
	if err1 != nil {
		log.Fatal("ListenAndServe: ", err1)
	}
}
