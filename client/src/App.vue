<template>
  <div id="app">
    <TheHeader :loginAccount="loginAccount" />
    <nav aria-label="breadcrumb">
      <ol class="breadcrumb">
        <li class="breadcrumb-item">
          <TagsSearcher />
        </li>
      </ol>
    </nav>
    <Home />
    <router-view @child-event="setLoginAccount" />
  </div>
</template>

<script>
import TheHeader from "@/components/organisms/TheHeader";
import TagsSearcher from "@/components/pages/TagsSearcher";
export default {
  data() {
    return {
      loginAccount: {}
    };
  },
  components: { TheHeader, TagsSearcher },
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