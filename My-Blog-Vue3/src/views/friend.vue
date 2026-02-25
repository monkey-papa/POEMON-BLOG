<template>
  <div class="friend-container">
    <!-- 背景图片 -->
    <div
      style="animation: header-effect 2s"
      :style="{ background: `${changeBgState}` }"
      class="background-image background-image-changeBg"
    ></div>
    <!-- 封面 -->
    <div class="friend-head myCenter">
      <h1>通讯录</h1>
    </div>
    <div class="friend-wrap">
      <div class="friend-main">
        <!-- 添加友链 -->
        <div @click="clickLetter()" class="form-wrap">
          <!-- 信封上面 -->
          <img class="before-img" src="../assets/file/letter-4.jpg" />
          <!-- 信 -->
          <div class="envelope" style="animation: hideToShow 2s">
            <div class="form-main shadow-box">
              <img class="envelope-img" src="../assets/file/friend_chain_bg.jpg" />
              <div class="envelope-content">
                <h3>有朋自远方来</h3>
                <div class="envelope-content-form">
                  <div class="myCenter form-friend">
                    <div class="user-title">
                      <div>名称：</div>
                      <div>简介：</div>
                      <div>封面：</div>
                      <div>头像：</div>
                      <div>网址：</div>
                    </div>
                    <div class="user-content">
                      <div>
                        <el-input maxlength="30" v-model="friend.title" />
                      </div>
                      <div>
                        <el-input maxlength="120" v-model="friend.introduction" />
                      </div>
                      <div>
                        <el-input maxlength="200" v-model="friend.cover" />
                      </div>
                      <div>
                        <el-input maxlength="200" v-model="friend.friendAvatar" />
                      </div>
                      <div>
                        <el-input maxlength="200" v-model="friend.url" />
                      </div>
                      <span class="friend-url-tip">请填写以http或https开头的有效地址</span>
                    </div>
                  </div>
                  <div class="submit-button-wrap myCenter">
                    <proButton
                      :info="'提交'"
                      @click.stop="submitFriend()"
                      :before="$constant.before_color_1"
                      :after="$constant.after_color_1"
                    />
                  </div>
                </div>
                <div class="envelope-footer">
                  <img class="envelope-footer-img" src="../assets/file/letter-3.jpg" />
                </div>
                <p class="envelope-footer-text">欢迎交换友链</p>
              </div>
            </div>
          </div>
          <img class="after-img" src="../assets/file/letter-2.jpg" />
        </div>
        <hr />
        <h2 class="friend-thanks-title"><i class="fa fa-address-card"></i>特别鸣谢</h2>
        <card
          class="recommendFriend"
          :resourcePathList="thanksFriendList"
          @clickResourcePath="clickFriend"
        />
        <h2 class="friend-title"><i class="fa fa-chain"></i>友情链接</h2>
        <card :resourcePathList="friendList" @clickResourcePath="clickFriend" />
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { notifySuccess } from "@/utils/notify";
import api from "@/api";
import { ref, onMounted, defineAsyncComponent, computed, type Ref, type ComputedRef } from "vue";
import { useStore } from "@/stores";
import { useGlobalProperties } from "@/composables/useGlobalProperties";
import { useFormValidation } from "@/composables/useFormValidation";
import type { ResourcePath, GetResourcePathListParams, SaveResourcePathParams } from "@/types";

interface FriendForm {
  title: string;
  introduction: string;
  cover: string;
  friendAvatar: string;
  url: string;
  type: string;
}

interface PaginationParams extends GetResourcePathListParams {
  resourceType: string;
  noPagination: boolean;
}

const store = useStore();
const { $constant } = useGlobalProperties();
const { validateRequiredWarn, validateUrl } = useFormValidation();

const changeBgState: ComputedRef<string> = computed(() => store.changeBg);

const card = defineAsyncComponent(() => import("./common/card.vue"));
const proButton = defineAsyncComponent(() => import("./common/proButton.vue"));

// 获取全部友链参数
const pagination: Ref<PaginationParams> = ref({
  resourceType: "friendUrl",
  noPagination: true,
});
const friendList: Ref<ResourcePath[]> = ref([]);
// 感谢友链列表（请替换为你自己的友链）
const thanksFriendList: Ref<ResourcePath[]> = ref([]);
const friend: Ref<FriendForm> = ref({
  title: "",
  introduction: "",
  cover: "",
  friendAvatar: "",
  url: "",
  type: "friendUrl",
});

onMounted(() => {
  getFriends();
});

const clickLetter = (): void => {
  if (document.body.clientWidth < 700) {
    $(".form-wrap").css({ height: "1000px", top: "-200px" });
  } else {
    $(".form-wrap").css({ height: "1150px", top: "-200px" });
  }
};

const submitFriend = async (): Promise<void> => {
  if (
    !validateRequiredWarn({
      [friend.value.title]: "写名称",
      [friend.value.introduction]: "写简介",
      [friend.value.cover]: "设置封面",
      [friend.value.friendAvatar]: "设置头像",
      [friend.value.url]: "写网址",
    })
  ) {
    return;
  }

  if (!validateUrl(friend.value.url, "请填写完整的有效网址，例如：http://****.com")) {
    return;
  }

  try {
    const resourcePath: SaveResourcePathParams = {
      title: friend.value.title,
      introduction: friend.value.introduction,
      cover: friend.value.cover,
      url: friend.value.url,
      type: friend.value.type,
    };
    await api.saveResourcePath(resourcePath);
    $(".form-wrap").css({ height: "447px", top: "0" });
    notifySuccess("提交成功，待管理员审核！");
  } catch {
    /* ignored */
  }
};

