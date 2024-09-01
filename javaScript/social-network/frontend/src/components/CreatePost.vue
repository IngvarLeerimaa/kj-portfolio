<template>
  <div class="modal" v-if="show">
    <div class="modal-content create">
      <span class="close" @click="close">&times;</span>
      <h2>Create post</h2>
      <form @submit.prevent="createPost">
        <div class="privacy-selectors">
          <label>
            <input type="radio" v-model="privacy" value="0" />
            Public:
          </label>
          <label>
            <input type="radio" v-model="privacy" value="1" />
            Followers:
          </label>
          <label>
            <input type="radio" v-model="privacy" value="2" />
            Specific followers: </label
          ><br />
        </div>
        <div>
          <input
            type="textarea"
            id="text"
            placeholder="Write your post"
            v-model="text"
          />
        </div>
        <div>
          <label for="image">Add image:</label>
          <input
            type="file"
            id="image"
            accept="image/*"
            @change="uploadImage"
          />
        </div>
        <button type="submit">Create post</button>
      </form>
    </div>
    <div class="modal-content followers" v-show="privacy === '2'">
      <div>
        <h2>Select Followers:</h2>
        <ul>
          <li v-for="follower in filtredUsers" :key="follower.id">
            <label class="follower">
              <img :src="follower.avatar" :alt="`${follower.id}`" />
              {{ follower.firstname }} {{ follower.lastname
              }}<input
                type="checkbox"
                :value="follower.id"
                v-model="checkedFollowers"
            /></label>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "CreatePost",
  props: {
    show: Boolean,
    users: Array,
  },
  data() {
    return {
      privacy: "0",
      text: "",
      image: null,
      followers: null,
      checkedFollowers: [],
    };
  },
  computed: {
    filtredUsers() {
      return this.users
        .slice()
        .filter((user) => user.follower || (user.following && !user.pending));
    },
  },
  methods: {
    close() {
      this.$emit("close");
      this.privacy = "0";
      this.text = "";
      this.image = null;
      this.followers = null;
      this.checkedFollowers = [];
    },

    uploadImage(event) {
      this.image = event.target.files[0];
    },

    async createPost() {
      if (this.text == "" && this.image == null) {
        alert("Please write a message or insert an image.");
        return;
      }

      if (this.privacy == "2" && this.checkedFollowers.length == 0) {
        alert("Please select at least one follower");
        return;
      }
      var formData = new FormData();
      formData.append("privacy", this.privacy);
      formData.append("followers", this.checkedFollowers);
      formData.append("text", this.text);
      formData.append("image", this.image);

      await fetch("http://localhost:3000/api/v1/post/create", {
        method: "POST",
        credentials: "include",
        body: formData,
      })
        .then(async (response) => {
          if (!response.ok) {
            throw new Error("Creating post failed");
          }
          alert("Post created successfully");
          this.$emit("post");
          this.close();
        })
        .catch((error) => {
          console.error(error);
        });
    },
  },
};
</script>

<style scoped>
.modal {
  display: flex;
  position: fixed;
  z-index: 3;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(130, 130, 130, 0.466);
  backdrop-filter: blur(3px);
}

.modal-content {
  background-color: rgb(200, 200, 200);
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 10px;
}

.create {
  width: 43%;
}

.followers {
  width: 20%;
  left: 75%;
  position: absolute;
  overflow: auto;
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}

img {
  width: 25px;
}

ul {
  list-style-type: none;
  padding: 0;
}

.follower:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}
</style>
