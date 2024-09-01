<template>
  <div class="event-container">
    <div class="event-author">
      <div class="user-info" @click="viewProfile(groupEvent.creatorId)">
        <div class="avatar-container">
          <img :src="this.user.avatar" :alt="`${groupEvent.creatorId}`" />
        </div>
        <div class="name-container">
          {{ this.user.firstname }} {{ this.user.lastname }}
        </div>
      </div>
      <div class="date">{{ this.user.createdTime }}</div>
    </div>
    <div class="event-info">
      <div class="event-title">Event: {{ groupEvent.title }}</div>
      <div class="event-description">
        Description: {{ groupEvent.description }}
      </div>
      <div class="event-date">Date: {{ this.eventDate }}</div>
    </div>
    <div class="attend-container">
      <div v-if="groupEvent.going">You are going to this event.</div>
      <div v-else-if="groupEvent.notgoing">
        You are not going to this event.
      </div>
      <div v-else>
        <button @click="going(true)">Going</button>
        <button @click="going(false)">Not going</button>
      </div>
    </div>
  </div>
</template>

<script>
import { days, months } from "@/assets/dateConsts.js";

export default {
  name: "EventComponent",
  data() {
    return {
      user: {
        firstname: "",
        lastname: "",
        avatar: "",
        createdTime: "",
      },
      eventDate: "",
    };
  },
  props: {
    groupEvent: Object,
  },
  mounted() {
    this.getUserData();
  },
  methods: {
    going(going) {
      this.$emit("going", { going: going, eventId: this.groupEvent.eventId });
    },
    async getUserData() {
      await fetch(
        `http://localhost:3000/api/v1/user?id=${this.groupEvent.creatorId}`,
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

          var timestamp = new Date(this.groupEvent.created);
          this.user.createdTime =
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

      var ts = new Date(this.groupEvent.date);
      this.eventDate =
        days[ts.getDay()] +
        ", " +
        months[ts.getMonth()] +
        " " +
        ts.getDate() +
        ", " +
        ts.getFullYear();
    },
    viewProfile(userId) {
      this.$emit("view-profile", userId);
    },
  },
};
</script>

<style scoped>
.event-container {
  background-color: rgb(200, 200, 200);
  border: 1px solid #888;
  border-radius: 10px;
  margin-top: 50px;
  margin-left: 25%;
  margin-right: 25%;
  overflow-wrap: break-word;
}

.event-author {
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

.event-title,
.event-date {
  padding: 10px;
  border-top: 1px solid #888;
  border-bottom: 1px solid #888;
}

.event-description,
.attend-container {
  padding: 10px;
}

.user-info:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}
</style>
