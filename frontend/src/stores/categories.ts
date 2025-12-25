import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { CreateCategory, GetCategories } from '../../wailsjs/go/main/App'
import { models } from '../../wailsjs/go/models'

export const useCategoriesStore = defineStore('categories', () => {
  const categories = ref<models.Category[]>([])
  const activeCat = ref<models.Category | null>(null)
  const isCreateModalOpen = ref(false)

  function getCategories() {
    return GetCategories().then((res) => {
      categories.value = res

      const firstCat = res[0] // может быть undefined
      if (!activeCat.value && firstCat) {
        activeCat.value = firstCat
      }
    })
  }


  function setActive(cat: models.Category) {
    activeCat.value = cat
  }

  function toggleIsCreateModalOpen() {
    isCreateModalOpen.value = !isCreateModalOpen.value
  }

  return {
    categories,
    activeCat,
    setActive,
    getCategories,
    toggleIsCreateModalOpen,
    isCreateModalOpen,
  }
})
