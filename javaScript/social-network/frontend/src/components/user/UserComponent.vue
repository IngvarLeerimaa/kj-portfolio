<template>
  <div class="user">
    <div class="profile-container" @click="profile">
      <div class="image-container">
        <img :src="user.avatar" :alt="`${user.firstname}`" />
      </div>
      <div class="name-container">{{ user.firstname }} {{ user.lastname }}</div>
    </div>
    <div class="follow-container">
      <span v-if="user.currentuser">(You)</span>
      <span v-else-if="user.pending">Pending</span>
      <button
        v-else-if="user.following"
        @click="follow(user.id, user.public, false)"
      >
        Unfollow
      </button>
      <button v-else @click="follow(user.id, user.public, true)">Follow</button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    user: {},
  },
  methods: {
    profile() {
      this.$emit("profile");
    },
    follow(userId, isPublic, following) {
      this.$emit("follow", {
        userId: userId,
        isPublic: isPublic,
        following: following,
      });
    },
  },
};
</script>

<style scoped>
.user {
  display: flex;
  border: 1px solid #888;
  border-radius: 10px;
  margin-bottom: 5px;
  padding: 10px;
  align-items: center;
}

.profile-container {
  display: flex;
  align-items: center;
  text-align: left;
  overflow: hidden;
}

.image-container {
  flex-shrink: 0;
  width: 35px;
  height: 35px;
  margin-right: 10px;
  border: 2px solid black;
  border-radius: 50%;
  overflow: hidden;
}

.image-container img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.name-container {
  margin-right: 10px;
  word-break: break-word;
}

.profile-container:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}
</style>
