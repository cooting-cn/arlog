<script setup>
import {useMessage} from "naive-ui";
import blogStore from "@/stores/arlog.js";
import api from "@/api/api.js";
import router from "@/router/index.js";
/*创建消息提示对象*/
const msg = useMessage()
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
  msg.loading("正在验证", {duration: 50e3})

  api.postLogin(logInfo).then(res => {
    // 无论成功还是失败，都会执行
    msg.destroyAll()
    switch (res.data.code) {
        /*登录成功*/
      case 200:
        /*存储token到本地*/
        st.token = res.data.data.token
        /*存储登录的用户名*/
        st.user = res.data.data.user.username
        // 显示成功消息
        msg.success(res.data.data.user.username + "登录成功")
        /*跳转到后台*/
        router.push("admin")
        break
        /*登录成功*/
      case 203:
        // 显示消息
        msg.error(res.data.msg)
        break
      case 204:
        // 显示消息
        msg.error(res.data.msg)
        break
    }
  })
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


</template>

<style scoped>

</style>