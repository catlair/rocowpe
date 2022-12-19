<script setup lang="ts">
import EditBoxBoss from './editBox/EditBoxBoss.vue';
import EditBoxSkill from './editBox/EditBoxSkill.vue';
import EditBoxOther from './editBox/EditBoxOther.vue';
import EditBoxFor from './editBox/EditBoxFor.vue';
import EditBoxChange from './editBox/EditBoxChange.vue';
import { useLocalStorage } from '@vueuse/core';

const emit = defineEmits(['change']);

const textarea = useLocalStorage('txt', '');
function handleForClick(type: 'start' | 'end', num: number) {
  console.log(type, num);
  if (type === 'start') {
    textarea.value += `开始循环${num}次\n`;
  } else {
    textarea.value += `结束循环\n`;
  }
}

function handleChangeClick(num: number) {
  console.log(num);
  textarea.value += `换到${num}宠\n`;
}

function handleOtherClick(num: number) {
  console.log(num);
  textarea.value += `${num}\n`;
}

function handleSkillClick(name: string, id: number) {
  console.log(name, id);
  textarea.value += `使用技能${name}\n`;
}

function handleBossClick(id: number, name: string) {
  console.log(id, name);
  if (name === '自定义') {
    textarea.value += `挑战${id}\n`;
  } else {
    textarea.value += `挑战${name}\n`;
  }
}

function handleTextareaChange(value: string) {
  emit('change', value);
}
</script>

<template>
  <el-row :gutter="20">
    <EditBoxBoss @click="handleBossClick" />
  </el-row>
  <EditBoxSkill @click="handleSkillClick" />
  <el-row :gutter="20">
    <EditBoxChange @click="handleChangeClick" />
  </el-row>
  <el-row :gutter="20">
    <EditBoxFor @click="handleForClick" />
  </el-row>
  <el-row :gutter="20">
    <EditBoxOther @click="handleOtherClick" />
  </el-row>
  <el-row :gutter="20">
    <el-input
      v-model="textarea"
      :autosize="{ minRows: 10 }"
      type="textarea"
      @change="handleTextareaChange"
      placeholder="Please input"
    />
    <el-button
      @click="() => handleTextareaChange(textarea)"
      style="margin-top: 10px"
      >生成</el-button
    >
  </el-row>
</template>

<style scoped></style>
