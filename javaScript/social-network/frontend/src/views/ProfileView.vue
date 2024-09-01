<template>
  <div class="modal" v-if="show">
    <div class="modal-content user">
      <span class="close" @click="close">&times;</span>
      <img :src="user.avatar" :alt="`${user.id}`" />
      <h2>{{ user.firstname }} {{ user.lastname }}</h2>
      <CurrentUser v-if="user.currentuser" :user="user" />
      <PublicUser
        v-else-if="user.public"
        :user="user"
        @updateFollow="updateFollow"
      />
      <PrivateUser v-else :user="user" @updateFollow="updateFollow" />
      <div v-if="posts.length > 0">
        <PostComponent
          v-for="(post, index) in posts"
          :key="index"
          :post="post"
        />
      </div>
      <div v-else>
        No posts by {{ user.firstname }} {{ user.lastname }} found.
      </div>
      <div class="modal-content followers" v-show="this.visible">
        <div>
          <h2>Followers:</h2>
          <ul>
            <li v-for="follower in followers" :key="follower.id">
              <label class="follower">
                <img :src="follower.avatar" :alt="`${follower.id}`" />
                {{ follower.firstname }} {{ follower.lastname }}</label
              >
            </li>
          </ul>
        </div>
      </div>
      <div class="modal-content following" v-show="this.visible">
        <div>
          <h2>Following:</h2>
          <ul>
            <li v-for="follower in following" :key="follower.id">
              <label class="follower">
                <img :src="follower.avatar" :alt="`${follower.id}`" />
                {{ follower.firstname }} {{ follower.lastname }}</label
              >
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import CurrentUser from "@/components/CurrentUser.vue";
import PublicUser from "@/components/PublicUser.vue";
import PrivateUser from "@/components/PrivateUser.vue";
import PostComponent from "@/components/PostComponent.vue";
import {
  getUser,
  followUser,
  getUserPosts,
  getUserFollowers,
} from "@/assets/fetchFunctions";

export default {
  components: {
    CurrentUser,
    PublicUser,
    PrivateUser,
    PostComponent,
  },
  props: {
    show: Boolean,
    userId: Number,
  },
  data() {
    return {
      user: {},
      posts: [],
      followers: [],
      following: [],
      visible: true,
    };
  },
  watch: {
    async show(newValue) {
      if (newValue) {
        try {
          const user = await getUser(this.userId);
          this.user = user;
          if (this.user.email.length > 0) {
            const posts = await getUserPosts(this.user.id);
            this.posts = posts;
            const followData = await getUserFollowers(this.user.id);
            this.followers = followData.followers;
            this.following = followData.following;
            this.visible = true;
          }
        } catch (error) {
          this.error = error.message;
        }
      }
    },
  },
  methods: {
    async updateFollow(payload) {
      try {
        await followUser(this.userId, payload.following);
        this.user.following = payload.following;
        this.user.pending = payload.pending;
        this.$emit("update", {
          userId: this.userId,
          attribute: "following",
          value: payload.following,
        });
        this.$emit("update", {
          userId: this.userId,
          attribute: "pending",
          value: payload.pending,
        });
      } catch (error) {
        this.error = error.message;
      }
    },
    close() {
      this.$emit("close");
      this.visible = false;
      this.user = {};
      this.posts = [];
    },
  },
};
</script>

<style scoped>
.modal {
  display: flex;
  position: fixed;
  z-index: 10;
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
  padding: 20px;
  border: 1px solid #888;
  border-radius: 10px;
  width: 50%;
  overflow-y: scroll;
}

.followers {
  width: 15%;
  top: 5%;
  left: 3%;
  max-height: 50vh;
  position: absolute;
  overflow: auto;
}
.following {
  width: 15%;
  top: 5%;
  left: 79%;
  max-height: 50vh;
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
  width: 50px;
}

ul {
  list-style-type: none;
  padding: 0;
}
</style>
