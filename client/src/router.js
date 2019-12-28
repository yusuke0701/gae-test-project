import Vue from 'vue'
import Router from 'vue-router'

import Comments from './components/Comments'
import SignedURL from './components/SignedURL'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/',
            name: 'コメント一覧画面',
            component: Comments
        },
        {
            path: '/urls',
            name: '署名付きURL',
            component: SignedURL
        }
    ]
})