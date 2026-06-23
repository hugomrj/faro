<script setup lang="ts">
import { watch, onMounted, onUnmounted } from 'vue'

interface Props {
  isOpen: boolean
  title: string
  size?: 'sm' | 'md' | 'lg'
}

const props = withDefaults(defineProps<Props>(), {
  size: 'md'
})

const emit = defineEmits<{
  'close': []
}>()

function handleEscape(e: KeyboardEvent) {
  if (e.key === 'Escape' && props.isOpen) {
    emit('close')
  }
}

onMounted(() => document.addEventListener('keydown', handleEscape))
onUnmounted(() => document.removeEventListener('keydown', handleEscape))

// Prevenir scroll del body cuando el modal está abierto
watch(() => props.isOpen, (val) => {
  document.body.style.overflow = val ? 'hidden' : 'auto'
})
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center p-4">
        <!-- Backdrop -->
        <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="$emit('close')"></div>
        
        <!-- Contenido -->
        <div 
          class="relative w-full card p-0 flex flex-col max-h-[90vh] transform transition-all"
          :class="{
            'max-w-sm': size === 'sm',
            'max-w-lg': size === 'md',
            'max-w-2xl': size === 'lg'
          }"
          @click.stop
        >
          <!-- Header -->
          <div class="flex items-center justify-between p-5 shrink-0" style="border-bottom: 1px solid var(--color-border-strong)">
            <h2 class="text-lg font-semibold" style="color: var(--color-text-primary)">{{ title }}</h2>
            <button @click="$emit('close')" class="p-1.5 rounded-lg hover:bg-[var(--color-sidebar-active)] transition-colors" style="color: var(--color-text-secondary)">
              <svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
          
          <!-- Body -->
          <div class="p-5 overflow-y-auto flex-1">
            <slot />
          </div>
          
          <!-- Footer (Opcional) -->
          <div v-if="$slots.footer" class="p-5 shrink-0 flex justify-end gap-3" style="border-top: 1px solid var(--color-border-strong)">
            <slot name="footer" />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: all 0.25s ease-out;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
.modal-enter-from .card,
.modal-leave-to .card {
  transform: scale(0.95) translateY(10px);
}
</style>