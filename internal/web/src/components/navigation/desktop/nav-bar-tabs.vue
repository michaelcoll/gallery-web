<script setup lang="ts">
import NavBarTab from "@/components/navigation/desktop/nav-bar-tab.vue";
import NavDaemonDropDown from "@/components/navigation/desktop/nav-daemon-dropdown.vue";
import { useAuth0 } from "@auth0/auth0-vue";

const { isAuthenticated } = useAuth0();
</script>

<template>
  <div class="nav-bar__tabs">
    <template v-if="isAuthenticated">
      <NavDaemonDropDown path="/daemon" />
      <NavBarTab path="/gallery" label="Gallery" />
    </template>
  </div>
</template>
