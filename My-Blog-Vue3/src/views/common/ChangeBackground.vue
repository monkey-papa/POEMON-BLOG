<template>
  <!-- 修改背景 -->
  <el-dialog
    class="custom-my-dialog changeBgBox"
    title="切换背景"
    v-model="dialogVisible"
    :modal="false"
    width="60%"
    align="left"
  >
    <div class="changeBgBox-content">
      <button class="defaultBtn" @click="handleDefaultBtn">
        <i class="fa fa-refresh"></i>恢复默认主题
      </button>
      <div class="customImg">
        <div class="customImg-item">设置自定义背景</div>
        <el-input
          v-model="httpInput"
          type="text"
          class="customImg-input"
          maxlength="1000"
          placeholder="请输入有效的图片链接，如 https://source.fomal.cc/img/home_bg.webp"
        />
        <button class="httpButton" @click="handleHttpInputBtn">🌈切换背景🌈</button>
      </div>
      <div class="customImg">
        <div class="customImg-item setting-color">设置主题色</div>
        <div class="color-box_contain myBetween">
          <div
            @click="handleSetColor(item.color)"
            class="color-box"
            :class="{ active: themeActive === item.color }"
            :style="{ borderColor: item.varColor }"
            v-for="(item, i) in colorList"
            :key="i"
          >
            <div class="gun" :style="{ backgroundColor: item.varColor }"></div>
            <div class="center">
              <div class="top" :style="{ color: item.varColor }">
                {{ item.name }}
              </div>
              <div class="bottom" :style="{ color: item.varColor }">
                {{ item.color }} | {{ item.rgb }}
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="theme-box" v-for="(items, index) in themeMap" :key="index">
        <h2 class="theme-box-title">
          {{ items.title }}
        </h2>
        <div class="theme-box-content">
          <span class="rotate showIcon">
            <i class="iconfont icon-fengche"></i>
          </span>
          <el-collapse class="theme-box-collapse" @change="handleChangeBg(items.handleVal, index)">
            <el-collapse-item :title="items.collapseTitle" name="1">
              <div class="bgBox">
                <a
                  @click="emit('changeBg', item)"
                  v-for="(item, i) in items.dataList"
                  :key="i"
                  href="javascript:;"
                  :class="items.class"
                  :style="
                    items.style === 'img'
                      ? { backgroundImage: `url(${item})` }
                      : items.style === 'gradient'
                      ? { background: `linear-gradient(${item})` }
                      : { background: `${item}` }
                  "
                ></a>
              </div>
            </el-collapse-item>
          </el-collapse>
        </div>
      </div>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, computed, type Ref, type ComputedRef } from "vue";

interface ColorItem {
  name: string;
  color: string;
  rgb: string;
  rgba: string;
  varColor: string;
}

import type { ThemeMapConfigItem } from "@/types";

interface Props {
  modelValue?: boolean;
  themeMap: ThemeMapConfigItem[];
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: false,
});

interface Emits {
  (e: "update:modelValue", value: boolean): void;
  (e: "defaultBtn"): void;
  (e: "httpInputBtn", value: string): void;
  (e: "setColor", color: string): void;
  (e: "changeBg", item: string): void;
  (e: "handleChangeBg", val: string, index: number): void;
}

const emit = defineEmits<Emits>();

const httpInput: Ref<string> = ref("");
const themeActive: Ref<string> = ref(localStorage.getItem("themeColor") || "#425aef");

const colorList: ColorItem[] = [
  {
    name: "兔子坦克形态",
    color: "#04597b",
    rgb: "rgb(4,89,123)",
    rgba: "var(--sDarkBlueRgb)",
    varColor: "var(--sDarkBlue)",
  },
  {
    name: "鳄鱼恶霸形态",
    color: "#b04fe6",
    rgb: "rgb(176,79,230)",
    rgba: "var(--sPurpleRgb)",
    varColor: "var(--sPurple)",
  },
  {
    name: "巨龙熔岩形态",
    color: "#ff7500",
    rgb: "rgb(255, 117, 0)",
    rgba: "var(--sOrangeRgb)",
    varColor: "var(--sOrange)",
  },
  {
    name: "向日癸形态",
    color: "#ffc848",
    rgb: "rgb(255, 200, 72)",
    rgba: "var(--sYellowRgb)",
    varColor: "var(--sYellow)",
  },
  {
    name: "自然精灵形态",
    color: "#6bdf8f",
    rgb: "rgb(107,223,143)",
    rgba: "var(--sGreenRgb)",
    varColor: "var(--sGreen)",
  },
  {
    name: "锦鲤粉形态",
    color: "#ec695c",
    rgb: "rgb(236,105,92)",
    rgba: "var(--sRedRgb)",
    varColor: "var(--sRed)",
  },
  {
    name: "中国红形态",
    color: "#d61010",
    rgb: "rgb(214, 16, 16)",
    rgba: "var(--sBigRedRgb)",
    varColor: "var(--sBigRed)",
  },
  {
    name: "至尊龙形态",
    color: "#425aef",
    rgb: "rgb(66, 90, 239)",
    rgba: "var(--sBlueRgb)",
    varColor: "var(--sBlue)",
  },
];

const dialogVisible: ComputedRef<boolean> = computed({
  get(): boolean {
    return props.modelValue;
  },
  set(val: boolean): void {
    emit("update:modelValue", val);
  },
});

const handleDefaultBtn = (): void => {
  themeActive.value = "#425aef";
  emit("defaultBtn");
};

const handleHttpInputBtn = async (): Promise<void> => {
  emit("httpInputBtn", httpInput.value);
  httpInput.value = "";
};

