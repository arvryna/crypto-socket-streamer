
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
    const newQuote = document.createElement('div')
    newQuote.className = 'newQuote'
    newQuote.innerHTML = `<b>${event.data}</b>`
    quotes.appendChild(newQuote)

    var elements = document.getElementsByClassName('newQuote')
    if(elements.length > 20){
        quotes.removeChild(elements[0])
    }

}
