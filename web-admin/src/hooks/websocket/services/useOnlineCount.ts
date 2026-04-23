import { ref } from "vue";

export function useOnlineCount() {
  const onlineUserCount = ref(0);
  const lastUpdateTime = ref(0);
  const isConnected = ref(false);
  const isConnecting = ref(false);

  return {
    onlineUserCount,
    lastUpdateTime,
    isConnected,
    isConnecting,
    initWebSocket: () => {},
    closeWebSocket: () => {},
  };
}
