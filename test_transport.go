package main

import (
	"fmt"
	"gestion_livraison/transports" // Remplace "gestion_livraison" par le bon nom de ton module
)

// TestTransport est une implémentation temporaire de TransportMethod
type TestTransport struct {
	ID string
}

// Implémentation de DeliverPackage pour TestTransport
func (t TestTransport) DeliverPackage(destination string) (string, error) {
	return fmt.Sprintf("TestTransport %s : Livraison simulée à %s", t.ID, destination), nil
}

// Implémentation de GetStatus pour TestTransport
func (t TestTransport) GetStatus() string {
	return "Test OK"
}

func main() {
	// Création d'un objet qui respecte TransportMethod
	var transport transports.TransportMethod = TestTransport{ID: "TEST123"}

	// Test des méthodes
	result, err := transport.DeliverPackage("Paris")
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println("État :", transport.GetStatus())
}
