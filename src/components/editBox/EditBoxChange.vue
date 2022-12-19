<script setup lang="ts">
import { ref } from 'vue';
const num = ref(2);

const emit = defineEmits(['click']);

function handleClick() {
  emit('click', num.value);
}
</script>

<template>
  <el-input-number v-model="num" :min="1" :max="6" style="margin-right: 5px" />
  <el-button @click="() => handleClick()">换宠</el-button>
</template>

<style scoped></style>
