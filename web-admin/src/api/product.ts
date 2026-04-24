import request from "@/utils/request";

export const ProductAPI = {
  findProductListApi(data?: any): Promise<any> {
    return request({
      url: "/admin-api/v1/product/find_product_list",
      method: "POST",
      data,
    });
  },
  addProductApi(data?: any): Promise<any> {
    return request({
      url: "/admin-api/v1/product/add_product",
      method: "POST",
      data,
    });
  },
  updateProductApi(data?: any): Promise<any> {
    return request({
      url: "/admin-api/v1/product/update_product",
      method: "PUT",
      data,
    });
  },
  deletesProductApi(data?: any): Promise<any> {
    return request({
      url: "/admin-api/v1/product/deletes_product",
      method: "DELETE",
      data,
    });
  },
};
