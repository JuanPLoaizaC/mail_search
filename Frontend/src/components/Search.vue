<script setup>
import { useEmailStore } from '../stores/emails';
import { storeToRefs } from 'pinia';
import { useGetData } from "@/composables/getData"; 

const useEmail = useEmailStore();
const { search } = storeToRefs(useEmail);
const { getData } = useGetData();

const searchData = async () => {
  getData(search);
};

</script>

<template>
  <div class="grid grid-flow-row auto-rows-max pr-8 pl-8 mt-3">
    <div class="grid grid-cols-12 gap-4">
      <div class="col-span-11">
        <label class="relative block">
          <span class="absolute inset-y-0 left-0 flex items-center pl-2">
            <svg class="h-5 w-5 fill-slate-300" viewBox="0 0 20 20"></svg>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
              <path strokeLinecap="round" strokeLinejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
            </svg>
          </span>
          <input class="placeholder:italic placeholder:text-slate-400 block bg-white w-full border border-slate-300 rounded-md py-2 pl-9 pr-3 shadow-sm focus:outline-none focus:border-sky-500 focus:ring-sky-500 focus:ring-1 sm:text-sm font-sans md:italic" placeholder="Search for anything ..." type="text" name="search" v-model="search" />
        </label>
      </div>
      <div>        
        <button type="button" class="text-white bg-blue-700 hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 font-medium rounded-full text-sm px-5 py-2.5 text-center mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800" :disabled="search.length < 2 || search.length > 150" @click="searchData" >Search</button>
      </div>
    </div>
    <div v-if="search.length < 2 || search.length > 150">
      <p class="text-left  text-red-700">Must be between 2 and 150 characters</p>
    </div>
  </div>
</template>