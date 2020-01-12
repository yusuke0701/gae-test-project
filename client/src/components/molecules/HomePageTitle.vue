<template>
  <div class="d-flex flex-row">
    <Logo :alt="alt" :src="src" />
  </div>
</template>

<script>
import Logo from "@/components/atoms/Logo";
export default {
  name: "HomePageTitle",
  components: { Logo },
  props: {
    alt: {
      type: String,
      required: true
    },
    src: {
      type: String,
      required: true
    }
  }
};
</script>