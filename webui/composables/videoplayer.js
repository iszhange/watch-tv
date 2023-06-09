export const useVideoPlayerStore = definePiniaStore(
  'videoplayer', 
  () => {
    const country = ref()
    const channel = ref()
    const channels = ref(Array)
    const stream = ref()
    const streams = ref(Array)

    return {country, channel, channels, stream, streams}
  },
  {
    persist: true,
    //storage: persistedState.localStorage,
  }
)