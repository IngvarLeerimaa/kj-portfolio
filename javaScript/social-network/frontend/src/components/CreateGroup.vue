<template>
  <div class="modal" v-if="show">
    <div class="modal-content create">
      <span class="close" @click="close">&times;</span>
      <h2>Create post</h2>
      <form @submit.prevent="createGroup">
        <div>
          <input
            type="textarea"
            id="title"
            placeholder="Title"
            v-model="title"
          />
        </div>
        <div>
          <input
            type="textarea"
            id="description"
            placeholder="Group Description"
            v-model="description"
          />
        </div>
        <button type="submit">Create Group</button>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "CreateGroup",
  props: {
    show: Boolean,
  },
  data() {
    return {
      title: "",
      description: "",
    };
  },
  methods: {
    close() {
      this.$emit("close");
      this.title = "";
      this.description = "";
    },

    async createGroup() {
      var formData = new FormData();
      formData.append("title", this.title);
      formData.append("description", this.description);

      await fetch("http://localhost:3000/api/v1/group/create", {
        method: "POST",
        credentials: "include",
        body: formData,
      })
        .then(async (response) => {
          if (!response.ok) {
            throw new Error("Creating group failed");
          }
          alert("Group created successfully");
          this.close();
          this.$emit("group");
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
