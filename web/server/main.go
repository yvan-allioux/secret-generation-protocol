package main

import (
    "log"
    "net/http"
    "sync"
    "math/rand"
    "time"
    "github.com/gorilla/websocket"
)

type SecretSharing struct {
    sync.Mutex
    Secret []int
    ClientsReady int // Compteur pour les clients prêts
    Clients []*websocket.Conn
}

var secretSharing SecretSharing

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func main() {
    rand.Seed(time.Now().UnixNano())

    http.HandleFunc("/ws", handleConnections)

    go handleMessages()

    log.Println("Serveur démarré sur :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    secretSharing.Lock()
    secretSharing.Clients = append(secretSharing.Clients, ws)
    secretSharing.Unlock()

    for {
        var msg map[string]int
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("error: %v", err)
            break
        }
        handleClientMessage(ws, msg)
    }
}

func handleClientMessage(client *websocket.Conn, msg map[string]int) {
    if valeur, ok := msg["nouvelle_valeur"]; ok {
        secretSharing.Lock()
        secretSharing.Secret = append(secretSharing.Secret, valeur)
        secretSharing.Unlock()
    } else if _, ok := msg["obtenir_tableau"]; ok {
        secretSharing.Lock()
        err := client.WriteJSON(secretSharing.Secret)
        secretSharing.Unlock()
        if err != nil {
            log.Printf("error: %v", err)
        }
    } else if _, ok := msg["reset_tableau"]; ok {
        secretSharing.Lock()
        secretSharing.Secret = nil
        secretSharing.Unlock()
    } else if _, ok := msg["client_pret"]; ok {
        secretSharing.Lock()
        secretSharing.ClientsReady++
        if secretSharing.ClientsReady >= 2 {
            secretSharing.ClientsReady = 0
            for _, c := range secretSharing.Clients {
                err := c.WriteJSON(true)
                if err != nil {
                    log.Printf("error: %v", err)
                }
            }
        }
        secretSharing.Unlock()
    }
}

func handleMessages() {
    for {
        // Ici, vous pouvez ajouter une logique pour gérer des messages globaux ou des événements
        // par exemple, envoyer des mises à jour périodiques à tous les clients connectés.
    }
}
