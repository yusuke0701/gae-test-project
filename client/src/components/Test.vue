<template>
  <div class="test">
    <h1>{{ msg }}</h1>
    <span v-if="seen">Now you see me</span>
    <ol>
      <li v-for="todo in todos" v-bind:key="todo.id">{{todo.text}}</li>
    </ol>
    <button @click="reverseMessage">Reverse Message</button>
  </div>
</template>

<script>
export default {
  name: "Test",
  data: function() {
    return {
      msg: "どう？",
      seen: true,
      todos: [
        { id: 1, text: "apple" },
        { id: 2, text: "orange" },
        { id: 3, text: "water melon" }
      ]
    };
  },
  methods: {
    reverseMessage: function() {
      this.msg = this.msg
        .split("")
        .reverse()
        .join("");
    }
  }
};
</script>

<style scoped>
</style>
