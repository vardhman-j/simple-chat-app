// api/index.js
var socket = new WebSocket("ws://localhost:8080/ws");

let connect = cb => {
  console.log("connecting");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    console.log(msg);
    if (cb) cb(msg);
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = msg => {
  console.log("sending msg: ", msg);
  if (socket.readyState === WebSocket.OPEN) {
    socket.send(msg);
  } else {
    console.log("Socket is not open yet");
  }
};

export { connect, sendMsg };