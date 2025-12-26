<script setup>
import { onMounted, ref } from 'vue'
import { Icon } from '@iconify/vue'
import CreateCatModal from './components/CreateCatModal.vue'
import { useCategoriesStore } from './stores/useCategoriesStore'
import CategorySideCard from './components/CategorySideCard.vue'
import Sidebar from './components/Sidebar.vue'
import TasksPanelComponent from './components/TasksPanelComponent.vue'
import { useTasksStore } from './stores/useTasksStore'
import CreateTaskModal from './components/CreateTaskModal.vue'
import { storeToRefs } from 'pinia'

const categoriesStore = useCategoriesStore()
const tasksStore = useTasksStore()

const { categories, activeCategory, isCreateCategoryModalOpen } = storeToRefs(categoriesStore)
const { loadCategories, setActiveCategory } = categoriesStore

onMounted(() => {
  categoriesStore.getCategories()
})
</script>

<template>
  <div class="flex h-screen bg-neutral-900 text-neutral-100">
    <Sidebar />

    <main class="flex-1 p-6 bg-neutral-900">
      <div
        v-if="categories.length > 0"
        class="h-full rounded-lg border border-neutral-800 bg-neutral-850 p-6"
      >
        <div class="flex flex-row justify-between">
          <h1 class="text-2xl font-semibold mb-2 font-unbounded">
            {{ activeCategory?.name }}
          </h1>
          <button
            @click="tasksStore.toggleIsCreateModalOpen"
            class="bg-neutral-900 px-1 py-0 rounded-full hover:bg-neutral-800 hover:cursor-pointer"
          >
            <Icon icon="ic:baseline-plus" />
          </button>
        </div>

        <!-- <p class="text-neutral-400">Контент для выбранной категории</p> -->
        <TasksPanelComponent v-if="activeCategory" :category_id="activeCategory.id" />
      </div>
      <div
        v-else
        class="h-full rounded-lg border border-neutral-800 bg-neutral-850 p-6 text-center"
      ></div>
    </main>
  </div>
  <CreateCatModal v-if="isCreateCategoryModalOpen" />
  <CreateTaskModal v-if="tasksStore.isCreateModalOpen" />
</template>
