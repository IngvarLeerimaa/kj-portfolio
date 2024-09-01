<template>
  <div class="group-invite">
    <div>You have been invited to join {{ group.title }}.</div>
    <div v-if="this.decision == null">
      <button @click="processInvite(true)">Accept</button>
      <button @click="processInvite(false)">Decline</button>
    </div>
    <div v-else>
      You have {{ decision ? "accepted" : "declined" }} the invitation to the
      group.
    </div>
  </div>
</template>

<script>
import { updateGroupUser } from "@/assets/fetchFunctions";
export default {
  props: {
    group: {},
  },
  data() {
    return {
      decision: null,
    };
  },
  methods: {
    async processInvite(decision) {
      try {
        await updateGroupUser({
          groupId: this.group.groupId,
          request: false,
          confirm: decision,
        });
        this.decision = decision;
        this.$emit("updateGroup", {
          groupId: this.group.groupId,
          attribute: "joined",
          value: decision,
        });
        this.$emit("updateGroup", {
          groupId: this.group.groupId,
          attribute: "invited",
          value: false,
        });
        this.$emit("done");
      } catch (error) {
        this.error = error.message;
      }
    },
  },
};
</script>

<style scoped>
.group-invite {
  max-width: 300px;
  overflow-wrap: break-word;
  background-color: rgb(200, 200, 200);
  margin: 5% auto;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 10px;
}
</style>
