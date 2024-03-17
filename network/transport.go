package network

type RPC struct {
	FROM string
	Payload []byte
}

type Transport interface {
	Consume() <-chain RPC
	Connect(Transport) error
}
