<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";

import PageLoader from "@/components/page-loader.vue";

const { isLoading } = useAuth0();
</script>

<template>
  <div v-if="isLoading" class="page-layout">
    <PageLoader />
  </div>
  <router-view v-else />
</template>
