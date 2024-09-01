<template>
  <div class="event-request">
    <div>
      <img :src="user.avatar" :alt="`${user.firstname}`" />
      {{ user.firstname }} {{ user.lastname }} in {{ group.title }} created an
      event:
    </div>
    <div class="event-title">Title: {{ groupEvent.title }}</div>
    <div class="event-description">
      Description: {{ groupEvent.description }}
    </div>
    <div class="event-date">Date: {{ groupEvent.date }}</div>
    <div v-if="this.decision == null">
      <button @click="processRequest(true)">Going</button>
      <button @click="processRequest(false)">Not going</button>
    </div>
    <div v-else>
      You are {{ decision ? "going" : "not going" }} to this event.
    </div>
  </div>
</template>

<script>
import { attendEvent } from "@/assets/fetchFunctions";
export default {
  props: {
    groupEvent: {},
    group: {},
    user: {},
  },
  data() {
    return {
      decision: null,
    };
  },
  methods: {
    async processRequest(decision) {
      try {
        await attendEvent({
          going: decision,
          eventId: this.groupEvent.eventId,
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
.event-request {
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
