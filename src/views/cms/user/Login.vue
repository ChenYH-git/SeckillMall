<template>
    <div class="app">
      <div class="login-header">
          管理员登录
      </div>
      <div class="login components-input-demo-presuffix">
        <a-form :model="form" :label-col="labelCol" :wrapper-col="{ span: 8}" ref="admin_form" :rules="rules">
          <a-form-item label="用户名" name="username">
            <a-input v-model:value="form.username" placeholder="请输入用户名">
              <template v-slot:prefix><mail-outlined type="mail"/></template>

              <template v-slot:suffix>
                <a-tooltip title="Extra information">
                  <unlock-outlined style="color: rgba(0,0,0,.45)" />
                </a-tooltip>
              </template>
            </a-input>
          </a-form-item>

          <a-form-item label="密码" name="password">
              <a-input-password v-model:value="form.password" type="password" placeholder="请输入密码">
                <template v-slot:prefix><safety-outlined type="safety"/></template>
              </a-input-password>
              
          </a-form-item>
          

          <a-form-item :wrapper-col=" { span: 8, offset: 3 }">
            <a-button type="primary" @click="onSubmit('admin_form')">
              提交
            </a-button>
            <a-button style="margin-left: 10px;">
              <router-link to="/">返回</router-link>
            </a-button>
          </a-form-item>
          
        </a-form>
      </div>
    </div>
</template>

<script>

import qs from "qs"
import { MailOutlined,SafetyOutlined ,UnlockOutlined } from '@ant-design/icons-vue';

export default {
  name: 'Login',
  components: {
    // UserOutlined,
    MailOutlined,
    UnlockOutlined,
    SafetyOutlined
  },
  data() {
    return {
      labelCol: { span: 4 },
      wrapperCol:  { span: 14, offset: 4 },
      form: {
        username: '',
        password: '',
      },
      rules:{
        username:[
          {required:true,message:"用户名不能为空",trigger:"blur"}
        ],
        password:[
          {required:true,message:"密码不能为空",trigger:"blur"},
          {min:6,message:"密码最小长度不能小于6",trigger:"blur"}
        ]

      }
    };
  },
  methods: {
    onSubmit(admin_form) {

      this.$refs[admin_form].validate().then(()=> {

        this.$axios.post("/user/admin_login",qs.stringify({
          username:this.form.username,
          password:this.form.password
        })).then((rep) => {
          if (rep.data.code == 200){
            console.log(rep);
            alert(rep.data.msg)
            this.$auth.setAdminAuth(rep.data.admin_token,rep.data.user_name)
            this.$router.push("/main")
          }else {
            alert(rep.data.msg)
          }

        }).catch((rep) => {
          alert(rep.data.msg)

        })

      }).catch(()=> {

      })

    
    },
  },
}
</script>

<style scoped>
  .app {
    margin: 50px 20%;
    padding: 20px;
    width: 100%;
    background: url("../../../assets/img/login.jpg");
    background-size: contain;
    background-repeat:no-repeat;
    /* background-size:100% 100%; */
    background-attachment: fixed;
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;


  }

  .login-header {
    margin-left: -35%;
    margin-bottom: 20px;
    font-size: 22px;
  }


  
</style>
