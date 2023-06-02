<template>
  <h1 className="text-3xl font-bold underline">Hello world!</h1>
  <div>
    <div>
      <!-- https://cd-live-stream.news.cctvplus.com/live/smil:CHANNEL1.smil/playlist.m3u8 -->
      <VideoPlayer title="测试" src="http://rzggis.com:8111/hls/57/index.m3u8" poster="" />
    </div>
  </div>
</template>
<script>
definePageMeta({ layout: 'default' })
export default {
}
</script>
