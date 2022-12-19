export const radixConvert = (num: string, fromRadix: number, toRadix: number) =>
  parseInt(num, fromRadix).toString(toRadix);

export const num2Chinese = (num: number) => {
  const chinese = '零一二三四五六七八九';
  return chinese[num];
};
