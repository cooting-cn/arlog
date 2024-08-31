<script setup>
</script>

<template>
  <!--  <n-layout class="h-80%" has-sider>
      <n-layout-sider>
        菜单
      </n-layout-sider>

      <n-layout>
        <n-layout-header class="h-5% ">导航</n-layout-header>

        <n-layout-content class="h-50%" content-style="padding: 24px;">
          内容
          &lt;!&ndash;加载路由&ndash;&gt;
          <RouterView/>

        </n-layout-content>

        <n-layout-footer class="h-5%">底下导航</n-layout-footer>

      </n-layout>

    </n-layout>-->

  <div class="h-full flex">
    <aside
        border-r="1px solid light_border dark:dark_border"
        class="flex-col flex-shrink-0 transition-width-300 w-220 "
    >
      <h1>菜单</h1>
    </aside>

    <article class="w-0 flex-col flex-1">
      <div class="h-5% min-h-5%">
        <h1>导航</h1>
      </div>

      <div class="h-95% min-h-95%">
        <h1>
          内容
          <router-view></router-view>
        </h1>
      </div>


    </article>
  </div>

</template>

<style scoped>


.n-layout-content {
  background-color: #f2f2f2;

}
</style>