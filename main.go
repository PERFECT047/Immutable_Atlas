package main

import (
	"log"

	"github.com/perfect047/immutable_atlas/p2p"
)

func main() {
	tr := p2p.TCPTransport{
		ListenAddress: ":3000",
		ShakeHand:     p2p.NOPHandshakeFunc,
		Decoder:       p2p.NOPDecoder{},
	}

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
