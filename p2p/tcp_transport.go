package p2p

import (
	"fmt"
	"net"
	"sync"
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

type TCPTransport struct {
	ListenAddress string
	ShakeHand     HandshakeFunc
	Decoder       DecoderI
	listener      net.Listener

	mutexLock sync.RWMutex
	peers     map[net.Addr]PeerI
}

func CreateTCPTransport(listernAddr string) *TCPTransport {
	return &TCPTransport{
		ShakeHand: NOPHandshakeFunc,
		ListenAddress: listernAddr,
	}
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
			fmt.Printf("TCP Accpt error: %s\n", err)
		}

		fmt.Printf("new connection connecting %+v\n", conn)

		go t.handleConnections(conn)
	}
}

func (t *TCPTransport) handleConnections(conn net.Conn) {

	peer := createTCPPeer(conn, true)

	if err := t.ShakeHand(peer); err != nil {
		fmt.Printf("TCP handshake error: %s\n", err)
		conn.Close()
		return 
	}

	msg := &Message{}

	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()

		fmt.Printf("message: %+v\n", msg)
	}

}
