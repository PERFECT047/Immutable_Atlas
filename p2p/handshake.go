package p2p

type HandshakeFunc func(PeerI) error

func NOPHandshakeFunc(PeerI) error { return nil }
