<script setup>
import vueDanmaku from 'vue3-danmaku'

const danmus = ref([{
  avatar: 'https://tg-image.com/file/b13dfda09561bf4318933.jpg',
  name: '三生',
  text: '6666'
},
  {
    avatar: 'https://tg-image.com/file/b13dfda09561bf4318933.jpg',
    name: '那末',
    text: '点赞！！！'
  }])

const types = ['primary', 'info', 'success', 'warning', 'error'];

function getRandomType() {
  const randomIndex = Math.floor(Math.random() * types.length);
  return types[randomIndex];
}
</script>

<template>
  <n-card class="h-100% min-w-900 ">
    <vue-danmaku v-model:danmus="danmus" :debounce="200" :speeds="100" class="w-100% h-550px " loop
                 randomChannel
                 useSlot>
      <template v-slot:dm="{ index, danmu }">
        <div class="flex items-center">
          <img :src="danmu.avatar" alt="用户头像" class="rounded-full w-30 h-30 mr-2">
          <n-gradient-text :type="getRandomType()">
            {{ danmu.name }}
          </n-gradient-text>
          :
          <n-gradient-text :type="getRandomType()">
            {{ danmu.text }}
          </n-gradient-text>

        </div>
      </template>
      <h1>6666666666</h1>
    </vue-danmaku>

  </n-card>

</template>

<style scoped>

</style>