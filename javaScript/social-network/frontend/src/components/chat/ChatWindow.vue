<template>
  <div class="chat-messages" v-if="user.id != undefined">
    <div class="user-info">
      <div :class="['image-container', { online: isOnline }]">
        <img :src="user.avatar" :alt="`${user.id}`" />
      </div>
      <div class="name-container">{{ user.firstname }} {{ user.lastname }}</div>
    </div>

    <div
      ref="messageContainer"
      class="message-container"
      @scroll="handleDebouncedScroll"
    >
      <div>Start of chat with {{ user.firstname }} {{ user.lastname }}.</div>
      <div
        v-for="(message, index) in messages"
        :key="index"
        :style="{ 'text-align': message.toId === user.id ? 'right' : 'left' }"
      >
        <div :class="['message-bubble', { green: message.toId == user.id }]">
          {{ message.message }}
        </div>
      </div>
    </div>
    <div class="message-input">
      <form @submit.prevent="newMessage">
        <span class="emoji-toggle" @click="toggleEmoji">🙂</span>
        <input type="text" v-model="message" required />
        <input type="submit" value="Send" />
      </form>
    </div>
    <EmojiPicker :show="emojiVisible" @insert="insertEmoji" />
  </div>
  <div v-else>Follow user/join a group to have a chat.</div>
</template>

<script>
import EmojiPicker from "@/components/EmojiPicker.vue";
import { sendMessage } from "@/assets/fetchFunctions";
import debounce from "lodash/debounce";
export default {
  components: {
    EmojiPicker,
  },
  props: {
    user: {},
    messages: Array,
    isOnline: Boolean,
    messagesAvailable: Boolean,
  },
  data() {
    return {
      emojiVisible: false,
      scrolling: false,
      message: "",
      handleDebouncedScroll: debounce(this.handleScroll, 100),
    };
  },
  watch: {
    messages: {
      handler() {
        if (!this.scrolling) this.$nextTick(this.scrollBottom);
      },
      deep: true,
    },
  },
  methods: {
    handleScroll() {
      this.scrolling = true;
      if (
        this.$refs.messageContainer.scrollTop === 0 &&
        this.messagesAvailable
      ) {
        this.$emit("moreMessages", this.user.id);
      }
      clearTimeout(this.scrollTimeout);
      this.scrollTimeout = setTimeout(() => {
        this.scrolling = false;
      }, 100);
    },
    newMessage() {
      sendMessage(this.user.id, this.message, true);
      this.$emit("newMessage", {
        toId: this.user.id,
        fromId: this.$userId,
        user: true,
        message: this.message,
        created: new Date(),
      });
      this.message = "";
    },
    toggleEmoji() {
      this.emojiVisible = !this.emojiVisible;
    },
    insertEmoji(emoji) {
      this.message += emoji;
    },
    scrollBottom() {
      const messageArea = this.$refs.messageContainer;
      messageArea.scrollTop = messageArea.scrollHeight;
    },
  },
};
</script>

<style scoped>
.chat-messages {
  height: 100%;
  overflow: hidden;
  overflow-wrap: break-word;
}

.user-info {
  display: flex;
  align-items: center;
  padding: 10px;
}

.image-container {
  flex-shrink: 0;
  width: 35px;
  height: 35px;
  margin-right: 10px;
  border: 2px solid black;
  border-radius: 50%;
  overflow: hidden;
}

.online {
  border: 2px solid rgb(0, 255, 0);
}

.image-container img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.message-container {
  height: 72%;
  border-top: 1px solid #888;
  border-bottom: 1px solid #888;
  overflow: auto;
}

.message-bubble {
  display: inline-block;
  padding: 10px;
  border-radius: 10px;
  margin: 5px;
  margin-left: 10px;
  color: black;
  background-color: #f1e2e291;
  max-width: 70%;
}

.green {
  background-color: rgb(18, 134, 24);
}

.message-input {
  height: 100%;
  flex-shrink: 0;
  align-items: center;
  padding: 10px;
}
.emoji-toggle {
  font-size: 20px;
  margin-right: 5px;
}

.emoji-toggle:hover {
  cursor: pointer;
}
</style>
