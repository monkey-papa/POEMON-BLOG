<template>
  <div v-if="!$common.isEmpty(props.treeHoleList)">
    <div class="process-line">
      <div class="process-item" v-for="(treeHole, index) in props.treeHoleList" :key="index">
        <div class="timeline-item-time">
          <span>
            {{ $common.getDateDiff(treeHole.createTime) }}
          </span>
          <span
            @click="deleteTreeHole(treeHole.id)"
            class="process-delete"
            v-if="currentUser && currentUser.id === treeHole.userId"
          >
            <i class="iconfont icon-shanchu"></i>
          </span>
        </div>
        <div class="timeline-item-content shadow-box" v-html="treeHole.content"></div>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { computed } from "vue";
import { useStore } from "@/stores";

interface TreeHoleItem {
  id: number;
  userId: number;
  content: string;
  createTime: string;
}

interface Props {
  treeHoleList?: TreeHoleItem[];
}

const store = useStore();

const currentUser = computed(() => store.currentUser);

const props = defineProps<Props>();

const emit = defineEmits<{
  deleteTreeHole: [id: number];
}>();

const deleteTreeHole = (id: number): void => {
  emit("deleteTreeHole", id);
};
</script>
<style lang="scss" scoped>
.process-line {
  border-left: 2px solid var(--blue2);
  padding: 50px 20px 10px;
  margin-left: 20px;
  position: relative;

  &::before {
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

  .process-item {
    position: relative;
    margin: 10px;
    color: var(--fontColor);

    .timeline-item-time {
      display: flex;
      color: var(--bigRed1);
      align-items: center;

      &::before {
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
  }
}
</style>
