import { boxBCR, gameDiv, gameOver, gamePaused, gameRunning, gameLost} from "./index.js";
import { addScore, hitShip } from "./ship.js";
const enemyDiv = document.querySelector(".enemies");
let enemyDirection, enemyX, enemyY, enemyBulletFire, mothershipLocation;
let enemyBulletFrequency = 3000;
let enemyBulletSpeed = 2;
export let windowFocused = true;
export let mothershipExists = false;
export let scoreMultiplier = 1

window.addEventListener('focus', () => {
    windowFocused = true;
});

window.addEventListener('blur', () => {
    windowFocused = false;
});

export function createEnemies(enemyCount) {
    clearInterval(enemyBulletFire)
    enemyDirection = 1;
    enemyX = boxBCR.width / 2 - 200
    enemyY = 0;
    enemyDiv.style.transform = `translate(${boxBCR.width / 2 - 200}px)`;
    enemyDiv.style.width = '400px';
    for (let i = 0; i < enemyCount; i++) {
        let x = i * 50 - Math.floor(i / 8) * 400;
        const enemy = document.createElement('img');
        enemy.setAttribute('id', i);
        enemy.setAttribute('class', 'enemy');
        enemy.src = `../static/img/enemy${Math.floor(Math.random() * 6)}.png`
        enemy.width = 45;
        enemy.style.transform = `translate(${x}px, ${70 + Math.floor(i / 8) * 50}px)`
        enemy.onload = ()=> {
            enemyDiv.appendChild(enemy);
        }
    }
    enemyBulletFire  = setInterval(selectEnemy, enemyBulletFrequency)
}

export function createMothership() {

    if (!mothershipExists){
    mothershipExists = true;
    const mothershipX = boxBCR.width - 50;
    mothershipLocation = mothershipX;
    const mothershipY = boxBCR.top + 20 ;
    const mothership = document.createElement('img');
    mothership.src = '../static/img/mothership.png';
    mothership.setAttribute("class", 'mothershipImg');
    mothership.style.transform = `translate(${mothershipX}px, ${mothershipY}px)`;
    mothership.onload = () => {
        gameDiv.appendChild(mothership);
    };}
}

    export function moveMothership() {

        if (mothershipExists) {
            let buf = mothershipLocation;
            if (gameRunning && !gamePaused && windowFocused) {
                if (mothershipLocation < 0) {
                    document.querySelector('.mothershipImg').remove();
                    mothershipLocation = buf;
                    mothershipExists = false;
                } else if (document.querySelector('.mothershipImg') === null){
                    return
                } else {   
                mothershipLocation -= 4;
                document.querySelector('.mothershipImg').style.transform = `translateX(${mothershipLocation}px)`; 
            }
        }
        }
    }

export function mothershipRestart() {
    if (mothershipExists) {
    mothershipExists = false;
    document.querySelector('.mothershipImg').remove();
    }
}



    export function moveEnemies() {
        if (gameRunning && !gamePaused && windowFocused) {
            if (enemyTouching()) {
                enemyDirection *= -1;
                enemyY += 20;
            }
            enemyX += enemyDirection;
        }
        enemyDiv.style.transform = `translate(${enemyX}px, ${enemyY}px)`
    }

export function mothershipDestroyed(bBCR) {
    let hit = false;
    if (document.querySelector('.mothershipImg') === null) return hit;
    const mothership = document.querySelector('.mothershipImg');
    const mothershipBCR = mothership.getBoundingClientRect();
    
    if (mothershipBCR.bottom >= bBCR.top && mothershipBCR.top <= bBCR.bottom && mothershipBCR.left <= bBCR.left && mothershipBCR.right >= bBCR.right) {
      console.log("mothership Down");
      hit = true;
      addScore(hit);
      mothershipExists = false;
      hit = true;
        mothership.remove();
    }
    return hit;
  }
  


export function enemyDestroyed(bBCR) {
    const enemies = document.querySelectorAll('.enemy');
    let hit = false;
    enemies.forEach((enemy) => {
        const eBCR = enemy.getBoundingClientRect();
        if (eBCR.top <= bBCR.top && eBCR.bottom >= bBCR.top && eBCR.left <= bBCR.left && eBCR.right >= bBCR.right) {
                enemy.remove();
                hit = true;
                addScore(false, enemy.id);
                if (enemies.length <= 1) {
                    addNewEnemies();
                }
        }
    })
    return hit;
}

export function selectEnemy() {
    if (!windowFocused || gamePaused || !gameRunning || gameOver) return;
    let id = Math.floor(Math.random() * 8) + 24;
    let enemy = document.getElementById(id);
    while (enemy === null && id >= 0) {
        id -= 8;
        enemy = document.getElementById(id);
    }
    id < 0 ? selectEnemy() : enemyFire(enemy);
}

function enemyFire(enemy) {
    const enemyBCR = enemy.getBoundingClientRect();
    let bulletX = enemyBCR.x - boxBCR.x + 22;
    let bulletY = enemyBCR.top - boxBCR.top;
    const enemyBullet = document.createElement('div');
    enemyBullet.setAttribute('class', 'enemy-bullet');
    enemyBullet.style.transform = `translate(${bulletX}px, ${bulletY}px)`
    gameDiv.appendChild(enemyBullet);
}

export function selectBulletsAndMove() {
    if (!gamePaused && windowFocused) {
        let bullets = document.querySelectorAll('.enemy-bullet');
        for (let bullet of bullets) { 
            moveEnemyBullet(bullet);
        }
    }
}

function moveEnemyBullet(enemyBullet) {
    if (gameRunning && !gamePaused) {
        const bulletBCR = enemyBullet.getBoundingClientRect();
        let bulletX = bulletBCR.x - boxBCR.x - 1;
        let bulletY = bulletBCR.y - boxBCR.y - 1;
    if (bulletBCR.bottom > boxBCR.bottom || hitShip(bulletBCR)) {
        enemyBullet.remove();
        return;
    }
    bulletY += enemyBulletSpeed
    enemyBullet.style.transform = `translate(${bulletX}px, ${bulletY}px)`
    }
}

function addNewEnemies() {
    if (enemyBulletFrequency > 1000) enemyBulletFrequency -= 100;
    enemyBulletSpeed += 0.1;
    scoreMultiplier *= 2;
    
    createEnemies(32);
}

function enemyTouching() {
    const enemies = document.querySelectorAll('.enemy');
    let touching = false;
    enemies.forEach((enemy) => {
        const enemyBCR = enemy.getBoundingClientRect();
        if (enemyBCR.bottom > boxBCR.bottom - 80) gameLost();
        if (enemyBCR.right >= boxBCR.right || enemyBCR.left <= boxBCR.left) touching = true;
    }) 
    return touching;
}