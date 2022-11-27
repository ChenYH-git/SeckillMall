import {createRouter,createWebHistory} from "vue-router"
import GoodsList from '@/views/front/GoodsList'
import GoodsDetail from "@/views/front/GoodsDetail";

const webHistory = createWebHistory()


const router = createRouter({
    history:webHistory,
    routes:[
        // front开始
        {
            path: '/',
            name: 'GoodsList',
            component: GoodsList
        },

        {
            path: "/detail/:id",
            name: "GoodsDetail",
            component: GoodsDetail
        },

        {
            path:"/login",
            name:"Login",
            component:() => import('./views/front/Login')
        },

        {
            path:"/register",
            name:"Register",
            component:() => import('./views/front/Register')
        },

        // front结束

        // cms开始

        {
            path: '/admin',
            name: 'Admin',
            component:() => import('./views/cms/user/Login')
        },


        
        {
            path:"/main",
            redirect:"/main/user_list",
        },


        {
            path:"/main",
            name:"Main",
            component:() => import('./views/Main'),
            // redirect:"/main/user_list",
            children:[
                {
                    path:"user_list",
                    name:"UserList",
                    component:() => import('./views/cms/user/UserList')
                },

                {
                    path:"product_list",
                    name:"ProductList",
                    component:() => import('./views/cms/product/ProductList')
                },

                {
                    path:"product_add",
                    name:"ProductAdd",
                    component:() => import('./views/cms/product/ProductAdd')
                },

                {
                    path:"product_edit/:id",
                    name:"ProductEdit",
                    component:() => import('./views/cms/product/ProductEdit')
                },

                {
                    path:"seckill_list",
                    name:"SeckillrList",
                    component:() => import('./views/cms/seckill/SeckillList')
                },
                {
                    path:"seckill_add",
                    name:"SeckillrAdd",
                    component:() => import('./views/cms/seckill/SeckillAdd')
                },
                {
                    path:"seckill_edit/:id",
                    name:"SeckillrEdit",
                    component:() => import('./views/cms/seckill/SeckillEdit')
                },
            ]
        },   
        

    ]
})



router.beforeEach((to,from,next) => {
    // to 满足条件的url放过
    if (to.path == "/" || to.path == "/login" || to.path == "/register" || to.path == "/admin"){
        next()  // 直接放过
    }else {


        // alert(to.path) // /main/user_list

        const reg = /\/main\/[a-zA-z]+[_-]?/
        // 做token校验
        if (reg.test(to.path)) {
            // 校验管理员，不受用户影响
            const admin_token = localStorage.getItem("admin_token")
            // alert("到main的url")
            if (admin_token == "" || admin_token == null || admin_token == '') {
                next("/admin")
            }else {
                next()
            }
        }else {
            // 校验用户的，不受管理员影响
            const front_token = localStorage.getItem("front_token")
             // front的需要权限的url
             if (front_token == "" || front_token == null || front_token == '') {
                next("/login")
            }else {
                next()
            }
        }

        
    }
})


export default router