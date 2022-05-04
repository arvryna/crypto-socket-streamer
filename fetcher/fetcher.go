package fetcher

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var (
	QuoteChan chan Quote
	TradeChan chan Trade
)

func fetchMessages(conn *websocket.Conn) {
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		var markets []marketData

		err = json.Unmarshal(data, &markets)
		if err != nil {
			fmt.Println("Parsing response failed", err)
		}

		market := markets[0]

		if market.Type == "q" {
			QuoteChan <- Quote{
				Symbol: market.Symbol,
				Ask:    market.Ask,
				Bid:    market.Bid,
			}
		}
		if market.Type == "t" {
			TradeChan <- Trade{
				Symbol: market.Symbol,
				Price:  market.Price,
				Size:   market.Size,
			}
		}
	}
}

const BUFFER = 10

func Init() {
	fmt.Println("Starting fetcher...")

	// Initiaze channels:
	QuoteChan = make(chan Quote, BUFFER)
	TradeChan = make(chan Trade, BUFFER)
	url := "wss://stream.data.alpaca.markets/v1beta1/crypto"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Can't connect to Market socket", err)
	}

	// Authenticate with Trading API
	auth := map[string]string{"action": "auth", "key": "PKYSDC0NHHAIYTKUID44", "secret": "o4an6wLWk5WKnd1FFeBcGxX1kfVBl1ZfckbJvaCq"}
	conn.WriteJSON(auth)

	// Subscribe to these events
	subscribe := `{"action": "subscribe", "trades":["ETHUSD"], "quotes":["ETHUSD"], "bars":["ETHUSD"]}`
	conn.WriteMessage(1, []byte(subscribe))
	fetchMessages(conn)
}
