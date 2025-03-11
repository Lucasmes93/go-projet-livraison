# 📌 Système de Gestion des Livraisons en Go

👥 **Groupe Go** : Ianis CHENNAF, Philippe Ivan MBARGA, Mateo OUDART, Salman Ali MADEC, Lucas MESSIA DOLIVEUX

🔗 **Dépôt GitHub** : [go-projet-livraison](https://github.com/Lucasmes93/go-projet-livraison)

---

## 📝 **Description du Projet**

Ce projet permet de gérer un système de livraison multi-transport en **Go**, avec trois types de moyens de transport : `Truck`, `Drone`, `Boat`.  
Chaque transport implémente une interface commune et les livraisons sont exécutées en **parallèle** grâce aux **Goroutines** et **Channels**.

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
   go mod init go-projet-livraison
   ```

3️⃣ **Vérifier la structure du projet** :
   ```
   go-projet-livraison/
   ├── factory.go
   ├── go.mod
   ├── main.go
   ├── README.md
   ├── tracking.go
   ├── transport.go
   ├── transports.go
   ```

4️⃣ **Lancer le programme**  
Dans le terminal, exécute :
   ```sh
   go run main.go
   ```

📌 **Le serveur tourne et affiche les livraisons réalisées.**

---

## 📌 **Fonctionnalités du Projet**

✅ Interface `TransportMethod` pour standardiser les livraisons  
✅ Trois moyens de transport : Camion 🚚, Drone 🚁, Bateau 🚢  
✅ Fabrique `GetTransportMethod()` pour créer dynamiquement les transports  
✅ Système de suivi des livraisons en temps réel avec **Goroutines** et **Channels**  
✅ Gestion des erreurs (ex: batterie faible du drone, mauvais temps pour le bateau)  

---

## 📖 **Explication du Code**

- `transport.go` : Définit l’interface `TransportMethod`.
- `transports.go` : Implémente trois transports (`Truck`, `Drone`, `Boat`).
- `factory.go` : Contient la fabrique `GetTransportMethod()` pour créer dynamiquement un transport.
- `tracking.go` : Gère le suivi des livraisons avec des **Goroutines** et **Channels**.
- `main.go` : Orchestre les livraisons en exécutant les méthodes en parallèle.

---

## ✍️ **Auteurs**

Ce projet a été réalisé dans le cadre du **Groupe Go** par :
- Ianis CHENNAF  
- Philippe Ivan MBARGA  
- Mateo OUDART  
- Salman Ali MADEC  
- Lucas MESSIA DOLIVEUX  

---

## 📝 **Licence**

Ce projet est sous **licence MIT** - voir le fichier `LICENSE` pour plus de détails.

✨ Développé avec **Go** ❤️ ✨
