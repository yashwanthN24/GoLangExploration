package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/YashwanthN24/GoProject/internal/config"
	"github.com/YashwanthN24/GoProject/internal/db"
)


func main() {
	db.InitDB()
	port := strings.TrimSpace(config.GetEnv("API_PORT", "8080"))

	
	fmt.Printf("🌐 Starting HTTP API Microservice on port %s...\n", port)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the API Microservice!")
	})
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
