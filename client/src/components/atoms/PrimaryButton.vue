<template>
  <b-button @click="onClick">{{msg}}</b-button>
</template>

<script>
export default {
  props: {
    msg: {
      type: String,
      required: true
    }
  },
  methods: {
    onClick(e) {
      this.$emit("click", e);
    }
  }
};
</script>