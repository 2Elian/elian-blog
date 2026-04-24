import type { RouteRecordRaw } from "vue-router";

const Layout = () => import("@/layouts/index.vue");

export default {
  path: "/news",
  name: "",
  component: Layout,
  redirect: "/news/comment",
  meta: {
    title: "消息管理",
    icon: "el-icon-message",
    rank: 2,
  },
  children: [
    {
      path: "/news/comment",
      component: () => import("@/views/admin/news/comment/Comment.vue"),
      name: "Comment",
      meta: { title: "评论管理", keepAlive: true },
    },
    {
      path: "/news/ai-assistant",
      component: () => import("@/views/admin/news/ai-assistant/AiAssistant.vue"),
      name: "AiAssistant",
      meta: { title: "AI助手", keepAlive: true },
    },
  ],
} satisfies RouteRecordRaw;
