import type { IModalConfig } from "@/components/CURD/types";
import { ProductAPI } from "@/api/product";

const modalConfig: IModalConfig<any> = {
  permPrefix: "product",
  component: "dialog",
  dialog: {
    title: "编辑产品",
    width: 650,
    draggable: true,
  },
  pk: "id",
  formAction: function (data) {
    return ProductAPI.updateProductApi(data);
  },
  formItems: [
    {
      label: "产品名称",
      prop: "name",
      rules: [{ required: true, message: "产品名称不能为空", trigger: "blur" }],
      type: "input",
      attrs: {
        placeholder: "请输入产品名称",
      },
    },
    {
      label: "产品类型",
      prop: "type",
      type: "input",
      attrs: {
        placeholder: "如: AI产品、工具、其他",
      },
    },
    {
      label: "价格",
      prop: "price",
      type: "input-number" as any,
      attrs: {
        placeholder: "请输入价格",
        min: 0,
        precision: 2,
      },
    },
    {
      label: "封面图",
      prop: "cover",
      type: "custom",
      slotName: "cover",
    },
    {
      label: "简介",
      prop: "description",
      type: "input" as any,
      attrs: {
        type: "textarea",
        placeholder: "请输入产品简介",
        rows: 3,
      },
    },
    {
      label: "排序",
      prop: "sort",
      type: "input-number" as any,
      attrs: {
        placeholder: "排序值",
        min: 0,
      },
    },
    {
      label: "状态",
      prop: "status",
      type: "select",
      attrs: {
        placeholder: "请选择状态",
      },
      options: [
        { label: "上架", value: 1 },
        { label: "下架", value: 0 },
      ],
    },
  ],
};

export default modalConfig;
