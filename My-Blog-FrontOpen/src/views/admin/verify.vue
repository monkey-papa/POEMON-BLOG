<template>
  <div class="content">
    <div class="left">
      <img src="https://www.qiniuyun.png" class="people p-animtion" alt="people" />
      <img src="https://www.qiniuyun.png" class="sphere s-animtion" alt="sphere" />
    </div>
    <div class="right">
      <div class="top">
        <div class="top-item">
          <span class="top-text" @click="$router.push({ path: '/' })">首页</span>
        </div>
        <div class="top-item">
          <span class="top-text" @click="$router.push({ path: '/user' })">前台用户登录</span>
        </div>
      </div>
      <div class="form-wrappepr">
        <h1>私生活后台</h1>
        <input type="text" class="inputs user" v-model="account" placeholder="请输入账号" />
        <input type="password" class="inputs pwd" v-model="password" @keyup.enter="login()" placeholder="请输入密码" />
        <proButton :info="'提交'" @click.native="login()" :before="$constant.before_color_2" :after="$constant.after_color_2">
        </proButton>
      </div>
    </div>
  </div>
</template>
<script>
const proButton = () => import('../common/proButton')
export default {
  components: {
    proButton
  },
  data() {
    return {
      redirect: this.$route.query.redirect,
      account: '',
      password: ''
    }
  },
  methods: {
    login() {
      if (this.$common.isEmpty(this.account) || this.$common.isEmpty(this.password)) {
        this.$message({
          message: '请输入账号或密码！',
          type: 'error'
        })
        return
      }
      let user = {
        account: this.account.trim(),
        password: this.$common.encrypt(this.password.trim()),
        isAdmin: true
      }
      this.$http
        .post(this.$constant.baseURL + '/user/login/', user, true, false)
        .then(res => {
          if (!this.$common.isEmpty(res.data)) {
            localStorage.setItem('adminToken', res.data.accessToken)
            this.$store.commit('loadCurrentAdmin', res.data)
            this.account = ''
            this.password = ''
            this.$router.push({ path: this.redirect })
          }
        })
        .catch(error => {
          this.$message({
            message: '账号异常，可能由于您不是管理员或者是账号密码错误，请重新输入',
            type: 'error'
          })
        })
    }
  }
}
</script>

<style scoped>
:root {
  font-size: 15px;
}

body {
  margin: 0;
  min-height: 100vh;
  background-color: #abc6f8;
  background-image: radial-gradient(closest-side, rgb(255, 255, 255), rgba(235, 105, 78, 0)), radial-gradient(closest-side, rgb(250, 203, 203), rgba(243, 11, 164, 0)), radial-gradient(closest-side, rgb(237, 252, 202), rgba(254, 234, 131, 0)),
    radial-gradient(closest-side, rgb(197, 248, 241), rgba(170, 142, 245, 0)), radial-gradient(closest-side, rgb(206, 200, 243), rgba(248, 192, 147, 0));
  background-size: 130vmax 130vmax, 80vmax 80vmax, 90vmax 90vmax, 110vmax 110vmax, 90vmax 90vmax;
  background-position: -80vmax -80vmax, 60vmax -30vmax, 10vmax 10vmax, -30vmax -10vmax, 50vmax 50vmax;
  background-repeat: no-repeat;
  animation: 10s movement linear infinite;
}

body::after {
  content: '';
  display: block;
  position: fixed;
  width: 100%;
  height: 100%;
  top: 0;
  left: 0;
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
}

