<script setup lang="ts">
import { mobileMenuStore } from "./mobile-menu.store";
</script>

<template>
  <span
    id="mobile-menu-toggle-button"
    class="mobile-nav-bar__toggle material-icons"
    @click="mobileMenuStore.toggleMobileMenu()"
  >
    {{ mobileMenuStore.mobileMenuIcon }}
  </span>
</template>
