<script setup lang="ts">
interface Header {
  key: string
  label: string
}

interface Props {
  headers: Header[]
  rows: Record<string, any>[]
  emptyMessage?: string
  loading?: boolean
}

withDefaults(defineProps<Props>(), {
  emptyMessage: 'No hay datos para mostrar.',
  loading: false
})
</script>

<template>
  <div class="glass-table overflow-x-auto">
    <table class="w-full text-sm text-left">
      <thead class="text-xs uppercase" style="color: var(--color-text-secondary); background: var(--color-card-hover)">
        <tr>
          <th v-for="header in headers" :key="header.key" class="px-5 py-3 font-medium">
            {{ header.label }}
          </th>
          <th v-if="$slots.actions" class="px-5 py-3 font-medium text-right">Acciones</th>
        </tr>
      </thead>
      <tbody>
        <!-- Estado de Carga -->
        <tr v-if="loading">
          <td :colspan="headers.length + 1" class="px-5 py-10 text-center" style="color: var(--color-text-secondary)">
            <svg class="w-6 h-6 animate-spin mx-auto mb-2" style="color: var(--color-primary)" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Cargando...
          </td>
        </tr>

        <!-- Estado Vacío -->
        <tr v-else-if="rows.length === 0">
          <td :colspan="headers.length + 1" class="px-5 py-10 text-center" style="color: var(--color-text-secondary)">
            {{ emptyMessage }}
          </td>
        </tr>

        <!-- Filas Normales -->
        <tr 
          v-else
          v-for="(row, index) in rows" 
          :key="index" 
          class="transition-colors duration-150 hover:bg-[var(--color-sidebar-active)]"
          style="border-top: 1px solid var(--color-border)"
        >
          <td v-for="header in headers" :key="header.key" class="px-5 py-3.5" style="color: var(--color-text-primary)">
            {{ row[header.key] }}
          </td>
          <td v-if="$slots.actions" class="px-5 py-3.5 text-right">
            <slot name="actions" :row="row" />
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>