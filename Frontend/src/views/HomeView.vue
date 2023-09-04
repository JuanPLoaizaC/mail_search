<script setup>
import Emails from '../components/Emails.vue';
import ListEmails from '../components/ListEmails.vue';
</script>

<template>
  <div class="grid grid-flow-row auto-rows-max pr-8 pl-8" style="margin-top: 2em">
    <label class="relative block">
      <span class="sr-only">Search</span>
      <span class="absolute inset-y-0 left-0 flex items-center pl-2">
        <svg class="h-5 w-5 fill-slate-300" viewBox="0 0 20 20"><!-- ... --></svg>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
          <path strokeLinecap="round" strokeLinejoin="round" d="M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z" />
        </svg>
      </span>
      <input class="placeholder:italic placeholder:text-slate-400 block bg-white w-full border border-slate-300 rounded-md py-2 pl-9 pr-3 shadow-sm focus:outline-none focus:border-sky-500 focus:ring-sky-500 focus:ring-1 sm:text-sm font-sans md:italic" placeholder="Search for anything..." type="text" name="search"/>
    </label>
  </div>
  <!--<Emails></Emails>-->
  <ListEmails></ListEmails>
</template>