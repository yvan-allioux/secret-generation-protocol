package main

import (
    "log"
    "net/http"
    "sync"
    "time"
    "github.com/gorilla/websocket"
)

// SecretItem représente un élément du secret avec sa valeur et le timestamp de sa réception.
type SecretItem struct {
    Value     int       // La valeur du secret.
    Timestamp time.Time // Le moment de la réception de la valeur.
}

// SecretSharing struct modifiée pour stocker des éléments de SecretItem.
type SecretSharing struct {
    sync.Mutex
    Secret []SecretItem               // Tableau pour stocker les secrets avec timestamps.
    ClientsReady int                  // Compteur pour les clients prêts.
    Clients []*websocket.Conn         // Tableau des connexions clients.
    ClientIDs map[*websocket.Conn]int // Stocke les identifiants des clients.
    NextClientID int                  // Génère le prochain identifiant.
    ClientDone map[int]bool           // Stocke les clients qui ont terminé.
}

// Instance globale de SecretSharing.
var secretSharing = SecretSharing{
    ClientIDs: make(map[*websocket.Conn]int),
    NextClientID: 0,
    ClientDone: make(map[int]bool), // Initialisation de la map.
}

// Configuration pour l'upgrade de HTTP vers WebSocket.
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true }, // Permet toutes les origines.
}

func main() {
    // Définit la route et le handler pour les connexions WebSocket.
    http.HandleFunc("/ws", handleConnections)

    // Lance une goroutine pour gérer les messages.
    go handleMessages()

    // Démarre le serveur sur le port 8080.
    log.Println("Serveur démarré sur :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

// Gère les nouvelles connexions WebSocket.
func handleConnections(w http.ResponseWriter, r *http.Request) {
    // Upgrade la requête HTTP vers une connexion WebSocket.
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal("Erreur d'upgrade WebSocket", err)
    }
    defer ws.Close()

    // Enregistre le nouveau client.
    secretSharing.Lock()
    secretSharing.NextClientID++
    clientID := secretSharing.NextClientID
    secretSharing.ClientIDs[ws] = clientID
    secretSharing.Clients = append(secretSharing.Clients, ws)
    secretSharing.Unlock()

    // Boucle pour lire les messages du client.
    for {
        var msg map[string]interface{}
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("Erreur de lecture JSON : %v", err)
            break
        }
        handleClientMessage(ws, msg)
    }
}

func sendConfirmationToAll() {
    println("Tous les clients ont terminé d'envoyer des requêtes.")

    
    for _, client := range secretSharing.Clients {
        err := client.WriteJSON(map[string]string{"confirmation": "Tous les clients ont terminé d'envoyer des requêtes."})
        if err != nil {
            log.Printf("Erreur lors de l'envoi de la confirmation : %v", err)
        }
    }
}

// Traite les messages reçus d'un client.
func handleClientMessage(client *websocket.Conn, msg map[string]interface{}) {
    // Ajoute une nouvelle valeur au secret partagé.
    if valeur, ok := msg["nouvelle_valeur"]; ok {
        secretSharing.Lock()
        // Création de l'élément SecretItem avec la valeur et le timestamp actuel.
        newItem := SecretItem{
            Value:     int(valeur.(float64)),
            Timestamp: time.Now(),
        }
        secretSharing.Secret = append(secretSharing.Secret, newItem)
        secretSharing.Unlock()  
    } else if _, ok := msg["obtenir_tableau"]; ok {
        // Envoie le tableau actuel de secrets au client.
        secretSharing.Lock()
        err := client.WriteJSON(secretSharing.Secret)
        secretSharing.Unlock()
        if err != nil {
            log.Printf("Erreur d'envoi JSON : %v", err)
        }
    } else if _, ok := msg["reset_tableau"]; ok {
        // Réinitialise le tableau de secrets.
        secretSharing.Lock()
        secretSharing.Secret = nil
        secretSharing.Unlock()
    } else if _, ok := msg["client_a_terminer_denvoyer"]; ok {
        secretSharing.Lock()
        clientID := secretSharing.ClientIDs[client]
        secretSharing.ClientDone[clientID] = true

        // Vérifie si tous les clients ont terminé d'envoyer des requêtes.
        allDone := true
        for _, done := range secretSharing.ClientDone {
            if !done {
                allDone = false
                break
            }
        }
        if allDone && len(secretSharing.ClientDone) >= 2 {
            sendConfirmationToAll()
            // Réinitialise la map.
            secretSharing.ClientDone = make(map[int]bool)
        }
        secretSharing.Unlock()
    } else if _, ok := msg["client_pret"]; ok {
        // Gère l'état 'prêt' des clients.
        secretSharing.Lock()
        secretSharing.ClientsReady++
        if secretSharing.ClientsReady >= 2 {
            secretSharing.ClientsReady = 0
            for _, c := range secretSharing.Clients {
                clientID, exists := secretSharing.ClientIDs[c]
                if !exists {
                    log.Printf("ID client introuvable pour une connexion")
                    continue
                }
                // Envoie l'identifiant unique à chaque client.
                err := c.WriteJSON(map[string]int{"client_id": clientID})
                if err != nil {
                    log.Printf("Erreur d'envoi JSON : %v", err)
                }
            }
        }
        secretSharing.Unlock()
    }
}

// Gère les messages globaux ou les événements.
func handleMessages() {
    // Logique pour gérer les messages globaux ou les événements.
}
