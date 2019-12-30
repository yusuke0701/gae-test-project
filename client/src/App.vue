<template>
  <div id="app">
    <div id="nav">
      <router-link to="/login">ログイン画面</router-link>|
      <router-link to="/comments">コメント一覧画面</router-link>|
      <router-link to="/urls">署名付きテスト</router-link>
    </div>
    <router-view @child-event="setLoginAccount" />
  </div>
</template>

<script>
export default {
  name: "app",
  data: function() {
    return {
      loginAccount: {}
    };
  },
  methods: {
    setLoginAccount: function(loginAccount) {
      this.loginAccount = loginAccount;
    }
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
