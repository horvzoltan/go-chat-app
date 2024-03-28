// document.addEventListener("DOMContentLoaded", function () {
//   var ws = new WebSocket("ws://localhost:8080/ws");
//   var log = document.getElementById("log");
//   var messageInput = document.getElementById("messageInput");
//   var sendButton = document.getElementById("sendButton");

//   ws.onopen = function (event) {
//     console.log("WebSocket connection opened");
//   };

//   ws.onclose = function (event) {
//     console.log("WebSocket connection closed");
//   };

//   ws.onerror = function (event) {
//     console.error("WebSocket error observed:", event);
//   };

//   ws.onmessage = function (event) {
//     console.log("onmessage called");
//     log.innerHTML += event.data + "<br>";
//   };

//   sendButton.onclick = function () {
//     var message = messageInput.value;
//     if (message) {
//       ws.send(JSON.stringify({ message: message }));
//       messageInput.value = "";
//     }
//   };
// });
document.getElementById("registerForm").onsubmit = async (event) => {
  event.preventDefault();

  const formData = new FormData(event.target);
  const data = Object.fromEntries(formData.entries());

  const response = await fetch("/api/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (response.ok) {
    console.log("Registration successful");
    // Redirect or update UI accordingly
  } else {
    console.error("Registration failed");
    // Handle error
  }
};
