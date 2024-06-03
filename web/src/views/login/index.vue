<script setup>

</script>

<template>
  <!--整个页面布局-->
  <div class="container">
    <!--登录框    -->
    <el-card class="login">

      <!--    显示头文字  -->
      <div style="text-align: center;">
        <div class="flex-item">
          <!--      版本标签    -->
          <el-tag type="primary">latest</el-tag>
        </div>
        <h1 style="margin-top: 10px">arlog登录</h1>
      </div>

      <!--    表单用户 密码  -->
      <!--   用户   -->
      <div class="sr">
        <el-input v-model="input" class="custom-input" placeholder="Please input"/>
      </div>
      <!--    密码  -->
      <div class="sr">
        <el-input v-model="input" class="custom-input" placeholder="Please input" size="large"/>
      </div>

      <!--  三方登录    -->
      <template #footer>三方登录</template>
    </el-card>
  </div>
</template>

<style scoped>

.custom-input {

  height: 45px; /* 自定义高度 */
  font-size: 16px; /* 自定义字体大小 */
}

.sr {
  margin-top: 20px;

}

.flex-item {
  position: absolute;
  top: 15px;
  right: 15%;
  padding: 10px;
}

/*登录框居中*/
.login {
  margin: 0 auto; /* 水平居中 */
  position: relative; /* 相对定位 */
  top: 20%;
  max-width: 400px;

  border-radius: 20px; /*弧度角度*/
}

/*容器布局高度100%*/
.container {
  height: 100%;
}
</style>