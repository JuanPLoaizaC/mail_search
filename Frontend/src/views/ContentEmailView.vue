<script setup>
import { onMounted } from 'vue';
import { PaperClipIcon } from '@heroicons/vue/20/solid';
import { useEmailStore } from '../stores/emails';
import { storeToRefs } from 'pinia';
import Spinner from '../components/Spinner.vue';

const useEmail = useEmailStore();
const { search, emailSelected, loading } = storeToRefs(useEmail);

onMounted(() => {
  highlightSearchTerm();
});

const highlightSearchTerm = () => {
  let ids = ['subject', 'date', 'from', 'to', 'content'];
  for (let index = 0; index < ids.length; index++) {        
    const contentElement = document.getElementById(ids[index]);
    if (contentElement) {
      const content = contentElement.textContent;
      const regex = new RegExp(`\\b(${search.value})\\b|(${search.value})`, 'gi');
      contentElement.innerHTML = content.replace(
        regex,
        "<span class='font-bold'>$&</span>"
      );
    }
  }
};
</script>

<template>
  <Spinner v-if="loading" />
  <div class="pr-8 pl-8 mb-2 mt-2">
    <div class="px-4 sm:px-0">
      <h3 class="text-base font-semibold leading-7 text-gray-900">Applicant Information</h3>
      <p class="mt-1 max-w-2xl text-sm leading-6 text-gray-500">Personal details and application.</p>
    </div>
    <div class="mt-6 border-t border-gray-100">
      <dl class="divide-y divide-gray-100">
        <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium leading-6 text-gray-900">Subject</dt>
          <dd id="subject" class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{{ emailSelected.subject }}</dd>
        </div>
        <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium leading-6 text-gray-900">Date</dt>
          <dd id="date" class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{{ emailSelected.date }}</dd>
        </div>
        <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium leading-6 text-gray-900">From</dt>
          <dd id="from" class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{{ emailSelected.from }}</dd>
        </div>
        <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium leading-6 text-gray-900">To</dt>
          <dd id="to" class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{{ emailSelected.to }}</dd>
        </div>
        <div class="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-0">
          <dt class="text-sm font-medium leading-6 text-gray-900">Content</dt>
          <dd id="content" class="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">{{ emailSelected.content }}</dd>
        </div>
      </dl>
    </div>
  </div>
</template>