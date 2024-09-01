import { createMothership, enemyDestroyed, mothershipDestroyed, scoreMultiplier, windowFocused } from "./enemy.js";
import { boxBCR, gameDiv, gameKeys, gameLost, gameOver, gamePaused, gameRunning } from "./index.js";

const scoreDiv = document.querySelector(".score");
const timeDiv = document.querySelector(".time");
let lives;
let score;
export let sec;
let min;
const ship = document.createElement('img')
let shipX, shipY;
export let bulletExists = false;
let bulletCount = 0;

export function createShip() {
    shipX = boxBCR.width/2 - 25
    shipY = boxBCR.height - 75
    ship.src = '../static/img/ship.png';
    ship.setAttribute('class', 'ship');
    ship.width = 50;
    ship.style.transform = `translate(${shipX}px, ${shipY}px)`;
    gameDiv.appendChild(ship);
}

export function moveShip() {
    if (gameRunning && !gamePaused) {
        if (gameKeys.ArrowLeft && shipX >= 2) shipX -= 5;
        if (gameKeys.ArrowRight && shipX < 900 - 52) shipX += 5;
    }
    ship.style.transform = `translate(${shipX}px, ${shipY}px)`
}

const bullet = document.createElement('div');

bullet.setAttribute('class', 'bullet');

let bulletX, bulletY;

export function fireBullet() {
    bulletExists = true;
    bulletX = shipX + 24;
    bulletY = shipY; 
    bullet.style.transform = `translate(${bulletX}px, ${bulletY}px)`
    gameDiv.appendChild(bullet);  

    bulletCount++;
    if (bulletCount === 13) {
        createMothership();
        bulletCount = 0;
    }
    
    
    moveBullet();
}

export function moveBullet() {
    if (gameOver) {
        bullet.remove();
        bulletExists = false;
        return;
    }
    if (gameRunning && !gamePaused) {
        const bulletBCR = bullet.getBoundingClientRect();
    if (bulletBCR.top < boxBCR.top || enemyDestroyed(bulletBCR) || mothershipDestroyed(bulletBCR)) {
        bullet.remove();
        bulletExists = false;
        return;
    }
    bulletY -= 10;
    bullet.style.transform = `translate(${bulletX}px, ${bulletY}px)`
    requestAnimationFrame(moveBullet)
    }
    
}


export function addLives() {
    const livesDiv = document.querySelector(".lives");
    lives = 3;
    let x = 0;
    for (let i = 0; i < lives; i++) {
        const life = document.createElement('img');
        life.src = '../static/img/life.png';
        life.setAttribute('id', `life-${i}`);
        life.style.transform = `translateX(${x}px)`
        life.onload = ()=> {
            livesDiv.appendChild(life);
        }
        x += 30
    }
    
}

export function hitShip(bulletBCR) {
    const shipBCR = document.querySelector('.ship').getBoundingClientRect();
    if (shipBCR.top <= bulletBCR.bottom && shipBCR.left <= bulletBCR.left
        && shipBCR.right >= bulletBCR.right) {
        lives--
        lives >= 0 ? document.getElementById(`life-${lives}`).remove() : gameLost();
        return true
    }
}

export function initTimeAndScore() {
    bulletCount = 0;
    score = 0;
    scoreDiv.innerHTML = `Score:${score}`;
    sec = 0;
    min = 0;
    timeDiv.innerHTML = `Time:${min.toString().padStart(2,"0")}:${sec.toString().padStart(2,"0")}`
}

// The 'id' parameter has been added to addScore so that the score can be calculated based on the enemy's ID.
// The scoring rules are defined based on the enemy's position and ID.
// For example, bottom 1-2 rows are worth 10 points, middle 1-2 rows are worth 20 points,
// top row is worth 30 points, and the mothership is worth 300 points.

export function addScore(isMothership, id) {
    if (isMothership === true) {
      score += 300;
      scoreDiv.innerHTML = `Score: ${score}`;
    } else if (isMothership === false) {
      if (id < 8) {
        score += 30 * scoreMultiplier;
      } else if (id < 16 && id >= 8) {
        score += 20 * scoreMultiplier;
      } else {
        score += 10 * scoreMultiplier;
      }
      scoreDiv.innerHTML = `Score: ${score}`;
    }
  }
  
export function addTime(){
    if (gameRunning && !gamePaused && windowFocused) {
        sec++;
        if (sec % 60 === 0) {
            sec = 0;
            min++;
        }
        timeDiv.innerHTML = `Time:${min.toString().padStart(2,"0")}:${sec.toString().padStart(2,"0")}`;
    }
}