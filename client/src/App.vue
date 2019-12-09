<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png" />
    <Test />
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  </div>
</template>

<script>
import Test from "./components/Test.vue";
import HelloWorld from "./components/HelloWorld.vue";

export default {
  name: "app",
  components: {
    Test,
    HelloWorld
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
