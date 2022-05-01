package main

import (
	"time"

	"github.com/arvryna/gridbot/emitter"
	"github.com/arvryna/gridbot/fetcher"
)

// Fetcher
func main() {
	emitter.Init()
	fetcher.Init()
	time.Sleep(100000 * time.Second)
}
