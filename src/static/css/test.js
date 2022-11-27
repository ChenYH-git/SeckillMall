// vuex/vuex.js中：
import Vuex from 'vuex';

const store = new Vuex.Store({
    state: {
        // 没有获取到设置为null
        token:localStorage.getItem('token')?localStorage.getItem('token') : null,
        user:localStorage.getItem('user')?localStorage.getItem('user') : null

    },
    mutations: {
        setAuth(state,token,user){
        state.token = token;
        localStorage.setItem('token', user.Authorization);
        },
        delAuth(state){
            state.token = null;
            state.user = null;
            localStorage.removeItem("token")
            localStorage.removeItem("user")

        }
    }

})

export default store



// router.js中：
// 定义导航守卫
router.beforeEach((to, from, next) => {
    if (to.path === '/login') {
    next();
    } else {
        let token = localStorage.getItem('token');

        if (token === 'null' || token === '') {
            next('/login');
        } else {
            next();
        }
    }
});