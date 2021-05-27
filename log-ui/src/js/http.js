// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
/*
import Vue from "vue";
import axios from "axios";

Vue.config.productionTip = false;

Vue.prototype.$axios = axios; //设置全局引用
Vue.prototype.$tip = TipWindow; //设置全局引用

//配置请求的设置
axios.defaults.headers.post["Content-type"] = "application/json";

export function PostRequest(uri, rspOk, rspErr, bodyData, isAuth) {
    console.log(LoginData);
    var sessionKey = window.sessionStorage.getItem("authKey");
    var ReqData = {
        params: JSON.stringify(bodyData),
        headers: { Authorization: sessionKey }
    };
    this.$axios
        .post(uri, ReqData)
        .then(response => {
            var rspHeader = response.headers;
            console.log(rspHeader);
            var rspData = response.data;
            console.log(rspData);
            //登录成功
            if (rspData.code === 0) {
                rspOk();
            }
            //登录失败
            else {
                rspErr();
                //console.log(response);
            }
        })
        .catch(error => {
            //console.log(error);
        });
}
*/