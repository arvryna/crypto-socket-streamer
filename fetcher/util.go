package fetcher

import (
	"strconv"
	"time"
)

func chanTest() {
	for i := 1; i < 100; i++ {
		time.Sleep(1 * time.Second)
		QuoteChan <- Quote{Symbol: strconv.Itoa(i)}
	}
}
