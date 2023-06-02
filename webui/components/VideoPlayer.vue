<template>
  <div>
    <div id="video-player"></div>
  </div>
</template>
<script setup>
import 'mui-player/dist/mui-player.min.css'
import MuiPlayer from 'mui-player'
import Hls from 'hls.js'
import { onMounted, onUnmounted, defineProps } from 'vue'

const propos = defineProps({
  title: String,
  src: String,
  poster: String,
})

let mp;

onMounted(() => {
  mp = new MuiPlayer({
    container: "#video-player",
    title: propos.title,
    src: propos.src,
    parse: {
      type: "hls",
      loader: Hls,
      config: {}
    }
  });
});
onUnmounted(() => {
  if (mp) {
    mp.destory();
  }
});
</script>