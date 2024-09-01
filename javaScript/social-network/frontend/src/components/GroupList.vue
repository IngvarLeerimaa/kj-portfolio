<template>
  <div class="groups-container">
    <h2>Groups:</h2>
    <div v-for="group in groups" :key="group.groupId">
      <GroupItem :group="group" @click="open(group.groupId)" />
    </div>
  </div>
</template>

<script>
import GroupItem from "@/components/GroupItem.vue";
export default {
  components: {
    GroupItem,
  },
  props: {
    groups: Array,
  },
  data() {
    return {};
  },
  methods: {
    open(groupId) {
      this.$emit("open", groupId);
    },
  },
};
</script>

<style scoped>
.groups-container {
  padding: 10px;
  overflow-wrap: break-word;
}
</style>
