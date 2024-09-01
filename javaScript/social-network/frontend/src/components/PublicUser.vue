<template>
  <div class="user-info">
    <div class="follow-info">
      <button
        class="unfollow-btn"
        v-if="user.following"
        @click="this.unfollow()"
      >
        Unfollow
      </button>
      <button class="follow-btn" v-if="!user.following" @click="this.follow()">
        Follow
      </button>
    </div>
    <div>Email: {{ user.email }}</div>
    <div>Date of Birth: {{ user.dateofbirth }}</div>
    <div v-if="user.nickname != ''">Nickname: {{ user.nickname }}</div>
    <div v-if="user.about != ''">About me: {{ user.about }}</div>
  </div>
</template>

<script>
export default {
  name: "PublicUser",
  props: {
    user: {},
  },
  emits: ["updateFollow"],
  methods: {
    follow() {
      this.$emit("updateFollow", { following: true, pending: false });
    },
    unfollow() {
      this.$emit("updateFollow", { following: false, pending: false });
    },
  },
};
</script>
