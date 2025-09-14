package p2p

type PeerI interface {
	Close() error
}

type TransportI interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}