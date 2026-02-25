<template>
  <div class="tree-hole-container">
    <h3 class="my-center">哔哔~~~</h3>
    <ol class="tree-hole-list" v-if="!$common.isEmpty(props.treeHoleList)">
      <li class="tree-hole-li" v-for="(treeHole, index) in props.treeHoleList" :key="index">
        <div
          class="tree-hole-content"
          :class="{
            leftTreeHole: index % 2 === 0 && !$common.mobile(),
            rightTreeHole: index % 2 !== 0 || $common.mobile(),
          }"
        >
          <el-avatar shape="square" class="avatar-img" :size="36" :src="props.avatar" />
          <div
            class="tree-hole-box shadow-box"
            :style="{
              background:
                $constant?.tree_hole_color[index % ($constant?.tree_hole_color?.length || 1)],
            }"
          >
            <div
              class="box-tag"
              v-if="index % 2 === 0 && !$common.mobile()"
              :style="{
                'border-color':
                  'transparent transparent transparent ' +
                  $constant.tree_hole_color[index % $constant.tree_hole_color.length],
              }"
            ></div>
            <div
              class="box-tag"
              v-if="index % 2 !== 0 || $common.mobile()"
              :style="{
                'border-color':
                  'transparent ' +
                  $constant.tree_hole_color[index % $constant.tree_hole_color.length] +
                  ' transparent transparent',
              }"
            ></div>
            <div class="box-content" v-html="treeHole.content"></div>
            <div class="box-content-footer">
              <div>😃 {{ formatTime(treeHole.createTime) }}</div>
              <div
                @click="deleteTreeHole(treeHole.id)"
                class="tree-hole-delete"
                v-if="currentUser && currentUser.id === treeHole.userId"
              >
                <i class="iconfont icon-shanchu"></i>
              </div>
            </div>
          </div>
        </div>
      </li>
    </ol>
    <div class="tree-hole-go">
      <i class="fa fa-paper-plane" @click="launch()"></i>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "@/stores";

const store = useStore();

const currentUser = computed(() => store.currentUser);

interface TreeHoleDisplayItem {
  id: number;
  message: string;
  avatar: string;
  createTime: string;
  userId: number;
  username?: string;
  content?: string;
}

interface Props {
  treeHoleList: TreeHoleDisplayItem[];
  avatar?: string;
}

const props = defineProps<Props>();

interface Emits {
  (e: "launch"): void;
  (e: "deleteTreeHole", id: number): void;
}

const emit = defineEmits<Emits>();

const launch = (): void => {
  emit("launch");
};

const deleteTreeHole = (id: number): void => {
  emit("deleteTreeHole", id);
};

const formatTime = (row: string): string => {
  const day = row.split(".")[0].split("T")[0];
  const time = row.split(".")[0].split("T")[1];
  return `${day} 日 ${time}`;
};
</script>

<style lang="scss" scoped>
.tree-hole-container {
  padding: 20px;
  margin: 0 auto;

  > h3 {
    color: var(--bigRed1);
  }

  .tree-hole-list {
    padding: 100px 0 20px;
    margin: 0;
    position: relative;
    list-style: none;

    &:before {
      content: "";
      width: 4px;
      border-radius: 50%;
      position: absolute;
      top: 0;
      bottom: 0;
      left: 50%;
      transform: translateX(-50%);
      background-color: var(--orange2);
    }

    &:after {
      content: "";
      width: 12px;
      height: 12px;
      border: 4px solid var(--maxLightRed);
      border-radius: 50%;
      position: absolute;
      top: 10px;
      left: 50%;
      transform: translateX(-50%);
      background-color: var(--white);
      animation: weiYanShadowFlashing 1.5s linear infinite;
    }

    .tree-hole-li {
      margin: 5px auto;

      .tree-hole-content {
        position: relative;
        width: 50%;

        &:before {
          content: "";
          width: 12px;
          height: 12px;
          border: 4px solid var(--blue2);
          border-radius: 50%;
          position: absolute;
          top: 10px;
          background-color: var(--white);
        }

        &.leftTreeHole {
          text-align: right;

          &:before {
            right: 0;
            transform: translateX(10px);
          }

          .avatar-img {
            right: 25px;
          }

          .tree-hole-box {
            margin-right: 90px;
          }

          .box-tag {
            right: -10px;
            border-width: 15px 0 5px 10px;
          }
        }

        &.rightTreeHole {
          margin-left: 50%;

          &:before {
            left: 0;
            transform: translateX(-10px);
          }

          .avatar-img {
            left: 25px;
          }

          .tree-hole-box {
            margin-left: 90px;
          }

          .box-tag {
            left: -10px;
            border-width: 15px 10px 5px 0;
          }
        }

        .avatar-img {
          position: absolute;
          top: 0;
          transition: all 0.3s ease;
          border: 1px solid var(--gray1);

          &:hover {
            border-color: var(--gray4);
            transform: translateY(-5px);
            box-shadow: 0 0 16px 3px var(--miniMask);
          }
        }

        .tree-hole-box {
          font-size: 16px;
          padding: 10px;
          width: 360px;
          border-radius: 5px;
          position: relative;
          letter-spacing: 0.1em;
          font-weight: 400;
          transition: all 0.3s ease;
          color: var(--black);
          text-align: left;
          display: inline-block;

          &:hover {
            transform: translateY(-5px);
            box-shadow: 0 0 16px 3px var(--miniMask);
          }

          .box-tag {
            content: "";
            position: absolute;
            border-style: solid;
          }

          .box-content {
            margin: 0 10px 10px;
            line-height: 30px;
          }

          .box-content-footer {
            display: flex;
            justify-content: space-between;
            color: var(--fontColor);
            padding: 10px 10px 0;
            border-top: 1px dashed var(--white);
            font-size: 14px;

            .tree-hole-delete {
              font-size: 14px;
              transition: all 0.3s ease;

              i {
                color: var(--brown1);
                vertical-align: -2px;
              }

              &:hover i {
                color: var(--brown2);
              }
            }
          }
        }
      }
    }
  }

  .tree-hole-go {
    color: var(--blue2);
    font-weight: 700;
    font-size: 25px;
    margin: 20px auto;
    text-align: center;

    i:hover {
      animation: scale 1s linear infinite;
    }
  }
}

@media screen and (max-width: 1000px) {
  .tree-hole-container {
    .tree-hole-list {
      .tree-hole-li {
        .tree-hole-content {
          .tree-hole-box {
            width: calc(100% - 90px);
          }
        }
      }
    }
  }
}

@media screen and (max-width: 600px) {
  .tree-hole-container {
    .tree-hole-list {
      .tree-hole-li {
        .tree-hole-content {
          margin-bottom: 50px;
          width: 100%;

          &.rightTreeHole {
            margin-left: unset;
          }
        }
      }

      &:before {
        left: 0;
      }

      &:after {
        left: 0;
      }
    }
  }
}
</style>
