<template>
  <div class="group-chat" @click="selectChat">
    <div class="image-container">
      <img
        src="http://localhost:3000/images/group.png"
        :alt="`${group.groupId}`"
      />
    </div>
    <div class="info-container">
      <div class="name-container">{{ group.title }}</div>
      <div class="message-container">
        {{
          lastMessages[group.groupId] != undefined
            ? lastMessages[group.groupId].message
            : "No chat history"
        }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    group: {},
    lastMessages: null,
  },

  data() {
    return {};
  },
  methods: {
    selectChat() {
      this.$emit("chat", this.group.groupId);
    },
  },
};
</script>

<style scoped>
.group-chat {
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

.group-chat:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}
</style>
