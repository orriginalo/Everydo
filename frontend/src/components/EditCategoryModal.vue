<script setup>
import { ref } from 'vue'
import { Icon } from '@iconify/vue'
import { useCategoriesStore } from '../stores/useCategoriesStore'
import { watch } from 'vue'

const props = defineProps({
  show: { type: Boolean, required: true },
})

const emits = defineEmits(['update:show'])

const showTooltip = ref(false)
const store = useCategoriesStore()

const name = ref('')
const exeName = ref('')

watch(
  () => store.toEditCategory,
  (val) => {
    if (val) {
      name.value = val.name ?? ''
      exeName.value = val.exe_name ?? ''
    }
  },
  { immediate: true },
)

function updateCategory() {
  if (name.value === '') {
    return
  }
  store.updateCategory(store.toEditCategory.id, name.value, exeName.value)
  store.toggleIsEditModalOpen()
}
</script>

<template>
  <transition name="fade-scale">
    <div v-if="show" class="fixed inset-0 z-20 flex items-center justify-center">
      <!-- Фон -->
      <div class="absolute inset-0 bg-black/60 backdrop-blur-sm"></div>

      <!-- Модалка -->
      <form
        @submit.prevent="updateCategory"
        class="relative bg-neutral-800 rounded-lg border border-neutral-800 p-6 w-105"
      >
        <h1 class="text-2xl font-semibold mb-2 text-white font-unbounded select-none">
          Изменение категории
        </h1>
        <p class="text-neutral-400 font-semibold select-none">Введите данные для категории</p>

        <div class="mt-4 gap-2 flex flex-col">
          <input
            v-model="name"
            type="text"
            placeholder="Название категории"
            class="w-full rounded-lg border border-neutral-700 bg-neutral-900 p-2 text-neutral-200 outline-none focus:border-neutral-500"
          />

          <div class="relative flex flex-row items-center gap-2">
            <input
              v-model="exeName"
              type="text"
              placeholder="Имя процесса (необязательно)"
              class="w-full rounded-lg border border-neutral-700 bg-neutral-900 p-2 text-neutral-200 outline-none focus:border-neutral-500"
            />

            <div
              class="relative select-none"
              @mouseenter="showTooltip = true"
              @mouseleave="showTooltip = false"
            >
              <Icon
                icon="mingcute:question-fill"
                class="text-neutral-400 transition-colors hover:text-neutral-200"
                width="28"
              />

              <transition name="fade-scale">
                <div
                  v-if="showTooltip"
                  class="absolute left-1/2 top-full z-30 mt-2 w-64 -translate-x-1/2 rounded-md bg-neutral-900 border border-neutral-700 p-2 text-sm text-neutral-300 shadow-xl"
                >
                  Укажите имя процесса (.exe), если хотите получать уведомления о имеющихся задачах.
                  <div
                    class="absolute -top-1 left-1/2 h-2 w-2 -translate-x-1/2 rotate-45 bg-neutral-900 border-l border-t border-neutral-700"
                  />
                </div>
              </transition>
            </div>
          </div>
        </div>

        <div class="mt-6 flex justify-end gap-2">
          <button
            type="button"
            @click="store.toggleIsEditModalOpen()"
            class="px-3 py-1.5 rounded-lg text-neutral-400 hover:text-neutral-200 hover:bg-neutral-700 transition"
          >
            Отмена
          </button>

          <button
            :disabled="name === ''"
            type="submit"
            class="bg-neutral-700 px-3 py-1.5 rounded-lg text-neutral-200 transition-all"
            :class="{
              'opacity-50 cursor-not-allowed': name.trim() === '',
              'hover:bg-neutral-600 font-bold': name.trim() !== '',
            }"
          >
            Изменить
          </button>
        </div>
      </form>
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
