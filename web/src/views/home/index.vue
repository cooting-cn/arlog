<script setup>

import {ChatLineRound, ChatSquare, Timer, User, View} from "@element-plus/icons-vue";

const fill = ref(true)
const txt = ref('接兼职，自建云waf防火墙，运维，监控，k8s等 联系博主QQ群:883574241 QQ号：953555501 备注加入原因')


</script>

<template>
  <div>
    <!-- 内容公告    -->
    <div style="max-width: 100%;">
      <el-check-tag checked style=" display: flex;align-items: center;" type="danger">
        <div style="width: 30px">
          <el-icon size="20">
            <ChatLineRound/>
          </el-icon>

        </div>
        <!--公告内容-->
        <el-text>
          {{ txt }}

        </el-text>
      </el-check-tag>
    </div>


    <!--文章内容-->
    <div style="margin-top: 20px;">

      <el-space :fill="fill" :size="10" style="width: 100%;" wrap>

        <el-card v-for="i in 6" :key="i" class="box-card">

          <!--     标题文字     -->
          <el-text size="large" style="width: 100%" tag="b">标题1</el-text>

          <!--     内容文字     -->
          <el-text style="width: 100%" truncated>
            Squeezed bSqueezed by
          </el-text>

          <!--     分割线     -->
          <el-divider style="height: 1px"/>

          <!--     用户     -->
          <el-space :size=30 wrap>
            <el-text size="small" type="info">
              <el-icon>
                <User/>
              </el-icon>
              admin
            </el-text>
            <!--     时间     -->
            <el-text size="small" tag="b" type="info">
              <el-icon>
                <Timer/>
              </el-icon>
              2018-04-12 20:46
            </el-text>
            <!-- 浏览量           -->
            <el-text size="small" tag="b" type="info">
              <el-icon>
                <View/>
              </el-icon>
              浏览量948
            </el-text>
            <!-- 评论量           -->
            <el-text size="small" tag="b" type="info">
              <el-icon>
                <ChatSquare/>
              </el-icon>
              评论量24
            </el-text>

          </el-space>
        </el-card>
      </el-space>
      <!-- 分页   -->
      <el-pagination :total="500"
                     background
                     class="mt-4"
                     layout="prev, pager, next"
                     small
                     style="margin-top: 20px;"
      />
    </div>


  </div>


</template>

<style scoped>

/**/

.el-divider--horizontal {
  margin: 5px 0;
}

.box-card {
  height: 110px;

}

</style>