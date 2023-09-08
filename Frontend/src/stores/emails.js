import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useEmailStore = defineStore('email', () => {
  const search = ref('');
  const emailsList = ref([]);
  const emailSelected = ref({});

  return { search, emailsList, emailSelected };
});
