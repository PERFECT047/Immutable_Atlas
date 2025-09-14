package main

import (
	"log"

	"github.com/perfect047/immutable_atlas/p2p"
)

func main() {
	tr := p2p.CreateTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select{}
}
