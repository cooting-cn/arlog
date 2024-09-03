<script setup>
import {useMessage} from "naive-ui";
import blogStore from "@/stores/arlog.js";
import api from "@/api/api.js";
import index from "@/router/index.js";
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
        index.push("admin")
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
      <div class="mt-20 flex  justify-center">

        <!--QQ图标-->
        <svg class="icon" height="40"
             viewBox="0 0 1024 1024" width="60">
          <path
              d="M849.621 619.904a1364.352 1364.352 0 0 0-28.8-80.47l-38.826-95.829c0-1.109 0.512-19.968 0.512-29.696 0-163.882-78.208-328.576-270.507-328.576S241.493 250.027 241.493 413.867c0 9.77 0.47 28.629 0.512 29.738l-38.826 95.83a1398.485 1398.485 0 0 0-28.8 80.469c-36.694 116.779-24.79 165.12-15.744 166.187 19.413 2.304 75.562-87.894 75.562-87.894 0 52.224 27.179 120.406 86.016 169.643-21.973 6.699-48.938 17.024-66.304 29.653-15.573 11.392-13.61 23.04-10.794 27.734 12.33 20.522 211.413 13.098 268.928 6.698 57.472 6.4 256.597 13.824 268.885-6.741 2.816-4.693 4.779-16.299-10.795-27.69-17.365-12.63-44.33-22.955-66.346-29.697 58.837-49.194 86.016-117.376 86.016-169.642 0 0 56.149 90.24 75.562 87.893 9.046-1.067 20.907-49.365-15.786-166.144"
          ></path>
        </svg>

        <!--微信图标-->
        <svg class="icon" height="40" viewBox="0 0 1024 1024" width="60">
          <path
              d="M1010.8 628c0-141.2-141.3-256.2-299.9-256.2-168 0-300.3 115.1-300.3 256.2 0 141.4 132.3 256.2 300.3 256.2 35.2 0 70.7-8.9 106-17.7l96.8 53-26.6-88.2c70.9-53.2 123.7-123.7 123.7-203.3zM618 588.8c-22.1 0-40-17.9-40-40s17.9-40 40-40 40 17.9 40 40c0 22-17.9 40-40 40z m194.3-0.3c-22.1 0-40-17.9-40-40s17.9-40 40-40 40 17.9 40 40-17.9 40-40 40z"
              fill="#00C800"></path>
          <path
              d="M366.3 106.9c-194.1 0-353.1 132.3-353.1 300.3 0 97 52.9 176.6 141.3 238.4l-35.3 106.2 123.4-61.9c44.2 8.7 79.6 17.7 123.7 17.7 11.1 0 22.1-0.5 33-1.4-6.9-23.6-10.9-48.3-10.9-74 0-154.3 132.5-279.5 300.2-279.5 11.5 0 22.8 0.8 34 2.1C692 212.6 539.9 106.9 366.3 106.9zM247.7 349.2c-26.5 0-48-21.5-48-48s21.5-48 48-48 48 21.5 48 48-21.5 48-48 48z m246.6 0c-26.5 0-48-21.5-48-48s21.5-48 48-48 48 21.5 48 48-21.5 48-48 48z"
              fill="#00C800"></path>
        </svg>

        <!--gitee图标-->
        <a href="https://gitee.com/">
          <svg class="icon" height="40" viewBox="0 0 1024 1024" width="60">
            <path
                d="M512 938.666667C276.352 938.666667 85.333333 747.648 85.333333 512S276.352 85.333333 512 85.333333s426.666667 191.018667 426.666667 426.666667-191.018667 426.666667-426.666667 426.666667z m215.978667-474.069334h-242.304a21.077333 21.077333 0 0 0-21.077334 21.077334v52.693333c0 11.605333 9.386667 21.077333 21.034667 21.077333h147.498667c11.648 0 21.077333 9.386667 21.077333 21.034667v10.538667c0 34.901333-28.288 63.189333-63.189333 63.189333H390.826667a21.077333 21.077333 0 0 1-21.034667-21.077333v-200.106667c0-34.944 28.245333-63.232 63.146667-63.232h294.997333a21.077333 21.077333 0 0 0 21.034667-21.077333l0.085333-52.693334a21.077333 21.077333 0 0 0-21.077333-21.077333h-295.04a158.037333 158.037333 0 0 0-157.994667 158.037333v294.997334c0 11.648 9.429333 21.077333 21.077333 21.077333h310.784a142.208 142.208 0 0 0 142.208-142.208v-121.173333a21.077333 21.077333 0 0 0-21.077333-21.077334z"
                fill="#C71D23"></path>
          </svg>
        </a>

      </div>
    </div>
  </div>


</template>

<style scoped>

</style>