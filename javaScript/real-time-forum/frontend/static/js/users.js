import { changeChatRoom } from "./chat.js";

export function appendUsers(users) {

    const categoryRow = document.getElementById('categories');
    let hasUserBtn = document.getElementById('online-user-button');
    console.log(users)
    // Ensure users is an object and has the expected structure
    if (typeof users !== 'object' || !users.payload || !Array.isArray(users.payload.users)) {
        hasOnlineUsers = false;
        if (hasUserBtn) {
            document.getElementById('online-users').setAttribute('hidden');
        categoryRow.style.width = '100%';
        }
        console.error("Invalid input. Users payload is not as expected.");
        return;
    }

    // Extract the array of user strings from the payload
    const userArray = users.payload.users;

    // Siin vaata kuidas seda loogikat teha.
    //cssis oleks vaja liigutada category rowt, kui on online users

    let hasOnlineUsers;
    
    if (userArray.length > 0) {
        hasOnlineUsers = true;
    } else {
        hasOnlineUsers = false;
        return;
    } 

    
    if (hasOnlineUsers) {
        categoryRow.style.width = '80%';
    } 
    hasUserBtn = true;


    // Assuming you want to create a div with ID 'user-list'
    const userListDiv = document.getElementById("online-users");
    userListDiv.innerHTML = '';
    // Assuming you want to append the div to the body

    // Create an <h2> element for "Online Users" text
    const heading = document.createElement("h1");
    heading.textContent = "Online Users";
    userListDiv.appendChild(heading);

    // Assuming you want to create an unordered list inside the div
    const usersList = document.createElement('ul');
    userListDiv.appendChild(usersList);

    // Clear the existing content of the users list
    usersList.innerHTML = '';

    // Append each user to the list
    userArray.forEach(user => {
        //console.log("user:", user)
        const li = document.createElement('button');
        li.appendChild(document.createTextNode(user.username));
        li.id = "user-" + user.userId;
        li.classList.add('online-user-button');
        li.style.backgroundColor = "white";
        li.onclick = () => {
            console.log("clicking it")

            changeChatRoom(user.username, user.userId);
            li.style.backgroundColor = "white";
        };
        

        usersList.appendChild(li);
    });
    
}
