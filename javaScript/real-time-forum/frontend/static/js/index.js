// index.js
import { linkFinder } from "./route.js";
import { Router } from "./route.js";
import { appendChatMessage, appendOldMessages } from "./chat.js";
import { appendCategories } from "./category.js";
// Now you can use Router in this file
import { appendUsers } from "./users.js";
import { appendThreads } from "./thread.js";
import { appendComments } from "./comment.js";
import { currentPostId } from "./thread.js";
import { currentThreadId } from "./category.js";



// Ws
export let conn;
export let isConn = false;
linkFinder();

class Event {
    constructor(type, payload) {
        this.type = type;
        this.payload = payload;
    }
}

class NewMessageEvent {
    constructor(message, from, sent) {
        this.message = message;
        this.from = from;
        this.sent = sent;
    }
}


export function login() {
        let formData = {
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
        return false;
    }
/* function logout() {
    const logoutBtn = document.createElement('button');
    logoutBtn.onclick = () => {
        conn = null;
    }
    document.querySelector(".nav").appendChild(logoutBtn)
} */
function connectWebsocket(otp, username) {
    // Check if the browser supports WebSocket
    if (window["WebSocket"]) {
        console.log("supports websockets");
        // Connect to websocket using OTP as a GET parameter
        conn = new WebSocket("wss://" + document.location.host + "/ws?otp=" + otp + "&username=" + username);
        const logoutBtn = document.createElement('button');
        logoutBtn.id = "logout";
        logoutBtn.innerHTML = "Logout";
        logoutBtn.onclick = () => {
        document.cookie = "sessionID=; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
        isConn = false;
        conn.close();
        Router("/login")
        document.querySelector(".nav").removeChild(logoutBtn);
        document.querySelector(".online-users").style.display = "none";
        document.querySelector(".nav").setAttribute("hidden", "true");
        }
    document.querySelector(".nav").removeAttribute("hidden");
    document.querySelector(".online-users").style.display = "block"
    document.querySelector(".nav").appendChild(logoutBtn)

        

        // Onopen
        conn.onopen = function (evt) {
            isConn = true;
        }

        conn.onclose = function (evt) {
            // Set disconnected
           if (!isSession) {
            alert("connection closed and you dont have an active session")
            document.cookie = "sessionID=; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
        isConn = false;
        conn.close();
        Router("/login")
        document.querySelector(".online-users").remove;
        document.querySelector(".nav").removeChild(logoutBtn);
        document.querySelector(".nav").setAttribute("hidden", "true");
        document.querySelector(".online-users").setAttribute("hidden", "true");

           }
        }

        // Add a listener to the onmessage event
        conn.onmessage = function (evt) {
          //console.log("before:",evt);
            // parse websocket message as JSON
            const eventData = JSON.parse(evt.data);
            // Assign JSON data to new Event Object
            const event = Object.assign(new Event, eventData);
            console.log("after:", event);
            // Let router manage message
            routeEvent(event);
        }

    } else {
        alert("Not supporting websockets");
    }
}

function routeEvent(event) {
    console.log("WE are at route event:", event)
    if (event.type === undefined) {
        alert("no 'type' field in event");
    }

    switch (event.type) {
        case "new_message":
            // Format payload
            const messageEvent = Object.assign(new NewMessageEvent, event.payload);
            appendChatMessage(messageEvent);
            break;

        case "get_categories":
            appendCategories(event);
            break;

        case "get_threads":
            appendThreads(event);
            break;

        case "post_threads":
            console.log("post_threads event:", event);
            alert("Post added");
            sendEvent("get_threads", currentThreadId);
            break;

        case "get_comments":
            console.log("get_comments event:", currentPostId);
            appendComments(event);
            break;

        case "post_comment":
            alert("Comment added");
            sendEvent("get_comments", currentPostId);
            break;

        case "get_users":
            appendUsers(event);
            break;
        case "get_messages":
            appendOldMessages(event);
            break;
        case "":
            Router("/");
            break;

        default:
            alert("unsupported message type");
            break;
    }
}

export function sendEvent(eventName, payload) {
        const event = new Event(eventName, payload);
        console.log("sendEvent funktsioon:", event);
        conn.send(JSON.stringify(event));
    }

    export function isSession() {
        let sessionID = null
        const cookies = document.cookie.split(";");
        cookies.forEach( cookie => {
            if (cookie.indexOf("sessionID") == 0){
                sessionID = cookie.substring(10)
            }
        });
        console.log(sessionID)
        if (sessionID == null) return false 
    
        let formData = {
            "sessionID": sessionID,
        }
    
        fetch("session", {
            method: 'post',
            body: JSON.stringify(formData),
            mode: 'cors',
        }).then((response) => {
            if (response.ok) {
                console.log(response)
                return response.json();
            } else {
                return false;
            }
        }).then((data) => {
            // Now we have a OTP, send a Request to Connect to WebSocket ")
            connectWebsocket(data.otp, data.username);
            console.log(data.otp)
            if (data.redirect) {
                isConn = true;
                Router(data.redirect)
            }
        }).catch((e) => { alert(e) });
        return true
    }

/**
 * Once the website loads
 * */
window.onload = function () {
    // Apply our listener functions to the submit event on both forms
    // we do it this way to avoid redirects
   /*  document.getElementById("chatroom-selection").onsubmit = changeChatRoom;
    document.getElementById("chatroom-message").onsubmit = sendMessage; */
   // document.getElementById("login-form").onsubmit = login;
     /* if (isConn == false && window.location.pathname != "/login") {
        alert('WebSocket is not connected. You will be redirected to the login page in 1 second.');
        setTimeout(function() {
           Router("/login")
        }, 1000);
    }
 */

};