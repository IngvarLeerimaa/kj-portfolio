<template>
  <div class="login-container">
    <div><h1>Social Network</h1></div>
    <form class="form-container" @submit.prevent="login">
      <div class="form-group">
        <label for="email">Email:</label>
        <input
          type="email"
          id="email"
          autocomplete="email"
          v-model="email"
          autofocus
          required
        />
      </div>
      <div class="form-group">
        <label for="password">Password:</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <div class="form-submit"><button type="submit">Login</button></div>
    </form>
    <router-link to="/register">Don't have an account? Sign up!</router-link>
  </div>
</template>

<script>
import WebSocketConnection from "../assets/websocket";
import app from "@/main";
export default {
  data() {
    return {
      email: "",
      password: "",
    };
  },
  methods: {
    async login() {
      await fetch("http://localhost:3000/api/v1/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email: this.email,
          password: this.password,
        }),
      })
        .then(async (response) => {
          if (!response.ok) {
            const msg = await response.text();
            throw new Error(msg);
          }
          return response.json();
        })
        .then((data) => {
          let date = new Date();
          date.setTime(date.getTime() + 1 * 24 * 60 * 60 * 1000);
          let expires = "expires=" + date.toUTCString();
          document.cookie =
            "sessionID=" + data.session_id + ";" + expires + ";path=/";
          if (WebSocketConnection.ws != null) WebSocketConnection.ws.close();
          this.$router.push("/");
          app.config.globalProperties.$userId = data.user_id;
          return;
        })
        .catch((error) => {
          alert("Login failed: " + error.message);
        });
    },
  },
};
</script>

<style scoped>
.login-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  height: 100vh;
  margin-top: 5%;
}

h1 {
  color: #2c3e50;
}

.form-group {
  display: flex;
  align-items: center;
  margin-bottom: 2px;
}

.form-submit {
  margin-top: 5px;
  margin-right: 20px;
}

label {
  width: 130px;
  text-align: right;
}

a {
  margin-top: 10px;
  font-weight: bold;
  color: #2c3e50;
}

a:hover {
  color: #707070;
}
</style>
