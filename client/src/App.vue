<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">コメント一覧画面</router-link>|
      <router-link to="/urls">署名付きテスト</router-link>
    </div>
    <router-view />
  </div>
</template>

<script>
export default {
  name: "app"
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
