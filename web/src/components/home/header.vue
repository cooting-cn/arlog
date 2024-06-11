<script setup>

import {ArrowDown, Search} from "@element-plus/icons-vue";

const input3 = ref('')

function ss() {
  console.log("搜索");
}
</script>

<template>
  <div class="hed">
    <div style="width: 60px;align-items: center;  display: flex;">
      <!--头像-->
      <el-avatar src="https://tg-image.com/file/8619addedd88492fc2da1.jpg"/>
      <!--首页    -->
    </div>


    <el-menu active-text-color="black"
             mode="horizontal"
             style="width: 600px"
    >
      <el-menu-item index="1">
        <span class="el-span"> 首页 </span>
      </el-menu-item>
      <el-menu-item index="2">
        <span class="el-span"> 分类 </span>
      </el-menu-item>
      <el-menu-item index="3">
        <span class="el-span"> github </span>
      </el-menu-item>
      <el-menu-item index="4">
        <span class="el-span"> gitee </span>
      </el-menu-item>

    </el-menu>
    <!--   文章搜索框 -->

    <el-input v-model="input3"

              placeholder="文章搜索"
              style="max-width: 300px;"
    >
      <!--搜索框的图标-->
      <template #append>
        <el-button :icon="Search"/>
      </template>
    </el-input>

    <!--右上角用户配置-->


    <el-dropdown class="el-dr" trigger="click">
      <!--  右上角用户  -->
      <el-check-tag checked type="warning">
        admin
      </el-check-tag>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>个人信息</el-dropdown-item>
          <el-dropdown-item>退出</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>

  </div>


</template>

<style scoped>


.el-dr {
  margin-left: 400px; /* 向右移动 100px */

}

.el-input {
  /*默认边框颜色*/
  --el-input-border-color: rgba(88, 87, 87, 0.51);
  /*背景颜色*/

  /*获取焦点后的边框颜色*/
  --el-input-focus-border-color: rgba(88, 87, 87, 0.51);
  /*鼠标悬停边框颜色*/
  --el-input-hover-border-color: rgba(88, 87, 87, 0.51);
}

/*搜索input框 */
:deep(.el-input__wrapper) {
  border-radius: 15px 0 0 15px; /* 左边圆角 */


}

/*搜索button按钮 */
:deep(.el-input-group__append) {
  border-radius: 0 15px 15px 0;;
}

/* 自定义鼠标悬停时的背景色 */
.el-menu--horizontal .el-menu-item:hover,
.el-menu--horizontal .el-menu-item.is-active {
  background-color: rgba(233, 231, 231, 0.52); /* 自定义背景色 */
}


/* 自定义鼠标悬停时的文字颜色 */
.el-menu--horizontal .el-menu-item:hover {
  color: #000000; /* 自定义悬停文字颜色 */

}

/*去掉选中下划线*/
.el-menu--horizontal > .el-menu-item {
  border-bottom: none;
  text-decoration: none;
}

/* 去掉 el-menu 的边框颜色 */
.el-menu--horizontal {
  border-bottom: none; /* 去掉底部边框 */
}


.el-span {
  font-size: 20px;
  font-weight: 700;
  color: rgba(52, 49, 49, 0.8);

}


.hed {
  display: flex; /*布局*/
  align-items: center; /* 垂直居中 */
  width: 100%;
  height: 100%;

}
</style>