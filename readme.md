Over the weekend, I built a side project, a web socket streaming server (in Golang) for real time crypto data using https://alpaca.markets API. I could not use their real API as they ask too much info, but i consumed their simulated paper API (chart looks misleading though). It was fun & interesting to build it.

# Flow of data
->> fetcher(FAN-IN) -> emitter(FAN-OUT) ->> client

# NOTE:
-should not be used as metric to see the code quality / consumption . this is just an experimental work and built for fun to push a bunch of live trade data in web socket, just watch the flow and enjoy : )

![screen](https://github.com/arvryna/crypto-socket-streamer/blob/main/docs/screen.png)
