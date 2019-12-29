<template>
  <div class="commentList">
    <li v-for="comment in commentList" v-bind:key="comment.id">{{comment.body}}</li>
  </div>
</template>

<script>
import { getAllComment } from "../../service/comments";
export default {
  name: "Comment list",
  data: function() {
    return { commentList: [] };
  },
  created() {
    getAllComment()
      .then(res => {
        if (res.status === 200) {
          res.json.then(jsonData => (this.commentList = jsonData));
        } else {
          res.text().then(data => window.alert(data));
        }
      })
      .catch(error => window.alert(error));
  }
};
</script>