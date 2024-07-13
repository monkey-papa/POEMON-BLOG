<template>
  <div class="tree-hole-container">
    <h3 class="my-center" style="color: var(--bigRed)">哔哔~~~</h3>
    <ol class="tree-hole-list" v-if="!$common.isEmpty(treeHoleList)">
      <li class="tree-hole-li" v-for="(treeHole, index) in treeHoleList" :key="index">
        <div class="tree-hole-content" :class="{
            leftTreeHole: index % 2 === 0 && !$common.mobile(),
            rightTreeHole: index % 2 !== 0 || $common.mobile(),
          }">
          <el-avatar shape="square" class="avatar-img" :size="36" :src="avatar"></el-avatar>
          <div class="tree-hole-box" :style="{
              background:
                $constant.tree_hole_color[
                  index % $constant.tree_hole_color.length
                ],
            }">
            <div class="box-tag" v-if="index % 2 === 0 && !$common.mobile()" :style="{
                'border-color':
                  'transparent transparent transparent ' +
                  $constant.tree_hole_color[
                    index % $constant.tree_hole_color.length
                  ],
              }"></div>
            <div class="box-tag" v-if="index % 2 !== 0 || $common.mobile()" :style="{
                'border-color':
                  'transparent ' +
                  $constant.tree_hole_color[
                    index % $constant.tree_hole_color.length
                  ] +
                  ' transparent transparent',
              }"></div>
            <div class="my-content" v-html="treeHole.content"></div>
            <div style="display: flex; justify-content: space-between">
              <div>😃 {{ treeHole.createTime | formatter }}</div>
              <div @click="deleteTreeHole(treeHole.id)" class="tree-hole-delete" v-if="
                  !$common.isEmpty($store.state.currentUser) &&
                  $store.state.currentUser.id === treeHole.userId
                ">
                <img style="vertical-align: -2px; width: 18px; height: 18px" src="../../assets/svg/trash.svg" />
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

<script>
export default {
  props: {
    treeHoleList: {
      type: Array
    },
    avatar: {
      type: String
    }
  },
  methods: {
    launch() {
      this.$emit('launch')
    },
    deleteTreeHole(id) {
      this.$emit('deleteTreeHole', id)
    }
  },
  filters: {
    formatter(row) {
      const day = row.split('.')[0].split('T')[0]
      const time = row.split('.')[0].split('T')[1]
      return `${day} 日 ${time}`
    }
  }
}
</script>

<style lang="scss" scoped>
.tree-hole-container {
  padding: 20px;
  margin: 0 auto;
}
.tree-hole-list {
  padding: 100px 0 20px;
  margin: 0;
  position: relative;
  list-style: none;
  &:before {
    content: '';
    width: 4px;
    border-radius: 50%;
    position: absolute;
    top: 0;
    bottom: 0;
    left: 50%;
    transform: translateX(-50%);
    background-color: var(--orange);
  }
  &:after {
    content: '';
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
}
.tree-hole-li {
  margin: 5px auto;
}
.tree-hole-content {
  position: relative;
  width: 50%;
  &:before {
    content: '';
    width: 12px;
    height: 12px;
    border: 4px solid var(--blue);
    border-radius: 50%;
    position: absolute;
    top: 10px;
    background-color: var(--white);
  }
}
.leftTreeHole {
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
.rightTreeHole {
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
  transition: all 0.3s ease-in-out;
}
.tree-hole-box {
  font-size: 16px;
  padding: 10px;
  width: 360px;
  border-radius: 5px;
  position: relative;
  letter-spacing: 0.1em;
  font-weight: 400;
  transition: all 0.3s ease-in-out;
  color: var(--black);
  text-align: left;
  display: inline-block;
  > div:last-child {
    color: var(--fontColor);
    padding: 10px 10px 0;
    border-top: 1px dashed var(--white);
    font-size: 14px;
  }
}
.tree-hole-box:hover,
.avatar-img:hover {
  transform: translateY(-5px);
  box-shadow: 0 0 16px 3px var(--miniMask);
}
.box-tag {
  content: '';
  position: absolute;
  border-style: solid;
}
.my-content {
  margin: 0 10px 10px;
  line-height: 30px;
}
.tree-hole-go {
  color: var(--blue);
  font-weight: 700;
  font-size: 25px;
  margin: 20px auto;
  text-align: center;
  i:hover {
    animation: scale 1s linear infinite;
  }
}
.tree-hole-delete {
  font-size: 14px;
}
@media screen and (max-width: 1000px) {
  .tree-hole-box {
    width: calc(100% - 90px);
  }
}
@media screen and (max-width: 600px) {
  .tree-hole-content {
    margin-bottom: 50px;
  }
  .tree-hole-list:after {
    left: 0;
  }
  .tree-hole-list:before {
    left: 0;
  }
  .rightTreeHole {
    margin-left: unset;
  }
  .tree-hole-content {
    width: 100%;
  }
}
</style>