const clickFriend = (path: string): void => {
  if (path.includes("http")) {
    window.open(path);
  }
};

const getFriends = async (): Promise<void> => {
  try {
    const res = await api.getClientResourcePathList(pagination.value);
    friendList.value = res.list || [];
  } catch {
    /* ignored */
  }
};
</script>
<style lang="scss" scoped>
.friend-container {
  .friend-head {
    height: 300px;
    position: relative;

    h1 {
      color: var(--blue2);
      z-index: 10;
      font-weight: 400;
      font-size: 45px;
    }

    &::before {
      position: absolute;
      width: 100%;
      height: 100%;
      content: "";
    }
  }

  .friend-wrap {
    .friend-main {
      margin: 40px auto 0;
      padding: 40px;
      background: var(--background);

      h2 {
        font-size: 20px;
      }

      .friend-thanks-title {
        color: var(--red);
        position: relative;

        .fa-address-card {
          color: var(--green7);
          font-size: 24px;
          margin-right: 5px;
        }
      }

      .friend-title {
        color: var(--darkBlue);
        position: relative;

        .fa-chain {
          color: var(--green7);
          font-size: 30px;
          margin-right: 5px;
        }
      }

      hr {
        position: relative;
        margin: 40px auto;
        border: 2px dashed var(--blue2);
        overflow: visible;

        &:before {
          position: absolute;
          top: -21px;
          left: 5%;
          color: var(--red);
          content: "\e673";
          font-size: 40px;
          line-height: 1;
          transition: all 1s ease-in-out;
          font-family: iconfont;
        }

        &:hover:before {
          left: calc(95% - 20px);
        }
      }

      .form-wrap {
        margin: 0 auto;
        overflow: hidden;
        width: 530px;
        height: 447px;
        position: relative;
        top: 0;
        transition: all 1s ease-in-out 0.3s;
        z-index: 0;

        .before-img {
          position: absolute;
          bottom: 126px;
          left: 0;
          background-repeat: no-repeat;
          width: 100%;
          height: 317px;
          z-index: -100;
        }

        .after-img {
          position: absolute;
          bottom: -2px;
          left: 0;
          background-repeat: no-repeat;
          width: 530px;
          height: 259px;
          z-index: 100;
          width: 100%;
        }

        .envelope {
          position: relative;
          margin: 0 auto;
          transition: all 1s ease-in-out 0.3s;
          padding: 200px 20px 20px;

          .form-main {
            background: var(--pink2);
            margin: 0 auto;
            border-radius: 10px;
            overflow: hidden;
            transition: all 0.3s ease;
            border: 1px solid var(--gray1);

            .envelope-img {
              width: 100%;
            }

            &:hover {
              border-color: var(--gray4);
            }

            .envelope-content {
              h3 {
                text-align: center;
              }

              .envelope-content-form {
                .form-friend {
                  padding: 20px 0;

                  .user-title {
                    text-align: right;
                    user-select: none;
                    color: var(--white);

                    div {
                      height: 55px;
                      line-height: 55px;
                      text-align: center;
                    }
                  }

                  .user-content {
                    text-align: left;

                    .friend-url-tip {
                      color: var(--bigRed1);
                      font-size: 14px;
                    }

                    > div {
                      height: 55px;
                      display: flex;
                      align-items: center;

                      &:first-child {
                        margin-top: 20px;
                      }
                    }

                    :deep(.el-input__wrapper) {
                      border: none;
                      height: 35px;
                      background: var(--whiteMask);
                    }
                  }
                }

                .submit-button-wrap {
                  margin-top: 5px;
                }
              }

              .envelope-footer {
                .envelope-footer-img {
                  width: 100%;
                  margin: 5px auto;
                }
              }

              .envelope-footer-text {
                font-size: 12px;
                text-align: center;
                color: var(--bigRed1);
              }
            }
          }
        }
      }

      :deep(.recommendFriend) {
        .card-item {
          &::before {
            line-height: 62px;
            content: "推荐";
            position: absolute;
            z-index: 10;
            color: var(--white);
            top: -15px;
            letter-spacing: 3px;
            left: 15px;
            font-size: 15px;
            width: 44px;
            height: 50px;
            display: flex;
            justify-content: center;
            background: linear-gradient(90deg, var(--orange), var(--orange6));
            border-radius: 0px 0px 12px 12px;
          }
        }
      }
    }
  }
}

@media screen and (max-width: 700px) {
  .friend-container {
    .friend-wrap {
      .friend-main {
        .form-wrap {
          width: auto;
        }
      }
    }
  }
}

@media screen and (max-width: 500px) {
  .friend-container {
    .friend-wrap {
      padding: 0 10px;

      .friend-main {
        padding: 40px 15px;
      }
    }
  }
}
</style>
