package transports

type TransportMethod interface {
	DeliverPackage(destination string, distance int) (string, error)
	GetStatus() string
}
