import simple from "../../framework/framework.js";
import { gridData, playerLocation, playerNumber, plantedBombs, gameGrid, cellSize } from "./gameData.js";
import { explosion, gridArray, loserElement, powerUp, powerUps } from "./visualElements.js";
import WSConnection from "./websocket.js";
import { isChatActive } from "./chat.js";
import { gameContainer } from "../index.js";

export var playerLives = 3;
export var players;
export var imgPaths = ["front.png", "front.png", "front.png", "front.png"]
export var keysDown = {};

var yourPlayer;
var chatActive = false;
var speed = 1, bombs = 1, explRange = 1, distance;
var playerX = 0, playerY = 0;
var bombTime = 0

export const downListener = new simple.Listen(document, "keydown", (e) => {
    keysDown[e.key] = true;
})

export const upListener = new simple.Listen(document, "keyup", (e) => {
    delete keysDown[e.key];
})

export const initGame = () => {
    players = [];
    players.push(document.getElementById("player-1"));
    players.push(document.getElementById("player-2"));
    const playerThree = document.getElementById("player-3");
    if (playerThree != null) players.push(playerThree);
    const playerFour = document.getElementById("player-4");
    if (playerFour != null) players.push(playerFour);

    yourPlayer = players[playerNumber - 1];
    distance = cellSize / 40;
    const chatInput = document.querySelector(".input-text");
    new simple.Listen(chatInput, "focus", () => {
        chatActive = true;
    }).listen();
    new simple.Listen(chatInput, "blur", () => {
        chatActive = false;
    }).listen();
}

var lCount = 0, rCount = 0, uCount = 0, dCount = 0;
var li = 1, ri = 1, ui = 1, di = 1; 
const movePlayer = () => {
    if (!chatActive) {
        if (keysDown["ArrowLeft"]) {    
            playerX -= move("left");
            lCount += speed;
            if (lCount > 10) {
                lCount = 0;
                li <= 7 ? li++ : li = 1;
            }
            const path = "left/" + li + ".png";
            if (yourPlayer != null) {
                yourPlayer.setAttribute("src", `./app/sprites/player/${playerNumber}/` + path)
            }
            sendLocation(path);
        }
        if (keysDown["ArrowRight"]) {
            playerX += move("right");
            rCount += speed;
            if (rCount > 10) {
                rCount = 0;
                ri <= 7 ? ri++ : ri = 1;
            }
            const path = "right/" + ri + ".png";
            if (yourPlayer != null) {
                yourPlayer.setAttribute("src", `./app/sprites/player/${playerNumber}/` + path)
            }
            sendLocation(path);
        } 
        if (keysDown["ArrowUp"]) {
            playerY -= move("up");
            uCount += speed;
            if (uCount > 10) {
                uCount = 0;
                ui <= 7 ? ui++ : ui = 1;
            }
            const path = "up/" + ui + ".png";
            if (yourPlayer != null) {
                yourPlayer.setAttribute("src", `./app/sprites/player/${playerNumber}/` + path)
            }
            sendLocation(path);
        } 
        if (keysDown["ArrowDown"]) {
            playerY += move("down");
            dCount += speed;
            if (dCount > 10) {
                dCount = 0;
                di <= 7 ? di++ : di = 1;   
            }
            const path = "down/" + di + ".png";
            if (yourPlayer != null) {
                yourPlayer.setAttribute("src", `./app/sprites/player/${playerNumber}/` + path)
            }
            sendLocation(path);
        }
        
        if (keysDown[" "] && bombs > 0 && !isChatActive) {
            const time = Date.now() 
            if (time - bombTime >= 300) {
                bomb();
                bombTime = time;
            }
        } 
    }

    if (yourPlayer != null) yourPlayer.style.transform = `translate(${distance * playerX}px, ${distance * playerY}px)`;
}

