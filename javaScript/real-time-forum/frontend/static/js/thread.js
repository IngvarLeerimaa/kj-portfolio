import AbstractView from "./AbstractView.js";
import { sendEvent } from "./index.js";
import { Router } from "./route.js";
import { currentThreadId } from "./category.js";
import { currentCategory } from "./category.js";
export let currentThread;
export let currentPostId;


export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Thread");
    }

    async getHtml() {
        return `
        <h1>Thread</h1>
        `;
    }

    async init() {
        sendEvent("get_users");
        sendEvent("get_threads", currentThreadId)
    }
}


export function appendThreads(threadEvent) {
    // Ensure threadEvent is an object with a payload property
    //console.log("appendThreads funktsioon:", threadEvent);
    
    addNewPost();

    if (typeof threadEvent !== 'object' || !threadEvent.payload) {
        console.error("Invalid threadEvent. It should be an object with a payload property.");
        return;
    }
    // Access the payload directly
    const threads = threadEvent.payload;
    currentThread = threads;
    // Check if the result is an array and it's not empty
    if (!Array.isArray(threads) || threads.length === 0) {
        console.error("Payload does not contain valid threads.");
        return;
    }
   // console.log("appendThreads funktsioon:", threads);
    let threadDiv = document.querySelector('main');
    threadDiv.id = 'thread';
    threads.forEach(thread => {
        const threadContainer = document.createElement('div');
        threadContainer.classList.add('thread-container');
    
        const threadLink = document.createElement('a');
        threadLink.textContent = thread.Title;
        threadLink.href = thread.Id;
        threadLink.classList.add('thread-link');
    
        const threadContent = document.createElement('span');
        threadContent.textContent = thread.Content;
        threadContent.classList.add('thread-content');
    
        threadContainer.appendChild(threadLink);
        threadContainer.appendChild(threadContent);
        threadDiv.appendChild(threadContainer);
        threadDiv.appendChild(document.createElement('br'));
    });
    
    


    //console.log("appendThread funktsioon:", threads.payload);
    
    const threadLinks = document.querySelectorAll('#thread a');

    threadLinks.forEach((link) => {
        link.addEventListener("click", (e) => {
            e.preventDefault();
           // const linkId = link.getAttribute("href");
            history.pushState(null, null, "/comment");
            currentPostId = link.getAttribute("href")
            Router("/comment");
        });
    });
}

class PostThreadEvent {
    constructor(title, content, category) {
      this.title = title;
      this.content = content;
      this.category = category;
    }
  }

function appendNewPost(e){
    e.preventDefault();
    const title = document.getElementById("title").value;
    const post = document.getElementById("newPost").value;

    // Get the keys of the checked checkboxes
    const checkboxes = document.querySelectorAll('input[type="checkbox"]');
    const checkedKeys = Array.from(checkboxes)
        .filter(checkbox => checkbox.checked)
        .map(checkbox => checkbox.value);
    if (checkedKeys.length === 0) {
        alert("Please select at least one category");
        return;
    }
    let completePost = new PostThreadEvent(title, post, checkedKeys);
    //let completePost = [title, post, ...checkedKeys];
    console.log("post:", completePost);
    document.getElementById("newPostForm").remove();
    document.getElementById("newPostbutton").remove();
    const threadContainersToRemove = document.querySelectorAll('.thread-container');
    threadContainersToRemove.forEach(container => {
    container.remove();
    });
    const brContainersToRemove = document.querySelectorAll('br');
    brContainersToRemove.forEach(container => {
    container.remove();
    });

    sendEvent("post_threads", completePost);
}

