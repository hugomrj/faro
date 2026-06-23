<script setup lang="ts">
interface Props {
  modelValue: string | number
  label?: string
  type?: string
  placeholder?: string
  error?: string
  disabled?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  type: 'text',
  placeholder: '',
  error: '',
  disabled: false
})

const emit = defineEmits<{
  'update:modelValue': [value: string | number]
}>()

function onInput(event: Event) {
  const target = event.target as HTMLInputElement
  emit('update:modelValue', target.value)
}
</script>

<template>
  <div class="w-full">
    <label v-if="label" class="block text-sm font-medium mb-1.5" style="color: var(--color-text-secondary)">
      {{ label }}
    </label>
    <div class="relative">
      <input
        :type="type"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        @input="onInput"
        class="input-field w-full"
        :class="{ 'opacity-50 cursor-not-allowed': disabled, 'border-red-500/50 focus:border-red-500 focus:ring-red-500/20': error }"
      />
    </div>
    <p v-if="error" class="mt-1.5 text-xs text-red-400">
      {{ error }}
    </p>
  </div>
</template>