<template>
  <div class="signinpanel" id="Login">
    <div class="col-md-12 col-md-offset-2">
      <div class="col-sm-8">
        <el-form id="signupForm">
          <h4
            class="no-margins"
            style="color: #415266;font-size:18pt;text-align:center"
          >
            日志系统
          </h4>
          <p class="m-t-md" style="color: #2d456b85;">&nbsp;</p>
          <div class="row">
            <div class="col-md-3 yonghutext">
              用户名:
            </div>
            <div class="col-md-9 yonghu">
              <input
                type="text"
                name="username"
                class="form-control uname"
                placeholder="用户名"
                v-model:value="username"
              />
            </div>
          </div>
          <div class="row ">
            <div class="col-md-3 yonghutext">
              密&nbsp;&nbsp;&nbsp;码:
            </div>
            <div class="col-md-9 yonghu">
              <input
                type="password"
                name="password"
                class="form-control pword m-b"
                placeholder="密码"
                v-model:value="passwd"
                autocomplete
              />
            </div>
          </div>

          <div class="row">
            <div class="col-md-3 yonghutext">
              验证码:
            </div>
            <div class="col-xs-5">
              <input
                type="text"
                name="validateCode"
                class="form-control code"
                v-model:value="userImgCode"
                placeholder="验证码"
                maxlength="6"
              />
            </div>
            <div class="col-xs-4">
              <a
                href="javascript:void(0);"
                title="点击更换验证码"
                v-on:click="onGetImgCode()"
              >
                <img
                  v-bind:src="imageCode"
                  id="imgcode"
                  class="imgcode"
                  width="100%"
                  height="30px"
                />
              </a>
            </div>
          </div>

          <el-button
            type="primary"
            id="btnSubmit"
            v-on:click="LoginIn()"
            data-loading="正在验证登录，请稍后..."
            style=" width: 30%"
          >
            登录
          </el-button>

          <el-button
            type="danger"
            id="btnSubReg"
            v-on:click="Register()"
            data-loading="打开注册页面，请稍后..."
            style=" width: 30%"
          >
            注册
          </el-button>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import md5 from "js-md5";

if (window.top !== window.self) {
  window.top.location = window.location;
}
export default {
  data() {
    return {
      imageCode: "",
      imageId: "",
      username: "",
      passwd: "",
      userImgCode: ""
    };
  },
  mounted() {
    this.onGetImgCode();
    //alert("est");
  },
  methods: {
    onLoginOk: function() {
      this.$router.push("/Main");
    },
    onGetImgCode: function() {
      console.log("start http");
      this.$axios
        .get("/api/captcha/f")
        .then(response => {
          var rspData = response.data;
          console.log(rspData);
          if (rspData.code === 0) {
            this.imageCode = rspData.data.img;
            this.imageId = rspData.data.id;
            console.log(this.imageCode);
          }
        })
        .catch(error => {
          console.log(error);
        });
    },

    LoginIn: function() {
      var LoginData = {
        account: this.username, //账号
        password: this.passwd, //密码
        code: this.userImgCode, //用户输入的验证码
        key: this.imageId //图片验证码返回的id
      };
      // console.log(LoginData);

      //alert("test");
      this.$axios
        .post("/api/account/f/login", JSON.stringify(LoginData))
        .then(response => {
          var rspHeader = response.headers;
          console.log(rspHeader);
          var rspData = response.data;
          console.log(rspData);

          //登录成功
          if (rspData.code === 0) {
            window.sessionStorage.setItem("authKey", rspHeader.authorization);
            var sessionKey = window.sessionStorage.getItem("authKey");
            console.log(sessionKey);
            this.$router.push("/Main");
          }
          //登录失败
          else {
            alert("登录失败");
            //console.log(response);
          }
        })
        .catch(error => {
          //console.log(error);
        });
    }
  } //end methods
};
</script>

<style>
/*引入css 样式文件*/
@import "../css/font-awesome.css";
@import "../css/login.css";
@import "../css/style.css";
@import "../css/bootstrap.min.css";
label.error {
  position: absolute;
  right: 48px;
  top: -2px;
  color: #ef392b;
  font-size: 12px;
}

.row {
  margin-right: 5px;
}
</style>
