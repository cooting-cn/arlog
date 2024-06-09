<script setup>

/*个性化信息*/

import TypeIt from "typeit";

const user = reactive({
  elUrl: 'https://tg-image.com/file/4693f7d40322f7b4c2e4b.jpg',
  sj: '人海未见之时，我亦独行在这城市， 料峭，春醒，酷暑，骤雨，寒意四起，大雁南飞后，大雪，寒风，斗转星移，人间寒暑， 如此之后， 才得蓦然回首，四目相对。',
  http: "https://arlog.cn",
  text: "一个做了十年的运维故事,2018年自学k8s, 2019年接触golng, 2022年开始学前端,到如今把过往所学开发个人blog,取名\"天兔座&\"寓意要做最靓的崽!",
})


/*刷新网页自动请求*/
onMounted(() => {
  new TypeIt("#myElement", {
    strings: user.text,
    cursor: false,/*在字符串末尾显示闪烁的光标*/
    lifeLike: true,/* 使打字速度不规则*/
  }).go();
})

import {MoreFilled} from '@element-plus/icons-vue'

const activities = [
  {
    content: 'arlog 天兔座& 阿尔法 项目成立',
    timestamp: '2024',
    size: 'large',
    type: 'primary',
    icon: MoreFilled,
  },
  {
    content: '想到自己实现一个博客',
    timestamp: '2022',
    color: '#0bbd87',
  },
  {
    content: '开始自学运维',
    timestamp: '2018',
    size: 'large',
  },
  {
    content: '开始参加工作',
    timestamp: '2017',
    type: 'primary',
    hollow: true,
  },
]

</script>

<template>
  <!--介绍页面-->
  <div class="center-block">
    <el-card style="max-width: 480px;">
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
            <br>
            <!--    文章 评论数据        -->
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
      <!--个人介绍-->
      <div class="mx-1">
        <div style="text-align: center; margin-top: -10px">
          <el-check-tag :checked=true type="info">
            简介
          </el-check-tag>

        </div>

        <el-text id="myElement" tag="b"></el-text>
      </div>


      <template #footer>
        <!--时间树-->
        <el-timeline style="max-width: 500px;">
          <div style="margin-left: -40px;">
            <el-timeline-item v-for="(activity, index) in activities"
                              :key="index"
                              :color="activity.color"
                              :hollow="activity.hollow"
                              :icon="activity.icon"
                              :size="activity.size"
                              :timestamp="activity.timestamp"
                              :type="activity.type"
            >
              {{ activity.content }}
            </el-timeline-item>
          </div>

        </el-timeline>


        <el-divider content-position="center" style="margin-top: -5px">友链推荐</el-divider>

        <!--友链推荐-->
        <el-space :size="20" alignment="center" class="alignment-container" wrap>
          <div v-for="i in 8" :key="i">

            <el-link href="https://www.arlog.cn">天兔座&</el-link>
          </div>
        </el-space>


      </template>
    </el-card>
  </div>


</template>

<style scoped>
/*友链*/
.alignment-container {
  height: 100px;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */

}


.mx-1 {
  height: 110px;

}

/*文章和评论的css*/
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
  height: 100%;
  overflow-y: hidden; /* 防止垂直滚动 */
  /*  margin-top: 30%; !*高度*!*/


}

</style>