package services;

import (
	"fmt"
	"net/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
  "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
  "stars-app/variables"
  pbLogin "stars-app/services/login"
	pbList "stars-app/services/list"
)


func Init(){
  variables.Mux = http.NewServeMux()
  opts := []grpc.ServerOption{
    grpc.Creds(credentials.NewClientTLSFromCert(variables.StarsAppCertPool, variables.Addr)),
  }
  variables.GrpcServer = grpc.NewServer(opts...)

  pbLogin.RegisterLoginServiceServer(variables.GrpcServer, new(pbLogin.AuthServices))
  pbList.RegisterListStarsServiceServer(variables.GrpcServer, new(pbList.GitHubServices))

  ctx := context.Background()
  dopts := []grpc.DialOption{grpc.WithTransportCredentials(variables.Creds)}
  gwmux := runtime.NewServeMux()

  err1 := pbLogin.RegisterLoginServiceHandlerFromEndpoint(ctx, gwmux, variables.Addr, dopts)
  if err1 != nil {
    fmt.Printf("serve: %v\n", err1)
    return
  }
  err1 = pbList.RegisterListStarsServiceHandlerFromEndpoint(ctx, gwmux, variables.Addr, dopts)
  if err1 != nil {
    fmt.Printf("serve: %v\n", err1)
    return
  }

  variables.Mux.Handle("/api/", gwmux)
  variables.Mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	variables.Mux.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("./public/views"))))
	variables.Mux.Handle("/modals/", http.StripPrefix("/modals/", http.FileServer(http.Dir("./public/modals"))))
	variables.Mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/static/css"))))
	variables.Mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./public/static/js"))))
	variables.Mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./public/static/images"))))
  variables.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./public/views/index.html")
  })
}