const move = (direction) => {
    const cRect = document.querySelector(".gameplay-container").getBoundingClientRect();
    const pRect = yourPlayer.getBoundingClientRect();

    switch (direction) {
        case "left":
            if (cRect.left > pRect.left - distance * speed) return parseInt((pRect.left - cRect.left) / distance)

            for (let i = 0; i < gridArray.length; i++) {
                if (gridArray[i] != "" && gridArray[i] != undefined) {
                    const elRect = gridArray[i].getBoundingClientRect();
                    const elClass = gridArray[i].getAttribute("class").replace(/game-cell /g, "");

                    if (elClass == "bomb" && ((elRect.left < pRect.left && elRect.right > pRect.left) ||
                     (elRect.left < pRect.right && elRect.right > pRect.right)) &&
                    elRect.top < pRect.top && elRect.bottom > pRect.bottom) {
                        if (elRect.left > pRect.left - distance * speed && gridArray[i - 1] != "" && i != 1 && i != 253) return 0;
                        return speed
                    };

                    if ((pRect.top >= elRect.top && pRect.top <= elRect.bottom)
                    || (pRect.bottom >= elRect.top && pRect.bottom <= elRect.bottom)) {
                        if (elRect.right >= pRect.left - distance * speed && elRect.right < pRect.right) {
                            if (elClass === "addbomb" || elClass === "addflame" || elClass === "addspeed") {
                                addPowerUp(elClass, i)
                                return speed;
                            }
                            return parseInt((pRect.left - elRect.right) / distance)
                        } 
                    }
                } 
            };
            return speed;
        case "right":
            if (cRect.right <= pRect.right + distance * speed) return parseInt((cRect.right - pRect.right) / distance);

            for (let i = 0; i < gridArray.length; i++) {
                if (gridArray[i] != "" && gridArray[i] != undefined) {
                    const elRect = gridArray[i].getBoundingClientRect();
                    const elClass = gridArray[i].getAttribute("class").replace(/game-cell /g, "");

                    if (elClass == "bomb" && ((elRect.left < pRect.left && elRect.right > pRect.left) ||
                    (elRect.left < pRect.right && elRect.right > pRect.right)) &&
                        elRect.top < pRect.top && elRect.bottom > pRect.bottom) {
                            if (elRect.right < pRect.right + distance * speed && gridArray[i + 1] != "" && i != 19 && i != 271) return 0;
                            return speed
                        };

                    if ((pRect.top >= elRect.top && pRect.top <= elRect.bottom)
                    || (pRect.bottom >= elRect.top && pRect.bottom <= elRect.bottom)) {
                        if (elRect.left <= pRect.right + distance * speed && elRect.left > pRect.left) {
                            if (elClass === "addbomb" || elClass === "addflame" || elClass === "addspeed") {
                                addPowerUp(elClass, i)
                                return speed;
                            }
                            return parseInt((elRect.left - pRect.right) / distance)  
                        } 
                    }
                } 
            };
            return speed;
        case "up":
            if (cRect.top > pRect.top - distance * speed) return parseInt((pRect.top - cRect.top) / distance); 
            
            for (let i = 0; i < gridArray.length; i++) {
                if (gridArray[i] != "" && gridArray[i] != undefined) {
                    const elRect = gridArray[i].getBoundingClientRect();
                    const elClass = gridArray[i].getAttribute("class").replace(/game-cell /g, "");

                    if (elClass == "bomb" && elRect.left < pRect.left && elRect.right > pRect.right &&
                        ((elRect.top < pRect.bottom && elRect.bottom > pRect.bottom) ||
                        (elRect.top < pRect.top && elRect.bottom > pRect.top))) {
                            if (elRect.top > pRect.top - distance * speed && gridArray[i - 21] != "" && i != 21 && i != 41) return 0;
                            return speed
                        };

                    if ((pRect.left >= elRect.left && pRect.left <= elRect.right)
                    || (pRect.right >= elRect.left && pRect.right <= elRect.right)) {
                        if (elRect.bottom >= pRect.top - distance * speed && elRect.bottom < pRect.bottom) {
                            if (elClass === "addbomb" || elClass === "addflame" || elClass === "addspeed") {
                                addPowerUp(elClass, i)
                                return speed;
                            }
                            return parseInt((pRect.top - elRect.bottom) / distance)
                        } 
                    }
                } 
            };
            return speed;
        case "down":
            if (cRect.bottom < pRect.bottom + distance * speed) return parseInt((cRect.bottom - pRect.bottom) / distance);

            for (let i = 0; i < gridArray.length; i++) {
                if (gridArray[i] != "" && gridArray[i] != undefined) {
                    const elRect = gridArray[i].getBoundingClientRect();
                    const elClass = gridArray[i].getAttribute("class").replace(/game-cell /g, "");

                    if (elClass == "bomb" && elRect.left < pRect.left && elRect.right > pRect.right &&
                        ((elRect.top < pRect.bottom && elRect.bottom > pRect.bottom) ||
                        (elRect.top < pRect.top && elRect.bottom > pRect.top))) {
                            if (elRect.bottom < pRect.bottom + distance * speed && gridArray[i + 21] != "" && i != 231 && i != 251) return 0;
                            return speed
                        };

                    if ((pRect.left >= elRect.left && pRect.left <= elRect.right)
                    || (pRect.right >= elRect.left && pRect.right <= elRect.right)) {
                        if (elRect.top <= pRect.bottom + distance * speed && elRect.top > pRect.top) {
                            if (elClass === "addbomb" || elClass === "addflame" || elClass === "addspeed") {
                                addPowerUp(elClass, i)
                                return speed;
                            }
                            return parseInt((elRect.top - pRect.bottom) / distance)
                        } 
                    }
                } 
            };
            return speed;
    }
}

