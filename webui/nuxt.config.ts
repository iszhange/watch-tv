// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  app: {
    head: {
      charset: 'utf-8',
      viewport: 'width=500, initial-scale=1',
      title: 'Watch TV',
      meta: [
        { name: 'description', content: 'Watch TV Web Site.' }
      ],
    },
  },
  ssr: false,
  modules: [
    '@nuxtjs/tailwindcss',
  ],
  css: ['@/assets/css/main.css'],
  runtimeConfig: {  
  },
  appConfig: {
    apiHost: "http://127.0.0.1:81"
  }
})
