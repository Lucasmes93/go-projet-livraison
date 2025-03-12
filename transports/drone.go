package transports

import (
	"errors"
	"fmt"
)

type Drone struct {
	ID       string
	Battery  int
	MaxRange int // Distance maximale que le drone peut parcourir
}

func (d *Drone) DeliverPackage(destination string, distance int) (string, error) {
	if d.Battery < 20 {
		return "", errors.New("🔋 Batterie faible, impossible de livrer")
	}
	if distance > d.MaxRange {
		return "", errors.New("📍 Distance trop grande, le drone ne peut pas livrer")
	}

	d.Battery -= 20 // Consomme de la batterie
	return fmt.Sprintf("🚁 Drone %s a livré à %s", d.ID, destination), nil
}

func (d Drone) GetStatus() string {
	return fmt.Sprintf("Batterie à %d%%, Portée max: %d km", d.Battery, d.MaxRange)
}
