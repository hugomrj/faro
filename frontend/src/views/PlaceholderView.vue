<!-- frontend/src/views/PlaceholderView.vue -->
<script setup lang="ts">
import { useRoute } from 'vue-router'

const route = useRoute()
const title = route.name?.toString().replace(/-/g, ' ').replace(/\b\w/g, l => l.toUpperCase()) || 'Módulo'
</script>

<template>
  <div class="flex flex-col items-center justify-center h-full p-6">
    <div class="text-6xl mb-4">🚧</div>
    <h1 class="text-2xl font-bold mb-2" style="color: var(--color-text-primary)">{{ title }}</h1>
    <p style="color: var(--color-text-secondary)">Este módulo está en construcción.</p>
  </div>
</template>
