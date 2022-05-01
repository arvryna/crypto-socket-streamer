
//https://alpaca.markets/docs/api-references/market-data-api/crypto-pricing-data/realtime/
// const url = "wss://stream.data.alpaca.markets/v1beta1/crypto";

const url = "ws://localhost:8080/wss";

const socket = new WebSocket(url);

console.log(socket)

// Setting up UI elements:
const quotes = document.getElementById('quotes')
const trades = document.getElementById('trades')

socket.onopen = () => {
    console.log("Client connected with socket...")
    socket.send("Ping")
}

socket.onclose = () => {
    console.log("Socket closed with socket...")
}

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
    }


}
