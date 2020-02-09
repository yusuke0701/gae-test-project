<template>
  <input
    :type="type"
    :size="size"
    :value="value"
    :placeholder="placeholder"
    required
    @keyup.enter="enterHandler"
  />
</template>

<script>
export default {
  props: {
    type: {
      type: String,
      required: true
    },
    size: {
      type: String,
      default: "20",
      required: false
    },
    value: {
      type: String,
      required: false
    },
    placeholder: {
      type: String,
      required: false
    }
  },
  methods: {
    enterHandler(e) {
      this.$emit("keyup", e);
    }
  }
};
</script>