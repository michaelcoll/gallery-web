<script setup lang="ts">
import NavBar from "@/components/navigation/desktop/nav-bar.vue";
import PageFooter from "@/components/page-footer.vue";
</script>

<template>
  <div class="page-layout">
    <NavBar />
    <div class="page-layout__content">
      <slot />
    </div>
    <PageFooter />
  </div>
</template>
