<template>
  <div>
    <div>
      <div class="handle-box">
        <el-select @clear="clear" clearable v-model="pagination.resourceType" placeholder="资源路径类型" class="handle-select mrb10">
          <el-option v-for="(item, i) in resourceTypes" :key="i" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
        <el-select @clear="clear" clearable v-model="pagination.status" placeholder="状态" class="handle-select mrb10">
          <el-option key="1" label="启用" :value="true"></el-option>
          <el-option key="2" label="禁用" :value="false"></el-option>
        </el-select>
        <el-button type="primary" icon="el-icon-search" @click="search()">搜索</el-button>
        <el-button type="primary" @click="addResourcePathDialog = true">新增资源路径</el-button>
      </div>
      <el-table :data="resourcePaths" border class="table" header-cell-class-name="table-header">
        <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
        <el-table-column prop="title" label="标题" align="center"></el-table-column>
        <el-table-column width="130" prop="classify" label="分类" align="center"></el-table-column>
        <el-table-column width="200" prop="introduction" label="简介" align="center"></el-table-column>
        <el-table-column label="封面" align="center">
          <template slot-scope="scope">
            <el-image lazy :preview-src-list="[scope.row.cover]" class="table-td-thumb" :src="scope.row.cover" fit="cover"></el-image>
          </template>
        </el-table-column>
        <el-table-column width="300" prop="url" label="链接" align="center"></el-table-column>
        <el-table-column prop="type" label="资源类型" align="center"></el-table-column>
        <el-table-column label="状态" align="center">
          <template slot-scope="scope">
            <el-tag :type="scope.row.status === false ? 'danger' : 'success'" disable-transitions>
              {{ scope.row.status === false ? "禁用" : "启用" }}
            </el-tag>
            <el-switch @click.native="changeStatus(scope.row)" v-model="scope.row.status"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="友人头像" align="center">
          <template slot-scope="scope" v-if="scope.row.type === 'friendUrl'">
            <el-image lazy :preview-src-list="[scope.row.friendAvatar]" class="table-td-thumb" :src="scope.row.friendAvatar" fit="cover"></el-image>
          </template>
        </el-table-column>
        <el-table-column :formatter="$common.formatter" prop="createTime" label="创建时间" align="center"></el-table-column>
        <el-table-column label="操作" width="180" align="center">
          <template slot-scope="scope">
            <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.row)">编辑</el-button>
            <el-button type="text" icon="el-icon-delete" style="color: var(--red)" @click="handleDelete(scope.row)">
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

    <el-dialog title="图片" :visible.sync="coverDialog" width="25%" :append-to-body="true" destroy-on-close center>
      <div>
        <uploadPicture :isAdmin="true" :ResourceType="resourcePath.type" @addPicture="addPicture" :maxSize="10" :maxNumber="1"></uploadPicture>
      </div>
    </el-dialog>

    <el-dialog title="文件" :visible.sync="uploadDialog" width="25%" :append-to-body="true" destroy-on-close center>
      <div>
        <uploadPicture :isAdmin="true" :ResourceType="resourcePath.type + 'Url'" @addPicture="addFile" :maxSize="10" :maxNumber="1" :listType="'text'" :accept="'image/*, video/*, audio/*'"></uploadPicture>
      </div>
    </el-dialog>

    <el-dialog title="资源路径" :visible.sync="addResourcePathDialog" width="50%" :before-close="clearDialog" :append-to-body="true" center>
      <div>
        <div>
          <div style="margin-bottom: 5px">标题：</div>
          <el-input maxlength="60" v-model="resourcePath.title"></el-input>
          <div style="margin-top: 10px; margin-bottom: 5px">分类：</div>
          <el-input :disabled="!['lovePhoto', 'funny', 'favorites'].includes(resourcePath.type)
            " maxlength="30" v-model="resourcePath.classify"></el-input>
          <div style="margin-top: 10px; margin-bottom: 5px">简介：</div>
          <el-input :disabled="!['friendUrl', 'favorites'].includes(resourcePath.type)" maxlength="1000" v-model="resourcePath.introduction"></el-input>
          <div style="margin-top: 10px; margin-bottom: 5px">封面：</div>
          <div style="display: flex">
            <el-input v-model="resourcePath.cover"></el-input>
            <div style="width: 66px; margin: 3.5px 0 0 10px">
              <proButton :info="'上传封面'" @click.native="addResourcePathCover()" :before="$constant.before_color_1" :after="$constant.after_color_1">
              </proButton>
            </div>
          </div>
          <div style="margin-top: 10px; margin-bottom: 5px">链接：</div>
          <div style="display: flex">
            <el-input :disabled="!['friendUrl', 'funny', 'favorites'].includes(resourcePath.type)
              " v-model="resourcePath.url"></el-input>
            <div style="width: 66px; margin: 3.5px 0 0 10px">
              <proButton :info="'上传文件'" @click.native="addResourcePathUrl()" :before="$constant.before_color_1" :after="$constant.after_color_1">
              </proButton>
            </div>
          </div>
          <div style="margin-top: 10px; margin-bottom: 5px">资源类型：</div>
          <el-select clearable v-model="resourcePath.type" placeholder="资源路径类型" class="handle-select mrb10">
            <el-option v-for="(item, i) in resourceTypes" :key="i" :label="item.label" :value="item.value">
            </el-option>
          </el-select>
          <div style="margin-top: 10px; margin-bottom: 5px">友人头像：</div>
          <div style="display: flex">
            <el-input :disabled="!['friendUrl'].includes(resourcePath.type)
              " v-model="resourcePath.friendAvatar"></el-input>
            <div style="width: 66px; margin: 3.5px 0 0 10px">
              <proButton :info="'上传文件'" @click.native="addResourcePathUrl()" :before="$constant.before_color_1" :after="$constant.after_color_1">
              </proButton>
            </div>
          </div>
        </div>
        <div style="display: flex; margin-top: 30px" class="myCenter">
          <proButton :info="'提交'" @click.native="addResourcePath()" :before="$constant.before_color_2" :after="$constant.after_color_2">
          </proButton>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const uploadPicture = () => import('../common/uploadPicture')
