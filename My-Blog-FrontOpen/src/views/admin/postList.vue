<template>
  <div>
    <div class="handle-box">
      <el-select @blur="closeOptions('recommendStatus')" ref="fuzzyRecommendStatus" v-model="pagination.recommendStatus" placeholder="是否推荐" style="width: 120px" class="mrb10">
        <el-option key="1" label="是" :value="true"></el-option>
        <el-option key="2" label="否" :value="false"></el-option>
      </el-select>
      <el-select @blur="closeOptions('recommendSort')" ref="fuzzySort" style="width: 140px" class="mrb10" v-model="pagination.sortId" placeholder="请选择分类">
        <el-option v-for="item in sorts" :key="item.id" :label="item.sortName" :value="item.id">
        </el-option>
      </el-select>
      <el-select @blur="closeOptions('recommendLabel')" ref="fuzzyLabel" style="width: 140px" class="mrb10" v-model="pagination.labelId" placeholder="请选择标签">
        <el-option v-for="item in labelsTemp" :key="item.id" :label="item.labelName" :value="item.id">
        </el-option>
      </el-select>
      <el-input v-model="pagination.searchKey" placeholder="文章标题" class="handle-input mrb10"></el-input>
      <el-button type="primary" icon="el-icon-search" @click="searchArticles()">搜索</el-button>
      <el-button type="danger" @click="clearSearch()">清除参数</el-button>
      <el-button type="primary" @click="$router.push({ path: '/postEdit' })">新增文章</el-button>
    </div>
    <el-table :data="articles" border class="table" header-cell-class-name="table-header">
      <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
      <el-table-column prop="updateBy" label="作者" align="center"></el-table-column>
      <el-table-column prop="articleTitle" label="文章标题" align="center"></el-table-column>
      <el-table-column prop="sort[0].sortName" label="分类" align="center"></el-table-column>
      <el-table-column prop="label[0].labelName" label="标签" align="center"></el-table-column>
      <el-table-column prop="viewCount" label="浏览量" align="center"></el-table-column>
      <el-table-column prop="likeCount" label="点赞数" align="center"></el-table-column>
      <el-table-column label="是否可见" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.viewStatus === false ? 'danger' : 'success'" disable-transitions>
            {{ scope.row.viewStatus === false ? "不可见" : "可见" }}
          </el-tag>
          <el-switch @click.native="changeStatus(scope.row, 1)" v-model="scope.row.viewStatus"></el-switch>
        </template>
      </el-table-column>
      <el-table-column label="封面" align="center">
        <template slot-scope="scope">
          <el-image :preview-src-list="[scope.row.articleCover]" lazy class="table-td-thumb" :src="scope.row.articleCover" fit="cover"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="是否启用评论" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.commentStatus === false ? 'danger' : 'success'" disable-transitions>
            {{ scope.row.commentStatus === false ? "否" : "是" }}
          </el-tag>
          <el-switch @click.native="changeStatus(scope.row, 2)" v-model="scope.row.commentStatus"></el-switch>
        </template>
      </el-table-column>
      <el-table-column label="是否推荐" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.recommendStatus === false ? 'danger' : 'success'" disable-transitions>
            {{ scope.row.recommendStatus === false ? "否" : "是" }}
          </el-tag>
          <el-switch @click.native="changeStatus(scope.row, 3)" v-model="scope.row.recommendStatus"></el-switch>
        </template>
      </el-table-column>
      <el-table-column prop="commentCount" label="评论数" align="center"></el-table-column>
      <el-table-column :formatter="$common.formatter" prop="createTime" label="创建时间" align="center"></el-table-column>
      <el-table-column :formatter="$common.formatter" prop="updateTime" label="最终修改时间" align="center"></el-table-column>
      <el-table-column label="操作" width="135" align="center">
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
</template>

