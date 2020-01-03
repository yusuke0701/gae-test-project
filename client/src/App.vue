<template>
  <div id="app">
    <Header :loginAccount="loginAccount" />
    <nav aria-label="breadcrumb">
      <ol class="breadcrumb">
        <li class="breadcrumb-item">
          <router-link to="/comments">コメント一覧画面</router-link>
        </li>
        <li class="breadcrumb-item active" aria-current="page">
          <router-link to="/urls">署名付きテスト</router-link>
        </li>
      </ol>
    </nav>
    <router-view @child-event="setLoginAccount" />
  </div>
</template>

<script>
import Header from "./components/Header";
export default {
  name: "app",
  data: function() {
    return {
      loginAccount: {}
    };
  },
  components: { Header },
  methods: {
    setLoginAccount: function(loginAccount) {
      this.loginAccount = loginAccount;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<!-- <style scoped> -->
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