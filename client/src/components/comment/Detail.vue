<template>
  <div>
    <input v-model="id" size="20" />
    <input v-model="body" size="20" />
    <button v-on:click="get">取得</button>
    <button v-on:click="post">投稿</button>
  </div>
</template>

<script>
import { insertComment, getComment } from "../service/comments";
export default {
  name: "CommentDetail",
  data: function() {
    return {
      id: "",
      body: ""
    };
  },
  methods: {
    get: function() {
      getComment(this.id)
        .then(res => {
          if (res.status === 200) {
            res.json.then(jsonData => (this.body = jsonData.body));
          } else {
            res.text().then(data => window.alert(data));
          }
        })
        .catch(error => window.alert(error));
    },
    post: function() {
      insertComment(this.id, this.body)
        .then(res => {
          if (res.status === 200) {
            // nop
          } else {
            res.text().then(data => window.alert(data));
          }
        })
        .catch(error => window.alert(error));
    }
  }
};
</script>