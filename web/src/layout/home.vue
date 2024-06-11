<script setup>


import Intr from "@/components/home/intr.vue";
import Header from "@/components/home/header.vue";
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="400px">
        <!--    这是个介绍页    -->
        <Intr/>
      </el-aside>

      <el-container>
        <!--这是个导航头-->
        <el-header class="header-container">

          <Header/>


        </el-header>
        <el-container>
          <!--内容页面-->
          <el-main style="background: white;margin-top: 15px;width: 100%;">

            <router-view></router-view>
          </el-main>

          <el-aside style="background: #c64444;margin-top: 15px;" width="400px">
            <!-- 右侧页面内容 -->
            <h1>6666</h1>
          </el-aside>

        </el-container>
      </el-container>

    </el-container>
  </div>


</template>

<style scoped>
.header-container {
  position: relative; /* 或者 absolute, fixed */
  top: 2px; /* 将容器向下移动20像素 */
  background: #fff;
}

.el-container {
  height: 100%;
}


.common-layout {
  height: 100%;
  /*background: rgba(240, 234, 241, 0.51);*/
  background-color: var(--el-bg-color-page);
}
</style>