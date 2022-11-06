<script setup lang="ts">
const props = defineProps<{
  path: string;
  label: string;
}>();
</script>

<template>
  <router-link :to="props.path" exact class="tab" active-class="tab-active">
    {{ props.label }}
  </router-link>
</template>
