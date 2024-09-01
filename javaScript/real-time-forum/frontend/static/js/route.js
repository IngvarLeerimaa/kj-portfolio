import register from "./register.js";
import loginPage from "./login.js";
import { isConn, isSession } from "./index.js";
import category from "./category.js";
import  thread  from "./thread.js";
import comment from "./comment.js";
import chat from "./chat.js";

export const Router = async (potentialPath) => {
    //console.log(potentialPath)
    const Routes = [
        { path: "/", component: category },
        { path: "/category", component: category },
        { path: "/register", component: register },
        { path: "/chat", component: chat},
        { path: "/login", component: loginPage },
       // { path: "/profile", component: profile},
        { path: "/comment", component: comment },
        { path: "/thread", component: thread }, 
        // Assuming NotFound is defined somewhere
       /*  { path: "/404", component: NotFound }, */
    ];

        if (isConn == true && (potentialPath == "/login" || potentialPath == "/register")) {
            potentialPath = "/";
        }

    if (isConn == false && potentialPath != "/login" && potentialPath != "/register") {
        // Siia vaja Sessioni check?
        if (isSession()) {

        } else {
            document.querySelector(".nav").setAttribute("hidden", "true");
       // setTimeout(function() {
           Router("/login")
        //}, 50);
        return
        }
    }
    const route = Routes.find(route => route.path === potentialPath);
    if (route) {
        const pageView = new route.component();
        const main = document.querySelector("main");
        main.innerHTML = '';
        main.innerHTML = await pageView.getHtml();
        pageView.init();
    } else {
    // Assuming main is an existing element in your HTML
    const main = document.querySelector("main");
    main.innerHTML = `<h1>Siin on kaka, else statement ${location.pathname}?</h1>`;
}
}
export function linkFinder() {
    const links = document.querySelectorAll(".Link");
    links.forEach((link) => {
        if (link.alreadyHasEventListener) return;
        link.alreadyHasEventListener = true;
        link.addEventListener("click", (e) => {
            e.preventDefault();
            const linkPath = link.getAttribute("href");
            history.pushState(null, null, linkPath);
            Router(linkPath);
        });
    });
    window.addEventListener("load", () => {
        const path = window.location.pathname;
        Router(path);
    });
    window.addEventListener("popstate", () => {
        const path = window.location.pathname;
        Router(path);
    });
}


    /* let formData = {
        "username": document.getElementById("username").value,
        "password": document.getElementById("password").value
    }
    // Send the request
    fetch("login", {
        method: 'post',
        body: JSON.stringify(formData),
        mode: 'cors',
    }).then((response) => {
        if (response.ok) {
            return response.json();
        } else {
            throw 'unauthorized';
        }
    }).then((data) => {
        // Now we have a OTP, send a Request to Connect to WebSocket
        connectWebsocket(data.otp, formData.username);
        if (data.redirect) {
            isConn = true;
            Router(data.redirect)
        }
    }).catch((e) => { alert(e) });
    return false; */

/* const urlRoute = (event) => {
    event = event || window.event;
    event.preventDefault();
    window.history.pushState({},"", event.target.href);
    urlLocationHandler();
}
const urlLocationHandler = async () => {
    let location = window.location.pathname; //takes correct url
    if (location.length === 0) {
        location = "/templates/index.html";
    }
    const route = urlRoutes[location] || urlRoutes[404];
    const html = await fetch(route.template).then((response) => response.text()); // Call the template function to get the template string
    document.getElementsByClassName("main").innerHTML = html;
    document.title = route.title;
    document.querySelector('meta[name="description"]').setAttribute("main", route.description);
};
 */