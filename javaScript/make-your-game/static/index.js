import { createEnemies, mothershipRestart, moveEnemies, moveMothership, selectBulletsAndMove} from "./enemy.js";
import { addLives, addTime, bulletExists, createShip, fireBullet, moveBullet, moveShip, initTimeAndScore } from "./ship.js";

export let boxBCR = document.querySelector(".box").getBoundingClientRect()
export const gameDiv = document.querySelector(".game");

const titleDiv = document.querySelector(".title");
const resumeBtn = document.querySelector("#resume");
const restartBtn = document.querySelector("#restart");
const tryAgainBtn = document.querySelector("#tryAgain");
let gamePausedByChecker = false;

//resumeBtn and restartBtn are buttons on the pause screen
resumeBtn.addEventListener('click', () => {
    pauseScreen.close();
    gamePaused = false;
    runGame();
    moveBullet();
})

restartBtn.addEventListener('click', () => {
    pauseScreen.close();
    gamePaused = false;
    restartGame();
    runGame();
    moveBullet();
})

//tryAgainBtn is a button on the game over screen
tryAgainBtn.addEventListener('click', () => {
    gameOverScreen.close();
    gameRunning = true;
    gameOver = false;
    restartGame();
    runGame();
    moveBullet();
})

window.addEventListener('resize', () => {
    boxBCR = document.querySelector(".box").getBoundingClientRect();
    checkScreen();
  });
  
function checkScreen(){
    if (tooSmallScreen() && gameRunning && !gamePaused && !gameOver) {
        gamePausedByChecker = true;
        gamePaused = true;
        isSmallScreen.show();
      } else if (!tooSmallScreen() && gameRunning && gamePaused && gamePausedByChecker) {
        gamePaused = false;
        isSmallScreen.close();
        runGame();
        moveBullet();
        gamePausedByChecker = false;
      }
      return;
}

export let gameRunning = false; 
export let gamePaused = false;
export let gameOver = false;
export let gameKeys = {
    ArrowLeft: false,
    ArrowRight: false,
    Space: false,
}

window.addEventListener('load', () => {
    createShip();
    createEnemies(32);
})

document.addEventListener("keydown", (e) => {
    if (e.code === "ArrowLeft") gameKeys['ArrowLeft'] = true;
    if (e.code === "ArrowRight") gameKeys['ArrowRight'] = true;
    
    if (e.code === 'Space') {
        if (gameRunning && !gamePaused && !bulletExists) {
            gameKeys['Space'] = true;
            checkScreen();
            
        }
        
        if (!gameRunning && !gamePaused && !gameOver) {
            titleDiv.remove();
            gameDiv.removeAttribute('hidden');
            gameRunning = true;
        }
    }
    if (e.code === 'Enter') {
        if (gameOver) {
            gameRunning = true;
            gameOver = false;
        }
    }
    if (e.code === 'Escape') {
        if (gameRunning && !gamePaused) {
            
            pauseScreen.show();
            gamePaused = true;
        } else if (gamePaused) {
            pauseScreen.close();
            gamePaused = false;
            runGame();
            moveBullet();
        } 
    }    
})

//
document.addEventListener('keyup', (e) => {
    if (e.code === 'ArrowLeft') gameKeys['ArrowLeft'] = false;
    if (e.code === 'ArrowRight') gameKeys['ArrowRight'] = false;
    if (e.code === 'Space') gameKeys['Space'] = false;
})

//runGame is the main game loop that handles movement and collision detection
function runGame() {

    if (!gamePaused && !gameOver) {
      moveEnemies();
      moveShip();
      moveMothership();
      selectBulletsAndMove();
      if (gameKeys.Space && !bulletExists) {
        fireBullet();
      }
      requestAnimationFrame(runGame);
    }
  }

//gameLost handles the game over screen
export function gameLost() {
    gameRunning = false;
    gameOver = true;
    gameOverScreen.show();
}

//restartGame handles the restart button on the pause screen
function restartGame() {
    let s = document.querySelector('.ship')
    s !== null ? s.remove() : null;
    document.querySelector('.enemies').innerHTML = '';
    document.querySelector('.lives').innerHTML = '';
    mothershipRestart();
    document.querySelector('.enemy-bullet') !== null ? document.querySelector('.enemy-bullet').remove() : null;
    document.querySelector('.bullet') !== null ? document.querySelector('.bullet').remove() : null;
    initTimeAndScore();
    addLives();
    addTime();
    createShip();
    createEnemies(32);
}

//isSmallScreen is used to determine if the game is being played on a too small screen
function tooSmallScreen() {
    return window.innerWidth < 900 || window.innerHeight < 600;
}

//addLives is called on load and restart
addLives();

//initTimeAndScore is called on load and restart
initTimeAndScore();

//setInterval is used to update the time every second
setInterval(addTime, 1000)

//runGame is called on load and restart
runGame();
