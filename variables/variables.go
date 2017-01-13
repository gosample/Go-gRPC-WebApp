package variables;
import (
  "net/http"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials"
  "crypto/tls"
  "crypto/x509"


)

var (
  GrpcServer *grpc.Server
  Mux *http.ServeMux
  StarsAppKeyPair  *tls.Certificate
  StarsAppCertPool *x509.CertPool
  Creds credentials.TransportCredentials
  Addr string = "localhost:8587"
)
