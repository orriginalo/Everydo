<script setup>
import { ref } from 'vue'
import { useCategoriesStore } from '../stores/useCategoriesStore'
import CategorySideCard from './CategorySideCard.vue'
import { Icon } from '@iconify/vue'
import { storeToRefs } from 'pinia'
import { OpenURL } from '../../wailsjs/go/main/App'

const categoriesStore = useCategoriesStore()
const { categories } = storeToRefs(categoriesStore)

// Состояние сайдбара: открыт или закрыт
const isOpen = ref(true)

const toggleSidebar = () => {
  isOpen.value = !isOpen.value
}

function openUrl(url) {
  OpenURL(url).then(() => {
    console.log('Opened URL')
  })
}
</script>

<template>
  <aside
    :class="[
      'bg-neutral-950 border-r border-neutral-800 transition-all duration-300 overflow-hidden flex flex-col',
      isOpen ? 'w-64' : 'w-12',
    ]"
  >
    <!-- Хедер -->
    <div
      class="flex justify-between items-center text-lg font-semibold tracking-wide"
      :class="isOpen ? 'p-4' : 'py-4 px-3 justify-center'"
    >
      <span v-if="isOpen" class="font-unbounded">Категории</span>
      <div class="flex items-center gap-2">
        <button
          @click="categoriesStore.toggleIsCreateModalOpen"
          v-if="isOpen"
          class="bg-neutral-900 px-1 py-1 rounded-full hover:bg-neutral-800"
        >
          <Icon icon="ic:baseline-plus" />
        </button>
        <button
          @click="toggleSidebar"
          class="bg-neutral-900 px-1 py-1 rounded-full hover:bg-neutral-800"
        >
          <Icon :icon="isOpen ? 'ic:baseline-chevron-left' : 'ic:baseline-chevron-right'" />
        </button>
      </div>
    </div>

    <!-- Список категорий -->
    <div v-if="isOpen" class="flex flex-col gap-1 px-2 flex-1 overflow-y-auto">
      <CategorySideCard v-for="cat in categories" :key="cat.id" :category="cat" />
      <div v-if="categories.length == 0" class="text-neutral-400 px-2 text-center">
        Вы ещё не добавили ни одной игры!
      </div>
    </div>

    <!-- Footer -->
    <div
      v-if="isOpen"
      class="text-xs text-neutral-500 px-4 py-3 border-t border-neutral-800 font-unbounded"
    >
      by
      <button
        class="underline hover:cursor-pointer hover:text-neutral-400 transition-all"
        @click="
          () => {
            openUrl(`https://github.com/orriginalo`)
          }
        "
      >
        orriginalo
      </button>
    </div>
  </aside>
</template>
