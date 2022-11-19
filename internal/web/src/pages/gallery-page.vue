<script setup lang="ts">
import PageLayout from "@/components/page-layout.vue";
import Gallery from "@/components/day-gallery.vue";
import type { Ref } from "vue";
import { onMounted, onUnmounted, ref } from "vue";
import { useDaemonStore } from "@/stores/daemon";
import { getMediaList } from "@/lib/media-api";
import type { GalleryImage } from "@/lib/gallery";
import { useAuth0 } from "@auth0/auth0-vue";
import type { Auth0VueClient } from "@auth0/auth0-vue/src/global";
import { BoltSlashIcon } from "@heroicons/vue/24/solid";
import { updateImageMap } from "@/lib/gallery";

const imagesMap: Ref<Map<string, Array<GalleryImage>>> = ref(new Map());
const currentPage = ref(0);
const lastPage = ref(false);
const infiniteList = ref(null);
const daemonStore = useDaemonStore();
const auth0Client = useAuth0();

const initGallery = async (auth0Client: Auth0VueClient, daemonId: string) => {
  const photos = await getMediaList(auth0Client, daemonId, 0, 25);

  updateImageMap(photos, imagesMap, daemonStore.id);
};

const addPage = async () => {
  currentPage.value++;

  const photos = await getMediaList(
    auth0Client,
    daemonStore.id,
    currentPage.value,
    25
  );

  if (photos.length > 0) {
    updateImageMap(photos, imagesMap, daemonStore.id);
  } else {
    lastPage.value = true;
    currentPage.value--;
  }
};

onMounted(() => {
  window.addEventListener("scroll", handleScroll);
});

onUnmounted(() => {
  window.addEventListener("scroll", handleScroll);
});

const handleScroll = () => {
  let element = infiniteList.value;
  if (
    element?.getBoundingClientRect()?.bottom < window.innerHeight &&
    !lastPage.value
  ) {
    addPage();
  }
};

daemonStore.$subscribe((mutation, state) => {
  if (state.id) {
    initGallery(auth0Client, state.id);
    currentPage.value = 0;
  }
});

if (daemonStore.id) {
  initGallery(auth0Client, daemonStore.id);
}
</script>

<template>
  <PageLayout>
    <div class="content-layout">
      <div class="content__body">
        <template v-if="daemonStore.active">
          <div ref="infiniteList">
            <Gallery
              v-for="[date, images] in imagesMap"
              :id="'g' + date"
              :key="date"
              :day="date"
              :images="images"
            />
          </div>
        </template>
        <template v-else>
          <div class="hero min-h-screen">
            <div class="hero-overlay bg-transparent"></div>
            <div class="hero-content text-center text-neutral-content">
              <div class="max-w-md">
                <BoltSlashIcon class="h-20 w-20 text-base-content m-auto" />
                <p class="mb-5">No active daemon</p>
              </div>
            </div>
          </div>
        </template>
      </div>
    </div>
  </PageLayout>
</template>
