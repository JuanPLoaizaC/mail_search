import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useEmailStore = defineStore('email', () => {
  const search = ref('');

  return { search };
})
