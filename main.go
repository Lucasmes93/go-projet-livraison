package main

import (
	"fmt"
	"go_projet_livraison/factory"
	"go_projet_livraison/tracking"
	"go_projet_livraison/transports"
)

func main() {
	deliveries := 0

	ch := make(chan string)

	truck, _ := factory.GetTransportMethod("truck")
	drone, _ := factory.GetTransportMethod("drone")
	boat, _ := factory.GetTransportMethod("boat")

	if b, ok := boat.(transports.Boat); ok && b.GetStatus() == "Météo actuelle : Clear" {
		go tracking.TrackDelivery(boat, "Paris", 300, ch)
		deliveries++
	}

	go tracking.TrackDelivery(truck, "New York", 200, ch)
	go tracking.TrackDelivery(drone, "Los Angeles", 80, ch)
	deliveries += 2


	for i := 0; i < deliveries; i++ {
		fmt.Println(<-ch)
	}
}
