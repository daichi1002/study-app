// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ["@nuxtjs/tailwindcss"],
  runtimeConfig: {
    baseUrl: process.env.BASE_URL || "http://localhost:8080/",
  },
});
