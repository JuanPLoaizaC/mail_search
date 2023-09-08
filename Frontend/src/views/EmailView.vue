<script setup>
import { useEmailStore } from '../stores/emails';
import { storeToRefs } from 'pinia';
import { useRouter } from "vue-router";

const router = useRouter();
const useEmail = useEmailStore();
const { emailsList, emailSelected } = storeToRefs(useEmail);

const redirectToContent = (email) => {
  router.push(`mails/${email.message_id}`);
  emailSelected.value = email;
};
</script>

<template>
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
              <tr class="hover:bg-slate-100"  v-for="email in emailsList" :key="email.message_id" @click="redirectToContent({ ... email })">
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
</template>