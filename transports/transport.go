package transports

type TransportMethod interface {
	DeliverPackage(destination string) (string, error)
	GetStatus() string
}
