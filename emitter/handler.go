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
	for {
		msgType, data, err := connection.ReadMessage()
		if err != nil {
			log.Println("Client error, terminating connection", err, connection.RemoteAddr())
			connection.Close()
			return
		}

		msg := string(data)
		fmt.Println("Received", string(msg), "from client", connection.RemoteAddr())

		// Send ping-pong in separate go routine as health checks
		if msg == "Ping" {
			err = connection.WriteMessage(msgType, []byte("Pong"))
			if err != nil {
				log.Println("Writing failed to client", err)
			}
		}

		// you need to check if trading data is requested in frontend, before streaming the data out
		// implement authentication
		// first lets try to run this not in a separate go routine
		for {
			select {
			case quote := <-fetcher.QuoteChan:
				quoteString := fmt.Sprintf("%v:[ASK]%v:[BID]%v", quote.Symbol, quote.Ask, quote.Bid)
				rawJSON := fmt.Sprintf(`{"type": 1, "payload": "%s"}`, quoteString)
				err := connection.WriteMessage(msgType, []byte(rawJSON))
				if err != nil {
					log.Println("Error sending data to connection", err, connection.RemoteAddr(), "Closing connection")
					connection.Close()
					return
				}
			case trade := <-fetcher.TradeChan:
				tradeString := fmt.Sprintf("%v:[Price]%v:[Size]%v", trade.Symbol, trade.Price, trade.Size)
				rawJSON := fmt.Sprintf(`{"type": 2, "payload": "%s"}`, tradeString)
				err := connection.WriteMessage(msgType, []byte(rawJSON))
				if err != nil {
					log.Println("Error sending data to connection", err, connection.RemoteAddr(), "Closing connection")
					connection.Close()
					return
				}
			}
		}
	}
}
