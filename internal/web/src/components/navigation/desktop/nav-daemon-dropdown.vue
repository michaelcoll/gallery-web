<template>
  <router-link
    :to="props.path"
    exact
    class="nav-bar__daemon"
    active-class="nav-bar__daemon--active"
  >
    Daemons
    <span class="daemon-name">{{ activeDaemonName }}</span>
  </router-link>
</template>

<script setup>
import { getDaemonList } from "@/services/daemon.service";
import { ref } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";

const props = defineProps({
  path: String,
});

const activeDaemonName = ref("No active daemon");

const getCurrentDaemon = async () => {
  const { getAccessTokenSilently } = useAuth0();
  const accessToken = await getAccessTokenSilently();
  const { data, error } = await getDaemonList(accessToken);

  if (data) {
    if (data.length > 0) {
      data.forEach((daemon) => {
        if (daemon.alive) {
          sessionStorage.currentDaemonName = daemon.name;
          sessionStorage.currentDaemonId = daemon.id;
          activeDaemonName.value = daemon.name;
        }
      });
    }
  }

  if (error) {
    console.log("error", error);
  }
};

getCurrentDaemon();
</script>
