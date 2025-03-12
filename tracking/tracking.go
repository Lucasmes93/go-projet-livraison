package tracking

import (
	"fmt"
	"go_projet_livraison/transports"
)

func TrackDelivery(transport transports.TransportMethod, destination string, distance int, ch chan string) {
	status, err := transport.DeliverPackage(destination, distance)
	if err != nil {
		ch <- fmt.Sprintf("❌ Livraison échouée : %v", err)
		return
	}
	ch <- status
}
