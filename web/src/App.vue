<script setup>
import {RouterView} from 'vue-router'

</script>

<template>
  <!--公共添加页面-->
  <RouterView/>
</template>

<style lang="css">
/*
公共样式
设置全屏 不滑动
*/
#app {
  width: 100%;
  height: 100%;
}

html {
  box-sizing: border-box;
  width: 100%;
  height: 100%;
  line-height: 1.5;
  tab-size: 4;
  text-size-adjust: 100%;
}

body {
  width: 100%;
  height: 100%;
  margin: 0;

}

</style>