function addNewPost(){
    
    const newPostbutton = document.createElement('button');
    newPostbutton.textContent = "Add new post";
    newPostbutton.id = "newPostbutton";
    newPostbutton.addEventListener("click", (e) => {
        e.preventDefault();
        const isForm = document.getElementById("newPostForm");
        if (isForm) {
            isForm.remove();
        } else {
        const newPostForm = document.createElement('form');
        newPostForm.id = "newPostForm";

        // Create checkboxes for each category
        const checkboxes = [...currentCategory.entries()].map(([key, value]) => `
            <input type="checkbox" id="category${key}" name="category${key}" value="${key}">
            <label for="category${key}">${value}</label>
        `).join('');

        newPostForm.innerHTML = `
        <label for="title">Title:</label>
        <br>
        <input type="text" id="title" name="title" Required><br>
        <label for="post">Write post below:</label>
        <br>
        <textarea id="newPost" name="post" Required></textarea><br>
        ${checkboxes}
        <br>
        <input type="submit" value="Submit">
        <br>
        `;

        
        newPostbutton.insertAdjacentElement('afterend', newPostForm);
        document.getElementById("newPostForm").onsubmit = appendNewPost;
        }
    });
    
    document.querySelector('main').appendChild(newPostbutton);

    const lineBreak = document.createElement('br');
    document.querySelector('main').appendChild(lineBreak);
}

/* export const threadTemplate = 
`<!--Header-->
<div class="container">
    <div class="header">
        <h1 class="center"> <a href="/">Forum</a> > {{ .CategoryName }} > {{ (index .Threads 0).Title }}</h1>
        {{ if .User.Id }}
        <div class="center">
                <a href="/" class="navBtn">Home</a>
                <a href="/profile" class="navBtn">Profile</a>
                <a href="/logout" class="navBtn">Logout</a> 
            <!-- <form action="/">
                <button type="submit">Home</button>
            </form> -->
            <!-- <form action ="/profile">
                <button type="submit">Profile</button>
            </form> --> 
                  <!-- <form action="/logout">
                <button type="submit">Logout</button>
            </form> -->
        </div>
        {{ else }}
        <div class="center">
             <a href="/" class="navBtn">Home</a>
            <a href="/login" class="navBtn">Login</a>
            <a href="/register" class="navBtn">Register</a>  
      <!--       <form action="/">
                <button type="submit">Home</button>
            </form>
            <form action="/login">
                <button type="submit">Login</button>
            </form>
            <form action="/register">
                <button type="submit">Register</button>
            </form>  -->
        </div>
        {{ end }}
    </div> <!-- close the header div -->
</div> <!-- close the container div -->

<!--THREAD-->
<main>  
  
    {{ range .Threads }}

        <div class="category-column">
            <h4>
                <a>{{ .Content }}</a>
            </h4>
            <div class="likes">
            {{ if $.User.Id }}
                {{ if .Liked }}
                    <button class="btn green" id="like" value="{{ .Id }}" data-thread="true">
                {{ else }}
                    <button class="btn" id="like" value="{{ .Id }}" data-thread="true">
                {{ end }}
                <i class="fa fa-thumbs-up fa-lg" aria-hidden="true"></i>    
                <label id="like-count">{{ .Likes }}</label>    
            </button>
                {{ if .Disliked }}
                    <button  class="btn red" id="dislike" value="{{ .Id }}" data-thread="true">
                {{ else }}
                    <button  class="btn" id="dislike" value="{{ .Id }}" data-thread="true">
                {{ end }}
                <i class="fa fa-thumbs-down fa-lg" aria-hidden="true"></i>  
                <label id="dislike-count">{{ .Dislikes }}</label>   
            </button>
             {{ else }}
                <i class="fa fa-thumbs-up fa-lg" aria-hidden="true"></i>    
                {{ .Likes }}   

                <i class="fa fa-thumbs-down fa-lg" aria-hidden="true"></i>  
                {{ .Dislikes}}
             {{ end }}
            </div>
            <!--username and date-->
            <br>
            <p>Created by {{ .Username }} at {{ .Created }}</p>
        </div>
        {{ end }}
        <br>
        
        <form action="/newComment">
            <input type="hidden" name="t" value="{{ (index .Threads 0).Id }}">
            <button class="comment" type="submit">New Comment</button>
            </form>
   {{ if .Comments}}
   <br>
   <p class="comment">Comments:</p>
   <br>
        {{ range .Comments }}
        <div class="category-column">
            <h4>
                <a>{{ .Text }}</a>
            </h4>

            <div class="likes">
            {{ if $.User.Id }}
                {{ if .Liked }}
                    <button class="btn green" id="like" value="{{ .Id }}" data-thread="false">
                {{ else }}
                    <button class="btn" id="like" value="{{ .Id }}" data-thread="false">
                {{ end }}
                <i class="fa fa-thumbs-up fa-lg" aria-hidden="true"></i>    
                <label id="like-count">{{ .Likes }}</label>    
            </button>
                {{ if .Disliked }}
                    <button  class="btn red" id="dislike" value="{{ .Id }}" data-thread="false">
                {{ else }}
                    <button  class="btn" id="dislike" value="{{ .Id }}" data-thread="false">
                {{ end }}
                <i class="fa fa-thumbs-down fa-lg" aria-hidden="true"></i>  
                <label id="dislike-count">{{ .Dislikes }}</label>   
            </button>
             {{ else }}
                <i class="fa fa-thumbs-up fa-lg" aria-hidden="true"></i>    
                {{ .Likes }}   

                <i class="fa fa-thumbs-down fa-lg" aria-hidden="true"></i>  
                {{ .Dislikes}}
             {{ end }}
            </div>
            <!--username and date-->
            <br>
            <p>Created by {{ .Username }} at {{ .Time }}</p>
        </div>
        {{ end }}
    {{ else }}
    <br>
        <p>There are no comments available.</p>
    {{ end}}
</main>`;  
 */
