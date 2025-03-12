package factory

import (
	"errors"
	"go_projet_livraison/transports"
)

func GetTransportMethod(method string) (transports.TransportMethod, error) {
	switch method {
	case "truck":
		return transports.Truck{ID: "T123", Capacity: 10}, nil
	case "drone":
		return &transports.Drone{ID: "D456", Battery: 100, MaxRange: 50}, nil
	case "boat":
		return transports.Boat{ID: "B789", Weather: "Stormy"}, nil
	default:
		return nil, errors.New("unknown transport method")
	}
}
