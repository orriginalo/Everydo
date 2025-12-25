import { defineStore } from 'pinia'
import { CreateTask, GetTasks } from '../../wailsjs/go/main/App'
import { models } from '../../wailsjs/go/models'
import { ref } from 'vue'

export const useTasksStore = defineStore('tasks', () => {
  const tasksByCategory = ref<Record<number, models.Task[]>>({})
  const isCreateModalOpen = ref(false)

  function getTasks(categoryId: number) {
    if (tasksByCategory.value[categoryId]) return

    GetTasks(categoryId).then((list) => {
      tasksByCategory.value[categoryId] = list
    })
  }

  function createTask(
    categoryId: number,
    payload: {
      name: string
      reloadType: string
      reloadEvery: number
      resetTime: string
      resetWeekday?: number
    },
  ) {
    return CreateTask(
      categoryId,
      payload.name,
      payload.reloadType,
      payload.reloadEvery,
      payload.resetTime,
      payload.resetWeekday,
    ).then((id) => {
      // Инициализация массива, если ещё нет
      if (!tasksByCategory.value[categoryId]) {
        tasksByCategory.value[categoryId] = []
      }

      tasksByCategory.value[categoryId].push(
        new models.Task({ id, category_id: categoryId, ...payload }),
      )
    })
  }

  return {
    tasksByCategory,
    getTasks,
    createTask,
    isCreateModalOpen,
  }
})
