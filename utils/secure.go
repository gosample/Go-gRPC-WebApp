package utils;
import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
  "io/ioutil"
  "google.golang.org/grpc/credentials"
  "stars-app/variables"
)


func Init() {

  keyBytes, err := ioutil.ReadFile("utils/certs/server.key")
  if err != nil {
    fmt.Print(err)
  }

  certBytes, err := ioutil.ReadFile("utils/certs/server.pem")
  if err != nil {
    fmt.Print(err)
  }

	pair, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		panic(err)
	}
	variables.StarsAppKeyPair = &pair
	variables.StarsAppCertPool = x509.NewCertPool()
	ok := variables.StarsAppCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("bad certs")
	}
  variables.Creds = credentials.NewTLS(&tls.Config{
    ServerName: variables.Addr,
    RootCAs:    variables.StarsAppCertPool,
  })
}
