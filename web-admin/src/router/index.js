import {createRouter, createWebHistory} from 'vue-router'
import blogStore from "@/stores/arlog.js";


const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'lg',
            component: () => import('@/views/login/login.vue'),
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/login/login.vue'),
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
                    /*文章页面*/
                    path: 'article',
                    name: 'article',
                    component: () => import('@/views/admin/article.vue')
                },
                {
                    /*标签页面*/
                    path: 'tag',
                    name: 'tag',
                    component: () => import('@/views/admin/tag.vue')
                },
                {
                    /*分类页面*/
                    path: 'sort',
                    name: 'sort',
                    component: () => import('@/views/admin/sort.vue')
                },
                {
                    /*图片页面*/
                    path: 'jpg',
                    name: 'jpg',
                    component: () => import('@/views/admin/jpg.vue')
                },
                {
                    /*留言页面*/
                    path: 'message',
                    name: 'message',
                    component: () => import('@/views/admin/message.vue')
                },
                {
                    /*ai页面*/
                    path: 'ai',
                    name: 'ai',
                    component: () => import('@/views/admin/ai.vue')
                },
                {
                    /*角色页面*/
                    path: 'role',
                    name: 'role',
                    component: () => import('@/views/admin/role.vue')
                },
                {
                    /*后台的用户页面*/
                    path: 'user',
                    name: 'user',
                    component: () => import('@/views/admin/user.vue')
                },
                {
                    /*用户资料*/
                    path: 'profile',
                    name: 'profile',
                    component: () => import('@/views/admin/profile.vue')
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
