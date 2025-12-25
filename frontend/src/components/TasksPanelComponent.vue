<script setup lang="ts">
import { computed, ref } from 'vue'
import type { models } from '../../wailsjs/go/models'
import { useTasksStore } from '@/stores/tasks'

const props = defineProps<{
  tasks: models.Task[]
}>()

const tasksStore = useTasksStore()

const tasksByType = computed(() => {
  const acc: Record<string, models.Task[]> = {
    daily: [],
    weekly: [],
    custom: [],
  }

  for (const task of props.tasks) {
    acc[task.reload_type]?.push(task)
  }

  return acc
})
</script>

<template>
  <div class="grid grid-cols-3 gap-2 font-unbounded">
    <div v-for="(tasks, type) in tasksByType" :key="type" class="bg-neutral-800 p-2 rounded-t-xl">
      <div class="font-bold mb-2">{{ type }}</div>

      <div v-for="task in tasks" :key="task.id" class="text-sm opacity-80">
        {{ task.name }}
      </div>
    </div>
  </div>
</template>
