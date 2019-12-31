<template>
  <div>
    <button @click="goHome">Homeへ</button>
    <button @click="goLogin">ログイン画面へ</button>
  </div>
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