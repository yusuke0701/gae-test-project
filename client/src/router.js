import Vue from 'vue'
import Router from 'vue-router'
import pages from '@/components/pages'
import Test from '@/components/test/Test'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/login',
            name: 'ログイン画面',
            component: pages.LoginPage
        },
        {
            path: '/registry',
            name: '新規アカウント登録画面',
            component: pages.AccountRegistry
        },
        {
            path: '/comments',
            name: 'コメント一覧画面',
            component: pages.CommentList
        },
        {
            path: '/comments/:id',
            name: 'コメント詳細画面',
            component: pages.CommentDetail,
        },
        {
            path: '/test',
            name: '動作確認用',
            component: Test
        }
    ]
})