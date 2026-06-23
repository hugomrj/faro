<script setup lang="ts">
import { ref, computed } from 'vue'

interface Header {
  key: string
  label: string
}

interface Props {
  headers: Header[]
  rows: Record<string, any>[]
  emptyMessage?: string
  loading?: boolean
  placeholder?: string // Placeholder personalizable para el buscador
}

const props = withDefaults(defineProps<Props>(), {
  emptyMessage: 'No hay datos para mostrar.',
  loading: false,
  placeholder: 'Buscar...'
})

// Estado interno del buscador (no contamina las vistas)
const searchQuery = ref('')

// Filtrado automático basándose en las columnas visibles
const filteredRows = computed(() => {
  if (!searchQuery.value) return props.rows
  
  const query = searchQuery.value.toLowerCase()
  return props.rows.filter(row => {
    // Busca en todas las columnas definidas en los headers
    return props.headers.some(header => {
      const value = row[header.key]
      if (value === null || value === undefined) return false
      return String(value).toLowerCase().includes(query)
    })
  })
})
</script>

<template>
  <div>
    <!-- ESTÁNDAR FARO: Toda tabla tiene buscador -->
    <div class="mb-4 flex justify-end">
      <div class="relative w-full max-w-xs">
        <!-- Icono Lupa -->
        <svg class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4" style="color: var(--color-text-secondary)" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
        <input 
          type="text" 
          v-model="searchQuery" 
          class="input-field pl-10 w-full" 
          :placeholder="placeholder"
        />
      </div>
    </div>

    <!-- Contenedor de la Tabla Glass -->
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

          <!-- Estado Vacío (usando las rows FILTRADAS) -->
          <tr v-else-if="filteredRows.length === 0">
            <td :colspan="headers.length + 1" class="px-5 py-10 text-center" style="color: var(--color-text-secondary)">
              {{ searchQuery ? 'No se encontraron resultados.' : emptyMessage }}
            </td>
          </tr>

          <!-- Filas Normales (Iteramos las rows FILTRADAS) -->
          <tr 
            v-else
            v-for="(row, index) in filteredRows" 
            :key="index" 
            class="transition-colors duration-150 hover:bg-[var(--color-sidebar-active)]"
            style="border-top: 1px solid var(--color-border)"
          >

            <td v-for="header in headers" :key="header.key" class="px-5 py-3.5" style="color: var(--color-text-primary)">
              <!-- Si la vista define un slot para esta columna (ej: #cell-price), lo usa. Si no, pone el texto plano -->
              <slot :name="`cell-${header.key}`" :row="row" :value="row[header.key]">
                {{ row[header.key] }}
              </slot>
            </td>


            <td v-if="$slots.actions" class="px-5 py-3.5 text-right">
              <slot name="actions" :row="row" />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>