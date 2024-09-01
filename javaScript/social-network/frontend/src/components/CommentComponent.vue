<template>
  <div class="comment-container">
    <div class="comment-author">
      <div class="user-info" @click="viewProfile(comment.userId)">
        <div class="avatar-container">
          <img
            class="avatar-image"
            :src="this.user.avatar"
            :alt="`${comment.userId}`"
          />
        </div>
        <div class="name-container">
          {{ this.user.firstname }} {{ this.user.lastname }}
        </div>
      </div>
      <div class="date">{{ this.user.dateTime }}</div>
    </div>
    <div class="content-container">
      <div class="text-container" v-if="comment.text != ''">
        <span class="comment-text">{{ comment.text }}</span>
      </div>
      <div class="image-container" v-if="comment.image != ''">
        <img
          class="comment-image"
          :src="comment.image"
          :alt="`${comment.userId}`"
        />
      </div>
    </div>
  </div>
</template>

<script>
import { days, months } from "@/assets/dateConsts.js";

export default {
  name: "CommentComponent",
  data() {
    return {
      user: {
        firstname: "",
        lastname: "",
        avatar: "",
        dateTime: "",
      },
    };
  },
  props: {
    comment: Object,
  },
  mounted() {
    this.getUserData();
  },
  methods: {
    async getUserData() {
      await fetch(
        `http://localhost:3000/api/v1/user?id=${this.comment.userId}`,
        {
          method: "GET",
          credentials: "include",
        }
      )
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

          var timestamp = new Date(this.comment.created);
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
    viewProfile(userId) {
      this.$emit("view-profile", userId);
    },
  },
};
</script>

<style scoped>
.comment-container {
  background-color: rgb(200, 200, 200);
  border: 1px solid #888;
  border-radius: 10px;
  margin-top: 20px;
  margin-left: 10%;
  margin-right: 10%;
}
.comment-author {
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
}

.text-container {
  text-align: left;
}

.comment-image {
  max-width: 90%;
  max-height: 50vh;
}

.user-info:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}
</style>
