<!-- components/ConfirmModal.vue -->
<script setup>
import { ref, defineProps, defineEmits } from 'vue'

const props = defineProps({
  title: { type: String, default: 'Подтвердите действие' },
  description: { type: String, default: '' },
  confirmText: { type: String, default: 'Да' },
  cancelText: { type: String, default: 'Отмена' },
  show: { type: Boolean, required: true },
})

const emits = defineEmits(['update:show', 'confirm', 'cancel'])

function close() {
  emits('update:show', false)
}

function onConfirm() {
  emits('confirm')
  close()
}

function onCancel() {
  emits('cancel')
  close()
}
</script>

<template>
  <transition name="fade-scale">
    <div
      v-if="show"
      class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm"
    >
      <div class="bg-neutral-800 rounded-lg border border-neutral-700 p-6 w-96">
        <h2 class="text-xl font-semibold font-unbounded text-white mb-2">{{ title }}</h2>
        <p class="text-neutral-400 mb-4">{{ description }}</p>

        <div class="flex justify-end gap-2">
          <button
            @click="onCancel"
            class="px-3 py-1.5 rounded-lg text-neutral-400 hover:text-neutral-200 hover:bg-neutral-700 transition"
          >
            {{ cancelText }}
          </button>

          <button
            @click="onConfirm"
            class="bg-red-700 px-3 py-1.5 rounded-lg text-white hover:bg-red-500 transition"
          >
            {{ confirmText }}
          </button>
        </div>
      </div>
    </div>
  </transition>
</template>

<style scoped>
.fade-scale-enter-active,
.fade-scale-leave-active {
  transition: all 0.15s ease-out;
}
.fade-scale-enter-from,
.fade-scale-leave-to {
  transform: scale(0.95);
  opacity: 0;
}
.fade-scale-enter-to,
.fade-scale-leave-from {
  transform: scale(1);
  opacity: 1;
}
</style>
