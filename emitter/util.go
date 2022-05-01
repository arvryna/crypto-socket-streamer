package emitter

import (
	"fmt"

	"github.com/arvryna/gridbot/fetcher"
)

func drainMarketData() {
	for {
		quote := <-fetcher.QuoteChan
		fmt.Println("Data from channel", quote.Symbol)
	}
}
