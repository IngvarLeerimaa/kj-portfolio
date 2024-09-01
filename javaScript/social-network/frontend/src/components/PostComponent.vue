<template>
  <div class="post-container">
    <div class="post-author">
      <div class="user-info" @click="viewProfile(post.userId)">
        <div class="avatar-container">
          <img :src="user.avatar" :alt="`${post.userId}`" />
        </div>
        <div class="name-container">
          {{ user.firstname }} {{ user.lastname }}
        </div>
      </div>
      <div class="date">{{ this.user.dateTime }}</div>
    </div>
    <div class="content-container">
      <div class="text-container" v-if="post.text != ''">
        <span class="post-text">{{ post.text }}</span>
      </div>
      <div class="image-container" v-if="post.image != ''">
        <img class="post-image" :src="post.image" :alt="`${post.userId}`" />
      </div>
    </div>
    <div>
      <div class="comment-toggle" @click="showComments()">Comments</div>
      <PostWithComments
        :show="commentsVisible"
        :post="post"
        :user="user"
        @view-profile="viewProfile"
        @close="closeComments"
      />
    </div>
  </div>
</template>

<script>
import PostWithComments from "@/components/PostWithComments.vue";
import { days, months } from "@/assets/dateConsts.js";

export default {
  name: "PostComponent",
  components: {
    PostWithComments,
  },
  data() {
    return {
      user: {
        firstname: "",
        lastname: "",
        avatar: "",
        dateTime: "",
      },
      commentsVisible: false,
    };
  },
  props: {
    post: Object,
  },
  mounted() {
    this.getUserData();
  },
  methods: {
    async getUserData() {
      await fetch(`http://localhost:3000/api/v1/user?id=${this.post.userId}`, {
        method: "GET",
        credentials: "include",
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Failed to get user");
          }
          return response.json();
        })
        .then((data) => {
          this.user.firstname = data.user.firstname;
          this.user.lastname = data.user.lastname;
          this.user.avatar = data.user.avatar;

          var timestamp = new Date(this.post.created);
          this.user.dateTime =
            days[timestamp.getDay()] +
            ", " +
            months[timestamp.getMonth()] +
            " " +
            timestamp.getDate() +
            ", " +
            timestamp.getFullYear() +
            " at " +
            (timestamp.getHours() < 10 ? "0" : "") +
            timestamp.getHours() +
            "." +
            (timestamp.getMinutes() < 10 ? "0" : "") +
            timestamp.getMinutes();
        })
        .catch((error) => {
          console.error(error);
        });
    },
    showComments() {
      this.commentsVisible = true;
    },
    closeComments() {
      this.commentsVisible = false;
    },
    viewProfile(userId) {
      this.$emit("view-profile", userId);
    },
  },
};
</script>

<style scoped>
.post-container {
  background-color: rgb(200, 200, 200);
  border: 1px solid #888;
  border-radius: 10px;
  margin-top: 50px;
  margin-left: 25%;
  margin-right: 25%;
  overflow-wrap: break-word;
}

.post-author {
  display: flex;
  flex-direction: column;
  text-align: left;
  padding: 10px;
}

.user-info {
  display: flex;
  align-items: center;
}

.avatar-container {
  flex-shrink: 0;
  width: 35px;
  height: 35px;
  margin-right: 10px;
  border: 2px solid black;
  border-radius: 50%;
  overflow: hidden;
}

.avatar-container img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.date {
  margin-left: 48px;
  font-size: smaller;
}

.content-container {
  padding: 10px;
  border-top: 1px solid #888;
  border-bottom: 1px solid #888;
}

.text-container {
  text-align: left;
}

.post-image {
  max-width: 90%;
  max-height: 50vh;
}

.comment-toggle {
  text-align: right;
  padding: 5px;
}

.user-info:hover,
.comment-toggle:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}
</style>
