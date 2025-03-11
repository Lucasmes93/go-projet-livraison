# 📌 Système de Gestion des Livraisons en Go

👥 **Groupe Go** : Ianis CHENNAF, Philippe Ivan MBARGA, Mateo OUDART, Salman Ali MADEC, Lucas MESSIA DOLIVEUX

🔗 **Dépôt GitHub** : [go-projet-livraison](https://github.com/Lucasmes93/go-projet-livraison)

---

## 📝 **Description du Projet**

Le projet **Système de Gestion des Livraisons** est une application développée en **Go**, permettant de gérer efficacement l’expédition de colis via différents moyens de transport :
- 🚚 **Camion (Truck)** : Idéal pour les longues distances sur route.
- 🚁 **Drone (Drone)** : Rapide mais limité en autonomie.
- 🚢 **Bateau (Boat)** : Convient aux expéditions maritimes.

Grâce aux **Goroutines** et **Channels**, ce système permet d’exécuter les livraisons en parallèle et d’optimiser leur suivi.

---

## 🚀 **Installation & Exécution**

### 🔹 **Prérequis**
- **Go installé** (`go version` pour vérifier)
- Un terminal (Invite de commande, PowerShell, ou Terminal Linux/macOS)

### 🔹 **Installation**

1️⃣ **Cloner le projet** :
   ```sh
   git clone https://github.com/Lucasmes93/go-projet-livraison.git
   cd go-projet-livraison
   ```

2️⃣ **Initialiser le module Go** (si ce n'est pas déjà fait) :
   ```sh
   go mod init go_projet_livraison
   ```

3️⃣ **Vérifier la structure du projet** :
   ```
  go-projet-livraison/
│── transports/               # Contient les méthodes de transport
│   ├── boat.go               # Bateau (Boat)
│   ├── drone.go              # Drone (Drone)
│   ├── truck.go              # Camion (Truck)
│   ├── transport.go          # Interface TransportMethod
│── factory/                  # Fabrique pour instancier les transports
│   ├── factory.go
│── tracking/                 # Suivi des livraisons avec Goroutines
│   ├── tracking.go
│── main.go                   # Point d'entrée du programme
│── go.mod                    # Fichier des dépendances Go
│── test_transport.go         # Tests unitaires
│── README.md                 # Documentation du projet

   ```

4️⃣ **Lancer le programme**  
Dans le terminal, exécute :
   ```sh
   go run main.go
   ```

📌 **Le programme affiche les livraisons en cours et leurs statuts en temps réel.**

---

## 📌 **Fonctionnalités du Projet**

✅ Interface `TransportMethod` pour standardiser les livraisons  
✅ Trois moyens de transport : **Camion** 🚚, **Drone** 🚁, **Bateau** 🚢  
✅ **Usine de création dynamique** : `GetTransportMethod()`  
✅ Gestion des livraisons **en parallèle** avec **Goroutines** et **Channels**  
✅ Gestion des erreurs :
   - 📉 **Batterie faible du drone**
   - 🌊 **Conditions météorologiques défavorables pour le bateau**
   
---

## 📖 **Explication du Code**

- **`transport.go`** : Définit l’interface `TransportMethod`.
- **`transports.go`** : Implémente les classes `Truck`, `Drone`, `Boat`.
- **`factory.go`** : Fournit `GetTransportMethod()` pour générer un transport.
- **`tracking.go`** : Gère le suivi des livraisons via **Channels**.
- **`main.go`** : Exécute les livraisons en **parallèle**.

---

## 📌 **Exemple d'Exécution**

```sh
$ go run main.go
🚚 Truck T123 delivered package to New York
🚁 Drone D456 delivered package to Los Angeles
🚢 Boat B789 delivered package to Paris
```

## ✍️ **Auteurs**

Ce projet a été réalisé dans le cadre du **Groupe Go** par :
- **Ianis CHENNAF**  
- **Philippe Ivan MBARGA**  
- **Mateo OUDART**  
- **Salman Ali MADEC**  
- **Lucas MESSIA DOLIVEUX**  

---

## 📝 **Licence**

Ce projet est sous **licence MIT** - voir le fichier `LICENSE` pour plus de détails.

✨ Développé avec **Go** ❤️ ✨
