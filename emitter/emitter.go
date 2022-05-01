package emitter

import (
	"fmt"
	"net/http"
)

/*
	Emitter:
	- It is responsibile for holding client connections and pushing data to all clients in FAN-OUT pattern
	- It also can handle retrying socket conns or if required killing existing inactive or un-recoverable connections
	- It takes data from fetcher and passes it to client
*/

func setupRoutes() {
	/* Get
	* Root endpoint */
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Root
	})

	/* GET /ping
	 * Health check */
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	/* GET /wss
	 * setup websocket connection, http -> websocket upgrade */
	http.HandleFunc("/wss", handleWsConn)
	http.ListenAndServe(":8080", nil)
}

func Init() {
	fmt.Println("Starting emitter...")
	go setupRoutes()
}
