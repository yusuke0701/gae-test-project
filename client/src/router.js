import Vue from 'vue'
import Router from 'vue-router'

import Login from './components/account/Login'
import Registry from './components/account/Registry'

import CommentList from './components/comment/List'

import SignedURL from './components/SignedURL'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/login',
            name: 'ログイン画面',
            component: Login
        },
        {
            path: '/registry',
            name: '新規アカウント登録画面',
            component: Registry
        },
        {
            path: '/comments',
            name: 'コメント一覧画面',
            component: CommentList
        },
        {
            path: '/urls',
            name: '署名付きURL',
            component: SignedURL
        }
    ]
})