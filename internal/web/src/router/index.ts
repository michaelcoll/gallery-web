import HomePage from "@/pages/home-page.vue";
import { authGuard } from "@auth0/auth0-vue";
import { createRouter, createWebHistory } from "vue-router";

const NotFoundPage = () => import("@/pages/not-found-page.vue");
const GalleryPage = () => import("@/pages/gallery-page.vue");
const DaemonPage = () => import("@/pages/daemon-page.vue");

const routes = [
  {
    path: "/",
    name: "home",
    component: HomePage,
  },
  {
    path: "/gallery",
    name: "gallery",
    component: GalleryPage,
    beforeEnter: authGuard,
  },
  {
    path: "/daemon",
    name: "daemon",
    component: DaemonPage,
    beforeEnter: authGuard,
  },
  {
    path: "/:catchAll(.*)",
    name: "Not Found",
    component: NotFoundPage,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
