const player = require('./player')

var mainDiv = document.getElementById("mainDiv")
mainDiv.addEventListener("mousedown", handleMouseDown);

console.log("index.js");
console.log(mainDiv);

var input = document.getElementById("input");
var output = document.getElementById("output");
var socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = function () {
    output.innerHTML += "Status: Connected\n";
    
};

socket.onmessage = function (e) {
    console.log(e.data);
    const sentSound = JSON.parse(e.data);
    //console.log(JSON.parse(e.data));
    player.testMusic(sentSound, socket);
};

function handleMouseDown(event) {
    console.log("handleMouseDown");
    console.log(event);
    const vel = event.clientY;
    const freq = event.clientX;
    send(vel, freq)
} 

function send(vel, freq) {
    const message = {
        Type: "START_SOUND",
        Sound: {
            Vel: vel,
            Freq: freq,
            Length: 30,
        }
    }

    socket.send(JSON.stringify(message));
}