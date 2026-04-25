import type { Directive, DirectiveBinding } from "vue";

import { useUserStore } from "@/store";
import { ROLE_ROOT } from "@/constants";

/**
 * 按钮权限
 */
export const hasPerm: Directive = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const requiredPerms = binding.value;

    if (!requiredPerms || (typeof requiredPerms !== "string" && !Array.isArray(requiredPerms))) {
      return;
    }

    const { roles, perms } = useUserStore().userInfo;

    // 超级管理员或通配符权限 "*" 拥有所有权限
    if (roles.includes(ROLE_ROOT) || roles.includes("admin") || perms.includes("*") || requiredPerms.includes("*:*:*")) {
      return;
    }

    // 检查权限
    const hasAuth = Array.isArray(requiredPerms)
      ? requiredPerms.some((perm) => perms.includes(perm))
      : perms.includes(requiredPerms);

    if (!hasAuth && el.parentNode) {
      el.parentNode.removeChild(el);
    }
  },
};

/**
 * 角色权限指令
 */
export const hasRole: Directive = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const requiredRoles = binding.value;

    if (!requiredRoles || (typeof requiredRoles !== "string" && !Array.isArray(requiredRoles))) {
      return;
    }

    const { roles } = useUserStore().userInfo;

    // admin 视为超级管理员
    if (roles.includes(ROLE_ROOT) || roles.includes("admin")) {
      return;
    }

    const hasAuth = Array.isArray(requiredRoles)
      ? requiredRoles.some((role) => roles.includes(role))
      : roles.includes(requiredRoles);

    if (!hasAuth && el.parentNode) {
      el.parentNode.removeChild(el);
    }
  },
};
