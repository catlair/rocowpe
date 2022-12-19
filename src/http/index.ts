import { createFetch } from '@vueuse/core';

export const useFetch = createFetch({
  baseUrl: '/api',
  fetchOptions: {
    mode: 'cors',
  },
});
