import { ref } from 'vue'
import { GetAll, GetByID, Create, Update, Delete } from '../../wailsjs/go/categories/Controller'
import type { Category, CategoryCreateRequest, CategoryUpdateRequest } from '../types'

export function useCategories() {
  const categories = ref<Category[]>([])
  const currentCategory = ref<Category | null>(null)
  const loading = ref(false)
  const error = ref('')

  async function fetchAll() {
    loading.value = true
    error.value = ''
    try {
      const result = await GetAll()
      categories.value = result || []
    } catch (err: any) {
      error.value = err.message || 'Error cargando categorías'
      categories.value = []
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id: number) {
    loading.value = true
    error.value = ''
    try {
      const result = await GetByID(id)
      currentCategory.value = result
      return result
    } catch (err: any) {
      error.value = err.message || 'Error cargando categoría'
      currentCategory.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  async function create(data: CategoryCreateRequest): Promise<number | null> {
    loading.value = true
    error.value = ''
    try {
      const id = await Create(data)
      return id
    } catch (err: any) {
      error.value = err.message || 'Error creando categoría'
      return null
    } finally {
      loading.value = false
    }
  }

  async function update(id: number, data: CategoryUpdateRequest): Promise<boolean> {
    loading.value = true
    error.value = ''
    try {
      await Update(id, data)
      return true
    } catch (err: any) {
      error.value = err.message || 'Error actualizando categoría'
      return false
    } finally {
      loading.value = false
    }
  }

  async function remove(id: number): Promise<boolean> {
    loading.value = true
    error.value = ''
    try {
      await Delete(id)
      return true
    } catch (err: any) {
      error.value = err.message || 'Error eliminando categoría'
      return false
    } finally {
      loading.value = false
    }
  }

  return {
    categories,
    currentCategory,
    loading,
    error,
    fetchAll,
    fetchById,
    create,
    update,
    remove
  }
}