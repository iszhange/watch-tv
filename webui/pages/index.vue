<template>
  <div>
    <!-- https://cd-live-stream.news.cctvplus.com/live/smil:CHANNEL1.smil/playlist.m3u8 -->
    <VideoPlayer title="测试" 
                 src="https://cd-live-stream.news.cctvplus.com/live/smil:CHANNEL1.smil/playlist.m3u8" 
                 poster="" 
                 :countries="countries"
                 :channels="channels"
                 :streams="channelStreams"
                 @update-channel="getChannels"
                 @update-stream="getStreams" />
  </div>
</template>
<script setup>
definePageMeta({ layout: 'default' })

const appConfig = useAppConfig()
const countries = ref(Array)
const channels = ref(Array)
const channelStreams = ref(Array)

await $fetch('/countries', {
  baseURL: appConfig.apiHost,
}).then(res => {
  if (res.code == 200) {
    countries.value = res.data.list
  }
}).catch(err => {
  console.log(err)
})

// 更新频道
function getChannels(code) {
  $fetch('/channels', {
    baseURL: appConfig.apiHost,
    query: {
      country: code,
      is_show: 1,
    }
  }).then(res => {
    if (res.code == 200) {
      channels.value = res.data.list
    }
  }).catch(err => {
    console.log(err)
  })
}

// 更新视频流
function getStreams(channel) {
  $fetch('/streams', {
    baseURL: appConfig.apiHost,
    query: {
      channel: channel,
      is_show: 1,
    }
  }).then(res => {
    if (res.code == 200) {
      channelStreams.value = res.data.list
    }
  }).catch(err => {
    console.log(err)
  })
}

</script>
