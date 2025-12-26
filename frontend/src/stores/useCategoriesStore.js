import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { CreateCategory, DeleteCategory, GetCategories, GetTasks } from '../../wailsjs/go/main/App'
import { models } from '../../wailsjs/go/models'

export const useCategoriesStore = defineStore('categories', () => {
  const categories = ref([])
  const activeCategory = ref(null)
  const isCreateCategoryModalOpen = ref(false)
  const isDeleteModalOpen = ref(false)

  const toDeleteCategory = ref(null)

  async function deleteCategory(category) {
    const tasks = await GetTasks(category.id)
    console.log('tasks', tasks)
    if (tasks.length === 0) {
      deleteCategoryForce(category)
    } else {
      isDeleteModalOpen.value = true
      toDeleteCategory.value = category
    }
  }
  function deleteCategoryForce(category) {
    isDeleteModalOpen.value = false
    toDeleteCategory.value = null
    DeleteCategory(category.id).then(() => {
      loadCategories()
    })
  }

  function createCategory(name, exeName) {
    CreateCategory(name, exeName).then((id) => {
      console.log('Category created')
      loadCategories(id)
    })
  }

  function toggleIsCreateModalOpen() {
    isCreateCategoryModalOpen.value = !isCreateCategoryModalOpen.value
  }

  async function loadCategories(activeCategoryId = null) {
    const res = await GetCategories()
    categories.value = res
    if (activeCategoryId) {
      activeCategory.value = categories.value.find((cat) => cat.id === activeCategoryId)
    } else {
      activeCategory.value = categories.value[0]
    }
    console.log('Categories loaded', res)
  }

  function setActiveCategory(category) {
    activeCategory.value = category
  }

  return {
    categories,
    activeCategory,
    isCreateCategoryModalOpen,
    isDeleteModalOpen,
    toDeleteCategory,
    deleteCategory,
    deleteCategoryForce,
    createCategory,
    loadCategories,
    setActiveCategory,
    toggleIsCreateModalOpen,
  }
})
