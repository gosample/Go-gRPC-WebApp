import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
  "io/ioutil"
)

var (
	starsAppKeyPair  *tls.Certificate
	starsAppCertPool *x509.CertPool
)

func init() {

  keyBytes, err := ioutil.ReadFile("stars-app/utils/certs/key.pem")
  if err != nil {
    fmt.Print(err)
  }

  certBytes, err := ioutil.ReadFile("stars-app/utils/certs/cert.pem")
  if err != nil {
    fmt.Print(err)
  }

	var err error
	pair, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		panic(err)
	}
	starsAppKeyPair = &pair
	starsAppCertPool = x509.NewCertPool()
	ok := starsAppCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("bad certs")
	}
}
