<template>
  <div>
    <!-- https://cd-live-stream.news.cctvplus.com/live/smil:CHANNEL1.smil/playlist.m3u8 -->
    <VideoPlayer title="测试" src="https://cd-live-stream.news.cctvplus.com/live/smil:CHANNEL1.smil/playlist.m3u8" poster="" />
  </div>
</template>
<script setup>
definePageMeta({ layout: 'default' })

const appConfig = useAppConfig()

const countries = await $fetch(appConfig.apiHost+'/countries', {})
console.log(111, countries)

</script>
