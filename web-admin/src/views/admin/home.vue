<script setup>
/*获取当前用户*/
/*声明本地持久化st*/
import blogStore from "@/stores/arlog.js";

import {use} from "echarts/core";
import {CanvasRenderer} from "echarts/renderers";
import {PieChart} from "echarts/charts";
import {LegendComponent, TitleComponent, TooltipComponent} from "echarts/components";
import VChart, {THEME_KEY} from "vue-echarts";
import {provide, ref} from "vue";

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent
]);

provide(THEME_KEY, "dark");

const option = ref({
  title: {
    text: "Traffic Sources",
    left: "center"
  },
  tooltip: {
    trigger: "item",
    formatter: "{a} <br/>{b} : {c} ({d}%)"
  },
  legend: {
    orient: "vertical",
    left: "left",
    data: ["Direct", "Email", "Ad Networks", "Video Ads", "Search Engines"]
  },
  series: [
    {
      name: "Traffic Sources",
      type: "pie",
      radius: "55%",
      center: ["50%", "60%"],
      data: [
        {value: 335, name: "Direct"},
        {value: 310, name: "Email"},
        {value: 234, name: "Ad Networks"},
        {value: 135, name: "Video Ads"},
        {value: 1548, name: "Search Engines"}
      ],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: "rgba(0, 0, 0, 0.5)"
        }
      }
    }
  ]
});
/*---------------------------------*/

const st = blogStore()
</script>

<template>
  <div class="flex">
    <n-card class="min-w-200 w-30%">
      <div class="flex items-center">
        <n-avatar :size="60" class="flex-shrink-0" round src="https://tg-image.com/file/b13dfda09561bf4318933.jpg"/>
        <div class="ml-20 flex-col">
            <span class="text-20 opacity-80">
              Hello,{{ st.user }}
            </span>

        </div>
      </div>

      <p class="mt-28 text-14 opacity-60">
        一个人几乎可以在任何他怀有无限热忱的事情上成功。
      </p>
      <p class="mt-12 text-right text-12 opacity-40">
        —— 查尔斯·史考伯
      </p>
    </n-card>
    <n-card class="ml-12 w-70%" title="✨ 欢迎使用 Vue Naive Admin 2.0">
      <template #header-extra>
        <a
            class="text-14 text-primary text-highlight hover:underline hover:opacity-80"
            href="https://isme.top"
            target="_blank"
            @click.prevent="message?.info('官网正在火速开发中...')"
        >
          isme.top
        </a>
      </template>

      <p class="opacity-60">
        这是一款极简风格的后台管理模板，包含前后端解决方案，前端使用 Vite + Vue3 + Pinia +
        Unocss，后端使用 Nestjs + TypeOrm +
        MySql，简单易用，赏心悦目，历经十几次重构和细节打磨，诚意满满！！
      </p>
      <footer class="mt-12 flex items-center justify-end">
        <n-button
            ghost
            href="https://isme.top"
            tag="a"
            target="__blank"
            type="primary"
        >
          开发文档
        </n-button>
        <n-button
            class="ml-12"
            href="https://github.com/zclzone/vue-naive-admin/tree/2.x"
            tag="a"
            target="__blank"
            type="primary"
        >
          代码仓库
        </n-button>
      </footer>
    </n-card>
  </div>
  <div class="mt-12 flex">
    <n-card class="w-50%" segmented title="💯 特性">
      <template #header-extra>
        <span class="opacity-90 text-highlight">👏 历经十几次重构和细节打磨</span>
      </template>

      <ul class="opacity-90">
        <li class="py-4">
          🆒 使用
          <b>Vue3</b>
          主流技术栈:
          <span class="text-highlight">Vite + Vue3 + Pinia</span>
        </li>
        <li class="py-4">
          🍇 使用
          <b>原子CSS</b>
          框架:
          <span class="text-highlight">Unocss</span>
          ，优雅、轻量、易用
        </li>
        <li class="py-4">
          🤹 使用主流的
          <span class="text-highlight">iconify + unocss</span>
          图标方案，支持自定义图标，支持动态渲染
        </li>
        <li class="py-4">
          🎨 使用 Naive UI，
          <span class="text-highlight">极致简洁的代码风格和清爽的页面设计</span>
          ，审美在线，主题轻松定制
        </li>
        <li class="py-4">
          👏 先进且易于理解的文件结构设计，多个模块之间
          <b>零耦合</b>
          ，单个业务模块删除不影响其他模块
        </li>
        <li class="py-4">
          🚀
          <span class="text-highlight">扁平化路由</span>
          设计，每一个组件都可以是一个页面，告别多级路由 KeepAlive 难实现问题
        </li>

        <li class="py-4">
          🍒
          <span class="text-highlight">基于权限动态生成路由</span>
          ，无需额外定义路由，
          <span class="text-highlight">403和404可区分</span>
          ，而不是无权限也跳404
        </li>
        <li class="py-4">
          🔐 基于Redis集成
          <span class="text-highlight">无感刷新</span>
          ，用户登录态可控，安全与体验缺一不可
        </li>
        <li class="py-4">
          ✨ 基于 Naive UI 封装
          <span class="text-highlight">message</span>
          全局工具方法，支持批量提醒，支持跨页面共享实例
        </li>
        <li class="py-4">
          ⚡️ 基于 Naive UI 封装常用的业务组件，包含
          <span class="text-highlight">Page</span>
          组件、
          <span class="text-highlight">CRUD</span>
          表格组件及
          <span class="text-highlight">Modal</span>
          组件，减少大量重复性工作
        </li>
      </ul>

      <n-divider class="mb-0! mt-12!">
        <p class="text-14 opacity-60">
          👉点击
          <b class="mx-2 transition hover:text-primary">
            <a href="https://isme.top" target="_blank">更多</a>
          </b>
          查看更多实用功能，持续开发中...
        </p>
      </n-divider>
    </n-card>

    <n-card class="ml-12 w-50%" segmented title="🛠️ 技术栈">

      <v-chart :option="option" class="chart"/>
    </n-card>
  </div>

  <n-card class="mt-12" segmented title="⚡️ 趋势">
    <div class="h-400">

    </div>
  </n-card>

</template>

<style scoped>

</style>