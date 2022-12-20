<script setup lang="ts">
import { ref } from 'vue';
import EditBox from './components/EditBox.vue';
import PreviewBox from './components/PreviewBox.vue';

const code = ref('');

function hanldeChange(value: string) {
  console.log(value);
  code.value = value;
}

// 获取视口宽度
const width = ref(window.innerWidth);
</script>

<template>
  <el-row v-if="width > 600" :gutter="20">
    <el-col :span="12"><EditBox @change="hanldeChange" /></el-col>
    <el-col :span="12"><PreviewBox :code="code" /></el-col>
  </el-row>
  <el-row v-else :gutter="20">
    <el-space direction="vertical">
      <el-col :span="24"><EditBox @change="hanldeChange" /></el-col>
      <el-col :span="24"><PreviewBox :code="code" /></el-col>
    </el-space>
  </el-row>
</template>

<style scoped></style>
