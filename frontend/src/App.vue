<!-- frontend/src/App.vue -->
<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <RouterView />
</template>

<style>
/* Reset básico */
html, body, #app {
  margin: 0;
  padding: 0;
  height: 100%;
  font-family: 'Inter', system-ui, -apple-system, sans-serif;
}
</style>