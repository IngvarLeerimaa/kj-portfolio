// Login page logic
import AbstractView from "./AbstractView.js";
import { login } from "./index.js";
import { Router } from "./route.js";
export default class extends AbstractView {
    constructor(params) {
        super(params);
        this.setTitle("Login");
    }
    async getHtml() {
        return `
        <div class="container">
            <div class="header">
                <h1 id="pageTitle">
                    <span >FORUM</span>
                </h1>
                <h4> <i>This time with websocket</i> </h4>
            </div>
            <div class="content">
                <form id="login-form">
                    <label for="username">Username/E-mail:</label>
                    <input type="text" id="username" name="username" autofocus>
                    <label for="password">Password:</label>
                    <input type="password" id="password" name="password">
                    <input id="login" type="submit" value="Login">
                </form>
                <button id="toReg">Sign up</button>                
            </div>
            <br>
        </div>
        `;
    }
    
    async init() {
        document.getElementById("login-form").onsubmit = login;
        document.getElementById("toReg").addEventListener("click", function() {
            //console.log("here")
            Router("/register");
        });
    }
}