import simple from "../../framework/framework.js";
import { plantedBombs, gameGrid, playerData, cellSize } from "./gameData.js";
import { playerLives } from "./gameplay.js";

export const titleElement = () => {
    const title = new simple.NewElement("div", {
        attrs: {
            class: "title-container"
        },
        children: [
            [
                "form", {
                    attrs: {
                        class: "nickname-form",
                        onsubmit: "return connectWS(event)",
                    },
                    children: [
                        [
                            "input", {
                                attrs: {
                                    type: "text",
                                    name: "nickname",
                                    placeholder: "Enter nickname",
                                    required: true,
                                    autofocus: true,
                                }
                            }
                        ],
                        [
                            "button", {
                                attrs: {
                                    type: "submit",
                                },
                                children: ["Join Game"]
                            }
                        ]
                    ]
                }
            ]
        ]
    }).create();

    return title;
}


export const lobbyElement = () => {
    const players = [];

    for (let i = 0; i < 4; i++) {
        if (playerData[i] == null) {
            players.push([
                "div", {
                    attrs: {
                        class: `player-${i+1}`
                    },
                    children: [
                        [
                            "div", {
                                attrs: {
                                    class: "waiting"
                                },
                                children: [`Waiting for player ${i+1}`, ["div", { attrs: {class: "count"}}]]
                            }
                        ]
                    ]
                }
            ])
        } else {
            players.push([
                "div", {
                    attrs: {
                        class: playerData[i].self ? `player-${i+1} self` : `player-${i+1}`
                    },
                    children: [
                        [
                            "img", {
                                attrs: {
                                    src: `./app/sprites/player/${i+1}/front.png`
                                }
                            }
                        ],
                        `Player ${i+1}:`,
                        ["br"],
                        `${playerData[i].name}`,
                    ]
                }
            ])
        }
    }

    const lobby = new simple.NewElement("div", {
        attrs: {
            class: "lobby-container"
        },
        children: [
            [
                "div", {
                    attrs: {
                        class: "player-grid"
                    },
                    children: players
                }
            ],
            ["div", { attrs: {class: "ready", hidden: true}}]
        ]
    }).create();
    return lobby;
}

export const gridArray = []

export const gridElement = (grid) => {
    const gameContainer = new simple.NewElement("div", {attrs: {class: "gameplay-container"}}).create();
    const rows = 13;
    const cols = 21;
    var cellX = 0;
    var cellY = 0;

    for (let i = 0; i < playerLives; i++) {
        const life = new simple.NewElement("img", {attrs: {
            class: "life", id: `life-${i + 1}`, src: "./app/sprites/life.png"
            }}).create();
        gameContainer.appendChild(life);
        life.style.left = cellSize * i - cellSize + cellSize * 0.14 + "px";
        life.style.top = -cellSize + "px";

    }

    for (const key in plantedBombs) {
        gameContainer.appendChild(bomb(plantedBombs[key].x, plantedBombs[key].y));
    }

    for (const key in powerUps) {
        const p = powerUps[key];
        gameContainer.appendChild(powerUp(p.id, p.x, p.y));
    }


    for (let i = 0; i < rows; i++) {
        cellX = 0;
        for (let j = 0; j < cols; j++) {
            switch (grid[i * cols + j]) {
                case "wall":
                    const wall = new simple.NewElement("img", {attrs: {
                        class: "game-cell wall", id: `cell-${i * cols + j}`, src: "./app/sprites/wall.png"
                        }}).create();
                    gameContainer.appendChild(wall);
                    wall.style.left = cellX + "px";
                    wall.style.top = cellY + "px";
                    gridArray.push(wall);
                    break;
                case "block":
                    const block = new simple.NewElement("img", {attrs: {
                        class: "game-cell block", id: `cell-${i * cols + j}`, src: "./app/sprites/block.png"
                        }}).create();
                    gameContainer.appendChild(block);
                    block.style.left = cellX + "px";
                    block.style.top = cellY + "px";
                    gridArray.push(block)
                    break;
                case "player": 
                var p
                if (i == 0 && j == 0) p = 1;
                if (i == rows -1 && j == cols - 1) p = 2;
                if (i == 0 && j == cols - 1) p = 3;
                if (i == rows - 1 && j == 0) p = 4;

                const player = new simple.NewElement("img", {attrs: {
                    class: "game-cell player", id: `player-${p}`, src: `./app/sprites/player/${p}/front.png`
                    }}).create();

                gameContainer.appendChild(player);
                player.style.left = (cellX + cellSize * 0.3) + "px";
                player.style.top = (cellY + cellSize * 0.1) + "px";
                gridArray.push(player);
                break;
                default: 
                gridArray.push("");
                break;
            }
            cellX += cellSize;
        }
        cellY += cellSize;
    }
    return gameContainer
}

export const bomb = (x, y) => {
    const cellId = x + 21 * y;
    const bomb = new simple.NewElement("img", {attrs: {
        class: "game-cell bomb", id: `cell-${cellId}`, src: "./app/sprites/bomb.png"
        }}).create();
    bomb.style.left = x * cellSize + "px";
    bomb.style.top = y * cellSize + "px";
    gridArray[cellId] = bomb;
    return bomb;
}

export const explosion = (x, y, i = 0) => {
    const img = new simple.NewElement("img", {
        attrs: {
            src: `./app/sprites/explosion/${i}.png`,
            class: "explosion"
        }
    }).create();
    img.style.left = x * cellSize + "px";
    img.style.top = y * cellSize + "px";
    gameGrid.appendChild(img);
    setTimeout(() => {
        gameGrid.removeChild(img);
        i < 6 ? explosion(x, y, i + 1) : null;
    }, 130)
}

export var powerUps = {}

export const powerUp = (id, x, y) => {
    const cellId = x + 21 * y;
    const pwUps = ["addbomb", "addflame", "addspeed"]
    const pUp = new simple.NewElement("img", {
        attrs: {
            src: `./app/sprites/powerup/${id}.png`,
            class: `game-cell ${pwUps[id]}`,
            id: `powerup-${cellId}`
        }
    }).create();
    pUp.style.left = x * cellSize + "px";
    pUp.style.top = y * cellSize + "px";
    gridArray[cellId] = pUp;
    powerUps[cellId] = {id: id,x: x,y: y};
    return pUp;
}

export const winnerElement = (id, name) => {
    const wContainer = new simple.NewElement("div", {
        attrs: {
            class: "winner-container"
        }, 
        children: [
            [
                "div", {
                    attrs: {
                        class: "winner"
                    },
                    children: [
                        `The winner is player ${id + 1}, ${name}`
                    ]   
                }
            ],
            [
                "button", {
                    attrs: {
                        onclick: "location.reload()"
                    },
                    children: ["Play again"]
                }
            ]
        ]
    }).create();

    return wContainer
}

export const logoElement = () => {
    const logo = new simple.NewElement("img", {
        attrs: {
            src: "./app/cosmetics/logo.png",
            class: "logo"
        }
    }).create();
    return logo;
}

export const loserElement = () => {
    const wContainer = new simple.NewElement("div", {
        attrs: {
            class: "winner-container"
        }, 
        children: [
            [
                "div", {
                    attrs: {
                        class: "winner",
                        id: "loser"
                    },
                    children: [
                        `You lost! Wait for results or play again!`
                    ]
                }
            ], 
            [
                "button", {
                    attrs: {
                        onclick: "location.reload()"
                    },
                    children: ["Play again"]
                }
            ]
        ]
    }).create();

    return wContainer
}