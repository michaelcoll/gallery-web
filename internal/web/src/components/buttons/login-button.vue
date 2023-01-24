<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";

const { loginWithRedirect } = useAuth0();

const handleLogin = () => {
  loginWithRedirect({
    authorizationParams: {
      prompt: "login",
    },
    appState: {
      target: "/gallery",
    },
  });
};
</script>

<template>
  <button class="btn btn-accent" @click="handleLogin">Log In</button>
</template>
