<script setup>

const fits = [
  'fill',
  'contain',
  'cover',
  'none',
  'scale-down',
]

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
          <el-main style="background: white;margin-top: 15px;min-width: 800px; ">

            <router-view></router-view>
          </el-main>

          <!-- 右侧页面内容 -->
          <el-aside style="margin-top: 15px;width: 25%;">
            <div style="margin-left: 10px; height: 100%">
              <el-card>
                <template #header>腾讯云推广</template>
                <a href="https://baidu.com">
                  <el-image fit="fill"
                            src="https://tg-image.com/file/3ed9dfaa8ae4d501e1f30.png"
                  />
                </a>
              </el-card>


            </div>

          </el-aside>

        </el-container>
      </el-container>

    </el-container>
  </div>


</template>

<style scoped>


/**/
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
  background-color: rgba(220, 218, 218, 0.4);
  /*  background-color: var(--el-bg-color-page);*/
  overflow-y: hidden; /*禁用滑轮y*/
  overflow-x: hidden; /*禁用滑轮x*/
}
</style>