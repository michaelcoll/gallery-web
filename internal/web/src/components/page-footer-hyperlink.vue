<script setup lang="ts">
const props = defineProps<{
  path: string;
}>();
</script>

<template>
  <a
    :href="props.path"
    class="page-footer__hyperlink"
    target="_blank"
    rel="noopener noreferrer"
  >
    <slot />
  </a>
</template>
