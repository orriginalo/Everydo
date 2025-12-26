import { defineStore } from 'pinia'
import { CreateTask, GetTasks } from '../../wailsjs/go/main/App'
import { models } from '../../wailsjs/go/models'
import { ref } from 'vue'

export const useTasksStore = defineStore('tasks', () => {
  const tasksByCategory = ref({})
  const isCreateModalOpen = ref(false)

  function toggleIsCreateModalOpen() {
    isCreateModalOpen.value = !isCreateModalOpen.value
  }

  function createTask(categoryId, name, reloadType, reloadEvery, resetTime, resetWeekday) {
    return CreateTask(categoryId, name, reloadType, reloadEvery, resetTime, resetWeekday).then(
      (id) => {
        if (!tasksByCategory.value[categoryId]) {
          tasksByCategory.value[categoryId] = {
            daily: [],
            weekly: [],
            custom: [],
          }
        }

        tasksByCategory.value[categoryId][reloadType].push({
          id: id,
          category_id: categoryId,
          name: name,
          reload_type: reloadType,
          reload_every: reloadEvery,
          reset_time: resetTime,
          reset_weekday: resetWeekday,
        })
      },
    )
  }

  return {
    toggleIsCreateModalOpen,
    createTask,
  }
})
