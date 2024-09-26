<script setup>

import Logo from "@/components/logo.vue";
import Menu from "@/components/menu.vue";
import Header from "@/components/header.vue";
</script>

<template>
  <div class="h-full flex ">
    <aside
        class="flex-col flex-shrink-0  w-220px bg-[#ffffffff] border-b-0 border-solid border-[rgba(209,213,219,0.61)] border-t-0 border-l-0 border-r">
      <!--    头像log  只添加了下边框-->
      <div class="h-60px  border-b border-solid border-[rgba(209,213,219,0.61)] border-t-0 border-l-0 border-r-0">
        <logo/>
      </div>

      <!--    菜单栏-->
      <div>
        <Menu/>
      </div>

    </aside>

    <article class="w-0 flex-col flex-1 ">
      <!--导航栏-->
      <div
          class="h-60px border-b border-solid border-[rgba(209,213,219,0.61)] border-t-0 border-l-0 border-r-0 bg-[#ffffffff] ">
        <Header/>
      </div>

      <!--实际网页内容-->
      <div class=" w-100% h-100% p-13">

        <router-view></router-view>


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