package main

import (
	"fmt"
	"go_projet_livraison/factory"
	"go_projet_livraison/tracking"
)

func main() {
	ch := make(chan string)

	truck, _ := factory.GetTransportMethod("truck")
	drone, _ := factory.GetTransportMethod("drone")
	boat, _ := factory.GetTransportMethod("boat")

	go tracking.TrackDelivery(truck, "New York", ch)
	go tracking.TrackDelivery(drone, "Los Angeles", ch)
	go tracking.TrackDelivery(boat, "Paris", ch)

	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
