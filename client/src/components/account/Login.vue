<template>
  <div class="login">
    <h2>ログイン画面</h2>
    <div>
      <p>
        ID:
        <input type="text" v-model="id" size="20" required />
      </p>
      <p>
        Password:
        <input type="password" v-model="password" size="20" required />
      </p>
      <button @click="doLogin">ログイン</button>
    </div>
    <div>
      アカウントを持ってない方はこちら
      <button @click="doRegistry">新規登録</button>
    </div>
  </div>
</template>

<script>
import { login } from "../../service/account";
export default {
  name: "Login",
  data: function() {
    return {
      id: "",
      password: ""
    };
  },
  methods: {
    doLogin: function() {
      login(this.id, this.password)
        .then(res => {
          if (res.status === 200) {
            this.$emit("child-event", res.data);
            this.loginAccount = res.data;
          } else {
            window.alert(res.data);
          }
        })
        .catch(error => window.alert(error));
    },
    doRegistry: function() {
      this.$router.push("/registry");
    }
  }
};
</script>