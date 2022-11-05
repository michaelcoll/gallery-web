<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";

const { logout } = useAuth0();

const handleLogout = () =>
  logout({
    returnTo: window.location.origin,
  });
</script>

<template>
  <button class="button__logout" @click="handleLogout">Log Out</button>
</template>
