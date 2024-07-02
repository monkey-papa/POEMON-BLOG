<template>
  <div>
    <div>
      <el-tag effect="dark" class="my-tag">
        <img style="vertical-align: -3px" src="../../assets/svg/edit.svg" />
        统计信息
      </el-tag>
      <!-- 总览 -->
      <div>
        <div class="history-title">总览</div>
        <div>
          <div
            style="
              width: 400px;
              margin: 0 auto;
              display: flex;
              justify-content: center;
            "
          >
            <div class="history-name" style="line-height: 35px">总访问量:</div>
            <div
              style="
                color: var(--maxLightRed);
                font-weight: bold;
                font-size: 30px;
                line-height: 35px;
              "
            >
              {{ total_sum }}
            </div>
          </div>
          <div class="history-info" style="width: 640px">
            <div style="margin-right: 40px">
              <div class="history-name">省份访问TOP10</div>
              <div>
                <el-table :data="historyInfo.province_all_top">
                  <el-table-column type="index" align="center" width="60">
                  </el-table-column>
                  <el-table-column
                    prop="province"
                    align="center"
                    label="省份"
                    width="140"
                  >
                  </el-table-column>
                  <el-table-column
                    prop="count"
                    align="center"
                    label="数量"
                    width="100"
                  >
                  </el-table-column>
                </el-table>
              </div>
            </div>
            <div>
              <div class="history-name">IP访问TOP10</div>
              <div>
                <el-table :data="historyInfo.ip_all_top">
                  <el-table-column type="index" align="center" width="60">
                  </el-table-column>
                  <el-table-column
                    prop="ip"
                    align="center"
                    label="IP"
                    width="150"
                  >
                  </el-table-column>
                  <el-table-column
                    prop="count"
                    align="center"
                    label="数量"
                    width="100"
                  >
                  </el-table-column>
                </el-table>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- 今日访问 -->
      <div>
        <div class="history-title">今日访问</div>
        <div>
          <div
            style="
              width: 250px;
              margin: 0 auto;
              display: flex;
              justify-content: center;
            "
          >
            <div class="history-name" style="line-height: 35px">
              今日访问量：
            </div>
            <div
              style="
                color: var(--maxLightRed);
                font-weight: bold;
                font-size: 30px;
                line-height: 35px;
              "
            >
              {{ today_sum }}
            </div>
          </div>
          <div class="history-info" style="width: 640px">
            <div style="margin-right: 40px">
              <div class="history-name">今日访问省份统计</div>
              <div>
                <el-table :data="historyInfo.province_today">
                  <el-table-column type="index" align="center" width="60">
                  </el-table-column>
                  <el-table-column
                    prop="province"
                    align="center"
                    label="省份"
                    width="140"
                  >
                  </el-table-column>
                  <el-table-column
                    prop="count"
                    align="center"
                    label="数量"
                    width="100"
                  >
                  </el-table-column>
                </el-table>
              </div>
            </div>
            <div>
              <div class="history-name">今日访问用户</div>
              <div class="history-avatar">
                <el-table :data="historyInfo.user_today">
                  <el-table-column align="center" label="头像" width="100">
                    <template slot-scope="scope">
                      <el-avatar
                        class="user-avatar"
                        :size="30"
                        :src="scope.row.avatar"
                      >
                      </el-avatar>
                    </template>
                  </el-table-column>
                  <el-table-column
                    prop="userName"
                    align="center"
                    label="用户"
                    width="200"
                  >
                  </el-table-column>
                </el-table>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- 昨日访问 -->
      <div>
        <div class="history-title">昨日访问</div>
        <div>
          <div
            style="
              width: 250px;
              margin: 0 auto;
              display: flex;
              justify-content: center;
            "
          >
            <div class="history-name" style="line-height: 35px">
              昨日访问量：
            </div>
            <div
              style="
                color: var(--maxLightRed);
                font-weight: bold;
                font-size: 30px;
                line-height: 35px;
              "
            >
              {{ yesterday_sum }}
            </div>
          </div>
          <div class="history-info" style="width: 300px">
            <div>
              <div class="history-name">昨日访问用户</div>
              <div class="history-avatar">
                <el-table :data="historyInfo.last_user">
                  <el-table-column align="center" label="头像" width="100">
                    <template slot-scope="scope">
                      <el-avatar
                        class="user-avatar"
                        :size="30"
                        :src="scope.row.avatar"
                      >
                      </el-avatar>
                    </template>
                  </el-table-column>
                  <el-table-column
                    prop="userName"
                    align="center"
                    label="用户"
                    width="200"
                  >
                  </el-table-column>
                </el-table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      historyInfo: {},
      yesterday_sum: 0,
      today_sum: 0,
      total_sum: 0,
    };
  },
  created() {
    this.postProvinceAndCity();
    this.getHistoryInfo();
  },
  methods: {
    getHistoryInfo() {
      this.$http
        .get(this.$constant.baseURL + "/list/ip/")
        .then((res) => {
          if (!this.$common.isEmpty(res.result[0])) {
            this.historyInfo = res.result[0].data;
            this.yesterday_sum = res.result[0].yesterday_sum;
            this.today_sum = res.result[0].today_sum;
            this.total_sum = res.result[0].total_sum;
          }
        })
        .catch((error) => {
          this.$message({
            message: error.message,
            type: "error",
          });
        });
    },
    async postProvinceAndCity() {
      const res = await this.$common.getIpAndCity(this);
      this.$http
        .post(this.$constant.baseURL + "/submit/", {
          province: res.weather.province,
          city: res.city,
          userId: this.$store.state.currentUser.id,
        })
        .then((res) => {})
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
.my-tag {
  width: 100%;
  text-align: left;
  background: var(--lightYellow);
  border: none;
  height: 40px;
  line-height: 40px;
  font-size: 16px;
  color: black;
}
.el-tag {
  margin: 10px;
}
.history-title {
  margin: 15px auto 15px;
  width: 120px;
  text-align: center;
  padding: 10px 20px;
  background: var(--lightGreen);
  color: white;
  font-weight: bold;
  border-radius: 5px;
}
.history-name {
  font-size: 18px;
  font-weight: bold;
  margin: 0 10px 10px 0;
  text-align: center;
}
.history-info {
  display: flex;
  text-align: center;
  margin: 20px auto 0;
}
.history-info >>> .el-table .cell {
  line-height: unset;
}
.history-avatar >>> .el-table .el-table__row .el-table__cell {
  padding: 3.5px 0;
}
.history-info >>> .el-table::before {
  height: unset;
}
</style>
