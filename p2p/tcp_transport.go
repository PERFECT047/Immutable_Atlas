package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn
	outbound bool
}

func createTCPPeer (conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mutexLock     sync.RWMutex
	peers         map[net.Addr]PeerI
}

func CreateTCPTransport (listernAddr string) *TCPTransport{
	return &TCPTransport{
		listenAddress: listernAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t * TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP Accpt error: %s\n", err)
		}

		go t.handleConnections(conn)
	}
}

func (t *TCPTransport) handleConnections(conn net.Conn) {

	peer := createTCPPeer(conn, true)
	fmt.Printf("Handling new Connection %+v\n", peer)
}
