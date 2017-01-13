package utils;
import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
  "io/ioutil"
  "google.golang.org/grpc/credentials"
)

var (
	StarsAppKeyPair  *tls.Certificate
	StarsAppCertPool *x509.CertPool
  Creds credentials.TransportCredentials
)

var(
  Addr string = "localhost:8587"
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
	StarsAppKeyPair = &pair
	StarsAppCertPool = x509.NewCertPool()
	ok := StarsAppCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("bad certs")
	}
  Creds = credentials.NewTLS(&tls.Config{
    ServerName: Addr,
    RootCAs:    StarsAppCertPool,
  })
}
