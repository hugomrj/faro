// frontend/src/composables/useProducts.ts
import { ref } from 'vue'
import { GetAll } from '../../wailsjs/go/categories/Controller' // <-- Nombre correcto
import { GetAll as GetProducts, GetByID, Create, Update, Delete } from '../../wailsjs/go/products/Controller' // <-- Nombres que coinciden con el Controller.go que acabamos de hacer
import type { Product, ProductCreateRequest, Category } from '../types'

export function useProducts() {
  const products = ref<Product[]>([])
  const categories = ref<Category[]>([])
  const currentProduct = ref<Product | null>(null)
  const loading = ref(false)
  const error = ref('')

  async function fetchCategories() {
    try {
      const result = await GetAll()
      categories.value = result || []
    } catch (err: any) {
      console.error(err)
    }
  }

  async function fetchAll() {
    loading.value = true
    error.value = ''
    try {
      const result = await GetProducts() // Usamos alias porque ya importamos un GetAll de categories
      products.value = result || []
    } catch (err: any) {
      error.value = err.message || 'Error cargando productos'
      products.value = []
    } finally {
      loading.value = false
    }
  }

  async function fetchById(id: number) {
    loading.value = true
    error.value = ''
    try {
      const result = await GetByID(id)
      currentProduct.value = result
      return result
    } catch (err: any) {
      error.value = err.message || 'Error cargando producto'
      currentProduct.value = null
      return null
    } finally {
      loading.value = false
    }
  }

  async function save(data: ProductCreateRequest, id?: number): Promise<boolean> {
    loading.value = true
    error.value = ''
    try {
      if (id) {
        await Update(id, data)
      } else {
        await Create(data)
      }
      return true
    } catch (err: any) {
      error.value = err.message || 'Error guardando producto'
      return false
    } finally {
      loading.value = false
    }
  }


  async function remove(id: number): Promise<boolean> {
    loading.value = true
    error.value = ''
    try {
      // Asegúrate de tener la función Delete en tu Controller.go de products
      await Delete(id) 
      // Actualizamos la lista local quitando el eliminado para que se sienta instantáneo
      products.value = products.value.filter(p => p.id !== id)
      return true
    } catch (err: any) {
      error.value = err.message || 'Error eliminando producto'
      return false
    } finally {
      loading.value = false
    }
  }




  return {
    products,
    categories,
    currentProduct,
    loading,
    error,
    fetchCategories,
    fetchAll,
    fetchById,
    save,
    remove
  }

  
}