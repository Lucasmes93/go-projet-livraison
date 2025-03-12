package transports

import "fmt"

type Boat struct {
	ID      string
	Weather string
}

func (b Boat) DeliverPackage(destination string, distance int) (string, error) {
	if b.Weather != "Clear" {
		return "", fmt.Errorf("🌊 Mauvais temps, impossible de livrer")
	}
	return fmt.Sprintf("⛵ Boat %s a livré à %s (distance: %d km)", b.ID, destination, distance), nil
}

func (b Boat) GetStatus() string {
	return fmt.Sprintf("Météo actuelle : %s", b.Weather)
}
