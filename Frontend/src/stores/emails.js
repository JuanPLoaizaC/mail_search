import { ref } from 'vue';
import { defineStore } from 'pinia';

export const useEmailStore = defineStore('email', () => {
  const search = ref('');
  const emailsList = ref([]);
  const emailSelected = ref({});
  const loading = ref(false);

  return { search, emailsList, emailSelected, loading };
});
