<script setup>
import blogStore from "@/stores/arlog.js";
import api from "@/api/api.js";
import router from "@/router/index.js";
import {NButton} from "naive-ui";

/*声明本地持久化st*/
const st = blogStore()
/*创建登录用户变量*/
const logInfo = reactive({
  username: "admin",
  password: "admin"
})

/*登录按钮*/
function login() {
  /*打印当前请求域名*/
  //console.log(window.location.origin)

  // 显示加载中的消息
  $message.loading("正在验证", {
    key: 'loadingMessage',  // 唯一的 key
    duration: 50e3          // 消息将持续 50 秒
  })
  api.postLogin(logInfo).then(res => {

    switch (res.data.code) {
        /*登录成功*/
      case 200:
        /*存储token到本地*/
        st.token = res.data.result.token
        /*存储登录的用户名*/
        st.user = res.data.result.user.username
        /*获取头像*/
        st.user_img = res.data.result.user["a-img"]
        // 显示成功消息
        $message.success(st.user + "登录成功")
        /*跳转到后台*/
        router.push("admin")
        break
        /*登录成功*/

    }
  })
}

/*绑定otp*/


const showModalOtp = ref(false)
const bindOtp = ref({
  username: "",
  coding: ""
})

/*登录otp*/
function showOtp() {
  showModalOtp.value = true
  bindOtp.value.username = logInfo.username
}


function loginOtp() {

  api.otp(bindOtp.value).then(
      res => {
        switch (res.data.code) {
            /*登录成功*/
          case 200:
            /*存储token到本地*/
            st.token = res.data.result.token
            /*存储登录的用户名*/
            st.user = res.data.result.user.username
            /*获取头像*/
            st.user_img = res.data.result.user["a-img"]
            // 显示成功消息
            $message.success(st.user + "登录成功")
            /*跳转到后台*/
            router.push("admin")
            break
            /*登录成功*/

        }
      }
  )
  //关闭弹窗
  showModalOtp.value = false
}
</script>

<template>


  <div
      class=" flex justify-center items-center min-h-screen bg-no-repeat bg-bottom bg-contain bg-[url(@/assets/images/login-background.png)]">

    <!--登录框-->
    <div class=" w-400px h-350px p-20 shadow-md bg-[#FFFFFFFF]" style="border-radius: 12px;">

      <h2 class="flex items-center justify-center text-center text-30 text-[#767474ff] font-normal">
        <img alt="" class="mr-3 h-55" src="@/assets/images/arlog.webp">
        arlog
      </h2>

      <form>
        <!--用户输入框-->
        <n-input
            v-model:value="logInfo.username"
            :input-props="{ autocomplete: 'username' ,id: 'username'}"
            :maxlength="8"
            class="mt-32 h-40 items-center"
            placeholder="用户"
            round
            size="large"
        >

          <template #prefix>
            <i class="i-ep-user mr-12 opacity-60"/>
          </template>
        </n-input>

        <!--密码输入框-->
        <n-input
            v-model:value="logInfo.password"
            :input-props="{ autocomplete: 'new-password',id: 'password' }"
            :maxlength="20"
            class="mt-20 h-40 items-center"
            placeholder="密码"
            round
            show-password-on="mousedown"
            size="large"
            type="password"
        >
          <template #prefix>
            <i class="i-ep-lock mr-12 opacity-60"/>
          </template>
        </n-input>
      </form>
      <!--登录按钮区域-->
      <div class="mt-30 flex items-center">
        <n-button
            class="h-40 flex-1 rounded-5 text-16"
            ghost
            type="warning"
            @click="showOtp"
        >
          otp一键登录
        </n-button>

        <n-button
            class="ml-32 h-40 flex-1 rounded-5 text-16"
            type="info"
            @click="login"
        >
          登录
        </n-button>
      </div>

      <!--三方图标登录-->
      <div class="mt-20 flex  justify-center gap-25 opacity-70">
        <n-icon class="i-simple-icons-gitee " size="35"/>
        <n-icon class="i-simple-icons-tencentqq " size="35"/>
        <n-icon class="i-simple-icons-wechat " size="35"/>
      </div>
    </div>
  </div>

  <!--动态弹出框 otp-->
  <n-modal v-model:show="showModalOtp">
    <n-card
        :bordered="false"
        aria-modal="true"
        class="w-15% min-w-200 "
        role="dialog"
        size="huge"
        title="登录top"
    >
      当前用户{{ bindOtp.username }}

      <div class="text-center">
        <!--动态码-->
        <n-form-item>
          <n-input v-model:value="bindOtp.coding" maxlength="6" placeholder="谷歌动态码"/>
        </n-form-item>

        <n-button type="info" @click="loginOtp">
          确定登录
        </n-button>
      </div>


    </n-card>
  </n-modal>
</template>

<style scoped>

</style>