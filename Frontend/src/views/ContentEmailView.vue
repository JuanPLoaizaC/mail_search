<script setup>
import { onMounted } from 'vue';
import { PaperClipIcon } from '@heroicons/vue/20/solid';
import { useEmailStore } from '../stores/emails';
import { storeToRefs } from 'pinia';
import Spinner from '../components/Spinner.vue';
import { useRouter } from 'vue-router';

const useEmail = useEmailStore();
const router = useRouter();
let { search, emailSelected, loading } = storeToRefs(useEmail);

onMounted(() => {
  highlightSearchTerm();
});

const backToEmails = () => {
  emailSelected = {};
  router.push('/mails');
};

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
  <!--<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
    <path strokeLinecap="round" strokeLinejoin="round" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18" />
  </svg>-->
  <div class="pr-8 pl-8 mb-2 mt-2">
    <div class="mb-4 sm:px-0 relative">
      <button @click="backToEmails()" >
        <span class="absolute inset-y-0 left-0 flex items-center">
          <svg class="h-10 w-10 fill-slate-300" viewBox="0 0 20 20"></svg>
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
            <path strokeLinecap="round" strokeLinejoin="round" d="M10.5 19.5L3 12m0 0l7.5-7.5M3 12h18" />
          </svg>
        </span>
      </button>
    </div>
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
  </div>
</template>