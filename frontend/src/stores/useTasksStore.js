import { defineStore } from 'pinia'
import { CompleteTask, CreateTask, GetTasks, UncompleteTask } from '../../wailsjs/go/main/App'
import { models } from '../../wailsjs/go/models'
import { ref } from 'vue'

export const useTasksStore = defineStore('tasks', () => {
  const tasksByCategory = ref({})
  const isCreateModalOpen = ref(false)

  function completeTask(task) {
    CompleteTask(task.id).then(() => {
      console.log('Task completed')
      loadTasks(task.category_id)
    })
  }

  function uncompleteTask(task) {
    UncompleteTask(task.id).then(() => {
      console.log('Task uncompleted')
      loadTasks(task.category_id)
    })
  }

  function loadTasks(categoryId) {
    GetTasks(categoryId)
      .then((tasks) => {
        console.log('tasks loaded')
        console.log(tasks)
        tasksByCategory.value[categoryId] = tasks
      })
      .catch((err) => {
        console.log(err)
      })
  }

  function toggleIsCreateModalOpen() {
    isCreateModalOpen.value = !isCreateModalOpen.value
  }

  function createTask(categoryId, name, reloadType, reloadEvery, resetTime, resetWeekday) {
    return CreateTask(categoryId, name, reloadType, reloadEvery, resetTime, resetWeekday).then(
      (id) => {
        tasksByCategory.value[categoryId].push({
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
    tasksByCategory,
    isCreateModalOpen,
    completeTask,
    uncompleteTask,
    toggleIsCreateModalOpen,
    createTask,
    loadTasks,
  }
})
