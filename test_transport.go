package main

import (
	"errors"
	"fmt"
	"gestion_livraison/transports" // Remplace "gestion_livraison" par le bon nom de ton module
	"math/rand"
	"time"
)

// TestTransport est une implémentation temporaire de TransportMethod
type TestTransport struct {
	ID string
}

type Truck struct {
	ID          string
	MaxCapacity int
}

func (t Truck) DeliverPackage(destination string) (string, error) {
	time.Sleep(3 * time.Second) // Simule un délai de livraison
	return fmt.Sprintf("🚛 Truck %s a livré à %s", t.ID, destination), nil
}

func (t Truck) GetStatus() string {
	return "En route avec une livraison fiable"
}

type Drone struct {
	ID       string
	Battery  int // Niveau de batterie en pourcentage
	MaxRange int // Distance max en km
}

func (d *Drone) DeliverPackage(destination string) (string, error) {
	if d.Battery < 20 {
		return "", errors.New("🔋 Batterie faible, impossible de livrer")
	}
	time.Sleep(1 * time.Second) // Livraison rapide
	d.Battery -= rand.Intn(30)  // Diminue la batterie de manière aléatoire
	return fmt.Sprintf("🚁 Drone %s a livré à %s", d.ID, destination), nil
}

func (d *Drone) GetStatus() string {
	return fmt.Sprintf("Batterie à %d%%", d.Battery)
}

type Boat struct {
	ID        string
	Capacity  int
	WeatherOK bool
}

func (b Boat) DeliverPackage(destination string) (string, error) {
	if !b.WeatherOK {
		return "", errors.New("🌊 Mauvais temps, impossible de livrer")
	}
	time.Sleep(5 * time.Second) // Simule un délai long
	return fmt.Sprintf("⛵ Boat %s a livré à %s", b.ID, destination), nil
}

func (b Boat) GetStatus() string {
	if b.WeatherOK {
		return "Météo favorable 🌤"
	}
	return "Météo défavorable 🌧"
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
