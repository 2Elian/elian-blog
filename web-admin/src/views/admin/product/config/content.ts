import type { IContentConfig } from "@/components/CURD/types";
import { ProductAPI } from "@/api/product";

const contentConfig: IContentConfig<any> = {
  pageTitle: "产品列表",
  permPrefix: "product",
  table: {
    border: true,
    highlightCurrentRow: true,
  },
  pagination: {
    background: true,
    layout: "prev,pager,next,jumper,total,sizes",
    pageSize: 10,
    pageSizes: [10, 20, 30, 50],
  },
  parseData: (res) => {
    return {
      total: res.data.total,
      list: res.data.list || [],
    };
  },
  deleteAction: function (ids: string) {
    return ProductAPI.deletesProductApi({
      ids: ids.split(",").map((id) => parseInt(id)),
    });
  },
  indexAction: function (params: any) {
    return ProductAPI.findProductListApi(params);
  },
  pk: "id",
  toolbar: [
    {
      name: "add",
      text: "新增产品",
      perm: "add",
      attrs: {
        icon: "plus",
        type: "success",
      },
    },
    {
      name: "delete",
      text: "删除",
      perm: "delete",
      attrs: {
        icon: "delete",
        type: "danger",
      },
    },
  ],
  defaultToolbar: ["refresh", "filter", "exports", "search"],
  cols: [
    {
      type: "selection",
      label: "批量操作",
      width: 50,
      align: "center",
    },
    {
      label: "ID",
      prop: "id",
      width: 70,
      align: "center",
      sortable: true,
    },
    {
      label: "封面",
      prop: "cover",
      width: 100,
      align: "center",
      templet: "image",
      imageWidth: 60,
    },
    {
      label: "产品名称",
      prop: "name",
      minWidth: 180,
      align: "center",
    },
    {
      label: "类型",
      prop: "type",
      width: 100,
      align: "center",
      templet: "custom",
    },
    {
      label: "价格",
      prop: "price",
      width: 100,
      align: "center",
    },
    {
      label: "简介",
      prop: "description",
      minWidth: 200,
      align: "center",
      templet: "custom",
    },
    {
      label: "排序",
      prop: "sort",
      width: 80,
      align: "center",
      sortable: true,
    },
    {
      label: "状态",
      prop: "status",
      width: 80,
      align: "center",
      templet: "custom",
    },
    {
      label: "创建时间",
      prop: "created_at",
      width: 170,
      align: "center",
      sortable: true,
      templet: "date",
      dateFormat: "YYYY/MM/DD HH:mm:ss",
    },
    {
      label: "操作栏",
      align: "center",
      fixed: "right",
      width: 220,
      templet: "tool",
      operat: [
        {
          name: "edit",
          text: "编辑",
          perm: "edit",
          attrs: {
            icon: "edit",
            type: "primary",
          },
        },
        {
          name: "editDetail",
          text: "编辑详情",
          perm: "edit",
          attrs: {
            icon: "document",
            type: "warning",
          },
        },
        {
          name: "delete",
          text: "删除",
          perm: "delete",
          attrs: {
            icon: "delete",
            type: "danger",
          },
        },
      ],
    },
  ],
};

export default contentConfig;
