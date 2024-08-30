<script setup></script>

<template>
  <n-layout class="h-80%" has-sider>
    <n-layout-sider>
      菜单
    </n-layout-sider>

    <n-layout>
      <n-layout-header class="h-5% ">导航</n-layout-header>

      <n-layout-content class="h-50%" content-style="padding: 24px;">
        内容
        <!--加载路由-->
        <RouterView/>

      </n-layout-content>

      <n-layout-footer class="h-5%">底下导航</n-layout-footer>

    </n-layout>

  </n-layout>

</template>

<style scoped>


.n-layout-content {
  background-color: #f2f2f2;

}
</style>