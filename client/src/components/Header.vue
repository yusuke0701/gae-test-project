<template>
  <header class="container-fluid">
    <div class="d-flex flex-row">
      <h1 class="flex-grow-1">掲示板システム</h1>
      <b-nav id="nav">
        <b-nav-item>
          <router-link to="/">Home</router-link>
        </b-nav-item>
        <b-nav-item v-if="Object.keys(loginAccount).length === 0">
          <router-link to="/login">Login</router-link>
        </b-nav-item>
        <b-nav-item v-else>
          <router-link to="/login">Admin</router-link>
        </b-nav-item>
      </b-nav>
    </div>
  </header>
</template>

<script>
export default {
  props: {
    loginAccount: {
      type: Object,
      required: true
    }
  }
};
</script>

<style scoped>
.container-fluid {
  top: 0;
  position: fixed;
  background-color: aliceblue;
  text-align: center;
}
</style>
