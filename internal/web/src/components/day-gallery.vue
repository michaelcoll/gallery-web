<!--
  - Copyright (c) 2022 Michaël COLL.
  -
  - Licensed under the Apache License, Version 2.0 (the "License");
  - you may not use this file except in compliance with the License.
  - You may obtain a copy of the License at
  -
  -      http://www.apache.org/licenses/LICENSE-2.0
  -
  - Unless required by applicable law or agreed to in writing, software
  - distributed under the License is distributed on an "AS IS" BASIS,
  - WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  - See the License for the specific language governing permissions and
  - limitations under the License.
  -->

<template>
  <span class="gallery-day">{{ dayjs(day).format("ddd D MMM") }}</span>
  <div :id="galleryID" class="gallery-container">
    <figure v-for="(image, key) in imagesData" :key="key" class="gallery-item">
      <a
        :href="image.largeURL"
        :data-pswp-width="image.width"
        :data-pswp-height="image.height"
        target="_blank"
        rel="noreferrer"
      >
        <img :src="image.thumbnailURL" alt="" />
        <figcaption class="pswp-caption-content">
          <strong>Image Info</strong><br />
          Date : {{ image.date }}
        </figcaption>
      </a>
    </figure>
  </div>
</template>

<script setup>
import { onMounted, onUnmounted, ref } from "vue";
import PhotoSwipeLightbox from "photoswipe/lightbox";
import PhotoSwipeDynamicCaption from "photoswipe-dynamic-caption-plugin";
import "photoswipe/style.css";
import "photoswipe-dynamic-caption-plugin/photoswipe-dynamic-caption-plugin.css";
import dayjs from "dayjs";
import "dayjs/locale/fr";

const props = defineProps({
  galleryID: String,
  day: String,
  images: Array,
});

const imagesData = ref([]);
const lightbox = ref();

dayjs.locale("fr");

onMounted(() => {
  if (!lightbox.value) {
    lightbox.value = new PhotoSwipeLightbox({
      gallery: "#" + props.galleryID,
      children: "figure",
      pswpModule: () => import("photoswipe"),
      spacing: 0.5,
    });

    new PhotoSwipeDynamicCaption(lightbox.value, {
      // Plugins options, for example:
      type: "auto",
    });

    lightbox.value.init();
  }
  imagesData.value = props.images;
});

onUnmounted(() => {
  if (lightbox.value) {
    lightbox.value.destroy();
    lightbox.value = null;
  }
});
</script>
