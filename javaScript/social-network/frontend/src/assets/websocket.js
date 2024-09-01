class WebSocketConnection {
  constructor() {
    this.ws = null;
  }

  connect(userID) {
    this.ws = new WebSocket(`ws://localhost:3000/ws?id=${userID}`);

    this.ws.onopen = () => {
      console.log("WebSocket connection established");
    };
    this.ws.onclose = () => {
      console.log("WebSocket connection closed");
    };
  }

  close() {
    if (this.ws) {
      this.ws.close();
      console.log("WebSocket connection closed");
    }
  }

  sendMessage(message) {
    if (this.ws.readyState === WebSocket.OPEN) {
      this.ws.send(message);
    } else {
      console.log("WebSocket connection not open.");
    }
  }
}

const WSConnection = new WebSocketConnection();

export default WSConnection;
