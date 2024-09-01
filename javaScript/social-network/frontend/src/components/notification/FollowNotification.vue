<template>
  <div class="follow-request">
    <div>
      <img :src="user.avatar" :alt="`${user.firstname}`" />
      {{ user.firstname }} {{ user.lastname }} sent you a follow request.
    </div>
    <div v-if="this.decision == null">
      <button @click="processRequest(true)">Accept</button>
      <button @click="processRequest(false)">Reject</button>
    </div>
    <div v-else>You {{ decision ? "accepted" : "rejected" }} the request.</div>
  </div>
</template>

<script>
import { updateFollow } from "@/assets/fetchFunctions";
export default {
  name: "FollowComponent",
  props: {
    user: {},
  },
  data() {
    return {
      decision: null,
    };
  },
  methods: {
    async processRequest(decision) {
      await updateFollow(this.user.id, decision);
      this.decision = decision;
      if (decision) {
        this.$emit("updateInfo", {
          userId: this.user.id,
          attribute: "follower",
          value: true,
        });
      }
      this.$emit("updateInfo", {
        userId: this.user.id,
        attribute: "pending",
        value: false,
      });
    },
  },
};
</script>

<style scoped>
.follow-request {
  background-color: rgb(200, 200, 200);
  margin: 5% auto;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 10px;
}
img {
  width: 25px;
}
</style>
