export const build = (num) => {
    let c = 1;
    let b = 2;
    let display = undefined; // to lock upportunity to break the flow before the columns are built
    const body = document.querySelector("body");
    const addBrick = setInterval(() => {
      if (display === undefined) {
        display = document.querySelector("section").style.display;
        document.querySelector("section").style.display = 'none';
      }
      const brick = document.createElement("div");
      brick.id = "brick-" + c;
      brick.innerHTML = c;
      if (b == c) { // central column
        brick.dataset.foundation = true;
        b = c + 3;
      }
      body.append(brick);
  
      if (c == num) {
        clearInterval(addBrick);
        document.querySelector("section").style.display = display;
      }
      c++;
    }, 100);
  };
  
  export const repair = (...ids) => {
    ids.forEach((val) => {
      const brick = document.querySelector("div#" + val);
      if (brick) brick.dataset.repaired = brick.hasAttribute("data-foundation") ? "in progress" : true;
    });
  };
  
  export const destroy = () => {
    let bricks = document.querySelectorAll("div");
    bricks[bricks.length - 1].remove();
  };