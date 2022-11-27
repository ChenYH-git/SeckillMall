<template>
  <div class="app">
    <div class="login-header">用户注册</div>
    <div class="login components-input-demo-presuffix">
      <a-form
        :model="form"
        :label-col="labelCol"
        :wrapper-col="{ span: 8 }"
        :rules="rules"
        ref="valid_form"
      >
        <a-form
        :model="form"
        :label-col="labelCol"
        :wrapper-col="{ span: 8 }"
        :rules="rules_email"
        ref="valid_emial_form"
      >
        <a-form-item label="邮箱地址" name="mail">
          <a-input v-model:value="form.mail" placeholder="请输入正确的邮箱地址">
            <template v-slot:prefix><mail-outlined type="mail" /></template>

            <template>
              <a-tooltip title="Extra information">
                <unlock-outlined style="color: rgba(0, 0, 0, 0.45)" />
              </a-tooltip>
            </template>
          </a-input>
        </a-form-item>

        <a-form-item label="验证码" name="captche">
          <a-input-search
            v-model:value="form.captche"
            placeholder="请输入邮箱验证码"
            size="large"
            @search="sendMail('valid_emial_form')"
          >
            <template v-slot:enterButton>
              <a-button> 发送验证码 </a-button>
            </template>
          </a-input-search>
        </a-form-item>

        </a-form>

        <a-form-item label="密码" name="password">
          <a-input-password
            v-model:value="form.password"
            type="password"
            placeholder="请输入密码"
          >
            <template v-slot:prefix><safety-outlined type="safety" /></template>
          </a-input-password>
        </a-form-item>

        <a-form-item label="确认密码" name="repassword">
          <a-input-password
            v-model:value="form.repassword"
            type="password"
            placeholder="请输入密码"
          >
            <template v-slot:prefix><safety-outlined type="safety" /></template>
          </a-input-password>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 8, offset: 3 }">
          <a-button type="primary" @click="onSubmit('valid_form')"> 提交 </a-button>
          <a-button style="margin-left: 10px">
            <router-link to="/">返回</router-link>
          </a-button>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script>
import qs from "qs";
import {
  MailOutlined,
  SafetyOutlined,
  UnlockOutlined,
} from "@ant-design/icons-vue";

export default {
  name: "Login",
  components: {
    // UserOutlined,
    MailOutlined,
    UnlockOutlined,
    // CodeOutlined,
    SafetyOutlined,
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
      wrapperCol: { span: 14, offset: 4 },
      form: {
        mail: "",
        password: "",
        repassword: "",
        captche: "",
      },
      rules: {
        mail: [
          { required: true, message: "邮箱不能为空", trigger: "blur" },
          {validator:validateEMail,trigger:"blur"}
          
        ],
        captche:[
          {required:true,message:"验证码不能为空",trigger:"blur"}
        ],
        password:[
          {required:true,message:"密码不能为空",trigger:"blur"},
          {min:6,message:"密码长度不能小于6",trigger:"blur"},
        ],
        repassword:[
          {required:true,message:"确认密码不能为空",trigger:"blur"},
          {min:6,message:"确认密码长度不能小于6",trigger:"blur"},
        ],
      },

      rules_email:{
        mail: [
          { required: true, message: "邮箱不能为空", trigger: "blur" },
          {validator:validateEMail,trigger:"blur"}
          
        ],
      }
    };
  },
  methods: {
    sendMail(valid_form) {

      this.$refs[valid_form].validate().then(()=> {
        this.$axios.post("/user/send_email",qs.stringify({
            email: this.form.mail,
          })
        )
        .then((rep) => {
          alert(rep.data.msg)
        })
        .catch((rep) => {
          console.log(rep);
        });
      }).catch(()=> {
        alert("邮箱格式不正确")
      })
      
    },

    onSubmit(valid_form) {
      this.$refs[valid_form].validate().then(()=>{
  
        this.$axios.post("/user/front_user_register",qs.stringify({
            email: this.form.mail,
            captche:this.form.captche,
            password:this.form.password,
            repassword:this.form.repassword
          })
        )
        .then((rep) => {
          if (rep.data.code == 200){
            alert(rep.data.msg)
            this.$router.push("/login")
          }else {
            alert(rep.data.msg)
          }
        })
        .catch((rep) => {
          console.log(rep);
        });


      }).catch(()=>{

      })
    },
  },
};
</script>

<style scoped>
.app {
  margin: 50px 20%;
  padding: 20px;
  width: 100%;
  background: url("../../assets/img/login.jpg");
  background-size: contain;
  background-repeat: no-repeat;
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
