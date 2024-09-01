<template>
  <div class="group" @click="openGroup">
    <div class="image-container">
      <img
        src="http://localhost:3000/images/group.png"
        :alt="`${group.groupId}`"
      />
    </div>
    <div class="info-container">
      <div class="title-container">{{ group.title }}</div>
      <div class="description-container">
        {{ group.description }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    group: {},
  },

  data() {
    return {};
  },
  methods: {
    openGroup() {
      this.$emit("open", this.group.groupId);
    },
  },
};
</script>

<style scoped>
.group {
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

.title-container {
  margin-bottom: 5px;
}

.description-container {
  font-style: italic;
  font-size: small;
}

.group:hover {
  cursor: pointer;
  color: rgb(150, 150, 150);
}
</style>