const sendLocation = (path = "front.png") => {
    const message = JSON.stringify({
        messageType: "location",
        x: playerX,
        y: playerY,
        message: path,
    });

   try {WSConnection.sendMessage(message)
    }catch(e){
    console.log(e)
};
}

const bomb = () => {
    bombs -= 1;
    const pRect = yourPlayer.getBoundingClientRect()
    const container = document.querySelector(".gameplay-container")
    const cRect = container.getBoundingClientRect()
    const pX = pRect.x - cRect.x + pRect.width / 2;
    const pY = pRect.y - cRect.y + pRect.height / 2;
    const currentCell = {x: Math.floor(pX / cellSize), y: Math.floor(pY / cellSize)}

    const message = JSON.stringify({
        messageType: "bomb",
        x: currentCell.x,
        y: currentCell.y,
    });

   try {WSConnection.sendMessage(message)
    }catch(e){
    console.log(e)
    };

    setTimeout(() => {
        var explCells = []
        explCells.push({type: "empty", x: currentCell.x, y: currentCell.y});
        const playersHit = playersInZone(currentCell.x, currentCell.y);
        if (playersHit.length > 0) explCells = explCells.concat(playersHit);
        const lCells = checkCell("left", explRange, currentCell.x - 1, currentCell.y);
        if (lCells) explCells = explCells.concat(lCells);
        const rCells = checkCell("right", explRange, currentCell.x + 1, currentCell.y);
        if (rCells) explCells = explCells.concat(rCells);
        const uCells = checkCell("up", explRange, currentCell.x, currentCell.y - 1);
        if (uCells) explCells = explCells.concat(uCells);
        const dCells = checkCell("down", explRange, currentCell.x, currentCell.y + 1);
        if (dCells) explCells = explCells.concat(dCells);


        const explMessage = JSON.stringify({
            messageType: "explosionCells",
            cells: explCells,
        });
    
       try {WSConnection.sendMessage(explMessage)
        }catch(e){
        console.log(e)
        };
        bombs += 1;
    }, 3000);

}

const playersInZone = (x, y) => {
    var plrs = [];
    const b = document.getElementById(`cell-${x + 21 * y}`)
        if (b == null) return plrs; 
        const bRect = b.getBoundingClientRect();
        players.forEach((p) => {
            if (p != null) {
                var inZone = false;
            const blastSize = cellSize * explRange
            const playerRect = p.getBoundingClientRect();
            if (playerRect.top <= bRect.bottom && playerRect.top >= bRect.top ||
                playerRect.bottom <= bRect.bottom && playerRect.bottom >= bRect.top) {
                if (playerRect.left <= bRect.right + blastSize && playerRect.left >= bRect.left - blastSize ||
                    playerRect.right <= bRect.right + blastSize && playerRect.right >= bRect.left - blastSize) {
                        inZone = true;
                    }
            }
            if (playerRect.left <= bRect.right && playerRect.left >= bRect.left ||
                playerRect.right <= bRect.right && playerRect.right >= bRect.left) {
                if (playerRect.top <= bRect.bottom + blastSize && playerRect.top >= bRect.top - blastSize ||
                    playerRect.bottom <= bRect.bottom + blastSize && playerRect.bottom >= bRect.top - blastSize) {
                        inZone = true;
                    }
            }
            if (inZone) plrs.push({type: "player", id: parseInt(p.getAttribute("id").replace(/player-/g, ""))});
            } 
        });
        return plrs;
}

