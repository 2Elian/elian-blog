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
      type: "select",
      attrs: {
        placeholder: "请选择类型",
      },
      options: [
        { label: "AI产品", value: 1 },
        { label: "工具", value: 2 },
        { label: "其他", value: 3 },
      ],
    },
    {
      label: "价格",
      prop: "price",
      type: "number",
      attrs: {
        placeholder: "请输入价格",
        min: 0,
        precision: 2,
      },
    },
    {
      label: "封面图",
      prop: "cover",
      type: "input",
      attrs: {
        placeholder: "封面图URL",
      },
    },
    {
      label: "产品链接",
      prop: "link",
      type: "input",
      attrs: {
        placeholder: "请输入产品链接",
      },
    },
    {
      label: "描述",
      prop: "description",
      type: "textarea",
      attrs: {
        placeholder: "请输入产品描述",
        rows: 4,
      },
    },
    {
      label: "排序",
      prop: "sort",
      type: "number",
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
