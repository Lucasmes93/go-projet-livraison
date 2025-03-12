package main

import (
	"go_projet_livraison/transports"
	"testing"
)

func TestTruckDelivery(t *testing.T) {
	truck := transports.Truck{ID: "T123", Capacity: 10}
	_, err := truck.DeliverPackage("Paris")
	if err != nil {
		t.Errorf("Truck failed to deliver")
	}
}

func TestDroneDelivery(t *testing.T) {
	drone := transports.Drone{ID: "D456", Battery: 100}
	_, err := drone.DeliverPackage("New York")
	if err != nil {
		t.Errorf("Drone failed to deliver")
	}
}

func TestBoatDelivery(t *testing.T) {
	boat := transports.Boat{ID: "B789", Weather: "Clear"}
	_, err := boat.DeliverPackage("London")
	if err != nil {
		t.Errorf("Boat failed to deliver")
	}
}