const handleSetColor = (color: string): void => {
  themeActive.value = color;
  emit("setColor", color);
};

const handleChangeBg = (val: string, i: number): void => {
  emit("handleChangeBg", val, i);
};

// 暴露方法给父组件使用（写在后面，否则会报错）
defineExpose({
  handleSetColor,
});
</script>

<style lang="scss">
// 自定义背景对话框
.changeBgBox {
  background: var(--favoriteBg);
  height: 500px;
  padding: 0;

  .el-dialog__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    top: 0;
    z-index: 999;
    position: sticky;
    overflow: hidden;
    padding: 5px 20px 10px;
    background: var(--blue9);
    color: var(--favoriteBg);

    .el-dialog__headerbtn {
      width: 24px;
      height: 24px;
      position: relative;

      i {
        font-size: 24px;
        color: var(--favoriteBg);
        position: relative;
        right: 0;
      }
    }

    .el-dialog__title {
      color: var(--favoriteBg);
    }
  }

  .el-dialog__body {
    padding: 10px 15px;

    .changeBgBox-content {
      text-align: center;

      .defaultBtn {
        background: var(--blue9);
        display: block;
        width: 100%;
        padding: 15px 0;
        border-radius: 6px;
        color: var(--white2);
        border: none;
        font-size: 18px;

        .fa-refresh {
          font-size: 18px;
          margin-right: 5px;
        }
      }

      .customImg {
        font-size: 18px;
        margin-top: 15px;
        display: flex;
        flex-direction: column;
        align-items: flex-start;

        &-item {
          color: var(--blue6);

          &.setting-color {
            margin-bottom: 15px;
          }
        }

        .color-box_contain {
          flex-wrap: wrap;

          .color-box {
            font-size: 14px;
            width: 49%;
            padding: 9px 10px 9px 4px;
            display: flex;
            align-items: center;
            border: 2px dashed;
            border-radius: 0.5rem;
            margin-bottom: 10px;
            transition: all 0.3s ease;
            position: relative;

            &.active:before {
              position: absolute;
              right: 8px;
              bottom: 22px;
              font-size: 24px;
              content: "🌻";
              transform: scale(1);
            }

            &:hover {
              transform: scale(0.9);
            }

            .gun {
              width: 8px;
              height: 50px;
              border-radius: 0.375rem;
            }

            .center {
              padding-left: 10px;

              .top {
                font-weight: 400;
                line-height: 24px;
                text-align: left;
                margin-bottom: 4px;
              }

              .bottom {
                font-weight: 400;
                line-height: 24px;
              }
            }
          }
        }

        .customImg-input {
          .el-input__wrapper {
            margin: 15px auto 0;
            width: 80%;
            border-radius: 30px;
            border: 1px solid var(--blue6);
            padding: 5px 10px 5px 10px;
            line-height: 2;
            outline: 1px solid var(--blue6);

            &:hover {
              outline-color: var(--blue5);
              border-color: var(--blue5);
            }

            &.is-focus {
              outline-color: var(--blue5);
              border-color: var(--blue5);
            }
          }
        }
      }

      .httpButton {
        min-width: 110px;
        margin: 10px auto 0;
        background: var(--blue5);
        display: block;
        width: 25%;
        padding: 15px 0;
        border-radius: 30px;
        color: var(--white2);
        border: none;
        font-size: 15px;

        &:hover {
          background-color: var(--orange4);
          transition: all 0.3s ease;
        }
      }

      .theme-box {
        .theme-box-title {
          color: var(--fontColor);
          text-align: left;
        }

        .theme-box-content {
          display: flex;
          align-items: center;

          .showIcon {
            animation: rotate 1s linear infinite;
            .icon-fengche {
              display: flex;
              height: 16px;
              width: 16px;
              color: var(--red);
            }
          }
        }

        .theme-box-collapse {
          flex: 1;
          margin-left: 10px;
          border: none;

          .el-collapse-item__header {
            background: var(--gray5);
            padding: 16px;
            border-radius: 6px;
            color: var(--fontColor);
            font-size: 0.875rem;
            font-weight: 400;
            position: relative;
            border-bottom-color: var(--toolbarBackground1);
            line-height: normal;
          }

          .el-collapse-item__wrap {
            border-radius: 6px;

            .bgBox {
              display: flex;
              flex-wrap: wrap;
              justify-content: space-between;

              .box {
                width: 166px;
                margin: 10px;
                height: 100px;
                border-radius: 6px;
                border-bottom: none;
                background-size: cover;
                transition: all 1s;

                &:hover {
                  transform: scale(1.1);
                }
              }

              .mobileBox {
                height: 240px;
              }
            }
          }
        }
      }
    }
  }
}

// 媒体查询
@media screen and (max-width: 611px) {
  .changeBgBox {
    .el-dialog__body {
      .changeBgBox-content {
        .theme-box {
          .theme-box-collapse {
            .el-collapse-item__wrap {
              .bgBox {
                .box {
                  height: 73px;
                  width: 135px;
                }

                .mobileBox {
                  height: 240px;
                  width: 135px;
                }
              }
            }
          }
        }
      }
    }
  }

  .rotate.showIcon {
    display: none;
  }
}

@media screen and (max-width: 438px) {
  .changeBgBox {
    .el-dialog__body {
      .changeBgBox-content {
        .theme-box {
          .theme-box-collapse {
            .el-collapse-item__wrap {
              .bgBox {
                .box {
                  background-size: contain;
                  background-repeat: round;
                  width: 100%;
                }

                .mobileBox {
                  width: 100%;
                }
              }
            }
          }
        }
      }
    }
  }
}
</style>
