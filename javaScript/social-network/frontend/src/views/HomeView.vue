<template>
  <div class="home">
    <nav class="nav-container">
      <router-link @click="scrollToTop" to="/">Social Network</router-link> |
      <button @click="openProfile(this.$userId)">Profile</button> |
      <button @click="showModal('createPost')">Create Post</button> |
      <button @click="showModal('createGroup')">Create Group</button> |
      <button @click="logout">Logout</button>
      <span class="notification-container" @click="showModal('notifications')">
        <img
          class="notification-img"
          src="http://localhost:3000/images/notification.png"
          alt="Notification"
        />
        <div v-if="newNotification" class="notification-badge"></div>
      </span>
    </nav>
    <div class="content">
      <NotificationView
        :show="modals['notifications']"
        :users="this.users"
        :groups="this.groups"
        :notifications="this.notifications"
        @update="updateUser"
        @updateGroup="updateGroup"
        @close="closeModal('notifications', $event)"
      />
      <ProfileView
        :show="modals['profile']"
        :userId="this.userId"
        @update="updateUser"
        @close="closeModal('profile')"
      />
      <CreatePost
        :show="modals['createPost']"
        :users="this.users"
        @post="reloadPosts"
        @close="closeModal('createPost')"
      />
      <CreateGroup
        :show="modals['createGroup']"
        @group="reloadGroups"
        @close="closeModal('createGroup')"
      />
      <GroupView
        :show="modals['groupView']"
        :users="this.users"
        :group="this.group"
        @view-profile="openProfile"
        @update="updateGroup"
        @close="closeModal('groupView')"
      />
      <div class="left">
        <div class="left-top">
          <UserList
            :users="this.users"
            @view-profile="openProfile"
            @update="updateUser"
          />
        </div>
        <div class="left-bottom">
          <GroupList :groups="groups" @open="openGroup" />
        </div>
      </div>
      <div class="middle">
        <PostComponent
          v-for="(post, index) in posts"
          :key="index"
          :post="post"
          @view-profile="openProfile"
        />
        <span>You have reached the end.</span>
      </div>
      <div class="right">
        <ChatView
          v-if="this.users.length > 0"
          :users="this.users"
          :groups="this.groups"
          :onlineUsers="this.onlineUsers"
          :newMessage="newMessage"
        />
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import {
  getUsers,
  getPosts,
  getNotifications,
  getGroups,
} from "@/assets/fetchFunctions";
import WSConnection from "@/assets/websocket";
import ProfileView from "@/views/ProfileView.vue";
import NotificationView from "@/views/NotificationView.vue";
import ChatView from "@/views/ChatView.vue";
import GroupView from "@/views/GroupView.vue";
import UserList from "@/components/UserList.vue";
import GroupList from "@/components/GroupList.vue";
import CreatePost from "@/components/CreatePost.vue";
import PostComponent from "@/components/PostComponent.vue";
import CreateGroup from "@/components/CreateGroup.vue";
import debounce from "lodash/debounce";

