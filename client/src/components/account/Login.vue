<template>
  <div>
    <p>ログイン/登録</p>
    <div id="login">
      <b-form>
        <p id="login-title">ログイン</p>
        <input type="text" v-model="id" size="20" placeholder="ID" required />
        <input type="password" v-model="password" size="20" placeholder="パスワード" required />
        <b-button @click="doLogin">ログイン</b-button>
      </b-form>
    </div>
    <div id="registry">
      <b-form>
        <p id="registry-title">はじめての方</p>
        <p>コメントを書き込むには、無料ID登録が必要です。</p>
        <b-button href="/registry">新規登録</b-button>
      </b-form>
    </div>
  </div>
</template>

<script>
import { login } from "../../service/account";
export default {
  name: "ログイン画面",
  data() {
    return {
      id: "",
      password: ""
    };
  },
  methods: {
    doLogin() {
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
    }
  }
};
</script>

<style scoped>
#login {
  background-color: aliceblue;
}
#login-title {
  background-color: azure;
}
#registry {
  background-color: aliceblue;
}
#registry-title {
  background-color: azure;
}
</style>