<script setup>
import { RouterView } from 'vue-router';
import Navbar from './components/Navbar.vue';
import Search from './components/Search.vue';
</script>

<template>
  <Navbar />
  <Search />
  <RouterView />
</template>
