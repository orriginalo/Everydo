<script setup>
import { computed, onMounted } from 'vue'
import { useTasksStore } from '@/stores/useTasksStore'

const props = defineProps({
  category_id: {
    type: Number,
    required: true,
  },
})

const typeTexts = {
  daily: 'ежедневно',
  weekly: 'еженедельно',
  custom: 'интервал',
}

const tasksStore = useTasksStore()

onMounted(async () => {
  await tasksStore.loadTasks(props.category_id)
})

// Только непустые категории
const tasksByTypeFiltered = computed(() => {
  const all = tasksStore.tasksByCategory[props.category_id] ?? {
    daily: [],
    weekly: [],
    custom: [],
  }

  // Оставляем только те, где есть таски
  const filtered = {}
  for (const type in all) {
    if (all[type].length > 0) {
      filtered[type] = all[type]
    }
  }
  return filtered
})
</script>

<template>
  <div class="grid grid-cols-2 gap-2 font-unbounded">
    <div
      v-for="(tasks, type) in tasksByTypeFiltered"
      :key="type"
      class="bg-neutral-800 p-2 rounded-xl"
    >
      <div class="font-bold mb-2">{{ typeTexts[type] }}</div>

      <div v-for="task in tasks" :key="task.id" class="text-sm opacity-80">
        {{ task.name }}
      </div>
    </div>
  </div>
</template>