const checkCell = (direction, count, x, y) => {
    if (x < 0 || y < 0 || x > 20 || y > 12 || count === 0) return false;
    var cells = [];
    var nextX, nextY
    const element = document.getElementById(`cell-${x + 21 * y}`)
    const cellType = element != null ? element.getAttribute("class").replace(/game-cell /g, "") : "empty";
    if (cellType === "wall") return false;
    cells.push({type: cellType, x: x, y: y})
    if (cellType === "block") return cells;
    switch (direction) {
        case "left":
            nextX = x - 1;
            nextY = y;
            break;
        case "right":
            nextX = x + 1;
            nextY = y;
            break;
        case "up":
            nextX = x;
            nextY = y - 1;
            break;
        case "down":
            nextX = x;
            nextY = y + 1;
            break;
    }
    const nextCell = checkCell(direction, count - 1, nextX, nextY);
    if (nextCell) cells = cells.concat(nextCell);
    return cells
}

export const createExplosions = (cells) => {
    cells.forEach((cell) => {
        const cID = cell.x + 21 * cell.y
        switch (cell.type) {
            case "empty": {
                const element = document.getElementById(`cell-${cID}`);
                if (element != null) element.remove();
                const pwr = document.getElementById(`powerup-${cID}`);
                if (pwr != null) pwr.remove();
                if (gridData[cID] == "block") gridData[cID] = "";
                if (cID != 0 && cID != 20 && cID != 252 && cID != 272) gridArray[cID] = "";
                delete powerUps[cID];
                delete plantedBombs[cID];
                explosion(cell.x, cell.y)
                break;
            }
            case "powerup": {
                const element = document.getElementById(`cell-${cID}`);
                if (element != null) element.remove();
                if (gridData[cID] == "block") gridData[cID] = "";
                delete plantedBombs[cID];
                explosion(cell.x, cell.y)
                gameGrid.appendChild(powerUp(cell.id, cell.x, cell.y));
                break;
            }
            case "bomb": {
                explosion(cell.x, cell.y)
                break;
            }
            case "life": {
                const lives = document.querySelectorAll(".life");
                if (lives.length > 0) {
                    lives[lives.length - 1].remove();
                    playerLives -= 1;
                } 
                if (playerLives == 0) {
                    downListener.remove();
                    upListener.remove();
                    for (const key in keysDown) {
                        keysDown[key] = false;
                    }
                    const message = JSON.stringify({
                        messageType: "dead",
                    });
                
                   try {WSConnection.sendMessage(message)
                    }catch(e){
                    console.log(e)
                };
                gameContainer.appendChild(loserElement());
                }
                break;
            }
        }
    })
}

const addPowerUp = (powerUp, id) => {
    if (powerUp === "addbomb") bombs += 1;
    if (powerUp === "addflame") explRange += 1;
    if (powerUp === "addspeed") speed += 1;
    document.getElementById(`powerup-${id}`).remove();
    gridArray[id] = "";
    delete powerUps[id];
    const message = JSON.stringify({
        messageType: "powerup",
        x: id,
    });

   try {WSConnection.sendMessage(message)
    }catch(e){
    console.log(e)
};
}

export const runGame  = () => {
    movePlayer();
    players.forEach((player, index) => {
        if (index != playerNumber - 1 && player != null) {
            player.style.transform = `translate(${distance * playerLocation[index].x}px, ${distance * playerLocation[index].y}px)`;
            player.setAttribute("src", `./app/sprites/player/${index + 1}/` + imgPaths[index])
        }
    })
    requestAnimationFrame(runGame);
}