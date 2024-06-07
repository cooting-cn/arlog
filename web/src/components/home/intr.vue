<script setup>

/*个性化信息*/

const user = reactive({
  elUrl: 'https://tg-image.com/file/4693f7d40322f7b4c2e4b.jpg',
  sj: '人海未见之时，我亦独行在这城市， 料峭，春醒，酷暑，骤雨，寒意四起，大雁南飞后，大雪，寒风，斗转星移，人间寒暑， 如此之后， 才得蓦然回首，四目相对。',
  http: "https://arlog.cn"
})
</script>

<template>
  <!--介绍页面-->
  <div class="center-block">
    <el-card style="max-width: 480px">
      <template #header>
        <div class="card-header">
          <!-- 头像 -->
          <div class="block">
            <el-avatar :size="150" :src="user.elUrl"/>
            <!--跳转链接-->
            <div>
              <el-link :href="user.http" type="warning">
                {{ user.http }}
              </el-link>
            </div>
            <!--      诗句      -->
            <div style="margin-top: 5%">
              <el-text line-clamp=3>
                {{ user.sj }}
              </el-text>
            </div>
            <!--    数据        -->
            <el-row class="custom-row">
              <el-col :span="6">
                <el-statistic :value="268500" title="文章数"/>
              </el-col>
              <!--分割线              -->
              <el-divider direction="vertical" style="height: 55px"></el-divider>
              <el-col :span="6">
                <el-statistic :value="268500" title="评论数"/>
              </el-col>
            </el-row>

          </div>
        </div>
      </template>
      介绍自己
      <p v-for="o in 4" :key="o" class="text item">{{ 'List item ' + o }}</p>
      <template #footer>
        <h1>友链推荐</h1>
      </template>
    </el-card>
  </div>


</template>

<style scoped>
.custom-row {
  display: flex;
  justify-content: space-between;
  padding: 0 60px;
}

.custom-row .el-col {
  display: flex;
  justify-content: center;
}

.block {
  text-align: center; /*居中*/
}

.center-block {
  margin-left: auto; /*居中*/
  margin-right: auto; /*居中*/
  width: 95%; /*宽度*/
  /*  margin-top: 30%; !*高度*!*/
}

</style>