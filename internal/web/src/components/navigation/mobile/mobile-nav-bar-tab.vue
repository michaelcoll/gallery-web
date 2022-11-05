<script setup lang="ts">
const props = defineProps<{
  path: string;
  label: string;
}>();
</script>

<template>
  <router-link
    :to="props.path"
    exact
    class="mobile-nav-bar__tab"
    active-class="mobile-nav-bar__tab--active"
  >
    {{ props.label }}
  </router-link>
</template>
