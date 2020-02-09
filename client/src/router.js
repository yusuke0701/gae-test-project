import Vue from 'vue'
import Router from 'vue-router'

import AccountRegistry from '@/components/AccountRegistry'
import CommentDetail from '@/components/CommentDetail'
import CommentList from '@/components/CommentList'
import Login from '@/components/pages/LoginPage'
import Test from '@/components/test/Test'

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
            component: AccountRegistry
        },
        {
            path: '/comments',
            name: 'コメント一覧画面',
            component: CommentList
        },
        {
            path: '/comments/:id',
            name: 'コメント詳細画面',
            component: CommentDetail,
        },
        {
            path: '/test',
            name: '動作確認用',
            component: Test
        }
    ]
})