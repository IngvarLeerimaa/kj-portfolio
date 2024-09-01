<template>
  <div class="modal" v-if="show">
    <div class="emoji-container">
      <span
        class="emoji"
        v-for="(emoji, index) in emojis"
        :key="index"
        @click="insert(emoji)"
        >{{ emoji }}</span
      >
    </div>
  </div>
</template>

<script>
export default {
  name: "EmojiPicker",
  props: {
    show: Boolean,
  },
  data() {
    return {
      emojis: [
        "😂",
        "❤️",
        "😍",
        "😊",
        "👍",
        "😭",
        "😘",
        "👏",
        "🙌",
        "😎",
        "😁",
        "🔥",
        "🎉",
        "💕",
        "😜",
        "👌",
        "😉",
        "🙏",
        "💪",
        "😅",
        "😱",
        "😆",
        "✨",
        "😋",
        "😏",
        "👀",
        "😄",
        "💖",
        "😇",
        "😡",
        "😢",
        "🤔",
        "😕",
        "💔",
        "😒",
        "😞",
        "😔",
        "😴",
        "😪",
        "😳",
      ],
    };
  },
  methods: {
    close() {
      this.$emit("close", this.remove);
    },
    insert(emoji) {
      this.$emit("insert", emoji);
    },
  },
};
</script>

<style scoped>
.modal {
  display: flex;
  position: fixed;
  z-index: 3;
  right: 12%;
  bottom: 6%;
  width: auto;
  max-width: 150px;
  min-height: fit-content;
}

.emoji-container {
  background-color: rgb(175, 175, 175);
  margin: 10% auto;
  padding: 5px;
  border: 1px solid #888;
  border-radius: 10px;
  font-size: 20px;
}

.emoji:hover {
  cursor: pointer;
}
</style>
