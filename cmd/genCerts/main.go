package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/lightningnetwork/lnd/cert"
)

var (
	tlsCertFile = flag.String("tlscertpath", "/Users/ellemouton/scratch/aperturetest/tls.cert", "tls cert file")
	tlsKeyFile  = flag.String("tlskeypath", "/Users/ellemouton/scratch/aperturetest/tls.key", "tls cert file")
)

func main() {
	flag.Parse()

	err := cert.GenCertPair("aperture test",
		*tlsCertFile,
		*tlsKeyFile,
		nil, nil, false, 14*30*24*time.Hour,
	)
	if err != nil {
		log.Fatalf("nope", err)
	}
	fmt.Println("yas")
}
