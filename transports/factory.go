package transports

import (
	"errors"
)

// GetTransportMethod crée dynamiquement une méthode de transport en fonction du type fourni.
func GetTransportMethod(method string) (TransportMethod, error) {
	switch method {
	case "truck":
		return &Truck{ID: "T123", MaxCapacity: 10}, nil
	case "drone":
		return &Drone{ID: "D456", Battery: 100, MaxRange: 50}, nil
	case "boat":
		return &Boat{ID: "B789", Capacity: 500, WeatherOK: true}, nil
	default:
		return nil, errors.New("méthode de transport inconnue")
	}
}
