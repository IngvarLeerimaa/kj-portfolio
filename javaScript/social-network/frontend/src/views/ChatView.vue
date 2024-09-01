<template>
  <div class="chat-users">
    <h2>Chat:</h2>
    <div v-for="(chat, index) in sortedChats" :key="index">
      <div v-if="chat.id != undefined">
        <ChatUser
          :user="chat"
          :isOnline="isOnline(chat.id)"
          :lastMessages="this.lastMessages"
          @chat="changeChat"
        />
      </div>
      <div v-else-if="chat.groupId != undefined">
        <GroupChat
          :group="chat"
          :lastMessages="this.lastGroupMessages"
          @chat="changeGroupChat"
        />
      </div>
    </div>
  </div>
  <div class="chat-window">
    <ChatWindow
      v-if="chatType == 'user'"
      :user="user"
      :messages="messages"
      :isOnline="isOnline(user.id)"
      :messagesAvailable="!fetching"
      @newMessage="appendMessage"
      @moreMessages="getMoreMessages($event, true)"
    />
    <GroupChatWindow
      v-if="chatType == 'group'"
      :group="group"
      :users="users"
      :messages="messages"
      :messagesAvailable="!fetching"
      @newMessage="appendMessage"
      @moreMessages="getMoreMessages"
    />
  </div>
</template>

<script>
import ChatUser from "@/components/chat/ChatUser.vue";
import ChatWindow from "@/components/chat/ChatWindow.vue";
import GroupChat from "@/components/chat/GroupChat.vue";
import GroupChatWindow from "@/components/chat/GroupChatWindow.vue";
import {
  getGroupMessages,
  getLastGroupMessage,
  getLastMessage,
  getMessages,
} from "@/assets/fetchFunctions";
export default {
  components: {
    ChatUser,
    ChatWindow,
    GroupChat,
    GroupChatWindow,
  },
  data() {
    return {
      chatType: "",
      user: {},
      group: {},
      offset: 0,
      fetching: false,
      messages: [],
      lastMessages: {},
      lastGroupMessages: {},
    };
  },
  props: {
    users: Array,
    groups: Array,
    onlineUsers: Object,
    newMessage: {},
  },
  watch: {
    newMessage(newVal) {
      if (newVal) {
        this.appendMessage(this.newMessage, true);
      }
    },
  },
  computed: {
    filteredUsers() {
      return this.users
        .slice()
        .filter((user) => user.follower || (user.following && !user.pending));
    },

    filteredGroups() {
      return this.groups.slice().filter((group) => group.joined);
    },

    sortedChats() {
      const combinedChats = [...this.filteredUsers, ...this.filteredGroups];
      return combinedChats.sort((a, b) => {
        const aDate = this.getDate(a.id || a.groupId, a.title);
        const bDate = this.getDate(b.id || b.groupId, b.title);

        if (aDate && bDate) {
          return bDate - aDate;
        } else if (aDate) {
          return -1;
        } else if (bDate) {
          return 1;
        } else {
          const aName = a.firstname || a.title;
          return aName.localeCompare(b.firstname || b.title);
        }
      });
    },
  },
  async mounted() {
    await this.fetchLastMessages();
    await this.fetchLastGroupMessages();
    if (this.sortedChats.length > 0) {
      if (this.sortedChats[0].id != undefined) {
        this.changeChat(this.sortedChats[0].id);
      } else {
        this.changeGroupChat(this.sortedChats[0].groupId);
      }
    }
  },
  methods: {
    async getMoreMessages(id, user) {
      this.fetching = true;
      let data;
      if (user) {
        data = await getMessages(id, this.offset);
      } else {
        data = await getGroupMessages(id, this.offset);
      }

      if (data.messages.length == 10) {
        this.offset += 10;
        this.fetching = false;
      }
      this.messages = data.messages.slice().reverse().concat(this.messages);
    },
    isOnline(id) {
      return id in this.onlineUsers ? this.onlineUsers[id] : false;
    },
    appendMessage(payload, incoming) {
      let msg = {
        toId: payload.toId,
        fromId: payload.fromId,
        message: payload.message,
        created: new Date(payload.created),
      };
      let id = incoming ? payload.fromId : payload.toId;
      payload.user
        ? (this.lastMessages[id] = msg)
        : (this.lastGroupMessages[id] = msg);

      if (this.chatType == "user") {
        if (
          (this.user.id == payload.toId || this.user.id == payload.fromId) &&
          payload.user
        ) {
          this.messages.push(msg);
          this.offset += 1;
        }
      } else if (this.chatType == "group") {
        if (this.group.groupId == payload.toId && !payload.user) {
          this.messages.push(msg);
          this.offset += 1;
        }
      }
    },
    async changeChat(userId) {
      this.offset = 0;
      this.messages = [];
      this.getMoreMessages(userId, true);
      this.user = this.users[this.users.findIndex((user) => user.id == userId)];
      this.chatType = "user";
    },
    async changeGroupChat(groupId) {
      this.offset = 0;
      this.messages = [];
      this.getMoreMessages(groupId, false);
      this.group =
        this.groups[this.groups.findIndex((group) => group.groupId == groupId)];
      this.chatType = "group";
    },
    async fetchLastMessages() {
      const usrs = this.filteredUsers;
      for (let i = 0; i < usrs.length; i++) {
        const lMessage = await getLastMessage(usrs[i].id);
        if (lMessage.message.toId > 0)
          this.lastMessages[usrs[i].id] = lMessage.message;
      }
    },
    async fetchLastGroupMessages() {
      const grps = this.filteredGroups;
      for (let i = 0; i < grps.length; i++) {
        const lMessage = await getLastGroupMessage(grps[i].groupId);
        if (lMessage.message.toId > 0)
          this.lastGroupMessages[grps[i].groupId] = lMessage.message;
      }
    },
    getDate(id, groupTitle) {
      if (groupTitle == undefined) {
        const message = this.lastMessages[id];
        return message ? new Date(message.created) : null;
      } else {
        const message = this.lastGroupMessages[id];
        return message ? new Date(message.created) : null;
      }
    },
  },
};
</script>

<style scoped>
.chat-users {
  background-color: rgb(200, 200, 200);
  border: 1px solid #888;
  border-radius: 10px;
  height: 50%;
  padding: 10px;
  margin-bottom: 5px;
  overflow-wrap: break-word;
  overflow-y: auto;
}

.chat-window {
  background-color: rgb(200, 200, 200);
  border: 1px solid #888;
  border-radius: 10px;
  height: 50%;
  margin-top: 5px;
}
</style>