<script>
export default {
  data() {
    return {
      pagination: {
        current: 1,
        size: 20,
        total: 0,
        searchKey: '',
        recommendStatus: null,
        sortId: null,
        labelId: null
      },
      articles: [],
      sorts: [],
      labels: [],
      labelsTemp: [],
      recommendCount: 0
    }
  },
  watch: {
    'pagination.sortId'(newVal) {
      this.pagination.labelId = null
      if (!this.$common.isEmpty(newVal) && !this.$common.isEmpty(this.labels)) {
        this.labelsTemp = this.labels.filter(l => l.sortId === newVal)
      }
    }
  },
  created() {
    this.getArticles()
    this.getSortAndLabel()
  },
  methods: {
    getSortAndLabel() {
      this.$http
        .get(this.$constant.baseURL + '/webInfo/listSortAndLabel/')
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.sorts = res.result[0].data[0].sorts
            this.labels = res.result[0].data[0].labels
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
    clearSearch() {
      this.pagination = {
        current: 1,
        size: 20,
        total: 0,
        searchKey: '',
        recommendStatus: null,
        sortId: null,
        labelId: null
      }
      this.getArticles()
    },
    getArticles() {
      this.$http
        .post(this.$constant.baseURL + '/admin/article/boss/list/', this.pagination, true)
        .then(res => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.articles = res.result[0].records.reverse()
            this.pagination.total = res.result[0].total
            this.articles.forEach(item => {
              if (item.recommendStatus) {
                this.recommendCount += 1
              }
            })
          } else if (res.result.length === 0) {
            this.articles = []
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
    handlePageChange(val) {
      if (this.pagination.searchKey || String(this.pagination.recommendStatus) != 'null' || String(this.pagination.sortId) != 'null' || String(this.pagination.labelId) != 'null') {
        this.$notify({
          type: 'error',
          title: '可恶🤬',
          message: '请清除查询参数再进行操作！',
          position: 'top-left',
          offset: 50
        })
        return
      }
      this.pagination.current = val
      this.getArticles()
    },
    searchArticles() {
      this.pagination.total = 0
      this.pagination.current = 1
      this.getArticles()
    },
    changeStatus(article, flag) {
      let param
      if (flag === 1) {
        param = {
          articleId: article.id,
          viewStatus: article.viewStatus
        }
      } else if (flag === 2) {
        param = {
          articleId: article.id,
          commentStatus: article.commentStatus
        }
      } else if (flag === 3) {
        if (article.recommendStatus) {
          this.recommendCount++
          if (this.recommendCount === 11) {
            this.$notify({
              type: 'error',
              title: '可恶🤬',
              message: '已到推荐文章数量上限！最多推荐10篇文章！',
              position: 'top-left',
              offset: 50
            })
            this.recommendCount = 10
            article.recommendStatus = false
            return
          }
        } else {
          this.recommendCount--
        }
        param = {
          articleId: article.id,
          recommendStatus: article.recommendStatus
        }
      }
      this.$http
        .get(this.$constant.baseURL + '/admin/article/changeArticleStatus/', param, true)
        .then(res => {
          if (flag === 1) {
            this.$notify({
              type: 'warning',
              title: '淘气👻',
              message: '修改成功！注意，文章不可见时必须设置密码才能访问！',
              position: 'top-left',
              offset: 50
            })
          } else {
            this.$notify({
              title: '可以啦🍨',
              message: '修改成功！',
              type: 'success',
              offset: 50,
              position: 'top-left'
            })
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
    handleDelete(item) {
      this.$confirm('确认要删除这篇文章吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true
      })
        .then(() => {
          this.$http
            .get(this.$constant.baseURL + '/article/deleteArticle/', { id: item.id }, true)
            .then(res => {
              this.pagination.current = 1
              this.getArticles()
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
      this.$router.push({ path: '/postEdit', query: { id: item.id } })
    },
    closeOptions(item) {
      if (item === 'recommendStatus') {
        this.$refs.fuzzyRecommendStatus.blur()
      } else if (item === 'recommendSort') {
        this.$refs.fuzzySort.blur()
      } else {
        this.$refs.fuzzyLabel.blur()
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.handle-box {
  margin-bottom: 20px;
}
.handle-input {
  width: 160px;
  display: inline-block;
}
.table {
  width: 100%;
  font-size: 14px;
}
::v-deep .el-table__body-wrapper::-webkit-scrollbar {
  height: 8px;
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
