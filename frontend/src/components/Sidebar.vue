<script setup>
import { useCategoriesStore } from '@/stores/useCategoriesStore'
import CategorySideCard from './CategorySideCard.vue'
import { Icon } from '@iconify/vue'

const categoriesStore = useCategoriesStore()
const tasksStore = useTasksStore()

const { categories, activeCategory, isCreateCategoryModalOpen } = storeToRefs(categoriesStore)
const { loadCategories, setActiveCategory } = categoriesStore
</script>

<template>
  <aside class="w-64 bg-neutral-950 border-r border-neutral-800">
    <div class="p-4 text-lg font-semibold tracking-wide flex justify-between">
      <span class="font-unbounded">Категории</span>
      <button
        @click="store.toggleIsCreateModalOpen"
        class="bg-neutral-900 px-1 py-0 rounded-full hover:bg-neutral-800 hover:cursor-pointer"
      >
        <Icon icon="ic:baseline-plus" />
      </button>
    </div>

    <div class="flex flex-col gap-1 px-2">
      <CategorySideCard v-for="cat in categories" :key="cat.id" :cat="cat" />
      <div v-if="categories.length == 0" class="text-neutral-400 px-2 text-center">
        Вы ещё не добавили ни одной игры!
      </div>
    </div>
  </aside>
</template>
