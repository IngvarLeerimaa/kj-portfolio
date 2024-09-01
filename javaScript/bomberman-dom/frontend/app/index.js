import simple from "../framework/framework.js";
import WSConnection from "./assets/websocket.js";
import { titleElement, logoElement } from "./assets/visualElements.js";
import { setCellSize } from "./assets/gameData.js";
import { chatElement, buttonClick } from "./assets/chat.js";


export const App = document.getElementById("app");
export const gameContainer = new simple.NewElement("div", {attrs: {class: "game-container"}}).create();

const title = titleElement();
const chat = chatElement();
const logo = logoElement();

App.appendChild(logo);
gameContainer.appendChild(title);
App.appendChild(gameContainer);
setCellSize();


window.connectWS = (event) => {
    const nickname = new FormData(event.target).get("nickname");
    // Quick fix et kõik ilusti gridi jääks
    if (nickname.length < 3 || nickname.length > 16) {
        alert("Nickname must be between 3-16 char") 
        getElementByClassName("nickname-form").reset();
        return false;
    } 
    if (WSConnection.ws === null) {WSConnection.connect(nickname)};
    gameContainer.removeChild(title);

    App.appendChild(chat);
    
    let button = document.querySelector(".send-button");
    button.addEventListener("click", buttonClick) 
    return false;
}