<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { models } from '../../wailsjs/go/models'
import { useTasksStore } from '../stores/useTasksStore'
import { storeToRefs } from 'pinia'
import { UncompleteTask, UpdateNextReset } from '../../wailsjs/go/main/App'
import { Icon } from '@iconify/vue'
import { useCategoriesStore } from '../stores/useCategoriesStore'
import { watch } from 'vue'

const props = defineProps({
  task: {
    type: models.Task,
    required: true,
  },
})

onMounted(() => {
  timerId = setInterval(() => {
    now.value = Date.now()
  }, 1000)
})

onUnmounted(() => {
  clearInterval(timerId)
})

const now = ref(Date.now())
let timerId = null

const timeRemainingConds = {
  daily: () => ms.value <= 6 * 60 * 60 * 1000,
  weekly: () => ms.value <= 48 * 60 * 60 * 1000,
  custom: () => ms.value <= 6 * 60 * 60 * 1000,
}

const store = useCategoriesStore()
const tasksStore = useTasksStore()
const { toggleTask, openedMenuTaskId } = storeToRefs(tasksStore)
const { completeTask, uncompleteTask, deleteTask, toggleTaskMenu, closeTaskMenu } = tasksStore

const toggleMenu = (e) => {
  e.stopPropagation()
  toggleTaskMenu(props.task.id)
}

const editTask = () => {
  closeTaskMenu()
  tasksStore.setEditTask(props.task)
  tasksStore.toggleIsEditModalOpen()
}

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
  return formatRemaining(ms.value)
})

watch(ms, async (newMs, oldMs) => {
  // Сработает, когда время только что стало <= 0 (перешло с >0 на <=0)
  if (oldMs > 0 && newMs <= 0) {
    // Неважно, выполнена задача или нет — обновляем
    if (props.task.is_completed) {
      await uncompleteTask(props.task)
    }
    // Всегда обновляем next_reset_at, если время вышло
    await UpdateNextReset(props.task.id)
    // Перезагружаем задачи
    tasksStore.loadTasks(store.activeCategory.id)
  }
})
</script>

<template>
  <div
    class="group flex flex-row items-center relative rounded-xl border pr-4 pl-1 py-3 transition-colors select-none"
    :class="
      task.is_completed
        ? 'bg-neutral-900 border-emerald-700'
        : 'bg-neutral-800 border-neutral-700 hover:bg-neutral-700'
    "
  >
    <div
      class="flex drag-handle left-2 h-full w-3 mr-1 cursor-grab opacity-0 group-hover:opacity-100 transition"
    >
      <div class="mx-auto h-full w-1 rounded-full bg-neutral-500"></div>
    </div>
    <div class="w-full">
      <!-- ВЕРХ: название + кнопка выполнить -->
      <div class="flex items-center justify-between gap-3">
        <h3
          class="text-wrap text-sm font-medium truncate"
          :class="task.is_completed ? 'text-emerald-300' : 'text-neutral-100'"
        >
          {{ task.name }}
        </h3>

        <button
          @click="task.is_completed ? uncompleteTask(task) : completeTask(task)"
          class="flex h-7 w-7 shrink-0 items-center justify-center rounded-lg border border-neutral-600 text-neutral-300 transition"
          :class="
            task.is_completed
              ? 'hover:bg-neutral-700'
              : 'hover:bg-emerald-700 hover:text-emerald-400'
          "
          :title="task.is_completed ? 'Сделать снова невыполненной' : 'Выполнить'"
        >
          {{ task.is_completed ? '⟳' : '✓' }}
        </button>
      </div>

      <!-- НИЗ: статус + меню -->
      <div class="mt-1 flex items-center justify-between text-xs text-neutral-400">
        <!-- статус -->
        <div class="whitespace-nowrap select-none">
          <template v-if="task.is_completed">
            Сброс через
            <span class="ml-1 text-neutral-200">
              {{ remainingTime }}
            </span>
          </template>

          <template v-else-if="timeRemainingConds[props.task.reload_type]()">
            Не выполнено
            <span class="ml-1 text-red-700">
              {{ remainingTime }}
            </span>
          </template>

          <span v-else>Не выполнено</span>
        </div>

        <!-- меню -->
        <div class="relative">
          <button
            @click="toggleTaskMenu(props.task.id)"
            class="flex h-7 w-7 items-center justify-center rounded-md text-neutral-400 hover:bg-neutral-600 transition"
          >
            <Icon
              icon="ic:baseline-more-vert"
              width="20"
              :class="openedMenuTaskId === props.task.id ? 'rotate-90 transition' : 'transition'"
            />
          </button>

          <transition name="fade-scale">
            <div
              v-if="openedMenuTaskId === props.task.id"
              class="absolute right-0 top-full mt-1 w-36 bg-neutral-900 border border-neutral-700 rounded-xl shadow-xl z-50 overflow-hidden font-raleway"
            >
              <button
                @click="editTask"
                class="w-full px-4 py-2 text-left text-sm text-white hover:bg-neutral-800 transition-colors"
              >
                Изменить
              </button>
              <button
                @click="
                  () => {
                    closeTaskMenu()
                    deleteTask(task)
                  }
                "
                class="w-full px-4 py-2 text-left text-sm text-red-400 hover:bg-red-950 transition-colors"
              >
                Удалить
              </button>
            </div>
          </transition>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped>
.fade-scale-enter-active,
.fade-scale-leave-active {
  transition: all 0.15s ease-out;
}
.fade-scale-enter-from,
.fade-scale-leave-to {
  transform: scale(0.95);
  opacity: 0;
}
.fade-scale-enter-to,
.fade-scale-leave-from {
  transform: scale(1);
  opacity: 1;
}
</style>
