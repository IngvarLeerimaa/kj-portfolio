import { colors } from "./fifty-shades-of-cold.data.js";

export const generateClasses = () => {
  let style = document.createElement("style");
  
  colors.forEach((elem) => {

    style.innerHTML += `.${elem} { background: ${elem}; }`;
  });

  document.head.append(style);
};

export const generateColdShades = () => {

  let myArr = [
    "aqua", "blue", "turquoise", "green", "cyan", "navy", "purple",
  ];

  colors.forEach((elem) => {
    let div = document.createElement("div");
    
    div.classList.add(elem);

    div.textContent = elem;
    if (myArr.some((el) => div.className.includes(el))) document.body.append(div);
  });
};

export const choseShade = (shade) => 
Array.from(document.querySelectorAll("div")).map((div) => 
div.className = shade);
