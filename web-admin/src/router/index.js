import {createRouter, createWebHistory} from 'vue-router'
import blogStore from "@/stores/arlog.js";


const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/login/index.vue'),
        },
        {
            /*添加后台布局*/
            path: '/admin',
            name: 'admin',
            component: () => import('@/layout/admin.vue'),
            meta: {requiresAuth: true}, /*只需要主路由配置需要检查的token*/
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
/*判断是否登录*/
router.beforeEach((to, from, next) => {
    /*获取本地token*/
    const authStore = blogStore()

    /*判断需要验证的路由地址*/
    if (to.matched.some(record => record.meta.requiresAuth)) {

        // 检查 token 是否存在
        if (!authStore.token) {
            // 如果 token 不存在，则跳转到登录页面
            next({name: 'login'})
        } else {
            next() // 继续导航
        }
    } else {
        next() // 如果不需要认证，则继续导航
    }
});

export default router
