<template>
  <div class="tags_searcher">
    <div class="container-fluid">
      <div class="d-flex flex-row">
        <div>
          <b-form-input
            type="search"
            v-model="thread_tag"
            @change="onChangeTag"
            size="20"
            placeholder="タグで検索"
          />
        </div>
        <div>
          <b-form-input
            type="search"
            v-model="thread_name"
            @change="onChange"
            size="20"
            placeholder="名前で検索"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      thread_tag: "",
      thread_name: ""
    };
  },
  methods: {
    onChangeTag(e) {
      this.changeResult = e.target.value;
    }
  }
};
</script>