export default {
  name: "HomeView",
  components: {
    ProfileView,
    NotificationView,
    UserList,
    ChatView,
    GroupView,
    GroupList,
    CreatePost,
    PostComponent,
    CreateGroup,
  },
  data() {
    return {
      userId: 0,
      group: {},
      users: [],
      onlineUsers: {},
      posts: [],
      groups: [],
      notifications: [],
      newMessage: {},
      fetching: false,
      newNotification: false,
      offset: 10,
      modals: {
        profile: false,
        createPost: false,
        createGroup: false,
        notifications: false,
        groupView: false,
      },
    };
  },
  async mounted() {
    try {
      const users = await getUsers();
      this.users = users;
      const posts = await getPosts(0);
      this.posts = posts;
      const groups = await getGroups();
      this.groups = groups;
      const notifications = await getNotifications();
      this.notifications = notifications;
      this.newNotification = this.notifications.length > 0;
    } catch (error) {
      this.error = error.message;
    }
    this.handleDebouncedScroll = debounce(this.handleScroll, 100);
    window.addEventListener("scroll", this.handleDebouncedScroll);
    WSConnection.connect(this.$userId);
    WSConnection.ws.onmessage = this.handleIncomingMessage;
  },
  methods: {
    async handleScroll() {
      if (
        !this.fetching &&
        document.documentElement.scrollTop +
          document.documentElement.clientHeight >=
          document.documentElement.scrollHeight
      ) {
        this.fetching = true;
        const posts = await getPosts(this.offset);
        this.posts = this.posts.concat(posts);
        if (posts.length == 10) {
          this.offset += 10;
          this.fetching = false;
        }
      }
    },
    showModal(modalName) {
      if (modalName == "notifications") {
        this.newNotification = false;
      }
      this.modals[modalName] = true;
    },
    closeModal(modalName, remove) {
      this.modals[modalName] = false;
      if (modalName == "notifications") {
        for (let i = 0; i < remove.length; i++) {
          this.notifications.splice(remove[i], 1);
        }
      }
    },
    openProfile(userId) {
      this.userId = userId;
      this.showModal("profile");
    },
    openGroup(groupId) {
      this.group = this.groups.find((g) => g.groupId == groupId);
      this.showModal("groupView");
    },
    updateUser(payload) {
      var i = this.users.findIndex((user) => user.id == payload.userId);
      this.users[i][payload.attribute] = payload.value;
    },
    updateGroup(payload) {
      var i = this.groups.findIndex(
        (group) => group.groupId == payload.groupId
      );
      this.groups[i][payload.attribute] = payload.value;
    },
    async reloadPosts() {
      this.posts = await getPosts(0);
      this.offset = 10;
    },
    async reloadGroups() {
      this.groups = await getGroups();
    },
    handleIncomingMessage(event) {
      let data = JSON.parse(event.data);
      switch (data.type) {
        case "newMessage": {
          data.message.user = true;
          this.newMessage = data.message;
          break;
        }
        case "newGroupMessage": {
          data.message.user = false;
          this.newMessage = data.message;
          break;
        }
        case "online": {
          this.onlineUsers[data.userId] = data.online;
          break;
        }
        case "follow": {
          var i = this.users.findIndex((user) => user.id == data.userId);
          if (data.decision) {
            if (data.request) {
              var cIndex = this.users.findIndex(
                (user) => user.id == this.$userId
              );
              if (this.users[cIndex].public) {
                this.users[i].follower = true;
              } else {
                this.notifications.unshift({
                  notificationType: "follow",
                  user: this.users[i],
                });
                this.newNotification = true;
              }
            } else {
              this.users[i].pending = false;
              this.users[i].follower = true;
            }
          } else {
            if (data.request) {
              this.users[i].follower = false;
            } else {
              this.users[i].following = false;
              this.users[i].pending = false;
            }
          }
          break;
        }
        case "notification": {
          this.notifications.unshift(data.notification);
          this.newNotification = true;
          break;
        }
        default: {
          console.error("invalid event type: " + data.type);
        }
      }
    },
    scrollToTop() {
      window.scrollTo({
        top: 0,
        behavior: "smooth",
      });
    },
    async logout() {
      await fetch("http://localhost:3000/api/v1/user/logout", {
        method: "DELETE",
        credentials: "include",
      });
      if (WSConnection.ws != null) WSConnection.ws.close();
      document.cookie =
        `sessionID=; Max-Age=0; path=/; domain=` + location.hostname;
      this.$router.push("/login");
    },
  },
};
</script>

<style scoped>
.home {
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.content {
  display: flex;
  flex: 1;
  position: relative;
}

.left,
.right {
  position: fixed;
  top: 0;
  bottom: 0;
}

.left {
  display: flex;
  flex-direction: column;
  left: 0;
  width: 20%;
  margin-top: 50px;
  margin-left: 20px;
  margin-bottom: 20px;
  padding: 10px;
}

.left-top,
.left-bottom {
  height: 50%;
  overflow-y: auto;
  background-color: rgb(200, 200, 200);
  border: 1px solid #888;
  border-radius: 10px;
}

.left-top {
  margin-bottom: 5px;
}

.left-bottom {
  margin-top: 5px;
}

.middle {
  flex: 1;
  overflow-y: auto;
}

.right {
  display: flex;
  flex-direction: column;
  right: 0;
  width: 20%;
  margin-top: 60px;
  margin-right: 20px;
  margin-bottom: 20px;
}

.nav-container {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 25px;
  background-color: rgb(160, 160, 160);
  border-bottom: 1px solid;
  padding: 10px;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
}

.nav-container * {
  pointer-events: auto;
}

.nav-container button,
.nav-container img {
  filter: none;
}

.notification-container {
  position: relative;
  margin-left: 20px;
}

.notification-img {
  width: 25px;
}

.notification-img:hover {
  filter: brightness(150%);
  cursor: pointer;
}

.notification-badge {
  position: absolute;
  top: 0;
  right: 0;
  width: 10px;
  height: 10px;
  background-color: red;
  border-radius: 50%;
}
</style>
