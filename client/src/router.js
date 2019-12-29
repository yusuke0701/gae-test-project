import Vue from 'vue'
import Router from 'vue-router'

import CommentList from './components/comment/List'
import SignedURL from './components/SignedURL'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/',
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