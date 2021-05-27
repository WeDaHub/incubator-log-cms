<template>
  <div id="LogList">
    <el-table :data="logArry" style="width: 100%">
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">
          <i class="el-icon-time"></i>
          <span style="margin-left: 10px">{{ scope.row.DateTime }}</span>
        </template>
      </el-table-column>
      <el-table-column label="日志" width="280">
        <template slot-scope="scope">
          <el-popover trigger="hover" placement="top">
            <p width="180">{{ scope.row.Content }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">{{ scope.row.Content }}</el-tag>
            </div>
          </el-popover>
        </template>
      </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
            >查看</el-button
          >
        </template>
      </el-table-column>
      <!--
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleEdit(scope.$index, scope.row)"
            >编辑</el-button
          >
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.$index, scope.row)"
            >删除</el-button
          >
        </template>
      
      </el-table-column>
        -->
    </el-table>
    <!--
    <el-tag
      :key="tag"
      v-for="tag in pageNum"
      :disable-transitions="false"
      @click="handlePage(tag)"
    >
      {{ tag }}
    </el-tag>
-->
    <br />
    <el-input-number
      v-model="curPage"
      @change="handlePage()"
      :min="1"
      :max="pageNum"
      label="页码选择"
    ></el-input-number>
  </div>
</template>

<style>
.el-popover {
  max-width: 300px;
}
</style>

<script>
export default {
  data() {
    return {
      logList: null,
      logArry: [],
      curPage: 1,
      totalPage: 0,
      evryPage: 10,
      pageNum: 1
    };
  },
  mounted() {
    this.getAllLogLists();
    //alert("est");
  },
  methods: {
    getAllLogLists() {
      // GET log/a/list/${page}/${size}
      console.log("start getAllLogLists");
      var sessionKey = window.sessionStorage.getItem("authKey");
      console.log(sessionKey);

      this.$axios
        .get("/api/log/a/list/" + this.curPage + "/" + this.evryPage, {
          headers: { authorization: sessionKey }
        })
        .then(response => {
          var rspData = response.data;
          console.log(rspData);
          if (rspData.code === 0) {
            this.totalPage = rspData.data.count;
            this.logArry = rspData.data.datas;
            this.pageNum = Math.ceil(this.totalPage / this.evryPage);
            //logList.push();
            //this.imageCode = rspData.data.img;
            //this.imageId = rspData.data.id;
            //console.log(this.imageCode);
            console.log(this.logArry);
          }
        })
        .catch(error => {
          console.log(error);
        });
    },
    handleEdit(index, row) {
      console.log(index, row);
    },

    handleDelete(index, row) {
      console.log(index, row);
    },
    handlePage() {
      if (!!!this.curPage) {
        this.curPage = 1;
      }
      //this.curPage = tag;
      this.getAllLogLists();
    }
  }
};
</script>
