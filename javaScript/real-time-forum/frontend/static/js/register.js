import AbstractView from "./AbstractView.js";

import { Router, linkFinder } from "./route.js";
export default class extends AbstractView { 
    constructor(params) {
        super(params);
        this.setTitle("Register");
    }

    async getHtml() {
        return `
        
        <div class="container">
        <div class="header">
        <h1 id="regTitle">
            Register
        </h1>
        </div>
        <form id="sign-up-form" onsubmit="return false;">
            Username: 
            <input type="text" id="username" placeholder="Username" required minlength="2" maxlength="64"> 
    
            First name: 
            <input type="text" id="first-name" placeholder="First name" required minlength="2" maxlength="64"> 
    
            Last name: 
            <input type="text" id="last-name" placeholder="Last name" required minlength="2" maxlength="64"> 
    
            Age: 
            <input type="number" id="age" placeholder="Age" required min="12" max="123"> 
    
            Gender: 
            
            <label for="gender-male">Male</label>
           <input type="radio" name="gender" id="gender-male" value="M" required> 
            <label for="gender-female">Female</label>
            <input type="radio" name="gender" id="gender-female" value="F"> 
            <br>
            E-mail: 
            <input type="email" id="email" placeholder="E-mail" required maxlength="64"> 
    
            Password: 
            <input type="password" id="password" placeholder="Password" minlength="7" maxlength="64" required> 
    
            Confirm password: 
            <input type="password" id="password-confirm" placeholder="Confirm Password" maxlength="64" required> 
    
            <div class="error" id="error-message"></div>
            <button id="signUp" type="submit">Sign up</button>
    
        </form>
        
        <button id="toLogin" class="Link">I already have an account</button>
        
    </div>
    
` }
async init() {
    document.getElementById("toLogin").addEventListener("click", function() {
        //console.log("here")
        Router("/login");
    });
    const signUpForm = document.getElementById("sign-up-form")
    
        signUpForm.addEventListener("submit", function () {
            const password = document.getElementById("password")
            const passwordConfirm = document.getElementById("password-confirm")

            if (password.value != passwordConfirm.value) {
                fkdUp("Passwords Don't Match")
            } else {
                let formData = {
                    username: document.getElementById("username").value,
                    firstName: document.getElementById("first-name").value,
                    lastName: document.getElementById("last-name").value,
                    age: document.getElementById("age").value,
                    gender: document.getElementById("gender-male").checked ? "M" : "F",
                    email: document.getElementById("email").value,
                    password: password.value,
                }
                console.log(formData)
                signUp(formData)
}
})
}}

const fkdUp = (err) => {
    console.log("Password don't match")
    const errorMessage = document.getElementById("error-message")
    errorMessage.innerText = err
}


function signUp(formData) {

    // Send the request
    fetch("register", {
        method: 'post',
        body: JSON.stringify(formData),
        mode: 'cors',
    }).then((response) => {
        if (response.ok) {
            return response.json();
        } else {
        console.log(response)
            throw 'user already exists';
        }
    }).then((data) => {
        if (data.redirect) {
            alert("User created successfully")
            Router(data.redirect)
        }
    }).catch((e) => { 
        alert(e) 
    });
    return false;
}