<template>
<div class="relative w-full vw-16 vh-9">
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

        <ul class="absolute hidden peer-hover:block hover:block z-10 max-h-56 overflow-y-auto scrollbar-hide rounded-md py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm bg-gray-300 bg-opacity-10" tabindex="-1" role="listbox" aria-labelledby="listbox-label" aria-activedescendant="listbox-option-3">
          <li class="text-gray-900 relative cursor-pointer select-none py-2 pl-3 pr-9 hover:scale-105" 
              v-for="(item, index) in countries" 
              :key="index" 
              role="option" 
              @click="changeCountry(index)" >
            <div class="flex items-center">
              <span :class="'fi fi-' + item.code.toLowerCase()"></span>
              <span class="font-normal ml-3 block truncate text-white">{{ item.name }}</span>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </div>
  <div class="absolute top-0 left-0 h-full w-52 cursor-pointer" id="side">
    <ul role="list" class="p-3 h-full w-full overflow-y-auto bg-gray-300 bg-opacity-10">
      <li class="flex py-4 hover:scale-105" 
          v-for="(item, index) in channels" 
          :key="index" 
          @click="changeChannel(index)" >
        <img class="h-10 w-10 rounded-full" :src=item.logo alt="">
        <div class="ml-3">
          <p class="text-sm font-medium text-white">{{ item.alt_names.length ? item.alt_names[0] : item.name }}</p>
          <p class="text-sm text-gray-500">{{ item.languages.join(" ") }}</p>
        </div>
      </li>
    </ul>
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
let mp;

// 国家
watch(curCountry, async (newVal, oldVal) => {
  emits('updateChannel', newVal.code)
})
if (propos.countries.length > 0) {
  curCountry.value = propos.countries[0]
}
function changeCountry(index) {
  curCountry.value = propos.countries[index];
}

// 频道
watch(curChannel, async (newVal, oldVal) => {
  emits('updateStream', newVal.id)
})
watch(() => propos.channels, async (newVal, oldVal) => {
  if (propos.channels.length>0 && !curChannel.value) {
    curChannel.value = propos.channels[0]
  }
})
function changeChannel(index) {
  curChannel.value = propos.channels[index]
}

// 视频流
function changeStream(index) {
  mp.reloadUrl(propos.streams[index].url)
}

onMounted(() => {

  mp = new MuiPlayer({
    container: "#video-player",
    //title: propos.title,
    src: "",
    autoFit: true,
    autoplay: true,
    objectFit: "cover",
    height: "100%",
    closeControlsTimer: 10000000000,
    live: true,
    parse: {
      type: "hls",
      loader: Hls,
      config: {}
    }
  });
  // 这个是为了等待播放器创建完成，不能省略
  mp.getControls();

  // 视频流
  watch(() => propos.streams, async (newVal, oldVal) => {
    
    if (propos.streams.length > 0) {
      mp.reloadUrl(propos.streams[0].url)
    } else {
      mp.showToast("未找到IPTV源")
    }  
  })

})

onUnmounted(() => {
  mp.destroy()
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
.mplayer-header,.mplayer-footer {
  display: none!important;
}
</style>