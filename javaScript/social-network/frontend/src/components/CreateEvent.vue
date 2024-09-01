<template>
  <div class="modal" v-if="show">
    <div class="modal-content create">
      <span class="close" @click="close">&times;</span>
      <h2>Create event</h2>
      <form @submit.prevent="create">
        <div>
          <input
            type="textarea"
            id="title"
            placeholder="Event title"
            required
            v-model="title"
          />
        </div>
        <div>
          <input
            type="textarea"
            id="description-"
            placeholder="Event description"
            required
            v-model="description"
          />
        </div>
        <div>
          <input type="date" id="date" v-model="date" required />
        </div>
        <div class="attending-selectors">
          <label>
            <input type="radio" v-model="attending" value="0" />
            Going
          </label>
          <label>
            <input type="radio" v-model="attending" value="1" />
            Not going
          </label>
        </div>
        <button type="submit">Create Event</button>
      </form>
    </div>
  </div>
</template>

<script>
import { createEvent } from "@/assets/fetchFunctions";
export default {
  name: "GroupEvent",
  props: {
    show: Boolean,
    groupId: Number,
  },
  data() {
    return {
      title: "",
      description: "",
      date: "",
      attending: "0",
    };
  },
  methods: {
    close() {
      this.$emit("close");
      this.title = "";
      this.description = "";
      this.date = "";
      this.attending = "0";
    },

    async create() {
      const data = {
        groupId: this.groupId,
        title: this.title,
        description: this.description,
        date: this.date,
        going: this.attending == "0" ? true : false,
      };

      try {
        await createEvent(data);
        alert("Event created successfully");
        this.$emit("post");
        this.close();
      } catch (error) {
        this.error = error.message;
      }
    },
  },
};
</script>

<style scoped>
.modal {
  display: flex;
  position: fixed;
  z-index: 3;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(130, 130, 130, 0.466);
  backdrop-filter: blur(3px);
}

.modal-content {
  background-color: rgb(200, 200, 200);
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  border-radius: 10px;
}

.create {
  width: 43%;
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}

img {
  width: 25px;
}
</style>
