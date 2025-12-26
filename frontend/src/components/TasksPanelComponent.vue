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
  <div class="flex flex-wrap flex-col gap-2 font-unbounded rounded-xl">
    <div
      v-if="tasksByType['daily'].length > 0"
      id="daily"
      class="bg-neutral-800 p-2 rounded-xl gap-1.5 flex flex-col"
    >
      <span class="font-bold text-xl pl-1">ежедневные</span>
      <TaskCard v-for="task in tasksByType['daily']" :key="task.id" :task="task" />
    </div>

    <div
      v-if="tasksByType['weekly'].length > 0"
      id="weekly"
      class="bg-neutral-800 p-2 rounded-xl gap-1.5 flex flex-col"
    >
      <span class="font-bold text-xl pl-1 max-h-min">еженедельные</span>
      <TaskCard v-for="task in tasksByType['weekly']" :key="task.id" :task="task" />
    </div>

    <div
      v-if="tasksByType['custom'].length > 0"
      id="custom"
      class="bg-neutral-800 p-2 rounded-xl gap-1.5 flex flex-col"
    >
      <span class="font-bold text-xl pl-1">интервалы</span>
      <TaskCard v-for="task in tasksByType['custom']" :key="task.id" :task="task" />
    </div>
  </div>
</template>
