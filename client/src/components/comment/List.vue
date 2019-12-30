<template>
  <div class="commentList">
    <li v-for="comment in commentList" v-bind:key="comment.id">{{comment.body}}</li>
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
          window.alert(res.data);
        }
      })
      .catch(error => window.alert(error));
  }
};
</script>