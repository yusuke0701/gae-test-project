<template>
  <b-form>
    <p id="login-title">ログイン</p>
    <input type="text" v-model="id" size="20" placeholder="ID" required />
    <input type="password" v-model="password" size="20" placeholder="パスワード" required />
    <b-button @click="doLogin">ログイン</b-button>
  </b-form>
</template>