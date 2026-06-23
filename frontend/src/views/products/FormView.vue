<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useProducts } from '../../composables/useProducts';
import type { Product, ProductCreateRequest } from '../../types';

// Átomos del Design System
import AppInput from '../../components/ui/AppInput.vue';
import AppButton from '../../components/ui/AppButton.vue';

const router = useRouter();
const route = useRoute();
// --- CORRECCIÓN AQUÍ: Nombres alineados con useProducts.ts ---
const { categories, fetchCategories, save, fetchById } = useProducts();

// Estado reactivo
const isLoading = ref(false);
const isEditing = computed(() => !!route.params.id);
const showSuccessAlert = ref(false);

// Tipado estricto usando la interfaz manual
const formData = ref<ProductCreateRequest>({
  name: '',
  category_id: 0,
  price: 0,
  stock: 0,
  status: 'Activo'
});





// Cargar datos iniciales
onMounted(async () => {
  isLoading.value = true;
  await fetchCategories();
  
  if (isEditing.value) {
    try {
      // Quitamos el ": Product" para que TS infiera el tipo real (Product | null)
      const product = await fetchById(Number(route.params.id));
      
      // Verificación de seguridad: Si el producto no existe, volvemos atrás
      if (!product) {
        console.error("Producto no encontrado");
        router.push({ name: 'products-list' });
        return;
      }

      // Si sí existe, llenamos el formulario
      formData.value = {
        name: product.name,
        category_id: product.category_id,
        price: product.price,
        stock: product.stock,
        status: product.status
      };
    } catch (error) {
      console.error("Error cargando producto", error);
      router.push({ name: 'products-list' }); // Si hay error de red, también volvemos
    } finally {
      isLoading.value = false;
    }
  } else {
    isLoading.value = false;
  }
});





// Manejo del envío estrictamente tipado
const handleSubmit = async () => {
  if (formData.value.category_id === 0 || formData.value.name.trim() === '') {
    alert("El nombre y la categoría son obligatorios.");
    return;
  }

  isLoading.value = true;
  try {
    const editId = isEditing.value ? Number(route.params.id) : undefined;
    await save(formData.value, editId);
    
    showSuccessAlert.value = true;
    setTimeout(() => {
      router.push({ name: 'products-list' }); // Asegúrate de que el router tenga este nombre
    }, 1200);
    
  } catch (error) {
    console.error("Error guardando producto", error);
    alert("Hubo un error al guardar el producto.");
  } finally {
    isLoading.value = false;
  }
};

const handleCancel = () => {
  router.back();
};
</script>

