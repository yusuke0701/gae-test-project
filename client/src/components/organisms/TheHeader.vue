<template>
  <div>
    <header class="container-fluid">
      <div class="d-flex flex-row">
        <HomePageTitle class="flex-grow-1" :alt="alt" :src="src" />
        <HomePageMenu :links="links" />
      </div>
    </header>
  </div>
</template>

<script>
import HomePageMenu from "@/components/molecules/HomePageMenu";
import HomePageTitle from "@/components/molecules/HomePageTitle";
export default {
  name: "TheHeader",
  components: { HomePageMenu, HomePageTitle },
  props: {
    loginAccount: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      alt: "logo",
      src: "./logo.png",
      links: [
        { msg: "Home", to: "/", isShow: true },
        {
          msg: "Login",
          to: "/login",
          isShow: Object.keys(this.loginAccount).length === 0
        },
        {
          msg: "Admin",
          to: "/admin",
          isShow: Object.keys(this.loginAccount).length !== 0
        }
      ]
    };
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
