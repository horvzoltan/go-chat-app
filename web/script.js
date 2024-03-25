document.addEventListener("DOMContentLoaded", function () {
  var ws = new WebSocket("ws://localhost:8080/ws");
  var log = document.getElementById("log");
  var messageInput = document.getElementById("messageInput");
  var sendButton = document.getElementById("sendButton");

  ws.onopen = function (event) {
    console.log("WebSocket connection opened");
  };

  ws.onclose = function (event) {
    console.log("WebSocket connection closed");
  };

  ws.onerror = function (event) {
    console.error("WebSocket error observed:", event);
  };

  ws.onmessage = function (event) {
    console.log("onmessage called");
    log.innerHTML += event.data + "<br>";
  };

  sendButton.onclick = function () {
    var message = messageInput.value;
    if (message) {
      ws.send(JSON.stringify({ message: message }));
      messageInput.value = "";
    }
  };
});
