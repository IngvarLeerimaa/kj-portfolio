import simple from "../../framework/framework.js";
import WSConnection from "./websocket.js";

export const chatElement = () => {
    const chatContainer = new simple.NewElement("div", {
        attrs: {
            class: "chat-container"
        },
        children: [
            // Incoming messages container
            ["div", {
                attrs: {
                    class: "incoming-messages"
                },
                children: [
                    ["div", {
                        attrs: {
                            placeholder: "Chat messages will appear here",
                            class: "messages"
                        }
                    }]],
            }],
            // Container for textarea and send button
            ["div", {
                attrs: {
                    class: "input-container"
                },
                children: [
                    ["textarea", {
                        attrs: {
                            placeholder: "Write smth...",
                            class: "input-text"
                        }
                    }],
                    ["button", {
                        attrs: {
                            id: "send-button",
                            type: "button",
                            class: "send-button"
                        },
                        children: ["Send"],
                    }]
                ]
            }]
        ]
    }).create();

    return chatContainer;
}


export function buttonClick() {
    let messageInput = document.querySelector(".chat-container textarea");
    let message = messageInput.value.trim();

    if (message === "") {
        return;
    }

    messageInput.value = ""; // Clear the input after sending

    message = JSON.stringify({
        messageType: "message",
        message: message
    });

   try {WSConnection.sendMessage(message)
    }catch(e){
        alert("You have to join a game before you can chat")
    console.log(e)
};
    
}

export class BubbleMessage {
    constructor(from, message, colour) {
        this.from = from;
        this.message = message;
        this.colour = colour;
    }
    createMessage() {

        const textColor = this.colour === "white" ? "black" : "white";
        const message = new simple.NewElement("div", {
            attrs: {
                class: "message-bubble",
                style: `background-color: ${this.colour}; color: ${textColor}`
            },
            children: [`${this.from}:${this.message}`]
        }).create();

        let chat = document.getElementsByClassName("incoming-messages");
        if (chat.length > 0) {
            chat[0].appendChild(message);
        }

        //Bad fix, but it works
        chat[0].appendChild(document.createElement("br"));

        return message; // If you need to use the created message later
    }
}

export const isChatActive = false;

const chatInput = document.getElementsByClassName("text-input");


export function addListener() {
    console.log("Adding chat listener");
    const chatInput = document.getElementsByClassName("text-input")[0];
    console.log(chatInput)
    if (chatInput !== undefined) {
        chatInput.addEventListener("focus", () => {
            console.log("Chat is active");
            isChatActive = true;
        });
        chatInput.addEventListener("blur", () => {
            isChatActive = false;
        });
    }
}
