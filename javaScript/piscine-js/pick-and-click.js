export const pick = () => {

    const hsl = document.createElement("div");
    hsl.className = "text hsl";
    hsl.style.top = `50%`;
    hsl.style.left = `50%`;
    hsl.style.textAlign = `center`;
    document.querySelector("body").append(hsl);
  
    const hue = document.createElement("div");
    hue.className = "text hue";
    document.querySelector("body").append(hue);
  
    const luminosity = document.createElement("div");
    luminosity.className = "text luminosity";
    document.querySelector("body").append(luminosity);
  
    const svg = document.createElementNS("http://www.w3.org/2000/svg", "svg");
  
    document.querySelector("body").append(svg);
    const axisX = document.createElementNS("http://www.w3.org/2000/svg", "line");
    axisX.id = "axisX";
  
    const axisY = document.createElementNS("http://www.w3.org/2000/svg", "line");
    axisY.id = "axisY";
  
    svg.appendChild(axisX);
    svg.appendChild(axisY);
    document.addEventListener("mousedown", (e) => navigator.clipboard.writeText(hsl.textContent));

    document.addEventListener("mousemove", (e) => {
      const x = Math.round((e.clientX / window.innerWidth) * 360);
      const y = Math.round((e.clientY / window.innerHeight) * 100);
 
      axisX.setAttributeNS(null, "x1", `${e.clientX}`);
      axisX.setAttributeNS(null, "x2", `${e.clientX}`);
      axisX.setAttributeNS(null, "y1", "0");
      axisX.setAttributeNS(null, "y2", "100%");
      axisY.setAttributeNS(null, "x1", "0");
      axisY.setAttributeNS(null, "x2", "100%");
      axisY.setAttributeNS(null, "y1", `${e.clientY}`);
      axisY.setAttributeNS(null, "y2", `${e.clientY}`);
      hue.textContent = `hue\n${x}`;
      luminosity.textContent = `${y}\nluminosity`;
      hsl.textContent = `hsl(${x},50%,${y}%)`;
      document.body.style.background = `hsl(${x},50%,${y}%)`;
    });
  };