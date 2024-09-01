<template>
  <div class="group-request">
    <div>
      <img :src="user.avatar" :alt="`${user.firstname}`" />
      {{ user.firstname }} {{ user.lastname }} requested to join
      {{ group.title }}.
    </div>
    <div v-if="this.decision == null">
      <button @click="processRequest(true)">Accept</button>
      <button @click="processRequest(false)">Decline</button>
    </div>
    <div v-else>
      You have {{ decision ? "accepted" : "declined" }} the request to the
      group.
    </div>
  </div>
</template>

<script>
import { updateGroupUser } from "@/assets/fetchFunctions";
export default {
  props: {
    user: {},
    group: {},
  },
  data() {
    return {
      decision: null,
    };
  },
  methods: {
    async processRequest(decision) {
      try {
        await updateGroupUser({
          groupId: this.group.groupId,
          userId: this.user.id,
          request: true,
          confirm: decision,
        });
        this.decision = decision;
        this.$emit("done");
      } catch (error) {
        this.error = error.message;
      }
    },
  },
};
</script>

<style scoped>
.group-request {
  max-width: 300px;
  overflow-wrap: break-word;
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
