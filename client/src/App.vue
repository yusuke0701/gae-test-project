<template>
  <div id="app">
    <Header :loginAccount="loginAccount" />
    <nav aria-label="breadcrumb">
      <ol class="breadcrumb">
        <li class="breadcrumb-item">
          <router-link to="/comments">コメント一覧画面</router-link>
        </li>
      </ol>
    </nav>
    <router-view @child-event="setLoginAccount" />
  </div>
</template>

<script>
import Header from "@/components/organisms/Header";
export default {
  data() {
    return {
      loginAccount: {}
    };
  },
  components: { Header },
  methods: {
    setLoginAccount(loginAccount) {
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
  color: #2c3e50;
  margin-top: 55px;
}
</style>