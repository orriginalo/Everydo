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
import ConfirmModal from './components/ConfirmModal.vue'
import EditCategoryModal from './components/EditCategoryModal.vue'
import EditTaskModal from './components/EditTaskModal.vue'

const categoriesStore = useCategoriesStore()
const tasksStore = useTasksStore()

const { categories, activeCategory, isCreateCategoryModalOpen, toDeleteCategory } =
  storeToRefs(categoriesStore)
const { loadCategories, setActiveCategory, deleteCategoryForce } = categoriesStore

const { deleteTaskForce } = tasksStore
onMounted(async () => {
  await loadCategories()
})
</script>

<template>
  <div class="flex h-screen bg-neutral-900 text-neutral-100 overflow-hidden scrollbar">
    <Sidebar />

    <main class="flex-1 p-6 bg-neutral-900">
      <div
        v-if="categories.length > 0"
        class="flex flex-col h-full rounded-lg border border-neutral-800 bg-neutral-850 p-6"
      >
        <div class="flex flex-row justify-between mb-2">
          <h1 class="text-2xl font-semibold mb-2 font-unbounded">
            {{ activeCategory?.name }}
          </h1>
          <button
            @click="tasksStore.toggleIsCreateModalOpen"
            class="flex flex-row items-center gap-2 font-unbounded bg-neutral-800 px-2 rounded-full hover:bg-neutral-700 hover:cursor-pointer"
          >
            <Icon icon="ic:baseline-plus" width="20" />
          </button>
        </div>

        <!-- <p class="text-neutral-400">Контент для выбранной категории</p> -->
        <TasksPanelComponent
          v-if="activeCategory"
          :category_id="activeCategory.id"
          class="flex-1 overflow-y-auto"
        />
      </div>
      <div
        v-else
        class="h-full rounded-lg border border-neutral-800 bg-neutral-850 p-6 text-center"
      ></div>
    </main>
  </div>
  <EditTaskModal v-model:show="tasksStore.isEditModalOpen" />
  <EditCategoryModal v-model:show="categoriesStore.isEditCategoryModalOpen" />
  <CreateCatModal v-model:show="isCreateCategoryModalOpen" />
  <CreateTaskModal v-model:show="tasksStore.isCreateModalOpen" />
  <ConfirmModal
    v-model:show="categoriesStore.isDeleteModalOpen"
    title="Удалить категорию?"
    description="Вместе с категорией удалятся все задачи из неё. Это действие нельзя будет отменить."
    confirmText="Удалить"
    cancelText="Отмена"
    @confirm="deleteCategoryForce(categoriesStore.toDeleteCategory)"
    @cancel="() => {}"
  />
  <ConfirmModal
    v-model:show="tasksStore.isDeleteModalOpen"
    title="Удалить задание?"
    description="Вы уверены что хотите удалить это задание? Это действие нельзя будет отменить."
    confirmText="Удалить"
    cancelText="Отмена"
    @confirm="deleteTaskForce(tasksStore.toDeleteTask)"
    @cancel="() => {}"
  />
</template>
