import type { RouteRecordRaw } from "vue-router";

const Layout = () => import("@/layouts/index.vue");

export default {
  path: "/social",
  name: "",
  component: Layout,
  redirect: "/social/friend",
  meta: {
    title: "社交管理",
    icon: "el-icon-link",
    rank: 3,
    hidden: true,
  },
  children: [
    {
      path: "/social/friend",
      component: () => import("@/views/admin/social/friend/Friend.vue"),
      name: "Friend",
      meta: { title: "友链管理", keepAlive: true },
    },
  ],
} satisfies RouteRecordRaw;