.content {
  width: 90vw;
  height: 90vh;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  z-index: 1;
  border-radius: 30px;
  background: rgba(255, 255, 255, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.18);
  display: flex;
}
.content .left {
  flex: 1;
  position: relative;
}
.content .left .sphere {
  position: absolute;
  left: 30%;
  width: 90%;
  z-index: 1;
  animation: sphereAnimation 2s;
  animation-fill-mode: forwards;
  animation-timing-function: ease;
}
.content .left .people {
  position: absolute;
  left: -50%;
  top: 20%;
  width: 70%;
  z-index: 2;
}
.content .left .p-animtion {
  animation: peopleAnimation 2s;
  animation-fill-mode: forwards;
  animation-timing-function: ease;
}
.content .left .p-other-animtion {
  animation-name: pOtherAnimation;
  animation-direction: alternate;
  animation-timing-function: linear;
  animation-iteration-count: infinite;
  animation-duration: 3s;
}
.content .left .s-animtion {
  animation: sphereAnimation 2s;
  animation-fill-mode: forwards;
  animation-timing-function: ease;
}
.content .left .s-other-animtion {
  animation-name: sOtherAnimation;
  animation-direction: alternate;
  animation-timing-function: linear;
  animation-iteration-count: infinite;
  animation-duration: 3s;
}
.content .right {
  flex: 1;
  position: relative;
  z-index: 12;
}
.content .right .top {
  width: 80%;
  margin-left: 38px;
  color: rgb(51, 52, 124);
  font-size: 20px;
  font-weight: 600;
  position: absolute;
  left: 50%;
  top: 5%;
  transform: translate(-50%, 0);
}
.content .right .top .top-item {
  color: var(--blue);
  float: left;
  width: 150px;
  height: 40px;
  line-height: 40px;
  text-align: center;
  margin-right: 10px;
  transition: 0.5s;
  border-radius: 50px;
  border: 2px solid var(--red);
}
.content .right .top .top-item:hover {
  border: 0;
  background-color: #fff;
  border-radius: 50px;
  box-shadow: -20px 10px 32px 1px rgba(182, 183, 185, 0.57);
}
.content .right .form-wrappepr {
  width: 60%;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  text-align: right;
}
.content .right .form-wrappepr h1 {
  float: left;
  margin: 30px 0;
}
.content .right .form-wrappepr .inputs {
  display: block;
  width: 100%;
  height: 70px;
  margin: 30px 0;
  border-radius: 10px;
  border: 0;
  background-color: rgb(210, 223, 237);
  color: rgb(80, 82, 84);
  outline: none;
  padding: 20px;
  box-sizing: border-box;
  font-size: 20px;
}
.content .right .form-wrappepr .tips {
  display: block;
  margin-top: -15px;
  color: rgb(160, 170, 182);
}
.content .right .form-wrappepr button {
  width: 100%;
  height: 50px;
  background-color: rgb(68, 96, 241);
  border-radius: 10px;
  font-size: 15px;
  color: #fff;
  border: 0;
  font-weight: 600;
  margin: 30px 0;
  box-shadow: -20px 28px 42px 0 rgba(62, 145, 255, 0.37);
}
.content .right .form-wrappepr .other-login .divider {
  width: 100%;
  margin: 20px 0;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.content .right .form-wrappepr .other-login .divider .line {
  display: inline-block;
  max-width: 35%;
  width: 35%;
  flex: 1;
  height: 1px;
  background-color: rgb(162, 172, 185);
}
.content .right .form-wrappepr .other-login .divider .divider-text {
  vertical-align: middle;
  margin: 0px 20px;
  display: inline-block;
  width: 150px;
  color: rgb(162, 172, 185);
  white-space: normal;
}
.content .right .form-wrappepr .other-login .other-login-wrapper {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.content .right .form-wrappepr .other-login .other-login-wrapper .other-login-item {
  width: 70px;
  padding: 10px;
  text-align: center;
  border-radius: 10px;
  font-weight: 600;
  color: rgb(51, 49, 116);
  margin: 0 10px;
  transition: 0.4s;
}
.content .right .form-wrappepr .other-login .other-login-wrapper .other-login-item img {
  width: 40px;
  height: 40px;
  vertical-align: middle;
}
.content .right .form-wrappepr .other-login .other-login-wrapper .other-login-item span {
  vertical-align: middle;
}
.content .right .form-wrappepr .other-login .other-login-wrapper .other-login-item:hover {
  width: 80px;
  height: 50%;
  background-color: #fff;
  border: 0;
  box-shadow: -20px 10px 32px 1px rgba(182, 183, 185, 0.37);
}

@keyframes movement {
  0%,
  100% {
    background-size: 130vmax 130vmax, 80vmax 80vmax, 90vmax 90vmax, 110vmax 110vmax, 90vmax 90vmax;
    background-position: -80vmax -80vmax, 60vmax -30vmax, 10vmax 10vmax, -30vmax -10vmax, 50vmax 50vmax;
  }
  25% {
    background-size: 100vmax 100vmax, 90vmax 90vmax, 100vmax 100vmax, 90vmax 90vmax, 60vmax 60vmax;
    background-position: -60vmax -90vmax, 50vmax -40vmax, 0vmax -20vmax, -40vmax -20vmax, 40vmax 60vmax;
  }
  50% {
    background-size: 80vmax 80vmax, 110vmax 110vmax, 80vmax 80vmax, 60vmax 60vmax, 80vmax 80vmax;
    background-position: -50vmax -70vmax, 40vmax -30vmax, 10vmax 0vmax, 20vmax 10vmax, 30vmax 70vmax;
  }
  75% {
    background-size: 90vmax 90vmax, 90vmax 90vmax, 100vmax 100vmax, 90vmax 90vmax, 70vmax 70vmax;
    background-position: -50vmax -40vmax, 50vmax -30vmax, 20vmax 0vmax, -10vmax 10vmax, 40vmax 60vmax;
  }
}
@keyframes sphereAnimation {
  0% {
    width: 10%;
  }
  100% {
    width: 90%;
    transform: translate(-30%, 5%);
  }
}
@keyframes peopleAnimation {
  0% {
    width: 40%;
  }
  100% {
    width: 70%;
    transform: translate(90%, -10%);
  }
}
@keyframes pOtherAnimation {
  0% {
    transform: translate(90%, -10%);
  }
  100% {
    transform: translate(90%, -15%);
  }
}
@keyframes sOtherAnimation {
  0% {
    transform: translate(-30%, 5%);
  }
  100% {
    transform: translate(-30%, 10%);
  }
}
</style>
