<template>
  <div class="login-page">
    <div class="login-card">
      <h2>{{ isLogin ? '登录' : '注册' }}</h2>

      <el-form ref="formRef" :model="form" :rules="rules" label-position="top" @submit.prevent="handleSubmit">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="form.username" placeholder="请输入用户名" />
        </el-form-item>

        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password" type="password" placeholder="请输入密码" show-password />
        </el-form-item>

        <el-form-item v-if="!isLogin" label="确认密码" prop="password2">
          <el-input v-model="form.password2" type="password" placeholder="再次输入密码" show-password />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" :loading="loading" block @click="handleSubmit">
            {{ isLogin ? '登录' : '注册' }}
          </el-button>
        </el-form-item>
      </el-form>

      <div class="switch-mode">
        <span v-if="isLogin">还没有账号？<a @click="isLogin = false">立即注册</a></span>
        <span v-else>已有账号？<a @click="isLogin = true">返回登录</a></span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { login, register } from '@/api'
import { useUserStore } from '@/store/user'
import { ElMessage } from 'element-plus'

const router = useRouter()
const userStore = useUserStore()

const isLogin = ref(true)
const loading = ref(false)
const formRef = ref()

const form = reactive({
  username: '',
  password: '',
  password2: ''
})

const rules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  password2: [{ required: true, message: '请确认密码', trigger: 'blur' }]
}

async function handleSubmit() {
  try {
    await formRef.value.validate()
  } catch {
    return
  }

  if (!isLogin.value && form.password !== form.password2) {
    ElMessage.error('两次密码不一致')
    return
  }

  loading.value = true
  try {
    const data = isLogin.value
      ? await login({ username: form.username, password: form.password })
      : await register({ username: form.username, password: form.password })

    userStore.setToken(data.token)
    ElMessage.success(isLogin.value ? '登录成功' : '注册成功')
    router.push('/')
  } catch {}
  finally { loading.value = false }
}
</script>

<style lang="scss" scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-card {
  width: 400px;
  background: var(--bg-card);
  border-radius: var(--radius);
  padding: 32px;
  box-shadow: var(--shadow-md);

  h2 {
    text-align: center;
    margin-bottom: 24px;
    font-size: 24px;
    color: var(--text-primary);
  }
}

.switch-mode {
  text-align: center;
  margin-top: 16px;
  font-size: 14px;
  color: var(--text-muted);

  a {
    color: var(--primary);
    cursor: pointer;
  }
}
</style>