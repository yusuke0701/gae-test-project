<template>
  <div class="test">
    <h1>{{ msg }}</h1>
    <span v-if="seen">Now you see me</span>
    <ol>
      <li v-for="todo in todos" v-bind:key="todo.id">{{todo.text}}</li>
    </ol>
    <button @click="reverseMessage">Reverse Message</button>
    <div>
      <button @click="onClick('ネイティブDOMイベントもとれる',$event)">click</button>
      {{clickResult}}
    </div>
    <div>
      <input @change="onChange" placeholder="change" />
      {{changeResult}}
    </div>
    <div>
      <input @input="onInput" placeholder="input" />
      {{inputResult}}
    </div>
    <div>
      <input @keyup="onKeyup" placeholder="keyup" />
      {{keyupResult}}
    </div>
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
      ],
      clickResult: "",
      changeResult: "",
      inputResult: "",
      keyupResult: ""
    };
  },
  methods: {
    reverseMessage() {
      this.msg = this.msg
        .split("")
        .reverse()
        .join("");
    },
    onClick(msg, e) {
      this.clickResult = "clicked";
      alert(msg);
      alert(e);
    },
    onChange(e) {
      this.changeResult = e.target.value;
    },
    onInput(e) {
      this.inputResult = e.target.value;
    },
    onKeyup(e) {
      this.keyupResult = e.target.value;
    }
  }
};
</script>

<style scoped>
</style>
