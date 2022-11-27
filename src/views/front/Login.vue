<template>
    <div class="app">
      <div class="login-header">
          用户登录
      </div>
      <div class="login components-input-demo-presuffix">
        <a-form :model="form" :label-col="labelCol" :wrapper-col="{ span: 8}" ref="loginForm" :rules="rules">
          <a-form-item label="邮箱地址" name="mail">
            <a-input v-model:value="form.mail" placeholder="请输入邮箱地址">
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
            <a-button type="primary" @click="onSubmit('loginForm')">
              登录
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
    // CodeOutlined,
    SafetyOutlined
  },
  data() {
    let validateEMail = async (rule, value) => {
      const reg = /^([a-zA-Z0-9]+[-_.]?)+@[a-zA-Z0-9]+\.[a-z]+$/;
      if (value == "" || value == undefined || value == null) {
        // callback(new Error("请输入邮箱"));
        return Promise.reject(new Error("请输入邮箱"));
      } else {
        if (!reg.test(value)) {
          return Promise.reject(new Error("请输入正确的邮箱"));
        } else {
          return Promise.resolve();
        }
      }
    };

    return {
      
      labelCol: { span: 4 },
      wrapperCol:  { span: 14, offset: 4 },
      form: {
        mail: '',
        password: '',
      },
      rules:{
        mail:[
          {required:true,message:"邮箱地址不能为空",trigger:"blur"},
          {validator:validateEMail,trigger:"blur"}
        ],
        password:[
          {required:true,message:"密码不能为空",trigger:"blur"},
          {min:6,message:"密码长度不能小于6",trigger:"blur"}

        ]

      }
    };
  },
  methods: {
    onSubmit(loginForm) {
      this.$refs[loginForm].validate().then(()=> {
        this.$axios.post('/user/front_user_login', qs.stringify({
            mail: this.form.mail,
            password: this.form.password,
        })).then((rep) => {
              if (rep.data.code == 200) {
                  console.log(rep);
                  this.$auth.setFrontAuth(rep.data.token,rep.data.username)
                  this.$router.push("/")
              }else {
                alert(rep.data.msg)
              }
            }).catch((rep) => {
                console.log(rep);
        });

      }).catch(()=> {

        alert("校验出错，请完善需要提交的信息")

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
    background: url("../../assets/img/login.jpg");
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
