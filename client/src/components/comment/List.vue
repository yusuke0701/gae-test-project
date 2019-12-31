<template>
  <div class="commentList">
    <li v-for="comment in commentList" :key="comment.id">
      {{comment.body}}
      <button @click="onClick(comment.id)">編集</button>
    </li>
  </div>
</template>

<script>
import { getAllComment } from "../../service/comment";
export default {
  name: "CommentList",
  data: function() {
    return { commentList: [] };
  },
  mounted() {
    getAllComment()
      .then(res => {
        if (res.status === 200) {
          this.commentList = res.data;
        } else {
          alert(res.data);
        }
      })
      .catch(error => alert(error));
  },
  methods: {
    onClick(id) {
      this.$router.push("/comments/" + id);
    }
  }
};
</script>