import { styles } from "./pimp-my-style.data.js";

export const pimp = () => {
  let butt = document.querySelector("button");
  let clicks = 0;

  if (butt.hasAttribute("data-clicks")) {
    clicks = parseInt(butt.dataset.clicks);
    console.log(clicks, butt.dataset.clicks);
  }

  butt.dataset.clicks = clicks;

  if (butt.classList.contains("unpimp")) {
    butt.classList.remove(styles[clicks]);
    clicks--;

    if (clicks == -1) {
      butt.classList.remove("unpimp");
      clicks++;
    }

  } else {
    butt.classList.add(styles[clicks]);
    clicks++;

    if (clicks == styles.length) {
      butt.classList.add("unpimp");
      clicks--;
    }

  }
  butt.dataset.clicks = clicks;
  console.log("set click", clicks);
};