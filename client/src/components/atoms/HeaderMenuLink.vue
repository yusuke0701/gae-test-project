<template>
  <router-link :to="to">{{msg}}</router-link>
</template>

<script>
export default {
  props: {
    to: {
      type: String,
      required: true
    },
    msg: {
      type: String,
      required: true
    }
  }
};
</script>