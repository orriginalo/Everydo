<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { models } from '../../wailsjs/go/models'
import { useTasksStore } from '../stores/useTasksStore'
import { storeToRefs } from 'pinia'
import { UncompleteTask } from '../../wailsjs/go/main/App'

const props = defineProps({
  task: {
    type: models.Task,
    required: true,
  },
})

const now = ref(Date.now())
let timerId = null

const tasksStore = useTasksStore()
const { toggleTask } = storeToRefs(tasksStore)
const { completeTask, uncompleteTask } = tasksStore

onMounted(() => {
  timerId = setInterval(() => {
    now.value = Date.now()
  }, 1000)
})

onUnmounted(() => {
  clearInterval(timerId)
})

function formatRemaining(ms) {
  if (ms <= 0) return 'обновляется'

  const totalSeconds = Math.floor(ms / 1000)
  const totalMinutes = Math.floor(totalSeconds / 60)
  const totalHours = Math.floor(totalMinutes / 60)
  const days = Math.floor(totalHours / 24)

  if (days > 0) {
    const hours = totalHours % 24
    return `${days}д ${hours}ч`
  }

  if (totalHours > 0) {
    const minutes = totalMinutes % 60
    return `${totalHours}ч ${minutes}м`
  }

  const minutes = totalMinutes
  const seconds = totalSeconds % 60
  return `${minutes}м ${seconds}с`
}

const ms = computed(() => {
  const resetAt = new Date(props.task.next_reset_at).getTime()
  return resetAt - now.value
})

const remainingTime = computed(() => {
  if (!props.task.next_reset_at) return null

  if (ms.value <= 0 && props.task.is_completed) {
    uncompleteTask(props.task)
  }
  return formatRemaining(ms.value)
})
</script>

<template>
  <div
    class="rounded-xl border px-4 py-3 transition-colors"
    :class="
      task.is_completed
        ? 'bg-neutral-900 border-emerald-700'
        : 'bg-neutral-800 border-neutral-700 hover:bg-neutral-700'
    "
  >
    <!-- Верх -->
    <div class="flex items-center justify-between gap-3">
      <h3
        class="text-sm font-medium truncate"
        :class="task.is_completed ? 'text-emerald-300' : 'text-neutral-100'"
      >
        {{ task.name }}
      </h3>

      <!-- Кнопка toggle -->
      <button
        @click="task.is_completed ? uncompleteTask(task) : completeTask(task)"
        class="flex h-7 w-7 items-center justify-center rounded-lg border border-neutral-600 text-neutral-300 transition hover:cursor-pointer"
        :class="
          task.is_completed
            ? 'hover:border-neutral-500 hover:text-neutral-200 hover:bg-neutral-700'
            : 'hover:border-emerald-600 hover:bg-emerald-700 hover:text-emerald-400'
        "
        :title="task.is_completed ? 'Сделать снова невыполненной' : 'Выполнить'"
      >
        {{ task.is_completed ? '⟳' : '✓' }}
      </button>
    </div>

    <!-- Низ -->
    <div class="mt-1 text-xs text-neutral-400">
      <template v-if="task.is_completed">
        Сброс через
        <span class="ml-1 text-neutral-200">
          {{ remainingTime }}
        </span>
      </template>

      <template v-else-if="!task.is_completed && ms <= 8 * 1000 * 60 * 60">
        Не выполнено
        <span class="ml-1 text-red-700">
          {{ remainingTime }}
        </span>
      </template>
      <span v-else> Не выполнено </span>
    </div>
  </div>
</template>
