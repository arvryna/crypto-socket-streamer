
//https://alpaca.markets/docs/api-references/market-data-api/crypto-pricing-data/realtime/
// const url = "wss://stream.data.alpaca.markets/v1beta1/crypto";

const url = "ws://localhost:8080/wss";

const socket = new WebSocket(url);

console.log(socket)

// Setting up UI elements:
const quotes = document.getElementById('quotes')
const trades = document.getElementById('trades')

const bars = [];

// Setting up graph
// https://jsfiddle.net/TradingView/yozeu6k1/
var chart = LightweightCharts.createChart(document.getElementById('chart'), {
    width: 600,
    height: 700,
    layout:{
        backgroundColor: '#000000',
        textColor: '#ffffff',
    },
    crosshair: {
        mode: LightweightCharts.CrosshairMode.Normal,
    },
    grid: {
		vertLines: {
			color: '#404040',
		},
		horzLines: {
			color: '#404040',
		},
	},
    priceScale: {
		borderColor: '#cccccc',
	},
	timeScale: {
		borderColor: '#cccccc',
		timeVisible: true,
	},
});

var candleSeries = chart.addCandlestickSeries();

var data = [
    { time: '2022-02-19', open: 54.62, high: 55.50, low: 54.52, close: 54.90 },
    { time: '2022-02-22', open: 55.08, high: 55.27, low: 54.61, close: 54.98 },
    { time: '2022-02-23', open: 56.09, high: 57.47, low: 56.09, close: 57.21 },
]

candleSeries.setData(data);

const currentBar = data[data.length - 1]


socket.onopen = () => {
    console.log("Client connected with socket...")
    socket.send("Ping")
}

socket.onclose = () => {
    console.log("Socket closed with socket...")
}

const ltime = 1645574400

socket.onmessage = function (event) {
    const LIMIT = 20
    const marketData = JSON.parse(event.data)

    if (marketData.type == 1) {
        // Append Quote data (Recycle after certain LIMIT)
        const newQuote = document.createElement('div')
        newQuote.className = 'newQuote'
        newQuote.innerHTML = `<b>${marketData.payload}</b>`
        quotes.appendChild(newQuote)

        var elements = document.getElementsByClassName('newQuote')
        if (elements.length > LIMIT) {
            quotes.removeChild(elements[0])
        }
    }

    if (marketData.type == 2) {
        // Append Trade data (Recycle after certain LIMIT)
        const newTrade = document.createElement('div')
        newTrade.className = 'newTrade'
        newTrade.innerHTML = `<b>${marketData.payload}</b>`
        trades.appendChild(newTrade)

        var elements = document.getElementsByClassName('newTrade')
        if (elements.length > LIMIT) {
            trades.removeChild(elements[0])
        }

        bars.push(marketData.p)

        //this is to update the last bar and showing activity
        var open = bars[0];
        var closed = bars[bars.length - 1];
        var high = Math.max(...bars)
        var low = Math.min(...bars)

        candleSeries.update({
            time: ltime +  60,
            open: open,
            high: high,
            low: low,
            close: closed
        })
    }


}
