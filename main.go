package main

import (
	"log"

	"github.com/perfect047/immutable_atlas/p2p"
)

func main() {
	tcpopts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		ShakeHand:     p2p.NOPHandshakeFunc,
		Decoder:       p2p.NOPDecoder{},
	}

	tr := p2p.CreateTCPTransport(tcpopts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
