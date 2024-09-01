export const generateLetters = (length = 120, steps = [300, 400, 600]) => {
    const ulet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    const body = document.body;
    const step = parseInt(length / steps.length);
    for (let i = 0; i < length; i++) {
      console.log(i, step, steps[parseInt(i / step)]);
      const letter = ulet[Math.floor(Math.random() * ulet.length)];
      const div = document.createElement("div");
      div.textContent = letter;
      div.style.position = "relative";
      div.style.fontSize = `${i + 11}px`;
      div.style.fontWeight = `${steps[parseInt(i / step)]}`;
      body.appendChild(div);
    }
  }