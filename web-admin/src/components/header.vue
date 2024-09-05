<script setup>

import {useRoute} from "vue-router";
import {useFullscreen} from '@vueuse/core'

const {isFullscreen, toggle} = useFullscreen()
const route = useRoute()
const activeKey = computed(() => route.name)

function handleLinkClick(link) {
  window.open(link)
}

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
                  @click="handleLinkClick('https://github.com/zclzone/vue-naive-admin/tree/2.x')"/>
        </template>
        跳转gitee
      </n-popover>

      <n-popover trigger="hover">
        <template #trigger>
          <n-icon class="i-simple-icons-github " size="25"
                  @click="handleLinkClick('https://github.com/zclzone/vue-naive-admin/tree/2.x')"/>
        </template>
        跳转github
      </n-popover>

      <n-popover trigger="hover">
        <template #trigger>
          <n-icon class="i-simple-icons-tencentqq " size="25"
                  @click="handleLinkClick('https://github.com/zclzone/vue-naive-admin/tree/2.x')"/>
        </template>
        联系QQ
      </n-popover>


      <n-popover trigger="hover">
        <template #trigger>
          <n-icon class="i-simple-icons-wechat" size="25"
                  @click="handleLinkClick('https://github.com/zclzone/vue-naive-admin/tree/2.x')"/>
        </template>
        联系微信
      </n-popover>


    </div>
    <!--个人资料区域   -->
    <div class=" mr-16">
      <n-dropdown :options="options">
        <n-tag :bordered="false">
          用户资料
        </n-tag>
      </n-dropdown>
    </div>


  </div>

</template>
<style scoped>

</style>
