<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { Icon } from '@iconify/vue'
import CreateCatModal from './components/CreateCatModal.vue'
import { useCategoriesStore } from './stores/categories'
import CategorySideCard from './components/CategorySideCard.vue'
import Sidebar from './components/Sidebar.vue'
import TasksPanelComponent from './components/TasksPanelComponent.vue'
import { useTasksStore } from './stores/tasks'

const activeId = ref<number | null>(1)
const store = useCategoriesStore()

const tasksStore = useTasksStore()

onMounted(() => {
  store.getCategories()
})
</script>

<template>
  <div class="flex h-screen bg-neutral-900 text-neutral-100">
    <Sidebar />

    <main class="flex-1 p-6 bg-neutral-900">
      <div
        v-if="store.categories.length > 0"
        class="h-full rounded-lg border border-neutral-800 bg-neutral-850 p-6"
      >
        <h1 class="text-2xl font-semibold mb-2 font-unbounded">
          {{ store.activeCat?.name }}
        </h1>

        <!-- <p class="text-neutral-400">Контент для выбранной категории</p> -->
        <TasksPanelComponent :tasks="tasksStore.tasksByCategory[store.activeCat.id] ?? []" />
      </div>
      <div
        v-else
        class="h-full rounded-lg border border-neutral-800 bg-neutral-850 p-6 text-center"
      ></div>
    </main>
  </div>
  <CreateCatModal v-if="store.isCreateModalOpen" />
</template>
