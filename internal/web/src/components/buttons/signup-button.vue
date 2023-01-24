<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";

const { loginWithRedirect } = useAuth0();

const handleSignUp = () => {
  loginWithRedirect({
    authorizationParams: {
      prompt: "login",
      screen_hint: "signup",
    },
    appState: {
      target: "/gallery",
    },
  });
};
</script>

<template>
  <button class="btn btn-outline" @click="handleSignUp">Sign Up</button>
</template>
