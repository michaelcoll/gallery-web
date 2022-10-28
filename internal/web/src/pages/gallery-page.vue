<template>
  <PageLayout>
    <div class="content-layout">
      <!--      <h1 id="page-title" class="content__title">Gallery</h1>-->
      <div class="content__body">
        <Gallery
          v-for="[date, images] in imagesMap"
          galleryID="my-test-gallery"
          :key="date"
          :day="date"
          :images="images"
        />
      </div>
    </div>
  </PageLayout>
</template>

<script setup>
import PageLayout from "@/components/page-layout.vue";
import Gallery from "@/components/day-gallery.vue";
import { ref } from "vue";
import { useAuth0 } from "@auth0/auth0-vue";
import { getMediaList } from "@/services/media.service";
import dayjs from "dayjs";

const apiServerUrl = import.meta.env.VITE_API_SERVER_URL;
const imagesMap = ref(new Map());

const getAllMedia = async () => {
  const { getAccessTokenSilently } = useAuth0();
  const accessToken = await getAccessTokenSilently();
  const { data, error } = await getMediaList(
    accessToken,
    sessionStorage.currentDaemonId,
    0,
    25
  );

  if (data) {
    data.forEach((photo) => {
      let parsedDate = dayjs(photo.dateTime);
      let day = parsedDate.format("YYYY-MM-DD");
      if (!imagesMap.value.has(day)) {
        imagesMap.value.set(day, []);
      }

      imagesMap.value.get(day).push({
        largeURL: `${apiServerUrl}/api/daemon/${sessionStorage.currentDaemonId}/media/${photo.hash}/content?access-token=${accessToken}`,
        thumbnailURL: `${apiServerUrl}/api/daemon/${sessionStorage.currentDaemonId}/media/${photo.hash}/thumbnail?access-token=${accessToken}`,
        width: photo.xDimension,
        height: photo.yDimension,
        date: parsedDate,
      });
    });
  }

  if (error) {
    console.log("error", error);
  }
};

getAllMedia();
</script>
