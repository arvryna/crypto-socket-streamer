
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

socket.onmessage = function(event){
    const newTrade = document.createElement('div')
    newTrade.className = 'newTrades'
    newTrade.innerHTML = `<b>data</b>${event}`
    trades.appendChild(newTrade)

    var elements = document.getElementsByClassName('newTrades')
    if(elements.length > 20){
        trades.removeChild(elements[0])
    }

}
