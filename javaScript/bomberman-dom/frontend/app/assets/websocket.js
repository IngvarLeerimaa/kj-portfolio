import {handleIncomingMessage} from "./gameData.js";

class WebSocketConnection {
    constructor() {
      this.ws = null;
    }
    connect(name) {
      this.ws = new WebSocket(`ws://localhost:3000/ws?name=${name}`);
      this.ws.onopen = () => {
        console.log("WebSocket connection established");
      };
      this.ws.onmessage = handleIncomingMessage;
      this.ws.onclose = (event) => {
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