<template>
  <div>
    <PageTitle :msg="title" />
    <div id="login"></div>
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
import PageTitle from "@/components/atoms/PageTitle";
import { login } from "@/service/account";
export default {
  components: { PageTitle },
  data() {
    return {
      id: "",
      password: "",
      title: "ログイン/登録"
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