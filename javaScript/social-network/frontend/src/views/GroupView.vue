<template>
  <div class="modal" v-if="show">
    <div class="modal-content group">
      <span class="close" @click="close">&times;</span>
      <img
        class="group-image"
        src="http://localhost:3000/images/group.png"
        :alt="`${group.groupId}`"
      />
      <h2>{{ group.title }}</h2>
      <h4>{{ group.description }}</h4>
      <div class="not-joined" v-if="!group.joined">
        <span v-if="group.requested"
          >Group admin has to approve your request
        </span>
        <button v-else-if="group.invited" @click="updateUserStatus(true)">
          Accept invite
        </button>
        <button v-else @click="joinGroup(false)">Join</button>
        <span>to see the content.</span>
      </div>
      <div class="joined" v-else>
        <button @click="toggleCreatePost">Create Post</button>
        <button @click="toggleCreateEvent">Create Event</button>
        <button @click="updateUserStatus(false)">Leave Group</button>
        <GroupPost
          :show="postForm"
          :groupId="group.groupId"
          @post="updateData"
          @close="toggleCreatePost"
        />
        <CreateEvent
          :show="eventForm"
          :groupId="group.groupId"
          @post="updateData"
          @close="toggleCreateEvent"
        />
        <div v-for="(item, index) in combinedFeed" :key="index">
          <PostComponent
            v-if="item.hasOwnProperty('postId')"
            :key="index"
            :post="item"
            @view-profile="viewProfile"
          />
          <EventComponent
            v-if="item.hasOwnProperty('eventId')"
            :key="index"
            :groupEvent="item"
            @going="goingToEvent"
            @view-profile="viewProfile"
          />
        </div>
        <div v-if="combinedFeed.length == 0">No posts/events found.</div>
        <div class="modal-content members" v-show="group.joined">
          <div>
            <h2>Members:</h2>
            <ul>
              <li v-for="member in members" :key="member.id">
                <label class="member">
                  <img :src="member.avatar" :alt="`${member.id}`" />
                  {{ member.firstname }} {{ member.lastname }}</label
                >
              </li>
            </ul>
          </div>
        </div>
        <div class="modal-content invite" v-show="group.joined">
          <div>
            <h2>Invite Users:</h2>
            <ul>
              <li v-for="user in notMembers" :key="user.id">
                <label class="user">
                  <img :src="user.avatar" :alt="`${user.id}`" />
                  {{ user.firstname }} {{ user.lastname }}</label
                >
                <span class="invited" v-if="user.invited">Invited</span>
                <button v-else @click="joinGroup(true, user.id)">Invite</button>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import GroupPost from "@/components/GroupPost.vue";
import PostComponent from "@/components/PostComponent.vue";
import CreateEvent from "@/components/CreateEvent.vue";
import EventComponent from "@/components/event/EventComponent.vue";
import {
  attendEvent,
  createGroupUser,
  getGroup,
  updateGroupUser,
} from "@/assets/fetchFunctions";

export default {
  components: {
    GroupPost,
    CreateEvent,
    PostComponent,
    EventComponent,
  },
  props: {
    show: Boolean,
    group: Object,
    users: Array,
  },
  data() {
    return {
      posts: [],
      events: [],
      members: [],
      notMembers: [],
      visible: false,
      postForm: false,
      eventForm: false,
    };
  },
  watch: {
    async show(newValue) {
      if (newValue) {
        this.updateData();
      }
    },
  },
  computed: {
    combinedFeed() {
      const postsAndEvents = [...this.posts, ...this.events];
      return postsAndEvents.sort((a, b) => {
        return new Date(b.created) - new Date(a.created);
      });
    },
  },
  methods: {
    async joinGroup(invite, userId) {
      try {
        await createGroupUser({
          groupId: this.group.groupId,
          userId: invite ? userId : 0,
          invite: invite,
        });
        if (invite) {
          this.notMembers.find((u) => u.id == userId).invited = true;
        } else {
          this.$emit("update", {
            groupId: this.group.groupId,
            attribute: "requested",
            value: true,
          });
        }
      } catch (error) {
        this.error = error.message;
      }
    },
    async updateUserStatus(decision) {
      try {
        await updateGroupUser({
          groupId: this.group.groupId,
          request: false,
          confirm: decision,
        });
        if (decision) {
          this.$emit("update", {
            groupId: this.group.groupId,
            attribute: "joined",
            value: true,
          });
          this.updateData();
        } else {
          this.$emit("update", {
            groupId: this.group.groupId,
            attribute: "joined",
            value: false,
          });
        }
        this.$emit("update", {
          groupId: this.group.groupId,
          attribute: "requested",
          value: false,
        });
      } catch (error) {
        this.error = error.message;
      }
    },
    async updateData() {
      if (this.group.joined) {
        try {
          const groupData = await getGroup(this.group.groupId);
          this.posts = groupData.posts;
          this.events = groupData.events;
          this.members = this.users.filter((user) =>
            groupData.members.includes(user.id)
          );
          this.notMembers = this.users.filter(
            (user) => !groupData.members.includes(user.id)
          );
          this.notMembers.forEach(
            (user) => (user.invited = groupData.invited.includes(user.id))
          );
        } catch (error) {
          this.error = error.message;
        }
      }
    },
    toggleCreatePost() {
      this.postForm = !this.postForm;
    },
    toggleCreateEvent() {
      this.eventForm = !this.eventForm;
    },
    viewProfile(userId) {
      this.$emit("view-profile", userId);
    },
    async goingToEvent(payload) {
      try {
        await attendEvent(payload);
        const e = this.events.find((event) => event.eventId == payload.eventId);
        payload.going ? (e.going = true) : (e.notgoing = true);
      } catch (error) {
        this.error = error.message;
      }
    },
    close() {
      this.$emit("close");
      this.visible = false;
      this.posts = [];
      this.events = [];
      this.members = [];
      this.notMembers = [];
    },
  },
};
</script>

<style scoped>
.modal {
  display: flex;
  position: fixed;
  z-index: 5;
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

.members {
  width: 15%;
  top: 5%;
  left: 3%;
  max-height: 50vh;
  position: absolute;
  overflow: auto;
}
.invite {
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
.group-image {
  width: 50px;
}
img {
  width: 35px;
}

ul {
  list-style-type: none;
  padding: 0;
}
</style>
