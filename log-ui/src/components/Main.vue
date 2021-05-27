<template>
  <div class="container" id="container">
    <aside class="left-bar" id="leftBar">
      <div class="title">
        <p>logcms首页</p>
      </div>
      <nav class="nav">
        <ul class="nav-item" id="navItem">
          <li>
            <a href="#tagwritelog" class="active"
              ><i class="iconfont icon-shequ"></i>写日志</a
            >
          </li>
          <li>
            <a href="#loghistory"
              ><i class="iconfont icon-blogger"></i>历史日报</a
            >
          </li>
          <li>
            <a href="#logcenter"><i class="iconfont icon-shequ"></i>常用功能</a>
          </li>
        </ul>
      </nav>
    </aside>
    <section class="main">
      <!-- 手机端菜单 -->
      <div id="menu"><a href="#">菜单</a></div>
      <div id="mainContent">
        <!-- 写日志 -->
        <div class="box">
          <a href="#" name="tagwritelog"></a>
          <div class="sub-category">
            <div><i class="iconfont icon-shequ"></i>写日志</div>
            <myLog style="height: 720px; width:100%"> </myLog>
          </div>
        </div>

        <!-- 个人中心 -->
        <div class="box">
          <a href="#" name="loghistory"></a>
          <div class="sub-category">
            <div><i class="iconfont icon-shequ"></i>历史日报</div>
            <myLogList style="height: 720px; width:100%"> </myLogList>
          </div>
        </div>

        <!-- 开发社区 -->
        <div class="box">
          <a href="#" name="logcenter"></a>
          <div class="sub-category" style="height:920px">
            <div><i class="iconfont icon-shequ"></i>常用功能</div>
            <div>
              <a target="_blank" v-on:click="writeLog()">
                <div class="item">
                  <el-button type="primary" icon="el-icon-edit-outline"
                    >写日志</el-button
                  >
                </div>
              </a>
              <a v-on:click="scanWeekLogs()">
                <div class="item">
                  <el-button type="primary" icon="el-icon-folder-opened"
                    >查看周报</el-button
                  >
                </div>
              </a>
              <a target="_blank" v-on:click="writeLog()">
                <div class="item">
                  <el-button type="primary" icon="el-icon-sell"
                    >日志导出</el-button
                  >
                </div>
              </a>
              <a v-on:click="scanWeekLogs()">
                <div class="item">
                  <el-button type="primary" icon="el-icon-s-order"
                    >批量管理</el-button
                  >
                </div>
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<style>
/*引入css 样式文件*/
@import "../css/geek/style.css";
</style>

<script>
import LogList from "@/components/LogList";
if (window.top !== window.self) {
  window.top.location = window.location;
}

export default {
  methods: {
    writeLog: function() {
      this.$router.push("/Log");
    },
    scanWeekLogs: function() {
      this.$router.push("/ScanWeekLogs");
    }
  }
};
</script>
