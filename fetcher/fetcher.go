package fetcher

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var QuoteChan chan Quote
var TradeChan chan Trade

type Quote struct {
	Symbol string `json:"S"`
}

type Trade struct {
	Symbol string `json:"S"`
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
	fmt.Println("Fetching markets..")
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
			}
		}
		// if market.Type == "t" {
		// 	TradeChan <- Trade{
		// 		Symbol: market.Symbol,
		// 	}
		// }
	}
}

const BUFFER = 10

func Init() {
	// Initiaze channels:
	QuoteChan = make(chan Quote, BUFFER)
	TradeChan = make(chan Trade, BUFFER)
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
	// fetchMessages(conn)
	chanTest()
}

func chanTest() {
	for i := 1; i < 100; i++ {
		time.Sleep(1 * time.Second)
		QuoteChan <- Quote{Symbol: strconv.Itoa(i)}
	}
}
