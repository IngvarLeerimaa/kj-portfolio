import { gameContainer } from "../index.js";
import { createExplosions, downListener, imgPaths, initGame, keysDown, players, runGame, upListener } from "./gameplay.js";
import { bomb, gridArray, gridElement, lobbyElement, powerUps, winnerElement } from "./visualElements.js";
import { BubbleMessage } from "./chat.js";

export var cellSize
export var playerData = []
export var playerLocation = [{x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}]
export var lobby = lobbyElement();
export var gameGrid;
export var plantedBombs = {};
export var playerNumber;

export var gridData;

export const handleIncomingMessage = (message) => {
  const data = JSON.parse(message.data);
  switch (data.type) {
    case "playerData": {
      gameContainer.contains(lobby) ? gameContainer.removeChild(lobby) : null;
      playerData = data.playerData;
      lobby = lobbyElement();
      gameContainer.appendChild(lobby);
      break;
    }
    case "twenty": {
      document.querySelectorAll(".count").forEach((el) => {
        el.innerHTML = data.count != -1 ? `${data.count}`: "";
      })
      break;
    }
    case "ten": {
      const ready = document.querySelector(".ready");
      if (data.count == 10) {ready.removeAttribute("hidden")}; 
      if (data.count == -1) {
        gameContainer.removeChild(lobby)
        gameContainer.appendChild(gameGrid);
        initGame();
        downListener.listen();
        upListener.listen();
        runGame();
        break;
      };
      const text = `Get ready! Game starting in ${data.count}`
      ready.innerHTML = text;
     
      break;
    }
    case "message": {
      let message = new BubbleMessage(data.from, data.message, data.colour).createMessage();
     
      console.log("message:", message)
      break;
    }
    case "prepareGame": {
      gridData = data.grid;
      gameGrid = gridElement(gridData);
      playerNumber = data.playerNumber;
      break;
    }
    case "playerLocation": {
      playerLocation[data.player].x = data.x;
      playerLocation[data.player].y = data.y;
      imgPaths[data.player] = data.path;
      break;
    }
    case "plantBomb": {
      gameGrid.appendChild(bomb(data.x, data.y));
      plantedBombs[data.x + 21 * data.y] = {x: data.x, y: data.y}
      break;
    }
    case "explosions": {
      createExplosions(data.cells)
      break;
    }
    case "rmPowerUp": {
      const p = document.getElementById(`powerup-${data.id}`)
      if (p != null) p.remove();
      gridArray[data.id] = "";
      delete powerUps[data.id];
      break;
    }
    case "rmPlayer": {
      players[data.id].remove();
      if (data.id == 0) gridData[0] = "";
      if (data.id == 1) gridData[272] = "";
      if (data.id == 2) gridData[20] = "";
      if (data.id == 3) gridData[252] = "";
      break;
    }
    case "winner": {
      downListener.remove();
      upListener.remove(); 
      for (const key in keysDown) {
        keysDown[key] = false;
      }

      const w = gameContainer.querySelector(".winner-container");
      if (w != null) w.remove();
    
      gameContainer.appendChild(winnerElement(data.id, data.name))
      break;
    }
    default: {
      console.error("invalid event type: " + data.type);
    }
  }
}

export const setCellSize = () => {
  cellSize = gameContainer.getBoundingClientRect().height / 15;
}

window.addEventListener("resize", () => {
  setCellSize();
  if (gameContainer.contains(gameGrid)) {
    gameContainer.removeChild(gameGrid);
    gameGrid = gridElement(gridData);
    gameContainer.appendChild(gameGrid);
    initGame(playerNumber)
  } else if (gridData != undefined){
    gameGrid = gridElement(gridData);
  }
})