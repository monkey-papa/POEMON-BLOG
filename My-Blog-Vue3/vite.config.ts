import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";
import { fileURLToPath } from "url";
import { visualizer } from "rollup-plugin-visualizer";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

const __dirname = fileURLToPath(new URL(".", import.meta.url));

// 插件：修复 vue3-seamless-scroll 组件的非被动 wheel 事件监听器
function fixVue3SeamlessScrollWheel() {
  return {
    name: "fix-vue3-seamless-scroll-wheel",
    enforce: "pre",
    transform(code: string, id: string) {
      if (id.includes("vue3-seamless-scroll")) {
        const fixedCode = code.replace(
          /addEventListener\s*\(\s*['"]wheel['"]\s*,\s*([a-zA-Z_$][a-zA-Z0-9_$]*)\s*\)/g,
          "addEventListener('wheel', $1, { passive: true })"
        );

        if (fixedCode !== code) {
          console.log(`[fix-vue3-seamless-scroll-wheel] 已修复文件 (transform): ${id}`);
          return {
            code: fixedCode,
            map: null,
          };
        }
      }
      return null;
    },
  };
}

export default defineConfig({
  plugins: [
    vue(),
    fixVue3SeamlessScrollWheel(), // 修复 vue3-seamless-scroll 组件的 wheel 事件
    // Element Plus 按需导入
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
    // 打包体积分析插件（类似 webpack-bundle-analyzer）
    // 只在构建时生成报告，不自动打开（可通过 npm run analyze 查看）
    visualizer({
      filename: "dist/stats.html", // 分析报告输出路径
      open: false, // 不自动打开，手动查看
      gzipSize: true, // 显示 gzip 压缩后的大小
      brotliSize: true, // 显示 brotli 压缩后的大小
      template: "treemap", // 使用 treemap 视图（类似 webpack-bundle-analyzer）
    }),
  ],
  resolve: {
    alias: {
      "@": resolve(__dirname, "src"),
    },
  },
  server: {
    port: 81,
    host: true,
    open: false,
    hmr: {
      overlay: false,
    },
  },
  esbuild: {
    drop: ["debugger"],
  },
  build: {
    outDir: "dist",
    assetsDir: "assets",
    sourcemap: false,
    minify: "esbuild",
    // 代码分割优化
    rollupOptions: {
      output: {
        chunkFileNames: "js/[name]-[hash].js",
        entryFileNames: "js/[name]-[hash].js",
        assetFileNames: "[ext]/[name]-[hash].[ext]",
        // 手动代码分割，将大型库单独打包
        manualChunks(id) {
          if (id.includes("node_modules")) {
            if (id.includes("echarts")) {
              return "charts";
            }
            if (id.includes("highlightjs-line-numbers")) {
              return "hljs-line-numbers";
            }
            return "vendor";
          }
        },
      },
    },
    // 提高构建性能
    chunkSizeWarningLimit: 1000, // 块大小警告限制（KB）
  },
  css: {
    devSourcemap: false, // 禁用开发环境的 CSS source map，避免缺失 map 文件的警告
    preprocessorOptions: {
      scss: {
        silenceDeprecations: ["legacy-js-api"],
      },
    },
  },
  optimizeDeps: {
    include: [
      "jquery",
      "animejs",
      "highlight.js",
      "wow.js",
      "clipboard",
      "nprogress",
      "echarts",
      "aplayer",
      "colorthief",
      "tocbot",
    ],
    // 强制重新构建 vue3-seamless-scroll，以便应用修复
    force: false,
  },
  // 自定义警告处理，忽略 source map 和 eval 相关的警告
  logLevel: "warn",
  customLogger: (() => {
    const warnedMessages = new Set<string>();
    return {
      hasWarned: false,
      hasErrorLogged: (error: Error | { message?: string }) =>
        warnedMessages.has(error.message ?? String(error)),
      warnOnce: (msg: string) => {
        if (!warnedMessages.has(msg)) {
          warnedMessages.add(msg);
          console.warn(msg);
        }
      },
      warn: (msg: string) => {
        if (
          msg.includes("source map") ||
          msg.includes(".map") ||
          msg.includes("ENOENT") ||
          msg.includes("Failed to load source map")
        ) {
          return;
        }
        console.warn(msg);
      },
      error: (msg: string) => {
        if (msg.includes("source map") || msg.includes(".map") || msg.includes("ENOENT")) {
          return;
        }
        console.error(msg);
      },
      info: (msg: string) => {
        console.info(msg);
      },
      clearScreen: () => {},
    };
  })(),
});
