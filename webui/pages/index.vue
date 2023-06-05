<template>
  <div>
    <!-- https://cd-live-stream.news.cctvplus.com/live/smil:CHANNEL1.smil/playlist.m3u8 -->
    <VideoPlayer title="测试" src="https://cd-live-stream.news.cctvplus.com/live/smil:CHANNEL1.smil/playlist.m3u8" poster="" :countries="countries" />
  </div>
</template>
<script setup>
definePageMeta({ layout: 'default' })

const appConfig = useAppConfig()
let countries = ref()

await $fetch('/countries', {
  baseURL: appConfig.apiHost,
}).then(res => {
  if (res.code == 200) {
    countries = res.data.list
  }
}).catch(err => {
  console.log(err)
})


</script>
