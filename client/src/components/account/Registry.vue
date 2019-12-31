<template>
  <div class="login">
    <h2>登録画面</h2>
    <div>
      <p>
        ID:
        <input type="text" v-model="id" size="20" required />
      </p>
      <p>
        メールアドレス:
        <input type="email" v-model="email" size="20" required />
      </p>
      <p>
        メールアドレス(確認用):
        <input type="email" v-model="email2" size="20" required />
      </p>
      <p>
        お名前:
        <input type="text" v-model="name" size="20" required />
      </p>
      <p>
        ニックネーム:
        <input type="text" v-model="nick_name" size="20" />
      </p>
      <p>
        Password:
        <input type="password" v-model="password" size="20" required />
      </p>
      <p>
        Password(確認用):
        <input type="password" v-model="password2" size="20" required />
      </p>
      <button v-on:click="regist">新規登録</button>
    </div>
  </div>
</template>

<script>
import { insertAccount } from "../../service/account";
export default {
  name: "Registry",
  data: function() {
    return {
      id: "",
      email: "",
      email2: "",
      name: "",
      nick_name: "",
      password: "",
      password2: ""
    };
  },
  methods: {
    regist: function() {
      if (this.email !== this.email2) {
        alert("メールアドレスが一致しません。");
        return;
      }
      if (this.password !== this.password2) {
        alert("パスワードが一致しません。");
        return;
      }

      insertAccount({
        id: this.id,
        email: this.email,
        name: this.name,
        nick_name: this.nick_name,
        password: this.password
      })
        .then(res => {
          if (res.status === 200) {
            alert("登録完了しました。\nログイン画面でログインをお願いします。");
            this.$router.push("/login");
          } else {
            alert(res.data);
          }
        })
        .catch(error => window.alert(error));
    }
  }
};
</script>