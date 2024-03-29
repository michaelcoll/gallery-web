import "./assets/css/styles.css";

import { createAuth0 } from "@auth0/auth0-vue";
import { createPinia } from "pinia";
import { createApp } from "vue";

import App from "./app.vue";
import router from "./router";

const pinia = createPinia();
const app = createApp(App);

app
  .use(router)
  .use(
    createAuth0({
      domain: import.meta.env.VITE_AUTH0_DOMAIN,
      clientId: import.meta.env.VITE_AUTH0_CLIENT_ID,
      authorizationParams: {
        redirect_uri: import.meta.env.VITE_AUTH0_CALLBACK_URL,
        audience: import.meta.env.VITE_AUTH0_API_AUDIENCE,
      },
      cacheLocation: "localstorage",
      useRefreshTokens: true,
    }),
  )
  .use(pinia)
  .mount("#root");
