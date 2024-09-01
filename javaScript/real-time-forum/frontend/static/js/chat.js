import AbstractView from "./AbstractView.js";
import { sendEvent } from "./index.js";
import { isConn } from "./index.js";
import { Router } from "./route.js";
class SendMessageEvent {
    constructor(message) {
        this.message = message;
    }
}

class ChangeChatRoomEvent {
        constructor(UserID, LastMessage, RoomChange, LastContent) {
            this.UserID = UserID;
            this.LastMessage = LastMessage
            this.RoomChange = RoomChange
            this.LastContent = LastContent
            
        }
    }
    
export default class extends AbstractView {
    constructor() {
        super();
        this.setTitle("Chat");
    }

    async getHtml() {
        return chatTemplate;
    }

    init() {
        document.getElementById("chatroom-message").onsubmit = sendMessage;
        sendEvent("get_users");
        document.getElementById("homeBtn").addEventListener("click", () => {
            selectedchat = "smth";
            //so it would
            sendEvent("change_room", new ChangeChatRoomEvent(0, 0, false));
        });
} 
}
let selectedchat;

export async function changeChatRoom(username, id) {
    
    console.log("changeChatRoom funktsioon", username, id, selectedchat)
    // Change Header to reflect the Changed chatroom
    if (username != null && username != selectedchat) {
        selectedchat = username;
       await Router("/chat");
       let header = document.getElementById("chat-header").innerHTML = "Currently in chat: " + selectedchat;

        let changeEvent = new ChangeChatRoomEvent(id, 0, true);
        sendEvent("change_room", changeEvent);
        /* let textarea = document.getElementById("chatmessages");
        textarea.innerHTML = `You changed room into: ${username}`; */
    }
    return false;
}

function sendMessage() {
var newmessage = document.getElementById("message");
if (newmessage != null && newmessage.value != "") {
    //username on hardcoded, vaja muuta
    let outgoingEvent = new SendMessageEvent(newmessage.value);
    sendEvent("send_message", outgoingEvent)
    newmessage.value = "";
}
return false;
}

let offset = 0;
let debounceTimer;
let eventSent = false;
let lastContent;

export function appendOldMessages(event) {
    if (event.payload.messages == null) {
        return;
    }
    let textarea = document.getElementById("chatmessages");
    offset = event.payload.lastmessage;

    event.payload.messages.forEach(message => {
       
        let createdDate = new Date(message.Created);
        let formattedDate = createdDate.toLocaleString();
        if (textarea.innerHTML == "") {
            textarea.innerHTML = formattedDate + " " + message.From + ":" + "\n" + message.Content;
        } else {
            textarea.innerHTML = formattedDate + " " + message.From + ":" + "\n" + message.Content + "\n" + textarea.innerHTML;
        }
        textarea.scrollTop = textarea.scrollTop + 1;

        lastContent = message.Content;
        if (lastContent == "You've reached the end of the chat") {
            return;
        }
    });

    const scrollListening = document.getElementById("chatmessages");
    scrollListening.addEventListener("scroll", function() {
        clearTimeout(debounceTimer);
        debounceTimer = setTimeout(() => {
            if (scrollListening.scrollTop == 0 && !eventSent && lastContent !== "You've reached the end of the chat") {
                console.log("scrolling to top");
                sendEvent("change_room", {LastContent: lastContent, offset: offset, roomChange: true});
                eventSent = true;
                setTimeout(() => {
                    if (lastContent !== "You've reached the end of the chat") {
                        eventSent = false;
                    }
                }, 500);
            }
        }, 500);
    });
}




export function appendChatMessage(messageEvent) {
    console.log(messageEvent)
    if (messageEvent.userId == undefined) {
        offset += 1;
        console.log(offset)
        var date = new Date(messageEvent.sent);
    // format message
    const formattedMsg = `${date.toLocaleString()} ${messageEvent.from}:\n${messageEvent.message}`;
    // Append Message
    let textarea = document.getElementById("chatmessages");
    textarea.innerHTML = textarea.innerHTML + "\n" + formattedMsg;
    textarea.scrollTop = textarea.scrollHeight;
    } else {
        document.getElementById("user-" + String(messageEvent.userId)).style.backgroundColor = "red";
    }
    
}


const chatTemplate =`
<div>
<h3 id="chat-header">Currently in chat with:</h3>

<textarea class="messagearea" id="chatmessages" readonly name="chatmessages"
    placeholder="Messages from others will appear"></textarea>
<!--
Chatroom-message form is used to send messages
-->
<form id="chatroom-message">
    <label for="message">Message:</label>
    <br>
    <input type="text" id="message" name="message">
    <input type="submit" value="Send">
</form>
</div>
<main></main>
        `;