<script setup>

/*获取当前用户*/
/*声明本地持久化st*/
import blogStore from "@/stores/arlog.js";

const st = blogStore()
/*扇形图*/
import VChart from "vue-echarts"
import {use} from 'echarts/core'
import {PieChart} from 'echarts/charts'
import {TooltipComponent, LegendComponent} from 'echarts/components'
import {CanvasRenderer} from 'echarts/renderers'
import api from "@/api/api.js";


use([TooltipComponent, LegendComponent, PieChart, CanvasRenderer,])


/*定义空数据*/
let echartsData = ref([])

/*定义自动执行函数*/
function lg() {
  $loadingBar.start()
  /*获取gitee的参数*/
  api.getCode().then(res => {

    /*转换成想要的map*/
    echartsData.value = res.data.result.languages.map(lang => ({

      value: lang.percent,   // 将 percent 作为 value
      name: lang.language    // 将 language 作为 name

    }))
    /*关闭加载*/
    $loadingBar.finish()
  })


}

/*自动执行*/
onMounted(lg)


const option = ref({


  tooltip: {
    trigger: 'item',
    formatter: '{b}: {d}%'
  },
  legend: {
    type: 'scroll',
    orient: 'vertical',
    left: 10,
    top: 'center',
    bottom: 20,
  },
  series: [
    {
      name: 'Access From',
      type: 'pie',
      radius: ['30%', '85%'],
      avoidLabelOverlap: false,
      center: ['50%', '55%'],

      itemStyle: {
        borderRadius: 10,
        borderColor: '#fff',
        borderWidth: 2
      },

      label: {
        show: true,  // 启用标签
        formatter: '{b}: {d}%',  // 在标签中显示名称和百分比
        position: 'outside',  // 标签显示在扇形图外部
        fontSize: 15,

      },
      labelLayout: {
        hideOverlap: true,  // 避免重叠
      },
      emphasis: {
        label: {
          show: true,
          fontSize: 20,
          fontWeight: 'bold'
        }
      },
      data: echartsData
    }
  ]
})


</script>

<template>
  <div class="flex">
    <!-- 登录头像用户   -->
    <n-card class="min-w-200 w-30% h-150px">
      <div class="flex items-center">
        <n-avatar :size="60" :src=st.user_img round/>
        <div class="ml-20 flex-col">
            <span class="text-20 opacity-90">
              Hello,{{ st.user }}
            </span>

        </div>
      </div>

      <p class="mt-15 text-14 opacity-80">
        如果你累了,可以停下脚步歇息,但是不要眷恋松弛的状态。
      </p>

    </n-card>

    <!-- 右边欢迎  -->
    <n-card class="ml-12 w-70% h-150px" title="✨ 欢迎使用  arlog">

      <p class="opacity-80">
        这是一款极简风格博客的后台管理，包含前后端解决方案，前端使用 Vite + Vue3 + Pinia +
        Unocss，后端使用Golang + Gin + MySql 8，简单易用，赏心悦目，诚意满满！！
      </p>
      <footer class="mt-12 flex items-center justify-end">
        <n-button
            ghost
            href="https://arlog.cn"
            tag="a"
            target="__blank"
            type="primary"
        >
          开发文档
        </n-button>
        <n-button
            class="ml-12"
            href="https://arlog.cn"
            tag="a"
            target="__blank"
            type="primary"
        >
          代码仓库
        </n-button>
      </footer>
    </n-card>
  </div>

  <!-- 特性介绍模块 -->
  <div class="mt-12 flex h-300px">
    <n-card class="w-50%" segmented>

      <span class="text-20 ">
              💯 特性
            </span>
      <ul class="opacity-90 mt-15">
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
          🎨 使用 Naive UI，
          <span class="text-highlight">极致简洁的代码风格和清爽的页面设计</span>
          ，审美在线，主题轻松定制
        </li>

        <li class="py-4">
          🚀
          <span class="text-highlight">扁平化路由</span>
          设计，每一个组件都可以是一个页面
        </li>
        <li class="py-4">
          ✨ 基于 Naive UI 封装
          <span class="text-highlight">message</span>
          全局工具方法，支持批量提醒，支持跨页面共享实例
        </li>

      </ul>

      <n-divider class="mb-0! mt-12!">
        <p class="text-14 opacity-60">
          👉点击
          <b class="mx-2 transition hover:text-primary">
            <a href="https://arlog.cn" target="_blank">更多</a>
          </b>
          查看更多实用功能，持续开发中...
        </p>
      </n-divider>
    </n-card>


    <n-card class="ml-12 w-50%">

      <span class="text-20  ">
               🛠️ 技术栈
            </span>
      <v-chart :option="option" class="mt--20px h-100% w-100%"/>

    </n-card>
  </div>

  <n-card class="mt-12 h-400px">
      <span class="text-20  ">
               ⚡️ 趋势
            </span>


  </n-card>

</template>

<style scoped>

</style>