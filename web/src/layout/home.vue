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
        <el-header>
          <Header/>
        </el-header>

        <el-main>Main
          <router-view></router-view>
        </el-main>

      </el-container>

    </el-container>
  </div>


</template>

<style scoped>


.el-container {
  height: 100%;
}


.common-layout {
  height: 100%;
}
</style>