package p2p

import (
	"errors"
	"fmt"
	"net"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func createTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddress string
	ShakeHand     HandshakeFunc
	Decoder       DecoderI
	OnPeer        func(PeerI) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC
}

func CreateTCPTransport(opt TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opt,
		rpcch:            make(chan RPC),
	}
}

// return read-only channel
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return
			}
			fmt.Printf("TCP accept error: %s\n", err)
			continue
		}

		fmt.Printf("new connection connecting %+v\n", conn)

		go t.handleConnections(conn)
	}
}

func (t *TCPTransport) handleConnections(conn net.Conn) {

	var err error

	defer func() {
		fmt.Printf("dropping peer connection: %s", err)
		conn.Close()
	}()

	peer := createTCPPeer(conn, true)

	if err := t.ShakeHand(peer); err != nil {
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}

	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			return
		}
	}

	rpc := RPC{}

	//Read Loop
	for {
		err := t.Decoder.Decode(conn, &rpc)

		if err != nil {
			fmt.Printf("TCP error: %s\n", err)
			return
		}

		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc

		fmt.Printf("message: %+v\n", rpc)
	}

}
