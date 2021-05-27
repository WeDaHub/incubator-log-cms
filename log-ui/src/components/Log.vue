<template>
  <div id="logDiv" style="width:480px">
    <!--
    <div style="height:80px;"></div>
    <textarea
      v-model="logdata"
      style="border:0;border-radius:5px;background-color:rgba(241,241,241,.98);width: 480px;height: 320px;padding: 10px;resize: none;"
      placeholder="请输入日志"
    ></textarea>

    <button
      id="btnSub"
      v-on:click="onWriteLog()"
      style="margin: 15px 0 0 10px;height:41px;width:10%; background-color: #2D456B"
      data-loading="打开注册页面，请稍后..."
    >
      提交
    </button>

    <button
      id="btnSave"
      v-on:click="onSaveLog()"
      style="margin: 15px 0 0 10px;height:41px;width:10%;background-color: rgb(211, 165, 165)"
      data-loading="正在验证登录，请稍后..."
    >
      保存为草稿
    </button>-->
    <el-form ref="form" label-width="80px">
      <el-form-item label="时间">
        <el-input :value="curDate" :disabled="true"> </el-input>
      </el-form-item>

      <el-form-item label="日志内容">
        <el-input type="textarea" id="loginput" v-model="logdata"></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" v-on:click="onWriteLog()" style=" width: 20%"
      >提交</el-button
    >
  </div>
</template>

<script>
import Vue from "vue";

export default {
  data() {
    return {
      logdata: null,
      curDate: "20210525",
      curLogDate: ""
    };
  },
  mounted() {
    const nowDate = new Date();
    const date = {
      year: nowDate.getFullYear(),
      month: nowDate.getMonth() + 1,
      date: nowDate.getDate()
    };
    const newmonth = date.month > 10 ? date.month : "0" + date.month;
    const day = date.date > 10 ? date.date : "0" + date.date;
    //this.curData = date.year + "-" + newmonth + "-" + day;
    this.curDate = date.year + "-" + newmonth + "-" + day;
    this.curLogDate = date.year + newmonth + day;
  },
  methods: {
    onWriteLog() {
      console.log("start onWriteLog");
      var sessionKey = window.sessionStorage.getItem("authKey");
      console.log(sessionKey);
      var bodyData = {
        content: this.logdata
      };
      var ReqData = {
        params: JSON.stringify(bodyData),
        headers: { authorization: sessionKey }
      };
      console.log(ReqData);

      this.$axios
        .put("/api/log/a/" + this.curLogDate, JSON.stringify(bodyData), {
          headers: { authorization: sessionKey }
        })
        .then(response => {
          var rspData = response.data;
          console.log(rspData);
          if (rspData.code === 0) {
            //this.imageCode = rspData.data.img;
            //this.imageId = rspData.data.id;
            //console.log(this.imageCode);
            this.$message({
              message: "日志写入成功",
              type: "success"
            });
          }
        })
        .catch(error => {
          console.log(error);
        });
    },
    onSaveLog() {
      alert("scanWeekLogs");
    }
  }
};
</script>

<style>
/*引入css 样式文件*/
@import "../css/font-awesome.css";
@import "../css/login.css";
@import "../css/style.css";
@import "../css/bootstrap.min.css";
@import "../css/geek/style.css";

#loginput {
  height: 211px;
}
</style>
