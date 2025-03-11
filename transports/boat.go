package transports

import (
	"fmt"
)

type Boat struct {
	ID      string
	Weather string
}

func (b Boat) DeliverPackage(destination string) (string, error) {
	if b.Weather != "Clear" {
		return "", fmt.Errorf("Boat cannot deliver due to bad weather (%s)", b.Weather)
	}
	return fmt.Sprintf("Boat %s delivered package to %s", b.ID, destination), nil
}

func (b Boat) GetStatus() string {
	return fmt.Sprintf("Boat weather condition: %s", b.Weather)
}
