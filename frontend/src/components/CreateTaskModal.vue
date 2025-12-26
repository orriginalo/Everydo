<script setup>
import { ref, computed } from 'vue'
import { Icon } from '@iconify/vue'
import { useTasksStore } from '@/stores/useTasksStore'
import { useCategoriesStore } from '@/stores/useCategoriesStore'

const tasksStore = useTasksStore()
const store = useCategoriesStore()

const typesMap = {
  daily: 'Ежедневно',
  weekly: 'Еженедельно',
  custom: 'Интервал',
}

const name = ref('')
const reloadType = ref('daily') // daily | weekly | custom
const resetTime = ref('03:00')
const resetWeekday = ref(0)
const reloadEvery = ref(1)

const isValid = computed(() => {
  if (name.value.trim() === '') return false
  if (reloadType.value === 'custom' && reloadEvery.value < 1) return false
  return true
})

function createTask() {
  if (!isValid.value) return

  tasksStore.createTask(store.activeCat.id, {
    name: name.value,
    reload_type: reloadType.value,
    reset_time: resetTime.value,
    reset_weekday: reloadType.value === 'weekly' ? resetWeekday.value : null,
    reload_every: reloadType.value === 'custom' ? reloadEvery.value : null,
  })

  tasksStore.loadTasks(store.activeCat.id)
  tasksStore.toggleIsCreateModalOpen()
}
</script>

<template>
  <div class="absolute inset-0 bg-black/60 backdrop-blur-sm"></div>

  <form @submit.prevent="createTask" class="fixed inset-0 z-20 flex items-center justify-center">
    <div class="bg-neutral-800 rounded-lg border border-neutral-700 p-6 w-105 transition-transform">
      <h1 class="text-2xl font-semibold mb-2 text-white font-unbounded">Создать задачу</h1>
      <p class="text-neutral-400 font-semibold">Настрой параметры задачи</p>

      <div class="mt-4 flex flex-col gap-3">
        <!-- NAME -->
        <input
          v-model="name"
          type="text"
          placeholder="Название задачи"
          class="w-full rounded-lg border border-neutral-700 bg-neutral-900 p-2 text-neutral-200 outline-none focus:border-neutral-500"
        />

        <!-- TYPE -->
        <div class="flex gap-2">
          <button
            v-for="type in ['daily', 'weekly', 'custom']"
            :key="type"
            type="button"
            @click="reloadType = type"
            class="flex-1 rounded-lg px-3 py-1.5 text-sm transition-all"
            :class="
              reloadType === type
                ? 'bg-neutral-600 text-white font-bold'
                : 'bg-neutral-900 text-neutral-400 hover:bg-neutral-700'
            "
          >
            {{ typesMap[type] }}
          </button>
        </div>

        <!-- RESET TIME -->
        <div>
          <label class="text-sm text-neutral-400 mb-1 block"> Время сброса </label>
          <input
            v-model="resetTime"
            type="time"
            class="w-full rounded-lg border border-neutral-700 bg-neutral-900 p-2 text-neutral-200"
          />
        </div>

        <!-- WEEKLY -->
        <div v-if="reloadType === 'weekly'">
          <label class="text-sm text-neutral-400 mb-1 block"> День недели </label>
          <select
            v-model="resetWeekday"
            class="w-full rounded-lg border border-neutral-700 bg-neutral-900 p-2 text-neutral-200"
          >
            <option :value="0">Воскресенье</option>
            <option :value="1">Понедельник</option>
            <option :value="2">Вторник</option>
            <option :value="3">Среда</option>
            <option :value="4">Четверг</option>
            <option :value="5">Пятница</option>
            <option :value="6">Суббота</option>
          </select>
        </div>

        <!-- CUSTOM -->
        <div v-if="reloadType === 'custom'">
          <label class="text-sm text-neutral-400 mb-1 block"> Интервал (дней) </label>
          <input
            v-model.number="reloadEvery"
            type="number"
            min="1"
            class="w-full rounded-lg border border-neutral-700 bg-neutral-900 p-2 text-neutral-200"
          />
        </div>
      </div>

      <!-- ACTIONS -->
      <div class="mt-6 flex justify-end gap-2">
        <button
          type="button"
          @click="tasksStore.toggleIsCreateModalOpen()"
          class="px-3 py-1.5 rounded-lg text-neutral-400 hover:text-neutral-200 hover:bg-neutral-700 transition"
        >
          Отмена
        </button>

        <button
          type="submit"
          :disabled="!isValid"
          class="px-3 py-1.5 rounded-lg transition-all"
          :class="
            isValid
              ? 'bg-neutral-700 hover:bg-neutral-600 font-bold text-white'
              : 'bg-neutral-700 opacity-50 cursor-not-allowed text-neutral-400'
          "
        >
          Создать
        </button>
      </div>
    </div>
  </form>
</template>
