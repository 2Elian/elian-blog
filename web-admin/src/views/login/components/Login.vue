<template>
  <div class="auth-panel-form">
    <h3 class="auth-panel-form__title" text-center>登 录</h3>
    <el-form
      ref="loginFormRef"
      :model="loginFormData"
      :rules="loginRules"
      size="large"
      :validate-on-rule-change="false"
    >
      <!-- 用户名 -->
      <el-form-item prop="username">
        <el-input v-model.trim="loginFormData.username" placeholder="用户名">
          <template #prefix>
            <el-icon><User /></el-icon>
          </template>
        </el-input>
      </el-form-item>

      <!-- 密码 -->
      <el-tooltip :visible="isCapsLock" content="大写锁定已打开" placement="right">
        <el-form-item prop="password">
          <el-input
            v-model.trim="loginFormData.password"
            placeholder="密码"
            type="password"
            show-password
            @keyup="checkCapsLock"
            @keyup.enter="handleLoginSubmit"
          >
            <template #prefix>
              <el-icon><Lock /></el-icon>
            </template>
          </el-input>
        </el-form-item>
      </el-tooltip>

      <div class="flex-x-between w-full">
        <el-checkbox v-model="rememberMe">记住我</el-checkbox>
      </div>

      <!-- 登录按钮 -->
      <el-form-item>
        <el-button :loading="loading" type="primary" class="w-full" @click="handleLoginSubmit">
          登 录
        </el-button>
      </el-form-item>
    </el-form>

    <!-- 注册入口 -->
    <div flex-center gap-10px>
      <div class="w-full h-[20px] flex justify-between items-center">
        <el-button
          class="w-full mt-4!"
          size="default"
          @click="toOtherForm('register')"
        >
          注册
        </el-button>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import type { FormInstance, FormRules } from "element-plus";

import router from "@/router";
import { useUserStore } from "@/store";
import { AuthStorage } from "@/utils/auth";
import { LoginReq } from "@/api/types";
import { User } from "@element-plus/icons-vue";

const userStore = useUserStore();
const route = useRoute();

const loginFormRef = ref<FormInstance>();
const loading = ref(false);
const isCapsLock = ref(false);
const rememberMe = ref(AuthStorage.getRememberMe());

watch(rememberMe, (val) => {
  AuthStorage.setRememberMe(val);
});

const loginFormData = ref<LoginReq>({
  username: "admin",
  password: "admin123",
});

const loginRules = computed<FormRules>(() => ({
  username: [{ required: true, trigger: "blur", message: "请输入用户名" }],
  password: [
    { required: true, trigger: "blur", message: "请输入密码" },
    { min: 6, message: "密码不能少于6位", trigger: "blur" },
  ],
}));

async function handleLoginSubmit() {
  try {
    const valid = await loginFormRef.value?.validate();
    if (!valid) return;

    loading.value = true;
    await userStore.login(loginFormData.value);
    const redirectPath = (route.query.redirect as string) || "/";
    await router.push(decodeURIComponent(redirectPath));
  } catch (error) {
    ElMessage.error("登录失败。" + error);
  } finally {
    loading.value = false;
  }
}

function checkCapsLock(event: KeyboardEvent) {
  if (event instanceof KeyboardEvent) {
    isCapsLock.value = event.getModifierState("CapsLock");
  }
}

const emit = defineEmits(["update:modelValue"]);

function toOtherForm(type: string) {
  emit("update:modelValue", type);
}
</script>

<style lang="scss" scoped>
.auth-panel-form {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.auth-panel-form__title {
  margin: 0 0 0.5rem;
  font-size: 1.125rem;
  font-weight: 600;
}
</style>
