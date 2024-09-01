<template>
  <div class="follow-info">
    <div class="follow-pending" v-if="user.pending">Follow request pending</div>
    <button
      class="unfollow-btn"
      v-else-if="!user.pending && user.following"
      @click="this.unfollow()"
    >
      Unfollow
    </button>
    <button class="follow-btn" v-else @click="this.follow()">Follow</button>
  </div>
  <div class="user-info follow" v-if="user.following && !user.pending">
    <div>Email: {{ user.email }}</div>
    <div>Date of Birth: {{ user.dateofbirth }}</div>
    <div v-if="user.nickname != ''">Nickname: {{ user.nickname }}</div>
    <div v-if="user.about != ''">About me: {{ user.about }}</div>
  </div>
  <div class="user-info not-follow" v-else>
    This is a private user, follow to see content.
  </div>
</template>

<script>
export default {
  name: "PrivateUser",
  props: {
    user: {},
  },
  emits: ["updateFollow"],
  methods: {
    follow() {
      this.$emit("updateFollow", { following: true, pending: true });
    },
    unfollow() {
      this.$emit("updateFollow", { following: false, pending: false });
    },
  },
};
</script>
