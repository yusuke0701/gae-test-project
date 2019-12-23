import Vue from 'vue'
import Router from 'vue-router'

import Home from './components/Home'
import Page1 from './components/Page1'
import Page2 from './components/Page2'

Vue.use(Router)

export default new Router({
    mode: 'history',
    base: process.env.BASE_URL,
    routes: [
        {
            path: '/',      // このパスにアクセスしたら
            name: 'home',
            component: Home // このコンポーネントを呼ぶ
        },
        {
            path: '/page1',
            name: 'page1',
            component: Page1
        },
        {
            path: '/page2',
            name: 'page2',
            component: Page2
        }
    ]
})