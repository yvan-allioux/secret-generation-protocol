package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static")) // Sert les fichiers dans le dossier 'static'
	http.Handle("/", fs) // Tous les chemins sont gérés par le FileServer

	log.Println("Server starting on port 8090...")
	err := http.ListenAndServe(":8090", nil) // Change le port si nécessaire
	if err != nil {
		log.Fatal(err)
	}
}
