package raven

import (
	"crypto/tls"
	"github.com/ravendb/ravendb-go-client"
	"log"
	"os"
)

var Store *ravendb.DocumentStore

func SetupStore() {

	if os.Getenv("PORT") != "" {
		//prod
		Store = ravendb.NewDocumentStoreWithUrlAndDatabase(os.Getenv("DB_HOST"), "")

		crt := os.Getenv("RAVEN_CERT")
		key := os.Getenv("RAVEN_KEY")

		cert, err := tls.X509KeyPair([]byte(crt), []byte(key))
		if err != nil {
			log.Println(err)
		}
		Store.SetCertificate(&ravendb.KeyStore{
			Certificates: []tls.Certificate{cert},
		})

	} else {
		//dev
		Store = ravendb.NewDocumentStoreWithUrlAndDatabase("http://localhost:8080/", "")

		cert, err := tls.LoadX509KeyPair("raven/vvanm.crt", "raven/vvanm.key")
		if err != nil {
			log.Println(err)
		}

		Store.SetCertificate(&ravendb.KeyStore{
			Certificates: []tls.Certificate{cert},
		})
	}

	Store.SetDatabase("rs-calcs")

	Store.Initialize()

}
