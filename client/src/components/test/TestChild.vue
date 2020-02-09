<template>
  <div>
    {{pval}}
    <button @click="addValueChild()">ADD</button>
  </div>
</template>

<script>
export default {
  props: ["pval"],
  methods: {
    addValueChild() {
      this.$emit("add", 10);
    }
  }
};
</script>