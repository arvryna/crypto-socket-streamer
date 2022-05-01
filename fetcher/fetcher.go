package fetcher

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var quoteChan chan string
var tradeChan chan string

type Quote struct {
}

type Trade struct {
}

type marketData struct {
	Type   string  `json:"T"`
	Symbol string  `json:"S"`
	Price  float64 `json:"p"`
	Size   float64 `json:"s"`
	Ask    float64 `json:"ap"`
	Bid    float64 `json:"bp"`
	Time   string  `json:"t"`
}

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
			fmt.Println("Fetching Quote", market.Symbol, market.Type)
		}
		if market.Type == "t" {
			fmt.Println("Fetching Trade", market.Symbol, market.Type)
		}
	}
}

func Init() {
	url := "wss://stream.data.alpaca.markets/v1beta1/crypto"

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Can't connect to Market socket", err)
	}

	// Authenticate with Trading API
	auth := map[string]string{"action": "auth", "key": "PK4WJX5BHJZBRYK4DGJQ", "secret": "yjdiafw56zgxVyR48LHrXCuCoSBTK7gAe5Vbrdv1"}
	conn.WriteJSON(auth)

	// Subscribe to these events
	subscribe := `{"action": "subscribe", "trades":["ETHUSD"], "quotes":["ETHUSD"], "bars":["ETHUSD"]}`
	conn.WriteMessage(1, []byte(subscribe))
	fetchMessages(conn)
}
