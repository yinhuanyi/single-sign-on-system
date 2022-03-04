<template>
  <div class="login-container">
    <el-form
      class="login-form"
      :model="loginForm"
      :rules="loginRules"
      ref="loginFormRef"
    >
      <!-- 标题 -->
      <div class="title-container">
        <h3 class="title">{{ $t('msg.login.title') }}</h3>
        <lang-select class="lang-select" effect="light" />
      </div>

      <!-- 用户名 -->
      <el-form-item prop="username">
        <span class="svg-container">
          <svg-icon icon="user" />
        </span>

        <el-input
          name="username"
          placeholder="username"
          type="text"
          v-model="loginForm.username"
        />
      </el-form-item>

      <!-- 密码 -->
      <el-form-item prop="password">
        <span class="svg-container">
          <!--   这里的svg-icon组件已经全局注册了   -->
          <svg-icon icon="password" />
        </span>
        <el-input
          name="password"
          placeholder="password"
          :type="passwordType"
          v-model="loginForm.password"
        />
        <span class="show-pwd svg-container" @click="onChangePwdType">
          <svg-icon
            :icon="passwordType === 'password' ? 'eye' : 'eye-open'"
          ></svg-icon>
        </span>
      </el-form-item>

      <!-- 登录按钮 -->
      <el-button
        type="primary"
        style="width: 100%; margin-bottom: 30px"
        :loading="loading"
        @click="handleLogin"
        >{{ $t('msg.login.loginBtn') }}</el-button
      >

      <!-- 添加一个文字说明  -->
      <div class="tips" v-html="$t('msg.login.desc')"></div>
    </el-form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { validatorPassword } from '@/views/login/rules'
import { useStore } from 'vuex'
import LangSelect from '@/components/LangSelect'
import { useI18n } from 'vue-i18n'

const store = useStore()
const i18n = useI18n()

const loginForm = ref({
  username: 'super-admin',
  password: '123456'
})

const loginRules = ref({
  username: [
    {
      required: true,
      trigger: 'change',
      message: i18n.t('msg.login.usernameRule')
    }
  ],
  // 给password指定了两个校验规则
  password: [
    {
      required: true,
      trigger: 'change',
      message: i18n.t('msg.login.passwordRule')
    },
    {
      validator: validatorPassword,
      trigger: 'change'
    }
  ]
})

const passwordType = ref('password')
const onChangePwdType = () => {
  // 使用ref声明的响应式数据，在script中使用时，需要加value来获取值，但是在视图模板template中不需要
  if (passwordType.value === 'password') {
    passwordType.value = 'text'
  } else {
    passwordType.value = 'password'
  }
}

const loading = ref(false)

const loginFormRef = ref(null)

const handleLogin = () => {
  // 1：进行表单校验
  loginFormRef.value.validate((valid) => {
    // 如果el-form的表单校验没有通过，直接return
    if (!valid) return
    // 2：触发登录动作
    loading.value = true
    // 由于vuex的action返回的是一个promise对象
    store
      .dispatch('user/login', loginForm.value)
      .then(() => {
        // 且关闭loading
        loading.value = false
      })
      .catch((error) => {
        console.log(error)
        // 如果请求过程有问题，那么关闭loading
        loading.value = false
      })
  })
}
</script>

<style scoped lang="scss">
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;
$cursor: #fff;

.login-container {
  //min-height: 100%;
  height: 100%;
  width: 100%;
  background-color: $bg;
  // 清除浮动、解决外边距塌陷
  overflow: hidden;

  .login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;

    // v-deep是深度选择器，由于el-form-item这个class是子组件的样式，要想选择到，需要使用深度选择器::v-deep，才能选择到子组件
    ::v-deep .el-form-item {
      border: 1px solid rgba(255, 255, 255, 0.1);
      background: rgba(0, 0, 0, 0.1);
      border-radius: 5px;
      color: #454545;
    }

    ::v-deep .el-input {
      display: inline-block;
      height: 47px;
      width: 85%;
      input {
        background: transparent;
        border: 0px;
        -webkit-appearance: none;
        border-radius: 0px;
        padding: 12px 5px 12px 15px;
        color: $light_gray;
        height: 47px;
        caret-color: $cursor;
      }
    }

    .tips {
      font-size: 16px;
      color: white;
      line-height: 24px;
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }

  .show-pwd {
    position: absolute;
    right: 10px;
    top: 7px;
    font-size: 16px;
    color: $dark_gray;
    cursor: pointer;
    user-select: none;
  }

  ::v-deep .lang-select {
    position: absolute;
    top: 4px;
    right: 0;
    background-color: #fff;
    font-size: 22px;
    padding: 4px;
    border-radius: 4px;
    cursor: pointer;
  }
}
</style>
