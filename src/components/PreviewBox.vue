<script setup lang="ts">
import { ref, watch } from 'vue';
import { num2Chinese, radixConvert } from '../utils';
import str2gbk from 'str2gbk';

const wpe = ref('');

const props = defineProps<{
  code: string;
}>();

const bossMap: Record<string, string> = {
  嘲风: 'C43E2',
  睚眦: 'C43E4',
  精卫: 'C43E5',
  狴犴: 'C43E3',
  飞廉: 'C43E9',
  毕方: 'C43F1',
  蠃鱼: 'C43F2',
  英招: 'C43F3',
  巴蛇: 'C43F4',
  驳: 'C4407',
  蛊雕: 'C4408',
  朱厌: 'C4409',
  夫诸: 'C440A',
};

const skillMap: { id: string; value: string }[] = JSON.parse(
  window.localStorage.getItem('skill')!
);

let cur = 1;
const skillUse: string[][] = [];
// 道具
const itemUse: string[] = [];
watch(
  () => props.code,
  () => {
    cur = 1;
    let value = `
封包数量┆16┆统一间隔┆200┆回合间隔┆0┆发送次数┆10┆返回剧院┆False┆\r\n封包┆GET/cgi-bin/mystery_uncharted?cmd=1&index=1&unkown=&skey=&platfrom=2┆间隔┆0┆\r\n`;
    props.code.split('\n').forEach((line) => {
      if (!line) {
        return;
      }
      if (line.startsWith(`使用技能`)) {
        const skill = line.replace('使用技能', '');
        const skillObj = skillMap.find((el) => el.value === skill);
        if (!skillObj) {
          return;
        }
        value += `封包┆95270000000B0003000027100000000000000006010${cur}00${radixConvert(
          skillObj.id,
          10,
          16
        )
          .toUpperCase()
          .padStart(6, '0')}┆间隔┆0┆\r\n封包┆95270000000B000600002710000000000000000400000000┆间隔┆0┆\r\n`;
        skillUse[cur] ? skillUse[cur].push(skill) : (skillUse[cur] = [skill]);
        return;
      }
      if (line.startsWith(`挑战`)) {
        const name = line.replace(`挑战`, '');
        if (/\d+/.test(name)) {
          value += `封包┆95270000000B000100002710000000000000000902000${radixConvert(
            name,
            10,
            16
          ).toUpperCase()}00000000┆间隔┆0┆\r\n封包┆95270000000B000200002710000000000000000400000000┆间隔┆0┆\r\n`;
          return;
        }
        const boss = bossMap[name];
        if (!boss) {
          return;
        }
        value += `封包┆95270000000B000100002710000000000000000902000${boss.toUpperCase()}00000000┆间隔┆0┆\r\n封包┆95270000000B000200002710000000000000000400000000┆间隔┆0┆\r\n`;
        return;
      }
      if (line.startsWith('换到') && line.endsWith('宠')) {
        const t = +line.replace('换到', '').replace('宠', '');
        value += `封包┆95270000000B0003000027100000000000000006020${cur}0000000${t}┆间隔┆0┆\r\n`;
        cur = t;
        return;
      }
      if (line.startsWith('使用小PP')) {
        value += `封包┆95270000000B0003000027100000000000000006030${cur}01020002┆间隔┆0┆\r\n封包┆95270000000B000600002710000000000000000400000000┆间隔┆0┆\r\n`;
        itemUse.push('小瓶PP泉水');
        return;
      }
      if (line.startsWith('使用中PP')) {
        value += `封包┆95270000000B0003000027100000000000000006030${cur}01020003┆间隔┆0┆\r\n封包┆95270000000B000600002710000000000000000400000000┆间隔┆0┆\r\n`;
        itemUse.push('中瓶PP泉水');
        return;
      }
      if (line.startsWith('使用大PP')) {
        value += `封包┆95270000000B0003000027100000000000000006030${cur}01020004┆间隔┆0┆\r\n封包┆95270000000B000600002710000000000000000400000000┆间隔┆0┆\r\n`;
        itemUse.push('大瓶PP泉水');
        return;
      }
      if (line.startsWith('使用小HP')) {
        value += `封包┆95270000000B0003000027100000000000000006030${cur}01010002┆间隔┆0┆\r\n封包┆95270000000B000600002710000000000000000400000000┆间隔┆0┆\r\n`;
        itemUse.push('小瓶HP药水');
        return;
      }
      if (line.startsWith('开始循环') && line.endsWith('次')) {
        value += `封包┆System_开始循环[${line
          .replace('开始循环', '')
          .replace('次', '')}]┆间隔┆0┆\r\n`;
        return;
      }
      if (line.startsWith('结束循环')) {
        value += `封包┆System_结束循环┆间隔┆0┆\r\n`;
        return;
      }
      if (line.startsWith('逃跑')) {
        value += `封包┆System_宠物逃跑┆间隔┆0┆\r\n`;
        return;
      }
    });
    value += `封包┆GET/cgi-bin/mystery_uncharted?cmd=2&unkown=&skey=&platfrom=2&WinUp=1┆间隔┆0┆\r\n`;
    let skillUseString = '';
    skillUse.forEach((v, i) => {
      v = [...new Set(v)];
      skillUseString += `${num2Chinese(i)}宠技能【${v.join(
        '/'
      )}】推荐精灵:暂无推荐\r\n`;
    });
    const itemUseSet = [...new Set(itemUse)];
    skillUseString += `道具要求【${itemUseSet.join('/')}】\r\n`;
    wpe.value =
      `<Hint>下面是神辅生成的脚本说明:\r\n${skillUseString}\r\n</Hint>` + value;
  }
);

function handleClick() {
  const blob = new Blob([str2gbk(wpe.value)], {
    type: 'text/plain',
  });
  const a = document.createElement('a');
  a.download = 'wpe.txt';
  a.href = URL.createObjectURL(blob);
  a.click();
}
</script>

<template>
  <el-input
    v-model="wpe"
    :autosize="{ minRows: 10, maxRows: 30 }"
    type="textarea"
    placeholder="Please input"
  />
  <el-button @click="handleClick" style="margin-top: 10px">下载</el-button>
</template>

<style scoped></style>
