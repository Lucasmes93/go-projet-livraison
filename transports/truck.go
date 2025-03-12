package transports

import "fmt"

type Truck struct {
	ID       string
	Capacity int
}

func (t Truck) DeliverPackage(destination string, distance int) (string, error) {
	return fmt.Sprintf("🚚 Truck %s a livré à %s (distance: %d km)", t.ID, destination, distance), nil
}

func (t Truck) GetStatus() string {
	return "En route avec une livraison fiable"
}
