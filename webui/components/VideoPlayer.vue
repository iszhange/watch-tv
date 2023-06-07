<template>
<div class="relative">
  <div id="video-player"></div>
  <div class="absolute top-0 w-full" id="header" >
    <div class="flex justify-center">
      <div v-if="curCountry" class="inline-block">
        <button type="button" class="relative peer w-full cursor-pointer rounded-md py-1.5 pl-3 pr-3 text-left text-gray-900 sm:text-sm sm:leading-6" aria-haspopup="listbox" aria-expanded="true" aria-labelledby="listbox-label">
          <span class="flex items-center">
            <span :class="'fi fi-' + curCountry.code.toLowerCase()"></span>
            <span class="ml-3 block truncate">{{ curCountry.name }}</span>
          </span>
        </button>

        <ul class="absolute hidden peer-hover:block hover:block z-10 max-h-56 overflow-y-auto scrollbar-hide rounded-md py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm" tabindex="-1" role="listbox" aria-labelledby="listbox-label" aria-activedescendant="listbox-option-3">
          <li class="text-gray-900 relative cursor-pointer select-none py-2 pl-3 pr-9" 
              v-for="(item, index) in countries" 
              :key="index" 
              role="option" 
              @click="changeCountry(index)" >
            <div class="flex items-center">
              <span :class="'fi fi-' + item.code.toLowerCase()"></span>
              <span class="font-normal ml-3 block truncate">{{ item.name }}</span>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
  <div class="absolute top-0 left-0 h-full w-60 border-solid bg-red-500" id="side">
    1111
  </div>
</div>  
</template>
<script setup>
import 'mui-player/dist/mui-player.min.css'
import "/node_modules/flag-icons/css/flag-icons.min.css"
import MuiPlayer from 'mui-player'
import Hls from 'hls.js'
import { onMounted, onUnmounted, ref, watch, nextTick } from 'vue'

const propos = defineProps({
  title: String,
  src: String,
  poster: String,
  countries: Array,
  channels: Array,
  streams: Array,
})
const emits = defineEmits([
  'updateChannel',
  'updateStream'
])

const curCountry = ref()
const curChannel = ref()
const curStream = ref()
let mp

// 国家
if (propos.countries.length > 0) {
  curCountry.value = propos.countries[0]
  emits('updateChannel', propos.countries[0].code)
}
function changeCountry(index) {
  curCountry.value = propos.countries[index];
}

// 频道
watch(curCountry, async (newVal, oldVal) => {
  emits('updateChannel', newVal.code)
})
watch(propos.channels, async (newVal, oldVal) => {
  if (propos.channels.length>0 && !curChannel) {
    curChannel.value = propos.channels[0];
    emits('updateStream', propos.channels[0].id)
  }
})
function changeChannel(index) {
  curChannel.value = propos.channels[index]
}

// 视频流
watch(propos.streams, async (newVal, oldVal) => {
  if (propos.streams.length>0 && !curStream) {
    curStream.value = propos.streams[0]
  }
})
watch(curChannel, async (newVal, oldVal) => {
  emits('updateStream', newVal.id)
})
function changeStream(index) {
  curStream.value = propos.streams[index]
}

onMounted(() => {
  mp = new MuiPlayer({
    container: "#video-player",
    title: propos.title,
    src: propos.src,
    autoFit: true,
    objectFit: "cover",
    //height: "auto",
    closeControlsTimer: 10000000000,
    parse: {
      type: "hls",
      loader: Hls,
      config: {}
    }
  });
  
  mp.getControls().forEach((ele) => {
    if (ele.parentNode.parentNode.id == "mplayer-header") {
      console.log(ele.parentNode.parentNode)

      ele.parentNode.parentNode.parentNode.removeChild(ele.parentNode.parentNode)
    }
  });

});
onUnmounted(() => {
  if (mp) {
    mp.destroy();
  }
});
</script>
<style>
/* 隐藏ul滚动条 */
ul::-webkit-scrollbar {
  display: none;
}
ul {
  scrollbar-width: none;
  -ms-overflow-style: none;
}
</style>