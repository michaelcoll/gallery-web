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

<template>
  <span class="m-2 text-sm">{{ dayjs(day).format("ddd D MMM") }}</span>
  <div :id="galleryID" class="gallery-container">
    <figure v-for="(image, key) in imagesData" :key="key" class="gallery-item">
      <a
        :href="image.largeURL"
        :data-pswp-width="image.width"
        :data-pswp-height="image.height"
        target="_blank"
        rel="noreferrer"
      >
        <img :src="image.thumbnailURL" alt="" loading="lazy" />
        <!--suppress HtmlUnknownTag -->
        <figcaption class="pswp-caption-content">
          <strong>Image Info</strong><br />
          Date : {{ image.date }}
        </figcaption>
      </a>
    </figure>
  </div>
</template>
