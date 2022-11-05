<script setup lang="ts">
import PageFooterHyperlink from "@/components/page-footer-hyperlink.vue";
</script>

<template>
  <footer class="page-footer">
    <div class="page-footer-grid">
      <div class="page-footer-grid__info">
        <div class="page-footer-info__message"></div>
        <div class="page-footer-info__resource-list"></div>
      </div>
      <div class="page-footer-grid__brand">
        <div class="page-footer-brand">
          <img
            class="page-footer-brand__logo"
            src="https://cdn.auth0.com/blog/hello-auth0/auth0-shield.svg"
            alt="Auth0"
            width="20"
            height="22.22"
          />
          <PageFooterHyperlink path="https://auth0.com/">
            Auth0 Inc
          </PageFooterHyperlink>
        </div>
      </div>
    </div>
  </footer>
</template>
