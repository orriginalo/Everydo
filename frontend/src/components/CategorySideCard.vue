<script setup lang="ts">
import { useCategoriesStore } from '@/stores/categories'
import { models } from '../../wailsjs/go/models'

const props = defineProps<{
  cat: models.Category
}>()

const store = useCategoriesStore()
</script>

<template>
  <button
    @click="store.activeCat = props.cat"
    class="group flex items-center gap-3 rounded-md px-3 py-2 text-left transition-all"
    :class="
      store.activeCat === props.cat
        ? 'bg-neutral-800 text-white shadow-inner'
        : 'text-neutral-400 hover:bg-neutral-800 hover:text-white'
    "
  >
    <!-- Иконка-заглушка -->
    <div class="h-8 w-8 rounded bg-neutral-700 flex items-center justify-center text-sm font-bold">
      {{ props.cat.name[0] }}
    </div>

    <span class="truncate">
      {{ props.cat.name }}
    </span>
  </button>
</template>
