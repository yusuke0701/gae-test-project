import Vue from 'vue'
import Router from 'vue-router'

import CommentList from './components/comment/List'
import Login from './components/account/Login'
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