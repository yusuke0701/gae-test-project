<template>
  <b-nav id="nav">
    <b-nav-item v-for="(link,index) in links" :key="index">
      <HeaderMenuLink v-if="link.isShow" :msg="link.msg" :to="link.to" />
    </b-nav-item>
  </b-nav>
</template>

<script>
import HeaderMenuLink from "@/components/atoms/HeaderMenuLink";
export default {
  name: "HomePageMenu",
  components: { HeaderMenuLink },
  props: {
    links: {
      type: Array,
      required: true
    }
  }
};
</script>