<script setup lang="ts">
import { MoonIcon } from "@heroicons/vue/24/solid";
import { SunIcon } from "@heroicons/vue/24/solid";
</script>

<template>
  <label class="swap swap-rotate btn">
    <!-- this hidden checkbox controls the state -->
    <input type="checkbox" />

    <!-- sun icon -->
    <SunIcon class="swap-on fill-current h-5 w-5 text-base-500" />

    <!-- moon icon -->
    <MoonIcon class="swap-off fill-current h-5 w-5 text-base-500" />
  </label>
</template>
