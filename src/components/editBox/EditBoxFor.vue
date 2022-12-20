<script setup lang="ts">
import { ref } from 'vue';
const num = ref(5);

const emit = defineEmits(['click']);

function handleClick(type: string) {
  emit('click', type, num.value);
}
</script>

<template>
  <el-space wrap>
    <el-input-number
      v-model="num"
      :min="1"
      :max="100"
      style="margin-right: 5px"
    />
    <el-button @click="() => handleClick('start')">开始循环</el-button>
    <el-button @click="() => handleClick('end')">结束循环</el-button>
  </el-space>
</template>

<style scoped></style>
