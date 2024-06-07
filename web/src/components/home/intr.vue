<script setup>
const state = reactive({
  circleUrl:
      'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
  squareUrl:
      'https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png',

})

const {circleUrl, squareUrl, sizeList} = toRefs(state)
</script>

<template>
  <!-- 头像 -->
  <div>


    <div class="block">
      <el-avatar :size="100" :src="circleUrl"/>
    </div>


  </div>

  <h1>这是个介绍页面</h1>
</template>

<style scoped>


.block {
  margin-top: 30%;
  text-align: center;
}


</style>