import vue from "@vitejs/plugin-vue";
import { type ConfigEnv, type UserConfig, loadEnv, defineConfig } from "vite";

import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

import UnoCSS from "unocss/vite";
import { resolve } from "path";
import { name, version } from "./package.json";

const __APP_INFO__ = {
  pkg: { name, version },
  buildTimestamp: Date.now(),
};

const pathSrc = resolve(__dirname, "src");

export default defineConfig(({ mode }: ConfigEnv): UserConfig => {
  const env = loadEnv(mode, process.cwd());
  return {
    base: "/",
    resolve: {
      alias: {
        "@": pathSrc,
      },
    },
    css: {
      preprocessorOptions: {
        scss: {
          additionalData: `@use "@/styles/variables.scss" as *;`,
        },
      },
    },
    server: {
      host: "0.0.0.0",
      port: 3001,
      open: true,
      proxy: {
        "/admin-api": {
          changeOrigin: true,
          target: "http://localhost:8080",
          rewrite: (path) => path,
        },
      },
    },
    plugins: [
      vue(),
      UnoCSS(),
      AutoImport({
        imports: ["vue", "@vueuse/core", "pinia", "vue-router"],
        resolvers: [ElementPlusResolver()],
        vueTemplate: true,
        dts: false,
      }),
      Components({
        resolvers: [ElementPlusResolver()],
        dirs: ["src/components", "src/**/components"],
        dts: false,
      }),
    ],
    optimizeDeps: {
      include: [
        "vue",
        "vue-router",
        "element-plus",
        "pinia",
        "axios",
        "@vueuse/core",
        "@element-plus/icons-vue",
        "nprogress",
        "qs",
        "echarts",
      ],
    },
    build: {
      chunkSizeWarningLimit: 2000,
      minify: "terser",
      terserOptions: {
        compress: {
          drop_console: true,
          drop_debugger: true,
        },
      },
    },
    define: {
      __APP_INFO__: JSON.stringify(__APP_INFO__),
    },
  };
});