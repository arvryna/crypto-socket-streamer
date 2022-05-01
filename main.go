package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
func fetchMarket() {
	url := "wss://stream.data.alpaca.markets/v1beta1/crypto"

	auth := map[string]string{"action": "auth", "key": "PK4WJX5BHJZBRYK4DGJQ", "secret": "yjdiafw56zgxVyR48LHrXCuCoSBTK7gAe5Vbrdv1"}
	subscribe := `{"action": "subscribe", "trades":["ETHUSD"], "quotes":["ETHUSD"], "bars":["ETHUSD"]}`
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Can't connect to Market socket", err)
	}

	conn.WriteJSON(auth)
	conn.WriteMessage(1, []byte(subscribe))

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}

func main() {
	go setupServer()
	go fetchMarket()
	time.Sleep(100000 * time.Second)
}
