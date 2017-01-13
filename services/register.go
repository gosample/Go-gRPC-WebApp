package services;

import (
	"fmt"
	"net/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
  "stars-app/utils"
  user "stars-app/messages/user"
  ghResponse "stars-app/messages/ghResponse"
  pbLogin "stars-app/services/login"
	pbList "stars-app/services/list"
)

type Services struct{}

func (m *Services) Login(c context.Context, s *user.User) (*user.User, error) {
	fmt.Printf("rpc request Echo(%q)\n", s.Username)
	return s, nil
}

func (m *Services) ListStars(c context.Context, s *ghResponse.List) (*ghResponse.List, error) {
	fmt.Printf("rpc request Echo\n")
	return s, nil
}

var (
  GrpcServer *grpc.Server
  Mux *http.ServeMux
)

func Init(){
  Mux = http.NewServeMux()
  opts := []grpc.ServerOption{
    grpc.Creds(credentials.NewClientTLSFromCert(utils.StarsAppCertPool, utils.Addr)),
  }
  GrpcServer = grpc.NewServer(opts...)
  pbLogin.RegisterLoginServiceServer(GrpcServer, new(Services))
  pbList.RegisterListStarsServiceServer(GrpcServer, new(Services))

  ctx := context.Background()
  dopts := []grpc.DialOption{grpc.WithTransportCredentials(utils.Creds)}
  gwmux := runtime.NewServeMux()

  err1 := pbLogin.RegisterLoginServiceHandlerFromEndpoint(ctx, gwmux, utils.Addr, dopts)
  if err1 != nil {
    fmt.Printf("serve: %v\n", err1)
    return
  }
  err1 = pbList.RegisterListStarsServiceHandlerFromEndpoint(ctx, gwmux, utils.Addr, dopts)
  if err1 != nil {
    fmt.Printf("serve: %v\n", err1)
    return
  }

  Mux.Handle("/", gwmux)
}
