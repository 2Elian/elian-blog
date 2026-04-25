import { ref, onMounted, onUnmounted } from "vue";
import { WebsiteAPI } from "@/api/website";

const POLL_INTERVAL = 15_000; // 15秒轮询一次

export function useOnlineCount() {
  const onlineUserCount = ref(0);
  const lastUpdateTime = ref(0);
  const isConnected = ref(false);
  const isConnecting = ref(false);

  let timer: ReturnType<typeof setInterval> | null = null;

  async function fetchCount() {
    try {
      isConnecting.value = true;
      const res = await WebsiteAPI.getOnlineCountApi();
      onlineUserCount.value = res.data.online_count;
      lastUpdateTime.value = res.data.timestamp * 1000;
      isConnected.value = true;
    } catch {
      isConnected.value = false;
    } finally {
      isConnecting.value = false;
    }
  }

  function startPolling() {
    fetchCount();
    timer = setInterval(fetchCount, POLL_INTERVAL);
  }

  function stopPolling() {
    if (timer) {
      clearInterval(timer);
      timer = null;
    }
  }

  onMounted(() => {
    startPolling();
  });

  onUnmounted(() => {
    stopPolling();
  });

  return {
    onlineUserCount,
    lastUpdateTime,
    isConnected,
    isConnecting,
    initWebSocket: startPolling,
    closeWebSocket: stopPolling,
  };
}
