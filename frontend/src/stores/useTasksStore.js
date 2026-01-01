import { defineStore } from 'pinia'
import {
  CompleteTask,
  CreateTask,
  GetTasks,
  UncompleteTask,
  DeleteTask,
  UpdateTask,
} from '../../wailsjs/go/main/App'
import { models } from '../../wailsjs/go/models'
import { ref } from 'vue'
import { UpdateTasksOrder } from '../../wailsjs/go/main/App'

export const useTasksStore = defineStore('tasks', () => {
  const tasksByCategory = ref({})
  const isCreateModalOpen = ref(false)
  const isEditModalOpen = ref(false)

  const toEditTask = ref(null)

  const isDeleteModalOpen = ref(false)
  const toDeleteTask = ref(null)

  const openedMenuTaskId = ref(null)

  function toggleTaskMenu(taskId) {
    openedMenuTaskId.value = openedMenuTaskId.value === taskId ? null : taskId
  }

  function closeTaskMenu() {
    openedMenuTaskId.value = null
  }

  async function deleteTask(task) {
    isDeleteModalOpen.value = true
    toDeleteTask.value = task
  }
  function deleteTaskForce(task) {
    isDeleteModalOpen.value = false
    toDeleteTask.value = null
    DeleteTask(task.id).then(() => {
      loadTasks(task.category_id)
    })
  }

  // function updateCategory(id, name, exeName) {
  //     UpdateCategory(id, name, exeName).then(() => {
  //       console.log('Category updated')
  //       loadCategories(id)
  //     })
  //   }
  // like this

  function updateTask(cat_id, id, name, reloadType, reloadEvery, resetTime, resetWeekday) {
    UpdateTask(id, name, reloadType, reloadEvery, resetTime, resetWeekday).then(() => {
      console.log('Task updated')
      loadTasks(cat_id)
    })
  }

  function toggleIsEditModalOpen() {
    isEditModalOpen.value = !isEditModalOpen.value
  }

  function setEditTask(task) {
    toEditTask.value = task
  }

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
        tasks.sort((a, b) => a.order - b.order)
        tasksByCategory.value[categoryId] = tasks
      })
      .catch((err) => {
        console.error(err)
      })
  }

  function reorderTasks(categoryId, reloadType, newOrder) {
    const tasks = tasksByCategory.value[categoryId]
    if (!tasks) return

    let order = 0

    // обновляем order только для нужного типа
    for (const task of newOrder) {
      const t = tasks.find((x) => x.id === task.id && x.reload_type === reloadType)
      if (t) {
        t.order = order++
      }
    }

    // пересортировать глобальный массив
    tasks.sort((a, b) => a.order - b.order)

    // ⬇️ СОХРАНЯЕМ В БЭК
    UpdateTasksOrder(
      newOrder.map((t, i) => ({
        id: t.id,
        order: i,
      })),
    )
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
    isDeleteModalOpen,
    toDeleteTask,
    isEditModalOpen,
    toEditTask,
    openedMenuTaskId,
    reorderTasks,
    toggleTaskMenu,
    closeTaskMenu,
    completeTask,
    uncompleteTask,
    toggleIsCreateModalOpen,
    createTask,
    loadTasks,
    toggleIsEditModalOpen,
    setEditTask,
    deleteTask,
    deleteTaskForce,
    updateTask,
  }
})
