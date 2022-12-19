<script setup lang="ts">
import { useLocalStorage } from '@vueuse/core';
import { onMounted, ref } from 'vue';
import { useFetch } from '../../http';

const { data, execute } = useFetch(`/`, { immediate: false }).get().json();

const emit = defineEmits(['click']);

interface SkillItem {
  id: string;
  value: string;
}

interface SkillFetchItem {
  id: string;
  name: string;
}

const state = ref('');
const skills = ref<SkillItem[]>([]);
const skillTags = useLocalStorage<SkillItem[]>('skillTags', []);

const loadError = () => {
  // @ts-ignore
  ElMessageBox.alert(
    '抱歉，未能成功加载数据库，所有功能都将无法使用',
    'Title',
    {
      confirmButtonText: '知道了',
      type: 'error',
    }
  );
};
const loadAll = () => {
  const localItem = window.localStorage.getItem('skill');
  if (!localItem) {
    fetchSkills();
    return [];
  } else {
    console.log(`使用 localStorage`);
    return JSON.parse(localItem);
  }
};

onMounted(() => {
  skills.value = loadAll();
});

const fetchSkills = () => {
  console.log(`加载数据库`);
  execute()
    .then(() => {
      console.log(`加载完成`);
      handleSkills();
      // @ts-ignore

      ElMessage.success(`更新完成`);
    })
    .catch((err) => {
      console.log(err);
      // @ts-ignore

      ElMessage.error(`更新失败，${err.message}`);
    });
};

const handleSkills = () => {
  const skillData: SkillFetchItem[] = data.value?.data?.skill;

  if (!skillData) {
    skills.value = [];
    throw new Error(`请求失败`);
  }
  const skillsValue = skillData
    .map((item) => {
      return {
        value: item.name,
        id: item.id,
      };
    })
    .filter((item) => {
      const id = Number(item.id);
      if (Number.isNaN(id)) {
        return false;
      }
      if (id < 0 || id > 5000) {
        return false;
      }
      const excludes = [2171, 2172, 2173, 2281, 2282, 2283, 2284];
      if (excludes.includes(id)) {
        return false;
      }
      if (item.value.endsWith('AI') || item.value.includes('测试')) {
        return false;
      }
      return true;
    });
  if (skillsValue.length > 2000) {
    window.localStorage.setItem('skill', JSON.stringify(skillsValue));
    skills.value = skillsValue;
  } else {
    loadError();
  }
};
const createFilter = (queryString: string) => {
  return (restaurant: SkillItem) => {
    return (
      restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) !== -1
    );
  };
};

const querySearchAsync = (queryString: string, cb: (arg: any) => void) => {
  const results = queryString
    ? skills.value.filter(createFilter(queryString))
    : skills.value;

  cb(results);
};

const handleSelect = (item: Record<keyof SkillItem, string>) => {
  console.log(item);
  state.value = '';
  skillTags.value.push(item);
};

const handleClose = (tag: string) => {
  skillTags.value.splice(
    skillTags.value.findIndex((el) => el.value === tag),
    1
  );
};

function handleClick(skill: SkillItem) {
  emit('click', skill.value, Number(skill.id));
}
</script>

<template>
  <el-row :gutter="20">
    <el-button @click="fetchSkills">强制更新数据库</el-button>
  </el-row>
  <el-row :gutter="20">
    <el-autocomplete
      v-model="state"
      :fetch-suggestions="querySearchAsync"
      placeholder="需要使用的技能"
      @select="handleSelect"
    />
  </el-row>
  <el-row :gutter="20">
    <el-tag
      v-for="tag in skillTags"
      :key="tag.id"
      class="mx-1"
      closable
      style="margin-left: 5px; cursor: pointer"
      @close="handleClose(tag.value)"
      @click="() => handleClick(tag)"
    >
      {{ tag.value }}
    </el-tag>
  </el-row>
</template>

<style scoped></style>
