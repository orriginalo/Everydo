<script setup>
import { computed, onMounted } from 'vue'
import { useTasksStore } from '../stores/useTasksStore'
import { getTasksByTypes } from '../utils'
import { storeToRefs } from 'pinia'
import { watch } from 'vue'
import TaskCard from './TaskCard.vue'
import draggable from 'vuedraggable'
import { ref } from 'vue'

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

const drag = ref(false)

watch(
  () => props.category_id,
  (id) => {
    loadTasks(id)
  },
  { immediate: true },
)
const dailyTasks = computed({
  get: () => tasksByType.value.daily,
  set: (val) => tasksStore.reorderTasks(props.category_id, 'daily', val),
})
const weeklyTasks = computed({
  get: () => tasksByType.value.weekly,
  set: (val) => tasksStore.reorderTasks(props.category_id, 'weekly', val),
})
const customTasks = computed({
  get: () => tasksByType.value.custom,
  set: (val) => tasksStore.reorderTasks(props.category_id, 'custom', val),
})
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
      <!-- TODO: добавить mb-1.5 -->
      <span class="font-bold text-xl pl-1 shrink-0 mb-1.5">ежедневные</span>

      <!-- <div class="flex-1 overflow-y-auto flex flex-col gap-1.5 min-h-0">
        <TaskCard v-for="task in tasksByType['daily']" :key="task.id" :task="task" />
      </div> -->
      <draggable
        v-model="dailyTasks"
        item-key="id"
        animation="300"
        @start="drag = true"
        @end="drag = false"
        handle=".drag-handle"
        ghost-class="opacity-40"
        force-fallback="true"
        class="flex-1 overflow-y-auto flex flex-col gap-1.5 min-h-0"
      >
        <template #item="{ element }">
          <TaskCard :task="element" />
        </template>
      </draggable>
    </div>

    <!-- WEEKLY -->
    <div
      v-if="tasksByType['weekly'].length > 0"
      class="bg-neutral-800 p-2 rounded-xl flex flex-col min-h-0"
    >
      <span class="font-bold text-xl pl-1 shrink-0 mb-2">еженедельные</span>

      <draggable
        v-model="weeklyTasks"
        item-key="id"
        animation="300"
        @start="drag = true"
        @end="drag = false"
        handle=".drag-handle"
        ghost-class="opacity-40"
        force-fallback="true"
        class="flex-1 overflow-y-auto flex flex-col gap-1.5 min-h-0"
      >
        <template #item="{ element }">
          <TaskCard :task="element" />
        </template>
      </draggable>
    </div>

    <!-- CUSTOM -->
    <div
      v-if="tasksByType['custom'].length > 0"
      class="bg-neutral-800 p-2 rounded-xl flex flex-col min-h-0"
    >
      <span class="font-bold text-xl pl-1 shrink-0">интервалы</span>

      <draggable
        v-model="customTasks"
        item-key="id"
        animation="300"
        @start="drag = true"
        @end="drag = false"
        handle=".drag-handle"
        ghost-class="opacity-40"
        force-fallback="true"
        class="flex-1 overflow-y-auto flex flex-col gap-1.5 min-h-0"
      >
        <template #item="{ element }">
          <TaskCard :task="element" />
        </template>
      </draggable>
    </div>
  </div>
</template>