const proButton = () => import('../common/proButton')

export default {
  components: {
    uploadPicture,
    proButton
  },
  data() {
    return {
      resourceTypes: [
        { label: '友链', value: 'friendUrl' },
        { label: '恋爱图片', value: 'lovePhoto' },
        { label: '音乐', value: 'funny' },
        { label: '收藏夹', value: 'favorites' }
      ],
      pagination: {
        current: 1,
        size: 10,
        total: 0,
        resourceType: '',
        status: null
      },
      resourcePaths: [],
      coverDialog: false,
      uploadDialog: false,
      addResourcePathDialog: false,
      isUpdate: false,
      resourcePath: {
        title: '',
        classify: '',
        introduction: '',
        cover: '',
        url: '',
        type: '',
        friendAvatar: ''
      }
    }
  },
  created() {
    this.getResourcePaths()
  },
  methods: {
    addPicture(res) {
      this.resourcePath.cover = res
      this.coverDialog = false
    },
    addFile(res) {
      this.resourcePath.url = res
      this.uploadDialog = false
    },
    addResourcePathUrl() {
      if (this.addResourcePathDialog === false) {
        return
      }
      if (!['funny'].includes(this.resourcePath.type)) {
        this.$notify({
          type: 'error',
          title: '可恶🤬',
          message: '请选择有效资源类型！',
          position: 'top-left',
          offset: 50
        })
        return
      }
      this.uploadDialog = true
    },
    addResourcePathCover() {
      if (this.addResourcePathDialog === false) {
        return
      }
      if (this.$common.isEmpty(this.resourcePath.type)) {
        this.$notify({
          type: 'error',
          title: '可恶🤬',
          message: '请选择资源类型！',
          position: 'top-left',
          offset: 50
        })
        return
      }
      this.coverDialog = true
    },
    addResourcePath() {
      if (this.$common.isEmpty(this.resourcePath.title) || this.$common.isEmpty(this.resourcePath.type)) {
        this.$notify({
          type: 'error',
          title: '可恶🤬',
          message: '标题和资源类型不能为空！',
          position: 'top-left',
          offset: 50
        })
        return
      }
      this.$http
        .post(this.$constant.baseURL + (this.isUpdate ? '/webInfo/updateResourcePath/' : '/webInfo/saveResourcePath/'), this.resourcePath, true)
        .then(res => {
          this.$notify({
            title: '可以啦🍨',
            message: '保存成功！',
            type: 'success',
            offset: 50,
            position: 'top-left'
          })
          this.addResourcePathDialog = false
          this.clearDialog()
          this.search(this.pagination.current)
        })
        .catch(error => {
          this.$notify({
            type: 'error',
            title: '可恶🤬',
            message: error.message,
            position: 'top-left',
            offset: 50
          })
        })
    },
    search(cur) {
      this.pagination.total = 0
      this.pagination.current = cur || 1
      this.getResourcePaths()
    },
    getResourcePaths() {
      this.$http
        .post(this.$constant.baseURL + '/webInfo/listResourcePath/', this.pagination, true)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.resourcePaths = res.result[0].records.reverse()
            this.pagination.total = res.result[0].total
          }
        })
        .catch(error => {
          this.$notify({
            type: 'error',
            title: '可恶🤬',
            message: error.message,
            position: 'top-left',
            offset: 50
          })
        })
    },
    changeStatus(item) {
      this.$http
        .post(this.$constant.baseURL + '/webInfo/updateResourcePath/', item, true)
        .then(res => {
          this.$notify({
            title: '可以啦🍨',
            message: '修改成功！',
            type: 'success',
            offset: 50,
            position: 'top-left'
          })
        })
        .catch(error => {
          this.$notify({
            type: 'error',
            title: '可恶🤬',
            message: error.message,
            position: 'top-left',
            offset: 50
          })
        })
    },
    handlePageChange(val) {
      this.pagination.current = val
      this.getResourcePaths()
    },
    handleDelete(item) {
      this.$confirm('确认删除？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      })
        .then(() => {
          this.$http
            .get(this.$constant.baseURL + '/webInfo/deleteResourcePath/', { id: item.id }, true)
            .then(res => {
              this.search(this.pagination.current)
              this.$notify({
                title: '可以啦🍨',
                message: '删除成功！',
                type: 'success',
                offset: 50,
                position: 'top-left'
              })
            })
            .catch(error => {
              this.$notify({
                type: 'error',
                title: '可恶🤬',
                message: error.message,
                position: 'top-left',
                offset: 50
              })
            })
        })
        .catch(() => {
          this.$notify({
            title: '可以啦🍨',
            message: '已取消删除！',
            type: 'success',
            offset: 50,
            position: 'top-left'
          })
        })
    },
    handleEdit(item) {
      this.resourcePath = JSON.parse(JSON.stringify(item))
      this.addResourcePathDialog = true
      this.isUpdate = true
    },
    clearDialog() {
      this.isUpdate = false
      this.addResourcePathDialog = false
      this.resourcePath = {
        title: '',
        classify: '',
        introduction: '',
        cover: '',
        url: '',
        type: '',
        friendAvatar: ''
      }
    },
    clear() {
      this.pagination.status = null
      this.getResourcePaths()
    }
  }
}
</script>

<style lang="scss" scoped>
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
::v-deep .el-table__body-wrapper::-webkit-scrollbar {
  height: 8px;
}
</style>
