document.addEventListener("DOMContentLoaded", function () {
    var ws = new WebSocket("ws://localhost:8080/ws");
    ws.onmessage = function (event) {
        console.log('onmessage called');
        var log = document.getElementById("log");
        log.innerHTML += event.data + "<br>";
    };

    ws.onopen = function (event) {
        ws.send(JSON.stringify({ message: "Hello Server!" }));
        console.log('WebSocket connection opened');
    };

    ws.onclose = function (event) {
        console.log('WebSocket connection closed');
    };

    ws.onerror = function (event) {
        console.error('WebSocket error observed:', event);
    };
});
