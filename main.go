package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWss(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connecting to WSS")
	// Allow cross origin
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// Upgrade this connection to websocket
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if err != nil {
		log.Fatal("Can't upgrade websocket", err)
	}

	fmt.Println("Client connected..", ws.RemoteAddr())

}

func setupServer() {
	fmt.Println("Connecting websockets..")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Welcome to Socfetcher")
	})

	http.HandleFunc("/wss", handleWss)

	// Setup server
	http.ListenAndServe(":8080", nil)

}

func main() {
	setupServer()
}