/* const likeBtns = document.querySelectorAll('#like');
const dislikeBtns = document.querySelectorAll('#dislike');
const likeCounts = document.querySelectorAll('#like-count');
const dislikeCounts = document.querySelectorAll('#dislike-count');

for (let i = 0; i < likeBtns.length; i++) {
const likeBtn = likeBtns[i];
const dislikeBtn = dislikeBtns[i];
var likeCount = 0;
var dislikeCount = 0;

likeBtn.addEventListener('click', function() {  
if (dislikeBtn.classList.contains('red')) {
dislikeBtn.classList.remove('red');
updateDislikes('remove', this.value, this.dataset.thread);
dislikeCount = parseInt(dislikeCounts[i].innerHTML, 10) - 1; 
dislikeCounts[i].innerHTML = dislikeCount;
} 
if (!this.classList.contains('green')) {
    updateLikes('add', this.value, this.dataset.thread);
    likeCount = parseInt(likeCounts[i].innerHTML, 10) + 1; 
} else {
    updateLikes('remove', this.value, this.dataset.thread);
    likeCount = parseInt(likeCounts[i].innerHTML, 10) - 1; 
}
likeCounts[i].innerHTML = likeCount;
this.classList.toggle('green'); 
});

dislikeBtn.addEventListener('click', function() { 
if (likeBtn.classList.contains('green')) {
likeBtn.classList.remove('green');
updateLikes('remove', this.value, this.dataset.thread);
likeCount = parseInt(likeCounts[i].innerHTML, 10) - 1; 
likeCounts[i].innerHTML = likeCount;
} 
if (!this.classList.contains('red')) {
    updateDislikes('add', this.value, this.dataset.thread);
    dislikeCount = parseInt(dislikeCounts[i].innerHTML, 10) + 1;
} else {
    updateDislikes('remove', this.value, this.dataset.thread);
    dislikeCount = parseInt(dislikeCounts[i].innerHTML, 10) - 1;
}
dislikeCounts[i].innerHTML = dislikeCount;
this.classList.toggle('red');
});
} */

/* function updateLikes(button, comment, thread) {
        fetch(`/like?l=${button}&c=${comment}&t=${thread}`);
        }         

function updateDislikes(button, comment, thread) {
        fetch(`/dislike?l=${button}&c=${comment}&t=${thread}`);
        }} */