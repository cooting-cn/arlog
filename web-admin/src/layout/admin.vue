<script setup>

import Logo from "@/components/logo.vue";
import Menu from "@/components/menu.vue";
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

  <div class="h-full flex ">
    <aside
        class="flex-col flex-shrink-0  w-220px bg-[#ffffffff] border-b-0 border-solid border-[rgba(209,213,219,0.61)] border-t-0 border-l-0 border-r">
      <!--    头像log  只添加了下边框-->
      <div class="h-60px  border-b border-solid border-[rgba(209,213,219,0.61)] border-t-0 border-l-0 border-r-0">
        <logo/>
      </div>


      <div>
        <Menu/>
      </div>

    </aside>

    <article class="w-0 flex-col flex-1 ">

      <div
          class="h-60px border-b border-solid border-[rgba(209,213,219,0.61)] border-t-0 border-l-0 border-r-0 bg-[#ffffffff] ">
        <h1>导航</h1>
      </div>

      <div class="h-95% min-h-95% ">
        <h1>
          内容
          <router-view></router-view>
        </h1>
      </div>


    </article>
  </div>

</template>

<style scoped>

/*暂时不应用的边框css*/
.button-class {
  /*下边框*/
  border-bottom: 1px solid rgba(209, 213, 219, 0.61);
  /*右边框*/
  border-right: 1px solid rgba(209, 213, 219, 0.61);
}

.n-layout-content {
  background-color: #f2f2f2;

}
</style>