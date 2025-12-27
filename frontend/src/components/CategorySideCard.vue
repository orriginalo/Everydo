<script setup>
import { ref } from 'vue'
import { useCategoriesStore } from '../stores/useCategoriesStore'
import { storeToRefs } from 'pinia'
import { models } from '../../wailsjs/go/models'
import { Icon } from '@iconify/vue'

const props = defineProps({
  category: {
    type: models.Category,
    required: true,
  },
})

const categoriesStore = useCategoriesStore()
const { activeCategory, openedMenuCategoryId } = storeToRefs(categoriesStore)
const { setActiveCategory, deleteCategory, closeCategoryMenu, toggleCategoryMenu } = categoriesStore

const toggleMenu = (e) => {
  e.stopPropagation() // чтобы клик по кнопке не выбирал категорию
  toggleCategoryMenu(props.category.id)
}

const editCategory = () => {
  closeCategoryMenu()
  categoriesStore.setEditCategory(props.category)
  categoriesStore.toggleIsEditModalOpen()
}
</script>

<template>
  <div class="relative">
    <button
      @click="setActiveCategory(props.category)"
      class="group flex rounded-md px-3 py-2 justify-between text-left transition-all w-full"
      :class="
        activeCategory.id === props.category.id
          ? 'bg-neutral-800 text-white shadow-inner'
          : 'text-neutral-400 hover:bg-neutral-900 hover:text-neutral-200'
      "
    >
      <div class="flex flex-row items-center gap-3">
        <div
          class="h-8 w-8 rounded bg-neutral-700 flex items-center justify-center text-sm font-bold"
        >
          {{ props.category.name[0] }}
        </div>

        <span class="truncate font-unbounded">
          {{ props.category.name }}
        </span>
      </div>

      <button @click="toggleMenu" class="hover:bg-neutral-600 transition-all rounded-xl py-1 px-0">
        <Icon
          icon="ic:baseline-more-vert"
          class="transition-all duration-300"
          :class="openedMenuCategoryId === props.category.id ? '' : 'rotate-90'"
          width="20"
        />
      </button>
    </button>

    <!-- Выпадающее меню с анимацией -->
    <transition name="fade-scale">
      <div
        v-if="openedMenuCategoryId === props.category.id"
        class="absolute right-2 top-10 w-36 bg-neutral-900 border border-neutral-700 rounded-xl shadow-xl z-50 overflow-hidden"
      >
        <button
          @click="editCategory"
          class="w-full text-left px-4 py-2 text-sm hover:bg-neutral-800 transition-colors"
        >
          Изменить
        </button>
        <button
          @click="
            () => {
              closeCategoryMenu()
              deleteCategory(props.category)
            }
          "
          class="w-full text-left px-4 py-2 text-sm hover:bg-red-950 text-red-400 transition-colors"
        >
          Удалить
        </button>
      </div>
    </transition>
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
