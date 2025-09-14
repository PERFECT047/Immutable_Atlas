package p2p

import "net"

// payload of any arbitrary type (will fix later) or error to be sent
type RPC struct {
	From    net.Addr
	Payload []byte
	Error   error
}