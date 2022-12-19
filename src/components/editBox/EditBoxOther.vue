<script setup lang="ts">
const emit = defineEmits(['click']);

function handleClick(type: string) {
  emit('click', type);
}
const buttons = ['逃跑', '使用小PP', '使用中PP', '使用大PP', '使用小HP'];
</script>

<template>
  <el-button v-for="btn of buttons" @click="() => handleClick(btn)">{{
    btn
  }}</el-button>
</template>

<style scoped></style>
