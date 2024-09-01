<template>
  <div class="register-container">
    <div><h1>Social Network</h1></div>
    <form class="form-container" @submit.prevent="register">
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
        <input
          type="password"
          id="password"
          autocomplete="new-password"
          v-model="password"
          required
        />
      </div>
      <div class="form-group">
        <label for="first-name">First Name:</label>
        <input
          type="text"
          id="first-name"
          autocomplete="given-name"
          v-model="firstName"
          required
        />
      </div>
      <div class="form-group">
        <label for="last-name">Last Name:</label>
        <input
          type="text"
          id="last-name"
          autocomplete="family-name"
          v-model="lastName"
          required
        />
      </div>
      <div class="form-group">
        <label for="date-of-birth">Date Of Birth:</label>
        <input type="date" id="date-of-birth" v-model="dateOfBirth" required />
      </div>
      <div class="form-group">
        <label for="avatar">Avatar(optional):</label>
        <input type="file" id="avatar" accept="image/*" @change="uploadImage" />
      </div>
      <div class="form-group">
        <label for="nickname">Nickname(optional):</label>
        <input type="text" id="nickname" v-model="nickname" />
      </div>
      <div class="form-group">
        <label for="about">About Me(optional):</label>
        <input type="textfield" id="about" v-model="about" />
      </div>
      <div class="form-submit"><button type="submit">Register</button></div>
    </form>
    <router-link to="/login">Already have an account? Log in!</router-link>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: "",
      password: "",
      firstName: "",
      lastName: "",
      dateOfBirth: "",
      avatar: null,
      nickname: "",
      about: "",
    };
  },
  methods: {
    uploadImage(event) {
      this.avatar = event.target.files[0];
    },
    async register() {
      var formData = new FormData();
      formData.append("email", this.email);
      formData.append("password", this.password);
      formData.append("firstname", this.firstName);
      formData.append("lastname", this.lastName);
      formData.append("dateofbirth", this.dateOfBirth);
      formData.append("avatar", this.avatar);
      formData.append("nickname", this.nickname);
      formData.append("about", this.about);
      await fetch("http://localhost:3000/api/v1/user/register", {
        method: "POST",
        body: formData,
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Register failed");
          }
          return response.json();
        })
        .then((data) => {
          alert(data.message);
          if (data.success) this.$router.push("/login");
        })
        .catch((error) => {
          console.error("register error: ", error);
        });
    },
  },
};
</script>

<style scoped>
.register-container {
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
  width: 200px;
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
