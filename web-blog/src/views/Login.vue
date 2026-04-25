<template>
  <div class="login-page">
    <div class="login-bg">
      <div class="bg-shape shape-1"></div>
      <div class="bg-shape shape-2"></div>
      <div class="bg-shape shape-3"></div>
    </div>

    <n-card class="login-card" :class="{ 'register-mode': !isLogin }">
      <div class="card-header">
        <h1 class="card-title">{{ siteName }}</h1>
        <p class="card-subtitle">{{ isLogin ? `欢迎来到${siteName}` : '创建账号' }}</p>
      </div>

      <n-form ref="formRef" :model="formValue" :rules="rules">
        <n-form-item label="用户名" path="username">
          <n-input
            v-model:value="formValue.username"
            placeholder="请输入用户名"
            round
            size="large"
          >
            <template #prefix>
              <n-icon><PersonOutline /></n-icon>
            </template>
          </n-input>
        </n-form-item>

        <n-form-item label="密码" path="password">
          <n-input
            v-model:value="formValue.password"
            type="password"
            placeholder="请输入密码"
            round
            size="large"
            show-password-on="click"
          >
            <template #prefix>
              <n-icon><LockClosedOutline /></n-icon>
            </template>
          </n-input>
        </n-form-item>

        <n-form-item v-if="!isLogin" label="确认密码" path="confirmPassword">
          <n-input
            v-model:value="formValue.confirmPassword"
            type="password"
            placeholder="请再次输入密码"
            round
            size="large"
            show-password-on="click"
          >
            <template #prefix>
              <n-icon><LockClosedOutline /></n-icon>
            </template>
          </n-input>
        </n-form-item>

        <template v-if="!isLogin">
          <n-form-item label="邮箱" path="email">
            <n-input
              v-model:value="formValue.email"
              placeholder="请输入邮箱"
              round
              size="large"
            >
              <template #prefix>
                <n-icon><MailOutline /></n-icon>
              </template>
            </n-input>
          </n-form-item>

          <n-form-item label="头像链接" path="avatar">
            <n-input
              v-model:value="formValue.avatar"
              placeholder="头像图片URL（可选）"
              round
              size="large"
            >
              <template #prefix>
                <n-icon><ImageOutline /></n-icon>
              </template>
            </n-input>
          </n-form-item>

          <n-form-item label="个人介绍" path="intro">
            <n-input
              v-model:value="formValue.intro"
              type="textarea"
              placeholder="介绍一下自己（可选，最多500字）"
              :maxlength="500"
              show-count
              :rows="3"
            />
          </n-form-item>

          <n-form-item label="个人网站" path="website">
            <n-input
              v-model:value="formValue.website"
              placeholder="个人网站URL（可选）"
              round
              size="large"
            >
              <template #prefix>
                <n-icon><GlobeOutline /></n-icon>
              </template>
            </n-input>
          </n-form-item>
        </template>
      </n-form>

      <n-button
        type="primary"
        block
        round
        size="large"
        :loading="loading"
        @click="handleSubmit"
        style="margin-top: 20px;"
      >
        {{ isLogin ? '登录' : '注册' }}
      </n-button>

      <div class="card-footer">
        <span>{{ isLogin ? '还没有账号？' : '已有账号？' }}</span>
        <n-button text type="primary" @click="toggleMode">
          {{ isLogin ? '立即注册' : '立即登录' }}
        </n-button>
      </div>
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import {
  NCard,
  NForm,
  NFormItem,
  NInput,
  NButton,
  NIcon,
  useMessage,
  type FormRules,
  type FormInst
} from 'naive-ui'
import { PersonOutline, LockClosedOutline, MailOutline, ImageOutline, GlobeOutline } from '@vicons/ionicons5'
import { login, register } from '@/api'
import { useUserStore } from '@/stores/user'
import { useSiteConfigStore } from '@/stores/siteConfig'

const router = useRouter()
const message = useMessage()
const userStore = useUserStore()
const siteConfig = useSiteConfigStore()
const siteName = computed(() => siteConfig.siteName)

onMounted(() => {
  siteConfig.fetchConfig()
})

const isLogin = ref(true)
const loading = ref(false)
const formRef = ref<FormInst | null>(null)

const formValue = reactive({
  username: '',
  password: '',
  confirmPassword: '',
  email: '',
  avatar: '',
  intro: '',
  website: ''
})

const rules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '用户名长度为3-20个字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码至少6个字符', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    {
      validator: (_rule, value) => {
        if (value !== formValue.password) {
          return new Error('两次密码不一致')
        }
        return true
      },
      trigger: 'blur'
    }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入有效的邮箱地址', trigger: 'blur' }
  ]
}

function toggleMode() {
  isLogin.value = !isLogin.value
  formValue.confirmPassword = ''
  formValue.email = ''
  formValue.avatar = ''
  formValue.intro = ''
  formValue.website = ''
}

async function handleSubmit() {
  try {
    await formRef.value?.validate()
  } catch {
    return
  }

  loading.value = true
  try {
    if (isLogin.value) {
      const res = await login({
        username: formValue.username,
        password: formValue.password
      }) as any
      userStore.setToken(res.data?.token || res.token)
      message.success('登录成功')
      router.push('/')
    } else {
      await register({
        username: formValue.username,
        password: formValue.password,
        email: formValue.email,
        avatar: formValue.avatar || undefined,
        intro: formValue.intro || undefined,
        website: formValue.website || undefined
      })
      message.success('注册成功，请登录')
      isLogin.value = true
    }
  } catch (e: any) {
    message.error(e.message || '操作失败')
  } finally {
    loading.value = false
  }
}
</script>

<style scoped lang="scss">
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
  position: relative;
  overflow: hidden;
}

.login-bg {
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, #0a0a0a 0%, #1a1a1a 100%);
  z-index: 0;
}

.bg-shape {
  position: absolute;
  border-radius: 50%;
  opacity: 0.1;

  &.shape-1 {
    width: 400px;
    height: 400px;
    background: white;
    top: -100px;
    right: -100px;
  }

  &.shape-2 {
    width: 300px;
    height: 300px;
    background: white;
    bottom: -50px;
    left: -50px;
  }

  &.shape-3 {
    width: 200px;
    height: 200px;
    background: white;
    top: 40%;
    left: 20%;
  }
}

.login-card {
  width: 100%;
  max-width: 400px;
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-lg);
  position: relative;
  z-index: 1;
  animation: fadeInUp 0.5s ease;

  &.register-mode {
    max-width: 480px;
  }

  :deep(.n-card__content) {
    padding: 32px;
  }
}

.card-header {
  text-align: center;
  margin-bottom: 30px;
}

.card-title {
  font-size: 28px;
  font-weight: 700;
  background: var(--accent-gradient);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 8px;
}

.card-subtitle {
  color: var(--text-muted);
  font-size: 15px;
}

.card-footer {
  text-align: center;
  margin-top: 24px;
  color: var(--text-muted);
  font-size: 14px;
}
</style>