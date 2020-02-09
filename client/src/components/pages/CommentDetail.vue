<template>
  <div>
    {{comment.id}}
    <input v-model="comment.body" size="20" />
    <button v-on:click="post">送信</button>
  </div>
</template>

<script>
import { getComment, updateComment } from "../service/comment";
export default {
  data() {
    return {
      comment: {}
    };
  },
  mounted() {
    getComment(this.$route.params.id)
      .then(res => {
        if (res.status === 200) {
          this.comment = res.data;
        } else {
          alert(res.data);
        }
      })
      .catch(error => alert(error));
  },
  methods: {
    post() {
      updateComment(this.comment)
        .then(res => {
          if (res.status === 200) {
            this.comment = res.data;
          } else {
            alert(res.data);
          }
        })
        .catch(error => alert(error));
    }
  }
};
</script>