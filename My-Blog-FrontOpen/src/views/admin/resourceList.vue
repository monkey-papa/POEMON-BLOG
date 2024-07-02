<template>
  <div>
    <div>
      <div class="handle-box">
        <el-select clearable v-model="pagination.resourceType" placeholder="资源类型" class="handle-select mrb10">
          <el-option key="1" label="用户头像" value="userAvatar"></el-option>
          <el-option key="2" label="文章封面" value="articleCover"></el-option>
          <el-option key="3" label="文章图片" value="articlePicture"></el-option>
          <el-option key="4" label="背景图片" value="webBackgroundImage"></el-option>
          <el-option key="5" label="网站头像" value="webAvatar"></el-option>
          <el-option key="6" label="随机头像" value="randomAvatar"></el-option>
          <el-option key="7" label="其他图片" value="randomCover"></el-option>
          <el-option key="8" label="画笔图片" value="graffiti"></el-option>
          <el-option key="9" label="评论图片" value="commentPicture"></el-option>
          <el-option key="10" label="表情包" value="internetMeme"></el-option>
          <el-option key="11" label="聊天群头像" value="im/groupAvatar"></el-option>
          <el-option key="12" label="群聊天图片" value="im/groupMessage"></el-option>
          <el-option key="13" label="朋友聊天图片" value="im/friendMessage"></el-option>
          <el-option key="14" label="音乐声音" value="funnyUrl"></el-option>
          <el-option key="15" label="音乐封面" value="funnyCover"></el-option>
          <el-option key="16" label="Love.Cover" value="love/bgCover"></el-option>
          <el-option key="17" label="Love.Man" value="love/manCover"></el-option>
          <el-option key="18" label="Love.Woman" value="love/womanCover"></el-option>
          <el-option key="19" label="收藏夹封面" value="favoritesCover"></el-option>
          <el-option key="20" label="照片墙" value="lovePhoto"></el-option>
          <el-option key="21" label="移动端照片" value="mobilePhoto"></el-option>
          <el-option key="22" label="风景背景图" value="sceneryPhoto"></el-option>
        </el-select>
        <el-button type="primary" icon="el-icon-search" @click="search()">搜索</el-button>
        <el-button type="primary" @click="addResources()">新增资源</el-button>
        <el-button type="danger" @click="clearSearch()">清除参数</el-button>
      </div>
      <el-table :data="resources" border class="table" header-cell-class-name="table-header">
        <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
        <el-table-column prop="userId" label="用户ID" align="center"></el-table-column>
        <el-table-column prop="mimeType" label="资源类型" align="center"></el-table-column>
        <el-table-column label="状态" align="center">
          <template slot-scope="scope">
            <el-tag :type="scope.row.status === false ? 'danger' : 'success'" disable-transitions>
              {{ scope.row.status === false ? "禁用" : "启用" }}
            </el-tag>
            <el-switch @click.native="changeStatus(scope.row)" v-model="scope.row.status"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="路径" align="center">
          <template slot-scope="scope">
            <template v-if="
                !$common.isEmpty(scope.row.mimeType)
              ">
              <el-image lazy :preview-src-list="[scope.row.path]" class="table-td-thumb" :src="scope.row.path" fit="cover"></el-image>
            </template>
            <template v-else>
              {{ scope.row.path }}
            </template>
          </template>
        </el-table-column>
        <el-table-column label="大小(KB)" align="center">
          <template slot-scope="scope">
            {{ Math.round(scope.row.size / 1024) }}
          </template>
        </el-table-column>
        <el-table-column prop="type" label="类型" align="center"></el-table-column>
        <el-table-column :formatter="$common.formatter" prop="createTime" label="创建时间" align="center"></el-table-column>
        <el-table-column label="操作" width="180" align="center">
          <template slot-scope="scope">
            <el-button type="text" icon="el-icon-delete" style="color: var(--orangeRed)" @click="handleDelete(scope.row)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination background layout="total, prev, pager, next" :current-page="pagination.current" :page-size="pagination.size" :total="pagination.total" @current-change="handlePageChange">
        </el-pagination>
      </div>
    </div>
    <el-dialog title="图片" :visible.sync="resourceDialog" width="50%" :append-to-body="true" @close="close" destroy-on-close center>
      <uploadPicture :isAdmin="true" :ResourceType="pagination.resourceType" @addPicture="addPicture" :maxSize="5" :maxNumber="2"></uploadPicture>
    </el-dialog>
  </div>
</template>
<script>
const uploadPicture = () => import('../common/uploadPicture')
export default {
  components: {
    uploadPicture
  },
  data() {
    return {
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        resourceType: ''
      },
      resources: [],
      resourceDialog: false
    }
  },
  created() {
    this.getResources()
  },
  methods: {
    close() {
      this.clearSearch()
    },
    clearSearch() {
      this.pagination = {
        current: 1,
        size: 10,
        total: 0,
        resourceType: ''
      }
      this.getResources()
    },
    handleDelete(item) {
      this.$confirm('确认删除资源？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      })
        .then(() => {
          this.$http
            .post(this.$constant.baseURL + '/delArticleImage/', { url: item.path, id: item.id }, true)
            .then(res => {
              this.pagination.current = 1
              this.getResources()
              this.$message({
                message: '删除成功！',
                type: 'success'
              })
            })
            .catch(error => {
              this.$message({
                message: error.message,
                type: 'error'
              })
            })
        })
        .catch(() => {
          this.$message({
            type: 'success',
            message: '已取消删除!'
          })
        })
    },
    addPicture(res) {},
    addResources() {
      if (this.$common.isEmpty(this.pagination.resourceType)) {
        this.$message({
          message: '请选择资源类型！',
          type: 'error'
        })
        return
      }
      this.resourceDialog = true
    },
    search() {
      this.pagination.total = 0
      this.pagination.current = 1
      this.getResources()
    },
    getResources() {
      this.$http
        .post(this.$constant.baseURL + '/resource/listResource/', this.pagination, true)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.resources = res.result[0].records.reverse()
            this.pagination.total = res.result[0].total
          }
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    changeStatus(item) {
      this.$http
        .get(
          this.$constant.baseURL + '/resource/changeResourceStatus/',
          {
            id: item.id,
            flag: item.status
          },
          true
        )
        .then(res => {
          this.$message({
            message: '修改成功！',
            type: 'success'
          })
        })
        .catch(error => {
          this.$message({
            message: error.message,
            type: 'error'
          })
        })
    },
    handlePageChange(val) {
      this.pagination.current = val
      this.getResources()
    }
  }
}
</script>
<style scoped>
.handle-box {
  margin-bottom: 20px;
}
.handle-select {
  width: 200px;
}
.table {
  width: 100%;
  font-size: 14px;
}
.mrb10 {
  margin-right: 10px;
  margin-bottom: 10px;
}
.table-td-thumb {
  display: block;
  margin: auto;
  width: 40px;
  height: 40px;
}
.pagination {
  margin: 20px 0;
  text-align: right;
}
.el-switch {
  margin: 5px;
}
</style>