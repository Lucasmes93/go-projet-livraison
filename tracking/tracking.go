package tracking

import (
	"fmt"
	"go_projet_livraison/transports"
)

func TrackDelivery(transport transports.TransportMethod, destination string, ch chan string) {
	status, err := transport.DeliverPackage(destination)
	if err != nil {
		ch <- fmt.Sprintf("Delivery failed: %v", err)
		return
	}
	ch <- status
}
