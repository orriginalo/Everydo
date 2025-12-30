<script setup>
import { computed, onMounted } from 'vue'
import { useTasksStore } from '../stores/useTasksStore'
import { getTasksByTypes } from '../utils'
import { storeToRefs } from 'pinia'
import { watch } from 'vue'
import TaskCard from './TaskCard.vue'

const props = defineProps({
  category_id: {
    type: Number,
    required: true,
  },
})

const tasksStore = useTasksStore()

const { tasksByCategory } = storeToRefs(tasksStore)
const { loadTasks } = tasksStore

const tasksByType = computed(() => {
  const tasks = tasksByCategory.value[props.category_id] ?? []
  return getTasksByTypes(tasks)
})

watch(
  () => props.category_id,
  (id) => {
    loadTasks(id)
  },
  { immediate: true },
)
</script>

<template>
  <div
    class="grid gap-2 h-full font-unbounded rounded-xl"
    style="grid-template-columns: repeat(auto-fit, minmax(260px, 1fr))"
  >
    <!-- DAILY -->
    <div
      v-if="tasksByType['daily'].length > 0"
      class="bg-neutral-800 p-2 rounded-xl flex flex-col min-h-0"
    >
      <span class="font-bold text-xl pl-1 shrink-0">ежедневные</span>

      <div class="flex-1 overflow-y-auto flex flex-col gap-1.5 min-h-0">
        <TaskCard v-for="task in tasksByType['daily']" :key="task.id" :task="task" />
      </div>
    </div>

    <!-- WEEKLY -->
    <div
      v-if="tasksByType['weekly'].length > 0"
      class="bg-neutral-800 p-2 rounded-xl flex flex-col min-h-0"
    >
      <span class="font-bold text-xl pl-1 shrink-0">еженедельные</span>

      <div class="flex-1 overflow-y-auto flex flex-col gap-1.5 min-h-0">
        <TaskCard v-for="task in tasksByType['weekly']" :key="task.id" :task="task" />
      </div>
    </div>

    <!-- CUSTOM -->
    <div
      v-if="tasksByType['custom'].length > 0"
      class="bg-neutral-800 p-2 rounded-xl flex flex-col min-h-0"
    >
      <span class="font-bold text-xl pl-1 shrink-0">интервалы</span>

      <div class="flex-1 overflow-y-auto flex flex-col gap-1.5 min-h-0">
        <TaskCard v-for="task in tasksByType['custom']" :key="task.id" :task="task" />
      </div>
    </div>
  </div>
</template>
