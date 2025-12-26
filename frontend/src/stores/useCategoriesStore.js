import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { CreateCategory, GetCategories } from '../../wailsjs/go/main/App'
import { models } from '../../wailsjs/go/models'

export const useCategoriesStore = defineStore('categories', () => {
  const categories = ref([])
  const activeCategory = ref(null)
  const isCreateCategoryModalOpen = ref(false)

  function loadCategories() {
    GetCategories().then((res) => {
      categories.value = res
    })
  }

  function setActiveCategory(category) {
    activeCategory.value = category
  }

  return {
    categories,
    activeCategory,
    isCreateCategoryModalOpen,
    loadCategories,
    setActiveCategory,
  }
})
