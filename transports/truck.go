package transports

import "fmt"

type Truck struct {
	ID       string
	Capacity int
}

func (t Truck) DeliverPackage(destination string) (string, error) {
	return fmt.Sprintf("Truck %s delivered package to %s", t.ID, destination), nil
}

func (t Truck) GetStatus() string {
	return "Truck ready"
}
