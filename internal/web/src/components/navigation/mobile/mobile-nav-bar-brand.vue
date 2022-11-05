<script setup lang="ts">
import { mobileMenuStore } from "./mobile-menu.store";
</script>

<template>
  <div class="mobile-nav-bar__brand" @click="mobileMenuStore.closeMobileMenu">
    <router-link to="/">
      <img
        class="mobile-nav-bar__logo"
        src="https://cdn.auth0.com/blog/hub/code-samples/hello-world/auth0-logo.svg"
        alt="Auth0 shield logo"
        width="82"
        height="24"
      />
    </router-link>
  </div>
</template>
