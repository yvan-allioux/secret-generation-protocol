package main

import (
    "encoding/json"
    "net/http"
    "sync"
	"math/rand"
	"time"
)

type SecretSharing struct {
    sync.Mutex
    Secret []int
    ClientsReady int // Compteur pour les clients prêts
}

var secretSharing SecretSharing

func main() {
	// Initialisez le générateur de nombres aléatoires
    rand.Seed(time.Now().UnixNano())

    http.HandleFunc("/nouvelle_valeur", NouvelleValeur)
    http.HandleFunc("/obtenir_tableau", ObtenirTableau)
    http.HandleFunc("/reset_tableau", ResetTableau)
    http.HandleFunc("/client_pret", ClientPret)

    http.ListenAndServe(":8080", nil)
}

func NouvelleValeur(w http.ResponseWriter, r *http.Request) {
    var valeur int
    // Parse la valeur depuis la requête
    err := json.NewDecoder(r.Body).Decode(&valeur)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Ajoute la nouvelle valeur au tableau secret
    secretSharing.Lock()
    secretSharing.Secret = append(secretSharing.Secret, valeur)
    secretSharing.Unlock()

    w.WriteHeader(http.StatusCreated)
}

func ObtenirTableau(w http.ResponseWriter, r *http.Request) {
    secretSharing.Lock()
    defer secretSharing.Unlock()

    // Renvoie le tableau secret en tant que réponse JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(secretSharing.Secret)
}

func ResetTableau(w http.ResponseWriter, r *http.Request) {
    secretSharing.Lock()
    defer secretSharing.Unlock()

    // Réinitialise le tableau secret
    secretSharing.Secret = nil

    w.WriteHeader(http.StatusOK)
}

func ClientPret(w http.ResponseWriter, r *http.Request) {
    secretSharing.Lock()
    defer secretSharing.Unlock()

    // Incrémente le compteur des clients prêts
    secretSharing.ClientsReady++
    if secretSharing.ClientsReady >= 2 {
        // Si les deux clients sont prêts, réinitialise le compteur et effectue le traitement
        secretSharing.ClientsReady = 0
        // Renvoie l'instruction de traitement au client'
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(true)
    }

    w.WriteHeader(http.StatusOK)
}


