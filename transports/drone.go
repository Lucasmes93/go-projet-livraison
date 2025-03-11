package transports

import (
	"errors"
	"fmt"
)

type Drone struct {
	ID      string
	Battery int
}

func (d *Drone) DeliverPackage(destination string) (string, error) {
	if d.Battery < 20 {
		return "", errors.New("Drone battery too low")
	}
	d.Battery -= 20
	return fmt.Sprintf("Drone %s delivered package to %s", d.ID, destination), nil
}

func (d Drone) GetStatus() string {
	return fmt.Sprintf("Drone battery: %d%%", d.Battery)
}
