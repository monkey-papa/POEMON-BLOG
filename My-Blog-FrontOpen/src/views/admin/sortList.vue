<template>
  <div>
    <div style="margin-bottom: 20px">
      <el-button type="primary" @click="openSortDialog">新增分类</el-button>
    </div>
    <!-- 上方表格 -->
    <div style="margin: 10px auto; width: 50px; text-align: center">分类</div>
    <el-table
      :data="sortInfo"
      border
      class="table"
      header-cell-class-name="table-header"
    >
      <el-table-column
        prop="id"
        label="ID"
        width="55"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="sortName"
        label="分类名称"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="sortDescription"
        label="分类描述"
        align="center"
      ></el-table-column>
      <el-table-column label="分类类型" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.sortType === 0">导航栏分类</span>
          <span v-else-if="scope.row.sortType === 1">普通分类</span>
        </template>
      </el-table-column>
      <el-table-column
        prop="priority"
        label="分类优先级"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="lengthOfLabel"
        label="标签总数"
        align="center"
      ></el-table-column>
      <el-table-column
        prop="countOfSort"
        label="文章总数"
        align="center"
      ></el-table-column>
      <el-table-column label="操作" width="380" align="center">
        <template slot-scope="scope">
          <el-button
            type="text"
            icon="el-icon-edit"
            @click="editSortAndLabel(scope.row, '编辑分类')"
          >
            编辑分类
          </el-button>
          <el-button
            type="text"
            icon="el-icon-edit"
            @click="seeLabel(scope.row)"
          >
            查看标签
          </el-button>
          <el-button
            type="text"
            icon="el-icon-edit"
            @click="insertLabel(scope.row)"
          >
            新增标签
          </el-button>
          <el-button
            type="text"
            icon="el-icon-delete"
            style="color: var(--orangeRed)"
            @click="deleteHandle(scope.row.id, 1)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 标签弹窗 -->
    <el-dialog width="80%" title="标签" :visible.sync="dialogTableVisible">
      <el-table
        v-if="!$common.isEmpty(sort)"
        :data="sort.labels"
        border
        class="table"
        style="margin-top: 40px"
        header-cell-class-name="table-header"
      >
        <el-table-column
          prop="id"
          label="ID"
          width="55"
          align="center"
        ></el-table-column>
        <el-table-column label="分类名称" align="center">
          <span>{{ sort.sortName }}</span>
        </el-table-column>
        <el-table-column
          prop="labelName"
          label="标签名称"
          align="center"
        ></el-table-column>
        <el-table-column
          prop="labelDescription"
          label="标签描述"
          align="center"
        ></el-table-column>
        <el-table-column
          prop="countOfLabel"
          label="文章总数"
          align="center"
        ></el-table-column>
        <el-table-column label="操作" width="320" align="center">
          <template slot-scope="scope">
            <el-button
              type="text"
              icon="el-icon-edit"
              @click="editSortAndLabel(scope.row, '编辑标签')"
            >
              编辑标签
            </el-button>
            <el-button
              type="text"
              icon="el-icon-delete"
              style="color: var(--orangeRed)"
              @click="deleteHandle(scope.row.id, 2)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
    <!-- 分类弹窗 -->
    <el-dialog
      :title="title"
      :visible.sync="Dialog"
      width="30%"
      :before-close="handleClose"
      :append-to-body="true"
      destroy-on-close
      center
    >
      <div class="my-dialog">
        <template v-if="flag === '分类'">
          <div class="myCenter">
            <el-radio-group v-model="sortForHttp.sortType">
              <el-radio-button :label="0">导航栏分类</el-radio-button>
              <el-radio-button :label="1">普通分类</el-radio-button>
            </el-radio-group>
          </div>
          <el-input placeholder="请输入分类名称" v-model="sortForHttp.sortName">
            <template slot="prepend">分类名称</template>
          </el-input>
          <el-input
            placeholder="请输入分类描述"
            v-model="sortForHttp.sortDescription"
          >
            <template slot="prepend">分类描述</template>
          </el-input>
          <el-input
            type="number"
            placeholder="请输入整数，数字小的在前面"
            v-model="sortForHttp.priority"
          >
            <template slot="prepend">分类优先级</template>
          </el-input>
        </template>
        <template v-else>
          <div class="my-dialog">
            <el-input
              placeholder="请输入标签名称"
              v-model="labelForHttp.labelName"
            >
              <template slot="prepend">标签名称</template>
            </el-input>
            <el-input
              placeholder="请输入标签描述"
              v-model="labelForHttp.labelDescription"
            >
              <template slot="prepend">标签描述</template>
            </el-input>
          </div>
        </template>
      </div>

      <span slot="footer" class="dialog-footer">
        <el-button @click="handleClose()">取 消</el-button>
        <el-button type="primary" @click="saveEdit(flag)">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      Dialog: false,
      sortInfo: [],
      sort: {},
      sortForHttp: {
        id: null,
        sortName: "",
        sortDescription: "",
        sortType: null,
        priority: null,
      },
      labelForHttp: {
        id: null,
        sortId: null,
        labelName: "",
        labelDescription: "",
      },
      //控制打开对话框类型
      flag: "",
      //对话框标题
      title: "",
      //控制标签弹窗
      dialogTableVisible: false,
    };
  },
  created() {
    this.getSortInfo();
  },
  methods: {
    deleteHandle(id, flag) {
      let url;
      if (flag === 1) {
        this.title = "分类";
        url = "/webInfo/deleteSort/";
      } else if (flag === 2) {
        this.title = "标签";
        url = "/webInfo/deleteLabel/";
      } else {
        return;
      }
      this.$confirm(
        `确认要删除这个${this.title}吗？删除${this.title}后请编辑文章，修改文章的${this.title}ID！！`,
        "提示",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
          center: true,
        }
      )
        .then(() => {
          this.$http
            .get(this.$constant.baseURL + url, { id: id }, true)
            .then((res) => {
              this.$message({
                message: "删除成功！",
                type: "success",
              });
              this.getSortInfo();
              this.sort = {};
              this.dialogTableVisible = false;
            })
            .catch((error) => {
              this.$message({
                message: error.message,
                type: "error",
              });
              this.dialogTableVisible = false;
            });
        })
        .catch(() => {
          this.$message({
            type: "success",
            message: "已取消删除!",
          });
          this.dialogTableVisible = false;
        });
    },
    //保存编辑
    saveEdit(flag) {
      if (flag === "分类") {
        if (
          this.$common.isEmpty(this.sortForHttp.sortType) ||
          this.$common.isEmpty(this.sortForHttp.priority) ||
          this.$common.isEmpty(this.sortForHttp.sortName) ||
          this.$common.isEmpty(this.sortForHttp.sortDescription)
        ) {
          this.$message({
            message: "请完善所有分类信息！",
            type: "error",
          });
          return;
        }
        let url;
        if (this.$common.isEmpty(this.sortForHttp.id)) {
          url = "/webInfo/saveSort/";
        } else {
          url = "/webInfo/updateSort/";
        }
        this.$http
          .post(this.$constant.baseURL + url, this.sortForHttp, true)
          .then((res) => {
            this.$message({
              message: "保存成功！",
              type: "success",
            });
            this.getSortInfo();
            this.handleClose();
          })
          .catch((error) => {
            this.$message({
              message: error.message,
              type: "error",
            });
          });
      } else {
        if (
          this.$common.isEmpty(this.labelForHttp.labelName) ||
          this.$common.isEmpty(this.labelForHttp.labelDescription)
        ) {
          this.$message({
            message: "请完善所有标签信息！",
            type: "error",
          });
          return;
        }
        let url;
        if (this.$common.isEmpty(this.labelForHttp.id)) {
          url = "/webInfo/saveLabel/";
        } else {
          url = "/webInfo/updateLabel/";
        }
        this.$http
          .post(this.$constant.baseURL + url, this.labelForHttp, true)
          .then((res) => {
            this.$message({
              message: "保存成功！",
              type: "success",
            });
            this.getSortInfo();
            this.handleClose();
            this.sort = {};
          })
          .catch((error) => {
            this.$message({
              message: error.message,
              type: "error",
            });
          });
      }
      this.dialogTableVisible = false;
    },
    //打开新增分类对话框
    openSortDialog() {
      this.Dialog = true;
      //默认选中导航栏分类
      this.sortForHttp.sortType = 0;
      this.flag = "分类";
      this.title = "新增分类";
    },
    //表格中编辑分类和标签事件
    editSortAndLabel(SortOrLabel, flag) {
      if (flag === "编辑分类") {
        this.title = "编辑分类";
        this.flag = "分类";
        this.sortForHttp.id = SortOrLabel.id;
        this.sortForHttp.sortName = SortOrLabel.sortName;
        this.sortForHttp.sortDescription = SortOrLabel.sortDescription;
        this.sortForHttp.sortType = SortOrLabel.sortType;
        this.sortForHttp.priority = SortOrLabel.priority;
      } else {
        this.title = "编辑标签";
        this.flag = "编辑标签";
        this.labelForHttp.id = SortOrLabel.id;
        this.labelForHttp.sortId = SortOrLabel.sortId;
        this.labelForHttp.labelName = SortOrLabel.labelName;
        this.labelForHttp.labelDescription = SortOrLabel.labelDescription;
      }
      this.Dialog = true;
    },
    //新增标签
    insertLabel(sort) {
      this.flag = "新增标签";
      this.title = "新增标签";
      this.labelForHttp.sortId = sort.id;
      this.Dialog = true;
    },
    //关闭对话框
    handleClose() {
      this.labelForHttp = {
        id: null,
        sortId: null,
        labelName: "",
        labelDescription: "",
      };
      this.sortForHttp = {
        id: null,
        sortName: "",
        sortDescription: "",
        sortType: null,
        priority: null,
      };
      this.Dialog = false;
    },
    //查看标签
    seeLabel(sort) {
      this.dialogTableVisible = true;
      this.flag = "查看标签";
      this.sort = sort;
    },
    //得到分类信息
    getSortInfo() {
      this.$http
        .get(this.$constant.baseURL + "/webInfo/getSortInfo/")
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            const sortInfo = res.result[0].data.filter((item) => {
              return item.id !== 11;
            });
            this.sortInfo = sortInfo;
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
  },
};
</script>

<style scoped>
.my-dialog > div {
  margin: 12px;
}
.my-dialog >>> input::-webkit-inner-spin-button {
  appearance: none;
}
::v-deep .el-table {
  margin-top: 10px !important;
}
</style>
