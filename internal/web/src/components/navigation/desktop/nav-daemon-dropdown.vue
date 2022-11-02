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

<script setup>
import { getDaemonList, daemonIsAlive } from "@/services/daemon.service";
import { useAuth0 } from "@auth0/auth0-vue";
import { useDaemonStore } from "@/stores/daemon";
import { onMounted, onUnmounted, ref } from "vue";

const props = defineProps({
  path: String,
});

const daemonStore = useDaemonStore();
const intervalId = ref();

const refreshCurrentDaemon = async () => {
  const { getAccessTokenSilently } = useAuth0();
  const accessToken = await getAccessTokenSilently();
  const { data, error } = await getDaemonList(accessToken);

  if (data) {
    if (data.length > 0) {
      data.forEach((daemon) => {
        if (daemon.alive) {
          daemonStore.useDaemon(daemon.id, daemon.name);
        }
      });
    }
  }

  if (error) {
    console.log("error", error);
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
