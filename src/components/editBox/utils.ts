import { SkillItem } from '../../types/skill';

/**
 * 过滤无效技能
 */
export function filterInvalidSkill(skills: SkillItem[]) {
  return skills.filter((item) => {
    const id = Number(item.id);
    if (Number.isNaN(id)) {
      return false;
    }
    if (id < 0 || id > 5000) {
      return false;
    }
    // 不是技能
    const excludes = [2171, 2172, 2173, 2281, 2282, 2283, 2284];
    if (excludes.includes(id)) {
      return false;
    }
    if (item.value.includes('AI') || item.value.includes('测试')) {
      return false;
    }
    return true;
  });
}
