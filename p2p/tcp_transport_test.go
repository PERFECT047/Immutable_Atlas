package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcp := TCPTransportOpts{
		ListenAddress: ":3000",
		ShakeHand:     NOPHandshakeFunc,
		Decoder:       NOPDecoder{},
	}
	tr := CreateTCPTransport(tcp)
	assert.Equal(t, tr.ListenAddress, ":3000")

	assert.Nil(t, tr.ListenAndAccept())

}
