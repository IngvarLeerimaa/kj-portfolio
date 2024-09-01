<template>
  <div class="chat-user" @click="selectChat">
    <div :class="['image-container', { online: isOnline }]">
      <img :src="user.avatar" :alt="`${user.id}`" />
    </div>
    <div class="info-container">
      <div class="name-container">{{ user.firstname }} {{ user.lastname }}</div>
      <div class="message-container">
        {{
          lastMessages[user.id] != undefined
            ? lastMessages[user.id].message
            : "No chat history"
        }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    user: {},
    isOnline: Boolean,
    lastMessages: null,
  },

  data() {
    return {};
  },
  methods: {
    selectChat() {
      this.$emit("chat", this.user.id);
    },
  },
};
</script>

<style scoped>
.chat-user {
  display: flex;
  border: 1px solid #888;
  border-radius: 10px;
  margin-bottom: 5px;
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

.info-container {
  text-align: left;
  overflow: hidden;
}

.name-container {
  margin-bottom: 5px;
}

.message-container {
  font-style: italic;
  font-size: small;
}

.chat-user:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}
</style>
