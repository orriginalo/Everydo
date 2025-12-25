<script setup lang="ts">
import { ref } from 'vue'
import { Icon } from '@iconify/vue'
import { useCategoriesStore } from '@/stores/categories'

const showTooltip = ref(false)
const store = useCategoriesStore()

const name = ref('')
const exeName = ref('')

function createCategory() {
  if (name.value === '') {
    return
  }
  store.createCategory(name.value, exeName.value)
  store.toggleIsCreateModalOpen()
}
</script>

<template>
  <div class="absolute inset-0 bg-black/60 backdrop-blur-sm"></div>

  <form
    @submit.prevent="createCategory"
    class="fixed inset-0 z-20 flex items-center justify-center"
  >
    <div class="bg-neutral-800 rounded-lg border border-neutral-800 p-6 w-105">
      <h1 class="text-2xl font-semibold mb-2 text-white font-unbounded">Создать категорию</h1>
      <p class="text-neutral-400 font-semibold">Введите данные для новой категории</p>

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

          <div class="relative" @mouseenter="showTooltip = true" @mouseleave="showTooltip = false">
            <Icon
              icon="mingcute:question-fill"
              class="text-neutral-400 hover:text-neutral-200"
              width="28"
            />

            <div
              v-if="showTooltip"
              class="absolute left-1/2 top-full z-30 mt-2 w-64 -translate-x-1/2 rounded-md bg-neutral-900 border border-neutral-700 p-2 text-sm text-neutral-300 shadow-xl"
            >
              Укажите имя процесса (.exe), если хотите получать уведомления о имеющихся задачах при
              запуске игры.

              <div
                class="absolute -top-1 left-1/2 h-2 w-2 -translate-x-1/2 rotate-45 bg-neutral-900 border-l border-t border-neutral-700"
              />
            </div>
          </div>
        </div>
      </div>

      <div class="mt-6 flex justify-end gap-2">
        <button
          @click="store.toggleIsCreateModalOpen()"
          class="px-3 py-1.5 rounded-lg text-neutral-400 hover:text-neutral-200 hover:bg-neutral-700 transition"
        >
          Отмена
        </button>

        <button
          :disabled="name === ''"
          @click="createCategory()"
          class="bg-neutral-700 px-3 py-1.5 rounded-lg text-neutral-200 transition-all"
          :class="{
            'opacity-50 cursor-not-allowed': name.trim() === '',
            'hover:bg-neutral-600 font-bold': name.trim() !== '',
          }"
        >
          Создать
        </button>
      </div>
    </div>
  </form>
</template>
