<template>
  <div class="modal" v-if="show">
    <div class="modal-content create">
      <span class="close" @click="close">&times;</span>
      <h2>Create comment</h2>
      <form @submit.prevent="createComment">
        <div>
          <input
            type="textarea"
            id="text"
            placeholder="Write your comment"
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
        <button type="submit">Create comment</button>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "CreateComment",
  props: {
    show: Boolean,
    postId: Number,
  },
  data() {
    return {
      text: "",
      image: null,
    };
  },
  methods: {
    close() {
      this.$emit("close");
      this.text = "";
      this.image = null;
    },

    uploadImage(event) {
      this.image = event.target.files[0];
    },

    async createComment() {
      if (this.text == "" && this.image == null) {
        alert("Please write a message or insert an image.");
        return;
      }

      var formData = new FormData();
      formData.append("text", this.text);
      formData.append("image", this.image);
      formData.append("postId", this.postId);

      await fetch("http://localhost:3000/api/v1/comment/create", {
        method: "POST",
        credentials: "include",
        body: formData,
      })
        .then(async (response) => {
          if (!response.ok) {
            throw new Error("Creating comment failed");
          }
          alert("Comment created successfully");
          this.$emit("reload");
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
</style>
