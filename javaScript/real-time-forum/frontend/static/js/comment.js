import AbstractView from "./AbstractView.js";
import { sendEvent } from "./index.js";
import { currentThreadId } from "./category.js";
import { currentPostId, currentThread } from "./thread.js";

export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Comment");
    }

    async getHtml() {
        return `
        <h1>Comment</h1>
        `;
    }

    async init() {
        sendEvent("get_comments", currentPostId);
        sendEvent("get_users"); 
    }
}

export function appendComments(commentEvent){

   
    appendCurrentThread(currentThread, currentThreadId);
    
    addNewCommentButton();
    console.log("commentEvent:", commentEvent, commentEvent.payload)
    if (typeof commentEvent !== 'object' || !commentEvent.payload) {
        console.error("Invalid commentEvent. It should be an object with a payload property.");
        return;
    }
    
    // Access the payload directly
    const comments = commentEvent.payload;
    
    // Check if the result is an array and it's not empty
    if (!Array.isArray(comments) || comments.length === 0) {
        console.error("Payload does not contain valid comments.");
        return;
    }
    
    console.log("comments before div:", comments);
    
    // Create a new div for the comments container
    const commentsDiv = document.querySelector('main');
    

  
    commentsDiv.id = 'comments';
   
    commentsDiv.appendChild(document.createElement('br'));
    comments.forEach(comment => {
        const commentDiv = document.createElement('div');
        commentDiv.classList.add('comment');
    
        const commentText = document.createElement('p');
        commentText.textContent = comment.Text;
        commentDiv.appendChild(commentText);
    
        const commentInfo = document.createElement('p');
        const postedTime = new Date(comment.Time).toLocaleString();
        const commenter = comment.Username ? comment.Username : "Anonymous";
        commentInfo.textContent = `Posted on: ${postedTime} by: ${commenter}`;
        commentDiv.appendChild(commentInfo);
    
        commentsDiv.appendChild(commentDiv);
    });
    


    
    //console.log("comments:", comments);
    // Create a new comment button

}

function appendCurrentThread(currentThread, currentThreadId) {
    // Ensure currentThreadId is within the bounds of the array
    console.log("currentThreadId:", currentThreadId, currentThread, currentThread.length);
    console.log("currentPostId:", currentPostId);
    let matchingThread = currentThread.find(thread => thread.Id == currentPostId);
    console.log("matchingThread:", matchingThread);

    // Ensure main element is correctly selected
    const mainElement = document.querySelector('main');
    if (!mainElement) {
        console.error("Main element not found.");
        return;
    }

    // Create the post div
    const postDiv = document.createElement('div');
    postDiv.id = 'post';

    const postTitle = document.createElement('h2');
    postTitle.textContent = matchingThread.Title;
    postDiv.appendChild(postTitle); 

    const postContent = document.createElement('p');
    postContent.textContent = matchingThread.Content;
    postDiv.appendChild(postContent); 

    const postInfo = document.createElement('p');
    const createdTime = new Date(matchingThread.Created).toLocaleString();
    const author = matchingThread.Username ? matchingThread.Username : "Anonymous";
    postInfo.textContent = `Posted on: ${createdTime} by: ${author}`;
    postDiv.appendChild(postInfo); 

   
    mainElement.appendChild(postDiv);
}




function appendNewComment(e){
    e.preventDefault(); // Prevent the form from being submitted normally
    const title = "igno/quickfix"/* document.getElementById("title").value; */
    const comment = document.getElementById("comment").value;
    if (comment === "") {
        alert("Comment cannot be empty.");
        console.error("Comment cannot be empty.");
        return;
    }
    let completeComment = [currentPostId, title, comment];
    console.log("Valmis comment:", completeComment);
    document.getElementById('newCommentButton').remove();
    document.getElementById('post').remove();
    document.getElementById("newCommentForm").remove();
    const commentsToRemove = document.querySelectorAll('.comment');
commentsToRemove.forEach(comment => {
    comment.remove();
});

    sendEvent("post_comment", completeComment);
}


function addNewCommentButton(){

    const newCommentButton = document.createElement('button');
    newCommentButton.id = "newCommentButton";
    newCommentButton.textContent = "Add new comment";
    newCommentButton.addEventListener("click", (e) => {
        e.preventDefault();
        let isForm = document.getElementById("newCommentForm");
        if (isForm) {
            isForm.remove();
        } else {
        const newCommentForm = document.createElement('form');
        newCommentForm.id = "newCommentForm";
        newCommentForm.innerHTML = `
        <label for="comment">Write comment below:</label>
        <br>
        <textarea id="comment" name="comment" Required></textarea><br>
        <input type="submit" value="Submit">
        <br>
        `;
        newCommentButton.insertAdjacentElement('afterend', newCommentForm);
        document.getElementById("newCommentForm").onsubmit = appendNewComment;
        }
    });

    // Append the button to the main element
    document.querySelector('main').appendChild(newCommentButton);

    // Create a line break element and append it to the main element
    const lineBreak = document.createElement('br');
    document.querySelector('main').appendChild(lineBreak);
}