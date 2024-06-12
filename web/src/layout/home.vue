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
          <el-main style="background: white;margin-top: 15px;overflow-y: hidden;min-width: 800px ">

            <router-view></router-view>
          </el-main>

          <!-- 右侧页面内容 -->
          <el-aside style="margin-top: 15px;width: 25%;overflow-y: hidden;overflow-x: hidden;">
            <div style="margin-left: 10px; height: 100%">
              <div class="demo-image">
                <span class="demonstration">腾讯云推广</span>
                <a href="https://baidu.com">
                  <el-image fit="fill"
                            src="https://tg-image.com/file/3ed9dfaa8ae4d501e1f30.png"
                  />
                </a>
              </div>
              <div class="demo-image" style="height: 400px">
                <span class="demonstration">腾讯云推广</span>
                <a href="https://baidu.com">
                  <el-image fit="fill"
                            src="https://tg-image.com/file/342fd8c1f8e473cc5580a.png"
                  />
                </a>
              </div>

            </div>

          </el-aside>

        </el-container>
      </el-container>

    </el-container>
  </div>


</template>

<style scoped>
.demo-image .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  display: inline-block;

  box-sizing: border-box;
  vertical-align: top;
}

.demo-image .block:last-child {
  border-right: none;
}

.demo-image .demonstration {
  display: block;
  text-align: center;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  margin-bottom: 20px;
}

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
  background: rgba(144, 141, 141, 0.18);
  /*  background-color: var(--el-bg-color-page);*/
}
</style>