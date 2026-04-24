import type { RouteRecordRaw } from "vue-router";

const Layout = () => import("@/layouts/index.vue");

export default {
  path: "/product",
  name: "",
  component: Layout,
  redirect: "/product/list",
  meta: {
    title: "产品管理",
    icon: "ep:goods",
    rank: 9,
  },
  children: [
    {
      path: "/product/list",
      component: () => import("@/views/admin/product/Product.vue"),
      name: "ProductList",
      meta: { title: "产品列表", keepAlive: true },
    },
  ],
} satisfies RouteRecordRaw;
