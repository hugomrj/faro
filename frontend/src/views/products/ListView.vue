<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useProducts } from '../../composables/useProducts'
import AppButton from '../../components/ui/AppButton.vue'
import AppTable from '../../components/ui/AppTable.vue'
import AppModal from '../../components/ui/AppModal.vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const { products, loading, fetchAll, remove } = useProducts()

const showDeleteModal = ref(false)
const selectedProduct = ref<any>(null)

const headers = [
  { key: 'name', label: 'Nombre' },
  { key: 'category_name', label: 'Categoría' },
  { key: 'price', label: 'Precio' },
  { key: 'stock', label: 'Stock' },
  { key: 'status', label: 'Estado' }
]

onMounted(() => fetchAll())

function confirmDelete(product: any) {
  selectedProduct.value = product
  showDeleteModal.value = true
}

async function handleDelete() {
  if (selectedProduct.value) {
    await remove(selectedProduct.value.id)
    showDeleteModal.value = false
  }
}
</script>

<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold" style="color: var(--color-text-primary)">Productos</h1>
        <p class="text-sm mt-1" style="color: var(--color-text-secondary)">Gestiona el inventario de tu negocio</p>
      </div>
      <AppButton @click="router.push('/products/create')">
        + Nuevo Producto
      </AppButton>
    </div>

    <!-- Tabla Glass -->
    <AppTable 
      :headers="headers" 
      :rows="products" 
      :loading="loading"
      empty-message="No hay productos creados aún."
    >
      <!-- FORMATEO ESPECIAL: Columna Precio (Se activa automáticamente por llamarse cell-price) -->
      <template #cell-price="{ value }">
        <span class="font-mono">$ {{ Number(value).toFixed(2) }}</span>
      </template>

      <!-- FORMATEO ESPECIAL: Columna Estado (Chip Glassmorphism) -->
      <template #cell-status="{ value }">
        <span 
          class="px-3 py-1 rounded-full text-xs font-semibold"
          :style="value === 'Activo' 
            ? 'background: rgba(16, 185, 129, 0.15); color: var(--color-primary-dark);' 
            : 'background: rgba(100, 116, 139, 0.15); color: var(--color-text-secondary);'"
        >
          {{ value }}
        </span>
      </template>

      <!-- Slot de acciones estándar -->
      <template #actions="{ row }">
        <div class="flex items-center justify-end gap-2">
          <AppButton variant="ghost" icon @click="router.push(`/products/edit/${row.id}`)">
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.862 4.487l1.687-1.688a1.875 1.875 0 112.652 2.652L10.582 16.07a4.5 4.5 0 01-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 011.13-1.897l8.932-8.931z" /></svg>
          </AppButton>
          <AppButton variant="ghost" icon @click="confirmDelete(row)">
            <svg class="w-4 h-4 text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0" /></svg>
          </AppButton>
        </div>
      </template>
    </AppTable>

    <!-- Modal de Eliminar -->
    <AppModal :is-open="showDeleteModal" title="Eliminar Producto" size="sm" @close="showDeleteModal = false">
      <p class="text-sm" style="color: var(--color-text-secondary)">
        ¿Estás seguro de eliminar <strong style="color: var(--color-text-primary)">{{ selectedProduct?.name }}</strong>? Esta acción no se puede deshacer.
      </p>
      <template #footer>
        <AppButton variant="secondary" @click="showDeleteModal = false">Cancelar</AppButton>
        <AppButton variant="danger" @click="handleDelete">Eliminar</AppButton>
      </template>
    </AppModal>
  </div>
</template>