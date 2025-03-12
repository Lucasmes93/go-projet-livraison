package main

import (
	"go_projet_livraison/transports"
	"testing"
)

// 🚚 Test pour le Truck
func TestTruckDelivery(t *testing.T) {
	truck := transports.Truck{ID: "T123", Capacity: 10}
	_, err := truck.DeliverPackage("Paris", 150) // ✅ Ajout de la distance
	if err != nil {
		t.Errorf("Truck failed to deliver: %v", err)
	}
}

// 🚁 Test pour le Drone
func TestDroneDelivery(t *testing.T) {
	drone := transports.Drone{ID: "D456", Battery: 100, MaxRange: 50}
	_, err := drone.DeliverPackage("New York", 40) // ✅ Ajout de la distance (doit être ≤ MaxRange)
	if err != nil {
		t.Errorf("Drone failed to deliver: %v", err)
	}
}

// ⛵ Test pour le Boat
func TestBoatDelivery(t *testing.T) {
	boat := transports.Boat{ID: "B789", Weather: "Clear"}
	_, err := boat.DeliverPackage("London", 500) // ✅ Ajout de la distance
	if err != nil {
		t.Errorf("Boat failed to deliver: %v", err)
	}
}
