<script setup lang="ts">
import { ref } from 'vue';
import { radixConvert } from '../../utils';

const value = ref('');
const code = ref();

const emit = defineEmits(['click']);

const options = [
  { value: '自定义', label: '自定义' },
  { value: 'C43E2', label: '嘲风' },
  { value: 'C43E4', label: '睚眦' },
  { value: 'C43E5', label: '精卫' },
  { value: 'C43E3', label: '狴犴' },
  { value: 'C43E9', label: '飞廉' },
  { value: 'C43F1', label: '毕方' },
  { value: 'C43F2', label: '蠃鱼' },
  { value: 'C43F3', label: '英招' },
  { value: 'C43F4', label: '巴蛇' },
  { value: 'C4407', label: '驳' },
  { value: 'C4408', label: '蛊雕' },
  { value: 'C4409', label: '朱厌' },
  { value: 'C440A', label: '夫诸' },
];

function handleChange() {
  if (value.value === '自定义') {
    return;
  }
  code.value = +radixConvert(value.value, 16, 10);
}

function handleClick() {
  emit(
    'click',
    code.value,
    options.find((el) => el.value === value.value)?.label
  );
}
</script>

<template>
  <el-space direction="horizontal" alignment="start" :size="30">
    <el-select
      v-model="value"
      @change="handleChange"
      class="m-2"
      placeholder="Select"
    >
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      />
    </el-select>
    <el-input v-model="code" style="width: 80px"></el-input>
    <el-button @click="handleClick">添加</el-button>
  </el-space>
</template>

<style scoped></style>
