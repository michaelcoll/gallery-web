import { authGuard } from "@auth0/auth0-vue";
import { createRouter, createWebHistory } from "vue-router";

import HomePage from "@/pages/home-page.vue";

const NotFoundPage = () => import("@/pages/not-found-page.vue");
const GalleryPage = () => import("@/pages/gallery-page.vue");
const ProfilePage = () => import("@/pages/profile-page.vue");

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
    path: "/profile",
    name: "profile",
    component: ProfilePage,
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
