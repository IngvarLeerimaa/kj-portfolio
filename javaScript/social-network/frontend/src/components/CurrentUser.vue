<template>
  <div class="privacy-select">
    <label for="account">Account type:</label>
    <select id="account" @change="changeAccount">
      <option :value="0" :selected="!user.public">Private</option>
      <option :value="1" :selected="user.public">Public</option>
    </select>
  </div>
  <div class="user-info">
    <div>Email: {{ user.email }}</div>
    <div>Date of Birth: {{ user.dateofbirth }}</div>
    <div v-if="user.nickname != ''">Nickname: {{ user.nickname }}</div>
    <div v-if="user.about != ''">About me: {{ user.about }}</div>
  </div>
</template>

<script>
import { updateUser } from "@/assets/fetchFunctions";
export default {
  name: "CurrentUser",
  props: {
    user: {},
  },
  methods: {
    async changeAccount(event) {
      try {
        await updateUser(event.target.value);
      } catch (error) {
        this.error = error.message;
      }
    },
  },
};
</script>
