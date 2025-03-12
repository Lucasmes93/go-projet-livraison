package main

import (
	"fmt"
	"go_projet_livraison/factory"
	"go_projet_livraison/tracking"
	"go_projet_livraison/transports"
)

func main() {
	ch := make(chan string)

	truck, _ := factory.GetTransportMethod("truck")
	drone, _ := factory.GetTransportMethod("drone")
	boat, _ := factory.GetTransportMethod("boat")

	// Vérifier la météo avant de lancer la livraison
	if b, ok := boat.(transports.Boat); ok && b.GetStatus() == "Météo actuelle : Clear" {
		go tracking.TrackDelivery(boat, "Paris", 300, ch)
	} else {
		fmt.Println("❌ Le bateau ne peut pas livrer à cause du mauvais temps.")
	}

	go tracking.TrackDelivery(truck, "New York", 200, ch)
	go tracking.TrackDelivery(drone, "Los Angeles", 80, ch)

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
