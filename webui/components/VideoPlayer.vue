<template>
  <div id="video-player"></div>
  <div ref="headerbox" class="flex justify-center">
    <div v-if="curCountry" class="inline-block">
      <button type="button" class="relative w-full cursor-default rounded-md bg-white py-1.5 pl-3 pr-3 text-left text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 focus:outline-none focus:ring-2 focus:ring-indigo-500 sm:text-sm sm:leading-6" aria-haspopup="listbox" aria-expanded="true" aria-labelledby="listbox-label">
      <span class="flex items-center">
        <span class="fi fi-gr"></span>
        <span class="ml-3 block truncate">{{ curCountry.name }}</span>
      </span>
    </button>

    <ul v-if="countries" class="absolute z-10 mt-1 max-h-56 overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm" tabindex="-1" role="listbox" aria-labelledby="listbox-label" aria-activedescendant="listbox-option-3">
      <!--
        Select option, manage highlight styles based on mouseenter/mouseleave and keyboard navigation.

        Highlighted: "bg-indigo-600 text-white", Not Highlighted: "text-gray-900"
      -->
      <li class="text-gray-900 relative cursor-default select-none py-2 pl-3 pr-9" id="listbox-option-{{ index }}" role="option" @click="(index)=>{this.curCountry = countries[index];}" v-for="(item, index) in countries" >
        <div class="flex items-center">
          <img src="https://images.unsplash.com/photo-1491528323818-fdd1faba62cc?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80" alt="" class="h-5 w-5 flex-shrink-0 rounded-full">
          <!-- Selected: "font-semibold", Not Selected: "font-normal" -->
          <span class="font-normal ml-3 block truncate">{{ item.name }}</span>
        </div>
        <span class="text-indigo-600 absolute inset-y-0 right-0 flex items-center pr-4">
          <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
            <path fill-rule="evenodd" d="M16.704 4.153a.75.75 0 01.143 1.052l-8 10.5a.75.75 0 01-1.127.075l-4.5-4.5a.75.75 0 011.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 011.05-.143z" clip-rule="evenodd" />
          </svg>
        </span>
      </li>

      <!-- More items... -->
    </ul>
    </div>
  </div>
</template>
<script setup>
import 'mui-player/dist/mui-player.min.css'
import "/node_modules/flag-icons/css/flag-icons.min.css"
import MuiPlayer from 'mui-player'
import Hls from 'hls.js'
import { onMounted, onUnmounted, defineProps, ref, watch } from 'vue'

const propos = defineProps({
  title: String,
  src: String,
  poster: String,
  countries: Array,
})

const headerbox = ref(null)
const curCountry = ref(null);
let mp;

watch(propos.countries, async (newVal) => {
  if (newVal.length > 0) {
    curCountry = newVal[0]
  }
})

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
    },
    custom: {
      headControls: [
        {
          slot:'castScreen', // 对应定义的 slot 值
          click: (e) => {
            let sele = e.srcElement
              ,value = sele.options[sele.selectedIndex].value;
            console.log(value);
          },
          style:{},
        }
      ]
    }
  });

  mp.getControls().forEach((ele) => {
    if (ele.id == "back-button") {
      ele.hidden = true
    }
  });

});
onUnmounted(() => {
  if (mp) {
    mp.destory();
  }
});
</script>