<script setup lang="ts">
import { useDaemonStore } from "@/stores/daemon";
import { onMounted, onUnmounted } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";
import { BoltIcon } from "@heroicons/vue/24/solid";
import { BoltSlashIcon } from "@heroicons/vue/24/solid";
import { CheckCircleIcon } from "@heroicons/vue/24/solid";
import { XCircleIcon } from "@heroicons/vue/24/solid";
import { registerWatcher, unregisterWatcher } from "@/lib/daemon-service";
import dayjs from "dayjs";
import duration from "dayjs/plugin/duration";
import relativeTime from "dayjs/plugin/relativeTime";

const daemonStore = useDaemonStore();
dayjs.extend(duration);
dayjs.extend(relativeTime);

onMounted(async () => {
  await registerWatcher(useAuth0(), daemonStore);
});

onUnmounted(() => {
  unregisterWatcher();
});
</script>

<template>
  <div class="dropdown dropdown-hover dropdown-bottom dropdown-end">
    <label tabindex="0" class="btn">
      <template v-if="daemonStore.active">
        <BoltIcon class="h-5 w-5 text-base-500" />
      </template>
      <template v-else>
        <BoltSlashIcon class="h-5 w-5 text-red-500" />
      </template>
    </label>
    <div
      tabindex="0"
      class="dropdown-content card card-compact w-72 shadow-xl mt-4 bg-info-content"
    >
      <div class="card-body">
        <h3 class="card-title text-accent text-sm">Daemon connection</h3>
        <div class="flex flex-row gap-2">
          <div class="flex h-5 w-5">
            <template v-if="daemonStore.active">
              <CheckCircleIcon
                class="h-5 w-5 text-green-500 absolute inline-flex animate-ping opacity-75"
              />
              <CheckCircleIcon
                class="relative inline-flex h-5 w-5 text-green-500"
              />
            </template>
            <template v-else>
              <XCircleIcon
                class="h-5 w-5 text-red-500 absolute inline-flex animate-ping opacity-75"
              />
              <XCircleIcon class="relative inline-flex h-5 w-5 text-red-500" />
            </template>
          </div>

          <div class="flex flex-col">
            <span>{{ daemonStore.name }}</span>
            <span v-if="daemonStore.hostname" class="text-xs"
              >{{ daemonStore.hostname }} • {{ daemonStore.version }}</span
            >
            <template v-if="!daemonStore.active && daemonStore.lastSeenStr">
              <span class="text-xs"
                >Last seen {{ daemonStore.lastSeenStr }}</span
              >
            </template>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
