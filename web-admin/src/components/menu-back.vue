<script setup>

import {RouterLink, useRoute} from "vue-router";

const route = useRoute()

const activeKey = computed(() => route.name)


const menuOptions = [
  {
    label: () => h(
        RouterLink,
        {
          to: {
            name: "home",
          }
        },
        {default: () => "首页看板"}
    ),
    key: "home",
    icon: () => h('i', {class: 'i-ep-histogram'}),
  },
  {
    label: "文章管理",
    key: "art",
    icon: () => h('i', {class: 'i-ep-notebook'}),
    children: [
      {
        label: () => h(
            RouterLink,
            {
              to: {
                name: "article",
              }
            },
            {default: () => "文章列表"}
        ),
        key: "article",
        icon: () => h('i', {class: 'i-ep-finished'}),
      },
      {
        label: () => h(
            RouterLink,
            {
              to: {
                name: "tag",
              }
            },
            {default: () => "标签列表"}
        ),
        key: "tag",
        icon: () => h('i', {class: 'i-ep-price-tag'}),
      },
      {
        label: () => h(
            RouterLink,
            {
              to: {
                name: "sort",
              }
            },
            {default: () => "分类列表"}
        ),
        key: "sort",
        icon: () => h('i', {class: 'i-ep-operation'}),
      }
    ]
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {
            name: "jpg",
          }
        },
        {default: () => "图片管理"}
    ),
    key: "jpg",
    icon: () => h('i', {class: 'i-ep-picture'}),

  },
  {
    label: () => h(
        RouterLink,
        {
          to: {
            name: "message",
          }
        },
        {default: () => "留言管理"}
    ),
    key: "message",
    icon: () => h('i', {class: 'i-ep-promotion'}),
  },
  {
    label: () => h(
        RouterLink,
        {
          to: {
            name: "ai",
          }
        },
        {default: () => "AI代理"}
    ),
    key: "ai",
    icon: () => h('i', {class: 'i-ep-help-filled'}),
  },
  {
    label: "系统管理",
    key: "system",
    icon: () => h('i', {class: 'i-ep-setting'}),
    children: [
      {
        label: "资源管理",
        key: "resource",
        icon: () => h('i', {class: 'i-ep-finished'}),
      },
      {
        label: () => h(
            RouterLink,
            {
              to: {
                name: "role",
              }
            },
            {default: () => "角色管理"}
        ),
        key: "role",
        icon: () => h('i', {class: 'i-ep-user-filled'}),
      },
      {
        label: () => h(
            RouterLink,
            {
              to: {
                name: "user",
              }
            },
            {default: () => "用户管理"}
        ),
        key: "user",
        icon: () => h('i', {class: 'i-ep-user'}),
      }
    ]
  }
]
</script>

<template>


  <n-menu

      :options="menuOptions"
      :value="activeKey"
      accordion
  />


</template>

<style scoped>

</style>