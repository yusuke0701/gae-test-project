<template>
  <header>
    <button class="btn btn-primary" @click="goHome">Homeへ</button>
    <button class="btn btn-primary" @click="goLogin">ログイン画面へ</button>
  </header>
</template>

<script>
export default {
  name: "Header",
  methods: {
    goHome() {
      if (location.pathname !== "/") {
        this.$router.push("/");
      }
    },
    goLogin() {
      if (location.pathname !== "/login") {
        this.$router.push("/login");
      }
    }
  }
};
</script>