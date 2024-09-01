<template>
  <div class="users">
    <h2>Users:</h2>
    <div v-for="user in users" :key="user.id">
      <UserComponent
        :user="user"
        @profile="viewProfile(user.id)"
        @follow="updateFollow"
      />
    </div>
  </div>
</template>

<script>
import UserComponent from "@/components/user/UserComponent.vue";
import { followUser } from "@/assets/fetchFunctions";
export default {
  components: {
    UserComponent,
  },
  props: {
    users: Array,
  },
  methods: {
    viewProfile(userId) {
      this.$emit("view-profile", userId);
    },
    async updateFollow(payload) {
      const userId = payload.userId;
      const isPublic = payload.isPublic;
      const following = payload.following;
      try {
        await followUser(userId, following);
        this.$emit("update", {
          userId: userId,
          attribute: "following",
          value: following,
        });
        this.$emit("update", {
          userId: userId,
          attribute: "pending",
          value: isPublic ? false : following,
        });
      } catch (error) {
        this.error = error.message;
      }
    },
  },
};
</script>

<style scoped>
.users {
  padding: 10px;
  overflow-wrap: break-word;
}
</style>