<template>
  <div class="p-6 flex flex-col gap-6 h-full overflow-y-auto">
    
    <!-- Cabecera -->
    <header class="flex items-center justify-between">
      <div>
        <h1 class="text-2xl font-bold text-[var(--text-primary)]">
          {{ isEditing ? 'Editar Producto' : 'Nuevo Producto' }}
        </h1>
        <p class="text-sm mt-1 text-[var(--text-secondary)]">
          Complete los detalles del producto para añadirlo al inventario.
        </p>
      </div>
      <AppButton 
        variant="secondary" 
        @click="handleCancel"
        :disabled="isLoading"
      >
        <span class="flex items-center gap-2">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
          Volver
        </span>
      </AppButton>
    </header>

    <!-- Alerta de Éxito (Glassmorphism) -->
    <Transition name="fade">
      <div 
        v-if="showSuccessAlert" 
        class="p-4 rounded-[16px] border text-sm font-medium"
        style="background: rgba(16, 185, 129, 0.1); border-color: var(--color-primary-dark); color: var(--color-primary-dark); backdrop-filter: blur(8px);"
      >
        Producto guardado correctamente. Redirigiendo...
      </div>
    </Transition>

    <!-- Contenedor Principal (Card) -->
    <div class="card flex-1 p-8 flex flex-col gap-8">
      
      <!-- Sección 1: Información General -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        
        <div class="flex flex-col gap-2 md:col-span-2">
          <label class="text-sm font-medium text-[var(--text-secondary)]">Nombre del Producto *</label>
          <AppInput 
            v-model="formData.name" 
            placeholder="Ej: Teclado Mecánico RGB"
            :disabled="isLoading"
          />
        </div>

        <div class="flex flex-col gap-2">
          <label class="text-sm font-medium text-[var(--text-secondary)]">Categoría *</label>
          <select 
            v-model.number="formData.category_id"
            class="input-field appearance-none cursor-pointer"
            :disabled="isLoading || categories.length === 0"
          >
            <option :value="0" disabled>Seleccionar categoría...</option>
            <option 
              v-for="cat in categories" 
              :key="cat.id" 
              :value="cat.id"
            >
              {{ cat.name }}
            </option>
          </select>
        </div>

        <div class="flex flex-col gap-2">
          <label class="text-sm font-medium text-[var(--text-secondary)]">Estado</label>
          <div class="flex gap-3 h-[46px] items-center">
            <button 
              type="button"
              @click="formData.status = 'Activo'"
              class="px-5 h-full rounded-[16px] text-sm font-medium transition-all duration-200 border"
              :style="formData.status === 'Activo' 
                ? 'background: rgba(16, 185, 129, 0.15); border-color: var(--color-primary-dark); color: var(--color-primary-dark); box-shadow: 0 0 15px rgba(16, 185, 129, 0.1);' 
                : 'background: var(--bg-input); border-color: var(--border-color); color: var(--text-secondary);'"
              :disabled="isLoading"
            >
              Activo
            </button>
            <button 
              type="button"
              @click="formData.status = 'Inactivo'"
              class="px-5 h-full rounded-[16px] text-sm font-medium transition-all duration-200 border"
              :style="formData.status === 'Inactivo' 
                ? 'background: rgba(239, 68, 68, 0.15); border-color: #ef4444; color: #ef4444; box-shadow: 0 0 15px rgba(239, 68, 68, 0.1);' 
                : 'background: var(--bg-input); border-color: var(--border-color); color: var(--text-secondary);'"
              :disabled="isLoading"
            >
              Inactivo
            </button>
          </div>
        </div>

      </div>

      <!-- Separador -->
      <div class="w-full h-px" style="background: var(--border-color);"></div>

      <!-- Sección 2: Inventario y Precios -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        
        <div class="flex flex-col gap-2">
          <label class="text-sm font-medium text-[var(--text-secondary)]">Precio Unitario *</label>
          <div class="relative">
            <span class="absolute left-4 top-1/2 -translate-y-1/2 text-sm text-[var(--text-secondary)]">$</span>
            <AppInput 
              v-model.number="formData.price" 
              type="number" 
              step="0.01"
              placeholder="0.00"
              class="pl-8"
              :disabled="isLoading"
            />
          </div>
        </div>

        <div class="flex flex-col gap-2">
          <label class="text-sm font-medium text-[var(--text-secondary)]">Stock Disponible *</label>
          <AppInput 
            v-model.number="formData.stock" 
            type="number" 
            placeholder="0"
            :disabled="isLoading"
          />
        </div>

      </div>

      <!-- Pie del Formulario -->
      <div class="mt-auto pt-6 border-t flex justify-end gap-4" style="border-color: var(--border-color);">
        <AppButton 
          variant="secondary" 
          @click="handleCancel"
          :disabled="isLoading"
        >
          Cancelar
        </AppButton>
        <AppButton 
          variant="primary" 
          @click="handleSubmit"
          :disabled="isLoading"
        >
          <span class="flex items-center gap-2">
            <svg v-if="isLoading" class="animate-spin h-4 w-4 text-current" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ isEditing ? 'Actualizar Cambios' : 'Guardar Producto' }}
          </span>
        </AppButton>
      </div>

    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}

select.input-field {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke='%2364748b' stroke-width='2'%3E%3Cpath stroke-linecap='round' stroke-linejoin='round' d='M19 9l-7 7-7-7'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 0.75rem center;
  background-size: 1.25rem;
  padding-right: 2.5rem;
}
</style>