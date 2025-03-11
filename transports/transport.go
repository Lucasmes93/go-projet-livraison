package transports

// TransportMethod définit une interface pour les méthodes de transport.
type TransportMethod interface {
	DeliverPackage(destination string) (string, error) // Transporter un colis
	GetStatus() string                                 // Renvoyer l’état actuel du transport
}
