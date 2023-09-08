<script setup>
import { ref } from 'vue';
import { useEmailStore } from '../stores/emails';
import { storeToRefs } from 'pinia';
import { useRouter } from "vue-router";
import Spinner from '../components/Spinner.vue';
import { ChevronLeftIcon, ChevronRightIcon } from '@heroicons/vue/20/solid';

let start = ref(0);
let end = ref(10);
let page = ref(1);

const router = useRouter();
const useEmail = useEmailStore();
const { emailsList, emailSelected, loading } = storeToRefs(useEmail);

const redirectToContent = (email) => {
  router.push(`mails/${email.message_id}`);
  emailSelected.value = email;
};

const previous = () => {
  page.value --;
  start.value -= 10;
  end.value -= 10;
};

const next = () => {
  page.value ++;
  start.value += 10;
  end.value += 10;
}; 
</script>

<template>
  <Spinner v-if="loading" />
  <div class="flex flex-col pr-8 pl-8 mt-1" v-if="emailsList.length > 0">
    <div class="overflow-x-auto">
      <div class="p-1.5 w-full inline-block align-middle">
        <div class="overflow-hidden border rounded-lg">
          <table class="min-w-full divide-y divide-gray-200 mb-5">
            <thead class="bg-gray-50">
              <tr>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase"
                >
                  Date
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase"
                >
                  Subject
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase"
                >
                  From
                </th>
                <th
                  scope="col"
                  class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase"
                >
                  To
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr class="hover:bg-slate-100"  v-for="email in emailsList.slice(start, end)" :key="email.message_id" @click="redirectToContent({ ... email })">
                <td class="px-6 py-4 text-sm font-medium text-gray-800 whitespace-nowrap">
                  {{ email.date }}
                </td>
                <td class="px-6 py-4 text-sm font-medium text-gray-800 whitespace-nowrap">
                  {{ email.subject }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  {{ email.from }}
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  {{ email.to }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  <div class="flex items-center justify-between border-t border-gray-200 bg-white px-4 py-3 sm:px-6 mt-4" v-if="emailsList.length > 0">
    <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
      <div>
        <p class="text-sm text-gray-700">
          Showing
          {{ ' ' }}
          <span class="font-medium">{{ start + 1 }}</span>
          {{ ' ' }}
          to
          {{ ' ' }}
          <span class="font-medium">{{ end <= emailsList.length ? end : emailsList.length }}</span>
          {{ ' ' }}
          of
          {{ ' ' }}
          <span class="font-medium">{{ emailsList.length }}</span>
          {{ ' ' }}
          results
        </p>
      </div>
      <div>
        <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
          <button class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0" :disabled="page === 1" @click="previous">
            <span class="sr-only">Previous</span>
            <ChevronLeftIcon class="h-5 w-5" aria-hidden="true" />
          </button>
          <p aria-current="page" class="relative z-10 inline-flex items-center bg-indigo-600 px-4 py-2 text-sm font-semibold text-white focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">{{ page }}</p>
          <button class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0" :disabled="(emailsList.length / end) <= 1" @click="next">
            <span class="sr-only">Next</span>
            <ChevronRightIcon class="h-5 w-5" aria-hidden="true" />
          </button>
        </nav>
      </div>
    </div>
  </div>
</template>