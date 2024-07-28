<template>
  <div v-if="!$common.isEmpty(treeHoleList)">
    <div class="process-line">
      <div
        class="process-item"
        v-for="(treeHole, index) in treeHoleList"
        :key="index"
      >
        <div class="timeline-item-time">
          <span>
            {{ $common.getDateDiff(treeHole.createTime) }}
          </span>
          <span
            @click="deleteTreeHole(treeHole.id)"
            class="process-delete"
            v-if="
              !$common.isEmpty($store.state.currentUser) &&
              $store.state.currentUser.id === treeHole.userId
            "
          >
            <i class="iconfont icon-shanchu"></i>
          </span>
        </div>
        <div
          class="timeline-item-content shadow-box"
          v-html="treeHole.content"
        ></div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    treeHoleList: {
      type: Array,
    },
  },
  methods: {
    deleteTreeHole(id) {
      this.$emit("deleteTreeHole", id);
    },
  },
};
</script>
<style lang="scss" scoped>
.process-line {
  border-left: 2px solid var(--blue2);
  padding: 50px 20px 10px;
  margin-left: 20px;
  position: relative;
  &:before {
    content: "";
    width: 8px;
    height: 8px;
    border: 4px solid var(--maxLightRed);
    border-radius: 50%;
    position: absolute;
    top: 15px;
    left: -1px;
    transform: translateX(-50%);
    background-color: var(--white);
    animation: weiYanShadowFlashing 1.5s linear infinite;
  }
}
.process-item {
  position: relative;
  margin: 10px;
  color: var(--fontColor);
}
.timeline-item-time::before {
  position: absolute;
  top: 5px;
  left: -37px;
  width: 6px;
  height: 6px;
  border: 3px solid var(--blue2);
  border-radius: 50%;
  background: var(--white);
  content: "";
}
.timeline-item-content {
  padding: 12px 15px;
  margin: 10px 0 15px;
  border-radius: 10px;
  background-color: var(--pink);
  border: 1px solid var(--gray1);
  transition: all 0.3s ease;
  &:hover {
    border-color: var(--gray4);
  }
}
.process-delete {
  margin-left: 10px;
  transition: all 0.3s ease;
  i {
    vertical-align: -3px;
    color: var(--brown1);
  }
  &:hover i {
    color: var(--brown2);
  }
}
</style>
