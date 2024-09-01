import { places } from "./where-do-we-go.data.js";

export const explore = () => {
  const p = sortPlaces(places);

  createHeading();
  createDirection();

  const i = Math.round(window.scrollY / window.innerHeight);
  let currentPlace;

  const setCurrentPlace = (index) => {
    currentPlace = p[index];
    refreshHeading(currentPlace);
  };

  setCurrentPlace(i);

  let lastPos = window.scrollY;
  window.addEventListener("scroll", (e) => {
    setCurrentPlace(Math.round(window.scrollY / window.innerHeight));
    document.querySelector(".direction").textContent = (lastPos < scrollY) ? "S" : "N";
    lastPos = scrollY;
  });

  renderSections(p);
};

const createDirection = () => {
  const el = document.createElement("div");
  el.classList.add("direction");
  el.textContent = "S";
  document.body.append(el);
};

const refreshHeading = (p) => {
  const h = document.querySelector("a.location");
  h.textContent = `${p.name}\n${p.coordinates}`;
  h.style.color = p.color;
  h.setAttribute("href", `https://www.google.com/maps/place/${p.coordinates}`);
};

const createHeading = () => {
  const el = document.createElement("a");
  el.classList.add("location");
  el.setAttribute("target", "_blank");
  el.setAttribute("href", "#");
  document.body.append(el);
};

const renderSections = (p) => {
  p.forEach((q) => {
    document.body.innerHTML += `<section style="background: url('./where-do-we-go_images/${q.name
      .split(",")[0]
      .toLowerCase()
      .replace(/[ ]/g, "-")}.jpg');"></section>`;
  });
};

const sortPlaces = (p) =>
  p.sort((a, b) => {
    const [lta] = DMS(a.coordinates);
    const [ltb] = DMS(b.coordinates);

    // console.log(lta, ltb);

    return ltb - lta;
  });

const DMS = (p) => {
  const [lat, long] = p.split(" ");

  let lt = DMShelper(lat.slice(0, -1));
  let ln = DMShelper(long.slice(0, -1));

  if (lat.slice(-1) === "S") { lt = -lt; }
  else if (long.slice(-1) === "W") { ln = -ln; }

  return [lt, ln];
};

const DMShelper = (co) => {
  const deg = +co.slice(0, co.indexOf("°"));
  co = co.slice(co.indexOf("°") + 1);

  const min = +co.slice(0, co.indexOf("'"));
  co = co.slice(co.indexOf("'") + 1);
  let sec = +co.slice(0, -1);

  return deg + min / 60 + sec / 60 / 60;
};