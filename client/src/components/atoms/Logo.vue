<template>
  <img class="img_logo" :alt="alt" :src="src" :height="55" />
</template>

<script>
export default {
  name: "Logo",
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