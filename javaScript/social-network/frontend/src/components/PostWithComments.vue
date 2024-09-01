<template>
  <div class="modal" v-if="show">
    <div class="modal-content post">
      <span class="close" @click="close">&times;</span>
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
        <div class="create-comment-toggle" @click="showCreateComment()">
          Add comment
        </div>
        <CreateComment
          :show="createCommentVisible"
          :postId="post.postId"
          @reload="getComments"
          @close="closeCreateComment"
        />
      </div>
      <div>
        <CommentComponent
          v-for="(comment, index) in this.comments"
          :key="index"
          :comment="comment"
          @view-profile="viewProfile"
        />
      </div>
    </div>
  </div>
</template>

<script>
import CreateComment from "@/components/CreateComment.vue";
import CommentComponent from "@/components/CommentComponent.vue";

export default {
  name: "PostWithComments",
  components: {
    CreateComment,
    CommentComponent,
  },
  props: {
    show: Boolean,
    post: Object,
    user: Object,
  },
  data() {
    return {
      comments: null,
      createCommentVisible: false,
    };
  },
  watch: {
    show(newVal) {
      if (newVal) {
        this.getComments();
      }
    },
  },
  methods: {
    close() {
      this.$emit("close");
      this.comments = null;
    },
    showCreateComment() {
      this.createCommentVisible = true;
    },
    closeCreateComment() {
      this.createCommentVisible = false;
    },
    async getComments() {
      await fetch(
        `http://localhost:3000/api/v1/comments?id=${this.post.postId}`,
        {
          method: "GET",
          credentials: "include",
        }
      )
        .then((response) => {
          if (!response.ok) {
            throw new Error("Failed to get posts");
          }
          return response.json();
        })
        .then((data) => {
          this.comments = data.comments;
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
  margin: 5% auto;
  border: 1px solid #888;
  border-radius: 10px;
  width: 50%;
  overflow-wrap: break-word;
  overflow-y: auto;
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

.create-comment-toggle {
  padding: 5px;
}

.user-info:hover,
.create-comment-toggle:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}

.close {
  color: #aaa;
  float: right;
  margin-right: 15px;
  margin-top: 10px;
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
