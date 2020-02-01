<template>
  <b-button @click="onClick">{{msg}}</b-button>
</template>

<script>
export default {
  name: "atom-button",
  props: ["msg"],
  methods: {
    onClick() {
      this.$emit("click");
    }
  }
};
</script>