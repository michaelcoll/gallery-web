<script setup lang="ts">
import { useDaemonStore } from "@/stores/daemon";
import { onMounted, onUnmounted, ref } from "vue";
import { daemonIsAlive, getDaemonList } from "@/lib/daemon-api";
import { useAuth0 } from "@auth0/auth0-vue";

const props = defineProps<{
  path: string;
}>();

const daemonStore = useDaemonStore();
const intervalId = ref();

const refreshCurrentDaemon = async () => {
  const daemons = await getDaemonList(useAuth0());

  if (daemons) {
    if (daemons.length > 0) {
      daemons.forEach((daemon) => {
        if (daemon.alive) {
          daemonStore.useDaemon(daemon.id, daemon.name);
        }
      });
    }
  }
};

const testDaemonIsAlive = async () => {
  const isAlive = await daemonIsAlive(daemonStore.id);

  if (!isAlive) {
    daemonStore.clearDaemon();
    clearInterval(intervalId.value);
  }
};

if (
  (daemonStore.hasActiveDaemon && !daemonIsAlive(daemonStore.id)) ||
  !daemonStore.hasActiveDaemon
) {
  refreshCurrentDaemon();
}

onMounted(() => {
  intervalId.value = setInterval(function () {
    testDaemonIsAlive();
  }, 10000);
});

onUnmounted(() => {
  clearInterval(intervalId.value);
});
</script>

<template>
  <router-link
    :to="props.path"
    exact
    class="nav-bar__daemon"
    active-class="nav-bar__daemon--active"
  >
    Daemons
    <span class="daemon-name">{{ daemonStore.name }}</span>
  </router-link>
</template>
