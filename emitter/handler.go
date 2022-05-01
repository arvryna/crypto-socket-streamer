package emitter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arvryna/gridbot/fetcher"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWsConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connection request from Client")
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

// These are internally spawned as individual go-routines
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
		fmt.Println("Received", string(msg), "from client", connection.RemoteAddr())

		if msg == "Ping" { // Send ping-pong in separate go routine as health checks
			err = connection.WriteMessage(msgType, []byte("Pong"))
			if err != nil {
				log.Println("Writing failed to client", err)
			}
		}

		// you need to check if trading data is requested in frontend, before streaming the data out
		if true {
			// first lets try to run this not in a separate go routine
			for {
				quote := <-fetcher.QuoteChan
				fmt.Println("Data from channel", quote.Symbol)
				err := connection.WriteMessage(msgType, []byte(quote.Symbol))
				if err != nil {
					log.Println("Error sending Quote data to client via socket", err)
				}
			}
		}

	}
}
