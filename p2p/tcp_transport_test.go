package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := CreateTCPTransport(listenAddr)
	assert.Equal(t, tr.ListenAddress, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())

}
