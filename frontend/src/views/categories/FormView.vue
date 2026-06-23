<!-- frontend/src/views/categories/FormView.vue -->
<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useCategories } from '../../composables/useCategories'
import type { CategoryCreateRequest, CategoryUpdateRequest } from '../../types'

const route = useRoute()
const router = useRouter()
const { currentCategory, loading, error, fetchById, create, update } = useCategories()

const isEdit = computed(() => route.params.id !== undefined)
const categoryId = computed(() => Number(route.params.id))

const form = ref({
  name: '',
  description: ''
})
const saving = ref(false)

async function loadCategory() {
  if (!isEdit.value) return
  
  const cat = await fetchById(categoryId.value)
  if (cat) {
    form.value = {
      name: cat.name,
      description: cat.description
    }
  }
}

async function save() {
  // Validación
  if (!form.value.name.trim()) {
    error.value = 'El nombre es obligatorio'
    return
  }
  if (form.value.name.trim().length < 2) {
    error.value = 'El nombre debe tener al menos 2 caracteres'
    return
  }

  saving.value = true
  error.value = ''
  
  try {
    if (isEdit.value) {
      const data: CategoryUpdateRequest = {
        name: form.value.name.trim(),
        description: form.value.description.trim()
      }
      const success = await update(categoryId.value, data)
      if (success) {
        router.push({ name: 'categories-list' })
      }
    } else {
      const data: CategoryCreateRequest = {
        name: form.value.name.trim(),
        description: form.value.description.trim()
      }
      const id = await create(data)
      if (id !== null) {
        router.push({ name: 'categories-list' })
      }
    }
  } finally {
    saving.value = false
  }
}

function goBack() {
  router.push({ name: 'categories-list' })
}

onMounted(loadCategory)
</script>

<template>
  <div class="p-6">
    <!-- Header -->
    <div class="flex items-center gap-4 mb-6">
      <button @click="goBack" class="btn-secondary !px-3">
        ←
      </button>
      <div>
        <h1 class="text-2xl font-bold" style="color: var(--color-text-primary)">
          {{ isEdit ? 'Editar Categoría' : 'Nueva Categoría' }}
        </h1>
        <p class="text-sm mt-1" style="color: var(--color-text-secondary)">
          {{ isEdit ? `Editando categoría #${categoryId}` : 'Completa el formulario para crear una nueva categoría' }}
        </p>
      </div>
    </div>

    <!-- Error -->
    <div v-if="error" class="card mb-4 !border-red-500">
      <p class="text-red-600 dark:text-red-400 text-sm">{{ error }}</p>
    </div>

    <!-- Loading -->
    <div v-if="loading" class="flex items-center justify-center py-20">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-500"></div>
    </div>

    <!-- Formulario -->
    <div v-else class="card max-w-lg">
      <form @submit.prevent="save" class="space-y-5">
        <div>
          <label class="block text-sm font-medium mb-2" style="color: var(--color-text-secondary)">
            Nombre <span class="text-red-500">*</span>
          </label>
          <input
            v-model="form.name"
            type="text"
            class="input-field"
            placeholder="Ej: Electrónica, Ropa, Alimentos..."
            maxlength="100"
            autofocus
          />
          <p class="text-xs mt-1" style="color: var(--color-text-secondary)">
            Mínimo 2 caracteres, máximo 100
          </p>
        </div>

        <div>
          <label class="block text-sm font-medium mb-2" style="color: var(--color-text-secondary)">
            Descripción
          </label>
          <textarea
            v-model="form.description"
            class="input-field min-h-[120px] resize-y"
            placeholder="Descripción opcional de la categoría..."
            maxlength="500"
          ></textarea>
          <p class="text-xs mt-1 text-right" style="color: var(--color-text-secondary)">
            {{ form.description.length }}/500
          </p>
        </div>

        <div class="flex gap-3 pt-2">
          <button 
            type="submit" 
            class="btn-primary" 
            :disabled="saving || !form.name.trim()"
          >
            <span v-if="saving" class="flex items-center gap-2">
              <span class="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></span>
              Guardando...
            </span>
            <span v-else>{{ isEdit ? 'Actualizar' : 'Crear Categoría' }}</span>
          </button>
          <button 
            type="button" 
            @click="goBack" 
            class="btn-secondary"
            :disabled="saving"
          >
            Cancelar
          </button>
        </div>
      </form>
    </div>
  </div>
</template>