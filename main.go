package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/arvryna/gridbot/fetcher"
	"github.com/gorilla/websocket"
)

// TODO:

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(connection *websocket.Conn) {
	// ping/pong
	for {

		msgType, data, err := connection.ReadMessage()
		if err != nil {
			log.Println("Client error, terminating connection", err, connection.RemoteAddr())
			connection.Close()
			return
		}

		msg := string(data)
		fmt.Println("Received", string(msg), "client", connection.RemoteAddr())

		if msg == "Ping" {
			err = connection.WriteMessage(msgType, []byte("Pong"))
			if err != nil {
				log.Println("Writing failed to client", err)
			}
		}

	}
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

	reader(ws)

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

// Fetcher
func main() {
	go setupServer()
	fetcher.Init()
	time.Sleep(100000 * time.Second)
}
