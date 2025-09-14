package p2p

type PeerI interface {

}

type TransportI interface {
	ListenAndAccept() error
}