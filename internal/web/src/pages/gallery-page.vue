<script setup lang="ts">
import PageLayout from "@/components/page-layout.vue";
import Gallery from "@/components/day-gallery.vue";
import type { Ref } from "vue";
import { ref } from "vue";
import dayjs from "dayjs";
import { useDaemonStore } from "@/stores/daemon";
import { getMediaList } from "@/lib/media-api";
import type { GalleryImage } from "@/lib/gallery";
import { buildImage } from "@/lib/gallery";
import { useAuth0 } from "@auth0/auth0-vue";
import type { Auth0VueClient } from "@auth0/auth0-vue/src/global";
import { BoltSlashIcon } from "@heroicons/vue/24/solid";

const imagesMap: Ref<Map<string, Array<GalleryImage>>> = ref(new Map());
const daemonStore = useDaemonStore();
const auth0Client = useAuth0();

const getAllMedia = async (
  auth0Client: Auth0VueClient,
  daemonId: string
): Promise<Map<string, Array<GalleryImage>>> => {
  const photos = await getMediaList(auth0Client, daemonId, 0, 25);

  let imagesMap = new Map();

  if (photos) {
    for await (const photo of photos) {
      let parsedDate = dayjs(photo.dateTime);
      let day = parsedDate.format("YYYY-MM-DD");

      let galleryImage = await buildImage(
        auth0Client,
        photo,
        daemonStore.id,
        parsedDate
      );
      let gallery = imagesMap.get(day);
      if (gallery) {
        gallery.push(galleryImage);
      } else {
        imagesMap.set(day, new Array<GalleryImage>(galleryImage));
      }
    }
  }

  return imagesMap;
};

daemonStore.$subscribe((mutation, state) => {
  if (state.id) {
    getAllMedia(auth0Client, state.id).then((map) => (imagesMap.value = map));
  }
});

if (daemonStore.id) {
  getAllMedia(auth0Client, daemonStore.id).then(
    (map) => (imagesMap.value = map)
  );
}
</script>

<template>
  <PageLayout>
    <div class="content-layout">
      <div class="content__body">
        <template v-if="daemonStore.active">
          <Gallery
            v-for="[date, images] in imagesMap"
            :key="date"
            gallery-i-d="my-test-gallery"
            :day="date"
            :images="images"
          />
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
