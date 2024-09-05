<script setup>

</script>

<template>
  <n-card>
    <n-space align="center">
      <n-avatar :size="100" round src="https://tg-image.com/file/b13dfda09561bf4318933.jpg"/>
      <div class="ml-20">
        <div class="flex items-center text-16">
          <span>用户名:</span>
          <span class="ml-12 opacity-80">666666</span>
          <n-button class="ml-32" text type="primary">
            <i class="i-fe:edit mr-4"/>
            修改密码
          </n-button>
        </div>
        <div class="mt-16 flex items-center">
          <n-button ghost type="primary">
            更改头像
          </n-button>
          <span class="ml-12 opacity-60">
              修改头像只支持在线链接，不提供上传图片功能，如有需要可自行对接！
            </span>
        </div>
      </div>
    </n-space>
  </n-card>


  <n-card class="mt-20" title="个人资料信息">
    <template #header-extra>
      <n-button text type="primary">
        <i class="i-fe:edit mr-4"/>
        修改资料
      </n-button>
    </template>

    <n-descriptions
        :column="1"
        :label-style="{ width: '200px', textAlign: 'center' }"
        bordered
        label-placement="left"
    >
      <n-descriptions-item label="昵称">
        666
      </n-descriptions-item>
      <n-descriptions-item label="性别">
        6666
      </n-descriptions-item>
      <n-descriptions-item label="地址">
        {6666
      </n-descriptions-item>
      <n-descriptions-item label="邮箱">
        6666
      </n-descriptions-item>
    </n-descriptions>
  </n-card>
</template>

<style scoped>

</style>