<script setup>
import { useCategoriesStore } from '@/stores/useCategoriesStore'
import { models } from '../../wailsjs/go/models'

const props = defineProps({
  cat: {
    type: models.Category,
    required: true,
  },
})

const categoriesStore = useCategoriesStore()
const tasksStore = useTasksStore()

const { categories, activeCategory, isCreateCategoryModalOpen } = storeToRefs(categoriesStore)
const { loadCategories, setActiveCategory } = categoriesStore
</script>

<template>
  <button
    @click="setActiveCategory(props.cat)"
    class="group flex items-center gap-3 rounded-md px-3 py-2 text-left transition-all"
    :class="
      activeCategory.id === props.cat.id
        ? 'bg-neutral-800 text-white shadow-inner'
        : 'text-neutral-400 hover:bg-neutral-800 hover:text-white'
    "
  >
    <!-- Иконка-заглушка -->
    <div class="h-8 w-8 rounded bg-neutral-700 flex items-center justify-center text-sm font-bold">
      {{ props.cat.name[0] }}
    </div>

    <span class="truncate font-unbounded">
      {{ props.cat.name }}
    </span>
  </button>
</template>
