<template>
  <div class="modal" v-if="show" @click="close">
    <div class="modal-wrapper" @click.stop>
      <div class="modal-content notification">
        <span class="close" @click="close">&times;</span>
        <h2>Notifications</h2>
        <div v-for="(n, index) in notifications" :key="index">
          <FollowNotification
            v-if="n.notificationType == 'follow'"
            :user="n.user"
            @updateInfo="updateInfo(index, $event)"
          />
          <InviteNotification
            v-if="n.notificationType == 'invite'"
            :group="findGroup(n.groupId)"
            @updateGroup="updateGroup"
            @done="addToRemove(index)"
          />
          <RequestNotification
            v-if="n.notificationType == 'request'"
            :group="findGroup(n.groupId)"
            :user="findUser(n.userId)"
            @updateGroup="updateGroup"
            @done="addToRemove(index)"
          />
          <EventNotification
            v-if="n.notificationType == 'event'"
            :groupEvent="n.groupEvent"
            :group="findGroup(n.groupEvent.groupId)"
            :user="findUser(n.groupEvent.creatorId)"
            @done="addToRemove(index)"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import FollowNotification from "@/components/notification/FollowNotification.vue";
import InviteNotification from "@/components/notification/InviteNotification.vue";
import RequestNotification from "@/components/notification/RequestNotification.vue";
import EventNotification from "@/components/notification/EventNotification.vue";
export default {
  name: "NotificationView",
  components: {
    FollowNotification,
    InviteNotification,
    RequestNotification,
    EventNotification,
  },
  props: {
    notifications: Array,
    users: Array,
    groups: Array,
    show: Boolean,
  },
  data() {
    return {
      remove: [],
    };
  },
  methods: {
    findGroup(id) {
      return this.groups.find((g) => g.groupId == id);
    },
    findUser(id) {
      return this.users.find((u) => u.id == id);
    },
    close() {
      this.$emit("close", this.remove);
      this.remove = [];
    },
    addToRemove(index) {
      this.remove.push(index);
    },
    updateInfo(index, payload) {
      if (payload.attribute == "pending") {
        this.addToRemove(index);
      }
      this.$emit("update", payload);
    },
    updateGroup(payload) {
      this.$emit("updateGroup", payload);
    },
  },
};
</script>

<style scoped>
.modal {
  position: fixed;
  z-index: 50;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-wrapper {
  display: flex;
  position: fixed;
  z-index: 50;
  right: 25%;
  top: 0;
  width: auto;
  min-width: 200px;
  max-height: 50vh;
}

.modal-content {
  background-color: rgb(200, 200, 200);
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 10px;
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
</style>
