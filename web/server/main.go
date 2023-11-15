package main

import (
    "log"
    "net/http"
    "sync"
    "math/rand"
    "time"
    "github.com/gorilla/websocket"
)

// SecretSharing struct stocke les informations sur le partage de secret entre les clients.
type SecretSharing struct {
    sync.Mutex         // Mutex pour gérer l'accès concurrent aux données.
    Secret []int       // Tableau pour stocker les valeurs secrètes partagées.
    ClientsReady int   // Compteur pour les clients prêts.
    Clients []*websocket.Conn // Liste des connexions client WebSocket.
}

var secretSharing SecretSharing // Instance globale de SecretSharing.

// Configuration pour l'amélioration des connexions HTTP vers WebSocket.
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true }, // Permet toutes les origines.
}

func main() {
    rand.Seed(time.Now().UnixNano()) // Initialise le générateur de nombres aléatoires.

    http.HandleFunc("/ws", handleConnections) // Gère les connexions WebSocket.

    go handleMessages() // Lance une goroutine pour gérer les messages.

    log.Println("Serveur démarré sur :8080")
    err := http.ListenAndServe(":8080", nil) // Démarre le serveur HTTP sur le port 8080.
    if err != nil {
        log.Fatal("ListenAndServe: ", err) 
    }
}

// Gère les nouvelles connexions WebSocket.
func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil) // Améliore la connexion HTTP en WebSocket.
    if err != nil {
        log.Fatal("47", err)
    }
    defer ws.Close() // Assure que la connexion WebSocket est fermée à la fin.

    secretSharing.Lock()
    secretSharing.Clients = append(secretSharing.Clients, ws) // Ajoute le client WebSocket à la liste.
    secretSharing.Unlock()

    for {
        var msg map[string]interface{} // Modification ici
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("error 59 : %v", err)
            break
        }
        handleClientMessage(ws, msg)
    }    
}

// Gère les messages des clients WebSocket.
func handleClientMessage(client *websocket.Conn, msg map[string]interface{}) {
    // Traite différents types de messages.
    if valeur, ok := msg["nouvelle_valeur"]; ok {
        if val, ok := valeur.(int); ok {
            secretSharing.Lock()
            secretSharing.Secret = append(secretSharing.Secret, val)
            secretSharing.Unlock()
        }
    } else if _, ok := msg["obtenir_tableau"]; ok {
        secretSharing.Lock()
        err := client.WriteJSON(secretSharing.Secret)
        secretSharing.Unlock()
        if err != nil {
            log.Printf("error 78 : %v", err)
        }
    } else if _, ok := msg["reset_tableau"]; ok {
        secretSharing.Lock()
        secretSharing.Secret = nil
        secretSharing.Unlock()
    } else if _, ok := msg["client_pret"]; ok {
        print("client_pret")
        secretSharing.Lock()
        secretSharing.ClientsReady++
        if secretSharing.ClientsReady >= 2 {
            secretSharing.ClientsReady = 0
            for _, c := range secretSharing.Clients {
                err := c.WriteJSON(true)
                if err != nil {
                    log.Printf("error 93 : %v", err)
                }
            }
        }
        secretSharing.Unlock()
    }
}


func handleMessages() {
    for {
        // Ici, vous pouvez ajouter une logique pour gérer des messages globaux ou des événements,
        // par exemple, envoyer des mises à jour périodiques à tous les clients connectés.
    }
}
