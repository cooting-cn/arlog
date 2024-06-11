<script setup>

import {ChatLineRound} from "@element-plus/icons-vue";

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


    <div style="margin-top: 20px">

      <el-space :fill="fill" style="width: 100%" wrap>
        <el-card v-for="i in 6" :key="i" class="box-card">

          <div v-for="o in 3" :key="o" class="text item">
            {{ 'List item ' + o }}
          </div>
        </el-card>
      </el-space>
    </div>


  </div>


</template>

<style scoped>

</style>