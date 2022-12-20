export const radixConvert = (num: string, fromRadix: number, toRadix: number) =>
  parseInt(num, fromRadix).toString(toRadix);

export const num2Chinese = (num: number) => {
  const chinese = '零一二三四五六七八九';
  return chinese[num];
};

/**
 * 下载
 */
export function downloadFile(data: string, filename: string) {
  const blob = new Blob([data], {
    type: 'text/plain',
  });
  const a = document.createElement('a');
  a.download = filename;
  a.href = URL.createObjectURL(blob);
  a.click();
}
