<script setup>
import {RouterView} from 'vue-router'
import {dateZhCN, zhCN} from 'naive-ui'
/*调整主题颜色*/
const naiveThemeOverrides = {
  common: {
    primaryColor: "#7a7f84",
    primaryColorHover: "#82909D",
    primaryColorPressed: "#484d77",
    primaryColorSuppl: "#82909D",
  },

}


</script>


<template>
  <!--设置全局属性主题配置  -->
  <n-config-provider
      :date-locale="dateZhCN"
      :locale="zhCN"
      :theme-overrides="naiveThemeOverrides"
  >

    <router-view/>
  </n-config-provider>
</template>


<style>

</style>
