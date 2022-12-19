<script setup lang="ts">
import { ref } from 'vue';
import EditBox from './components/EditBox.vue';
import PreviewBox from './components/PreviewBox.vue';

const code = ref('');

function hanldeChange(value: string) {
  console.log(value);
  code.value = value;
}
</script>

<template>
  <el-row :gutter="20">
    <el-col :span="12"><EditBox @change="hanldeChange" /></el-col>
    <el-col :span="12"><PreviewBox :code="code" /></el-col>
  </el-row>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
