import {createRouter, createWebHistory} from 'vue-router'


const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            /*首页*/
            path: '/',
            name: 'home',
            component: () => import('@/views/home/index.vue'),
        },
        {
            /*登录页面*/
            path: '/login',
            name: 'login',
            component: () => import('@/views/login/index.vue'),
        },
        {
            /*添加后台布局*/
            path: '/admin',
            name: 'admin',
            component: () => import('@/layout/admin.vue'),
            children: [

                {
                    /*后台的首页*/
                    path: '',
                    name: 'home',
                    component: () => import('@/views/admin/home.vue')
                },

                {
                    /*后台的用户页面*/
                    path: 'user',
                    name: 'user',
                    component: () => import('@/views/admin/user.vue')
                },
            ]
        },
    ]
})

export default router
