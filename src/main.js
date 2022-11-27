import { createApp } from 'vue';
import ant from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import App from './App.vue';

import router from './router'
import axios from './axios'

import auth from './auth/auth'



// createApp(App).user(ant).mount('#app')

const app = createApp(App);

app.use(ant);
app.use(router)





app.mount('#app')

app.config.globalProperties.$router = router
app.config.globalProperties.$axios = axios
app.config.globalProperties.$auth = auth



