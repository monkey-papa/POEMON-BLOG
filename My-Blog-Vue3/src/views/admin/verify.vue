<template>
  <div class="content transformCenter">
    <div class="left">
      <img src="../../assets/file/verify2.png" class="people p-animation" alt="people" />
      <img src="../../assets/file/verify1.png" class="sphere s-animation" alt="sphere" />
    </div>
    <div class="right">
      <div class="top">
        <div class="top-item">
          <span class="top-text" @click="router.push({ path: '/' })">首页</span>
        </div>
        <div class="top-item">
          <span class="top-text" @click="router.push({ path: '/user' })">前台用户登录</span>
        </div>
      </div>
      <div class="form-wrapper transformCenter">
        <h1>私生活后台</h1>
        <input type="text" class="inputs user" v-model="account" placeholder="请输入账号" />
        <input
          type="password"
          class="inputs pwd"
          v-model="password"
          @keyup.enter="login()"
          placeholder="请输入密码"
        />
        <proButton
          :info="'提交'"
          @click="login()"
          :before="$constant?.before_color_1"
          :after="$constant?.after_color_1"
        />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import api from "@/api";
import { ref, type Ref } from "vue";
import { defineAsyncComponent, type Component } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { useFormValidation } from "@/composables/useFormValidation";
import type { CurrentAdmin } from "@/types";

const store = useStore();
const { $route, $constant, $common } = useGlobalProperties();
const router = useRouter();
const { validateRequired } = useFormValidation();

const loadCurrentAdmin = (value: CurrentAdmin): void => {
  store.loadCurrentAdmin(value);
};

const proButton: Component = defineAsyncComponent(() => import("../common/proButton.vue"));

const account: Ref<string> = ref("");
const password: Ref<string> = ref("");

const login = async (): Promise<void> => {
  if (
    !validateRequired({
      [account.value]: "账号",
      [password.value]: "密码",
    })
  ) {
    return;
  }
  const user = {
    account: account.value.trim(),
    password: $common.encrypt(password.value.trim()),
    isAdmin: true,
  };

  try {
    const res = await api.adminLogin(user);
    loadCurrentAdmin(res);
    account.value = "";
    password.value = "";
    // 从路由参数获取 redirect，如果为空则默认跳转到后台首页
    const redirectPath = ($route.query.redirect as string) || "/backendMain";
    router.push({ path: redirectPath });
  } catch {
    /* ignored */
  }
};
</script>

<style lang="scss" scoped>
.content {
  width: 90vw;
  height: 90vh;
  z-index: 1;
  border-radius: 30px;
  background: rgba(255, 255, 255, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.18);
  display: flex;

  .left {
    flex: 1;
    position: relative;

    .sphere {
      position: absolute;
      left: 30%;
      width: 90%;
      z-index: 1;
      animation: sphereAnimation 2s;
      animation-fill-mode: forwards;
      animation-timing-function: ease;
    }

    .people {
      position: absolute;
      left: -50%;
      top: 20%;
      width: 70%;
      z-index: 2;
    }

    .p-animation {
      animation: peopleAnimation 2s;
      animation-fill-mode: forwards;
      animation-timing-function: ease;
    }

    .s-animation {
      animation: sphereAnimation 2s;
      animation-fill-mode: forwards;
      animation-timing-function: ease;
    }
  }

  .right {
    flex: 1;
    position: relative;
    z-index: 12;

    .top {
      width: 80%;
      margin-left: 38px;
      color: rgb(51, 52, 124);
      font-size: 20px;
      font-weight: 600;
      position: absolute;
      left: 50%;
      top: 5%;
      transform: translate(-50%, 0);

      .top-item {
        color: var(--blue2);
        float: left;
        width: 150px;
        height: 40px;
        line-height: 40px;
        text-align: center;
        margin-right: 10px;
        transition: 0.1s;
        border-radius: 50px;
        border: 2px solid var(--red);

        &:hover {
          border: 0;
          background-color: #fff;
          box-shadow: -20px 10px 32px 1px rgba(182, 183, 185, 0.57);
        }
      }
    }

    .form-wrapper {
      width: 60%;
      text-align: right;

      h1 {
        float: left;
        margin: 30px 0;
        color: var(--black5);
      }

      .inputs {
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
    }
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
</style>
