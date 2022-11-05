<script setup lang="ts">
import MobileNavBarTab from "@/components/navigation/mobile/mobile-nav-bar-tab.vue";
import { useAuth0 } from "@auth0/auth0-vue";

const { isAuthenticated } = useAuth0();
</script>

<template>
  <div class="mobile-nav-bar__tabs">
    <template v-if="isAuthenticated">
      <MobileNavBarTab path="/profile" label="Profile" />
      <MobileNavBarTab path="/gallery" label="Gallery" />
      <MobileNavBarTab path="/daemon" label="Daemons" />
    </template>
  </div>
</template>
