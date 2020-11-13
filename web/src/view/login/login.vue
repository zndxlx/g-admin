<template>
  <div class="login">
    <div class="container">
      <div class="logo_login">
        <img src="@/assets/logo_login.png" alt="" />
        <p class="title">G-Admin</p>
      </div>
      <el-form ref="form" :model="form" class="login-form">
        <el-form-item>
          <el-input
            v-model="loginForm.username"
            placeholder="请输入账号"
            suffix-icon="el-icon-user"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-input
            v-model="loginForm.password"
            placeholder="请输入密码"
            suffix-icon="el-icon-lock"
          ></el-input>
        </el-form-item>

        <el-form-item style="position: relative" class="form-item-code">
          <el-input
            v-model="loginForm.captcha"
            placeholder="请输入验证码"
            suffix-icon="el-icon-lock"
            style="width: 60%"
          ></el-input>
          <div class="vPic">
            <img
              v-if="picPath"
              :src="picPath"
              width="100%"
              height="100%"
              alt="验证码"
              @click="loginVefify()"
            />
          </div>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="login" style="width: 100%"
            >登陆</el-button
          >
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { captcha } from "@/api/user";
export default {
  name: "login",
  components: {},
  data() {
    return {
      msg: "hello",
      loginForm: {
        username: "admin",
        password: "111111",
        captcha: "",
        captchaId: "",
      },
      picPath: "",
    };
  },
  computed: {},
  methods: {
    ...mapActions("user", ["LoginIn"]),
    async login() {
      await this.LoginIn(this.loginForm);
      this.loginVefify();
    },
    loginVefify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath;
        this.loginForm.captchaId = ele.data.captchaId;
      });
    },
  },
  created() {
    this.loginVefify();
  },
  mounted() {},
};
</script>

<style scoped lang="scss">
.container {
  background: #f0f2f5 url(~@/assets/background.svg) no-repeat 50%;
  background-size: 100%;
  height: 100vh;

  .logo_login {
    padding-top: 80px;
    text-align: center;
    img {
      width: 100px;
    }
    .title {
      font-size: 33px;
      font-weight: 600;
    }
  }

  .login-form {
    // text-align: center;
    margin: 20px auto;
    width: 368px;
    .form-item-code {
      .vPic {
        right: 0;
        top: 0;
        position: absolute;
        width: 33%;
        height: 38px;
        background: pink;
      }
    }
  }
}
</style>