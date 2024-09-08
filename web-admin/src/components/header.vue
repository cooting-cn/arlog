<script setup>

import {useRoute} from "vue-router";
import {useFullscreen} from '@vueuse/core'
import blogStore from "@/stores/arlog.js";
import router from "@/router/index.js";


/*全屏 功能*/
const {isFullscreen, toggle} = useFullscreen()

/*面包屑 获取当前路由*/
const route = useRoute()
const activeKey = computed(() => route.name)

/*打开链接*/
function handleLinkClick(link) {
  window.open(link)
}

/*获取当前用户*/
/*声明本地持久化st*/
const st = blogStore()

/*右上角头像菜单*/
const options = [
  {
    label: "用户资料",
    key: "profile",
    icon: () => h('i', {class: 'i-ep-user-filled'})
  },
  {
    label: "编辑资料",
    key: "editProfile",
    icon: () => h('i', {class: 'i-ep-edit'})
  },
  {
    label: "退出登录",
    key: "logout",
    icon: () => h('i', {class: 'i-ep-switch-button'})
  }
]

/*编辑资料弹出变量*/
const showModal = ref(false)

/*菜单点击选项功能*/
function handleSelect(key) {
  switch (key) {
    case 'profile':
      router.push('/admin/profile')
      break
    case 'editProfile':
      showModal.value = true
      break
    case 'logout':
      $dialog.warning({
        title: "用户注销！",
        content: "你确定退出？",
        positiveText: "确定",
        negativeText: "取消",
        onPositiveClick: () => {

          st.token = ""
          $message.warning(st.user + "已经注销")
          router.push('/login')

        },
        onNegativeClick: () => {
          $message.success("取消")
        }
      })

      break
  }
}


const formValue = ref({
  user: {
    name: "",
    age: ""
  },
  phone: ""
})


</script>

<template>
  <!--图标-->
  <div class="flex items-center  min-w-1600px">
    <!-- 左侧图标 -->
    <div class="flex justify-center items-center h-60px w-50px">
      <n-icon class="i-ep-expand " color="rgba(87, 89, 87, 0.9)" size="25"/>
    </div>
    <!--面包屑-->
    <div>
      <n-tag :bordered="false" type="warning">
        {{ activeKey }}
      </n-tag>
    </div>


    <div class="flex mr-16 cursor-pointer  gap-16 ml-auto  w-200px" style="color: rgba(87, 89, 87, 0.9);">
      <!--  右边功能区域  -->
      <span class="mx-6  ">|</span>
      <!--    全屏功能  -->
      <n-popover trigger="hover">
        <template #trigger>
          <n-icon
              :class="isFullscreen ? 'i-ep-close-bold' : 'i-ep-full-screen'"
              size="25"
              @click="toggle"
          />
        </template>
        全屏
      </n-popover>


      <!--     跳转gitee -->
      <n-popover trigger="hover">
        <template #trigger>
          <n-icon class="i-simple-icons-gitee " size="25"
                  @click="handleLinkClick('https://gitee.com/cooting/arlog')"/>
        </template>
        跳转gitee
      </n-popover>

      <n-popover trigger="hover">
        <template #trigger>
          <n-icon class="i-simple-icons-github " size="25"
                  @click="handleLinkClick('https://gitee.com/cooting/arlog')"/>
        </template>
        跳转github
      </n-popover>

      <n-popover trigger="hover">
        <template #trigger>
          <n-icon class="i-simple-icons-tencentqq " size="25"
                  @click="handleLinkClick('https://gitee.com/cooting/arlog')"/>
        </template>
        联系QQ
      </n-popover>


      <n-popover trigger="hover">
        <template #trigger>
          <n-icon class="i-simple-icons-wechat" size="25"
                  @click="handleLinkClick('https://gitee.com/cooting/arlog')"/>
        </template>
        联系微信
      </n-popover>


    </div>

    <!--个人资料区域   -->
    <div class=" mr-16">

      <n-dropdown :options="options" @select="handleSelect">
        <div class="flex cursor-pointer items-center">
          <n-avatar :size="48" round src="https://tg-image.com/file/b13dfda09561bf4318933.jpg"/>
          <div class="ml-12 flex flex-col ">

            <span class="text-20 opacity-70 text-center h-25px">{{ st.user }}</span>

            <span class="text-12 opacity-50 self-center text-center ">[超级管理员]</span>

          </div>
        </div>
      </n-dropdown>
    </div>


  </div>

  <!--动态弹出框-->
  <n-modal v-model:show="showModal">
    <n-card
        :bordered="false"
        aria-modal="true"
        role="dialog"
        size="huge"
        style="width: 500px"
        title="修改资料"
    >
      <n-form ref="formRef" :model="formValue" label-placement="left">
        <n-form-item
            label="姓名"
            label-align="left"
            path="user.name"
        >
          <n-input v-model:value="formValue.user.name" placeholder="输入姓名"/>
        </n-form-item>

        <n-form-item label="年龄" path="user.age">
          <n-input v-model:value="formValue.user.age" placeholder="输入年龄"/>
        </n-form-item>

        <n-form-item label="电话号码" path="user.phone">
          <n-input v-model:value="formValue.phone" placeholder="电话号码"/>
        </n-form-item>

      </n-form>

      内容
      <template #footer>
        尾部
      </template>
    </n-card>
  </n-modal>


</template>
<style scoped>

</style>
