<script setup>
import {useMessage} from "naive-ui";

const nameRef = ref(1)

const message = useMessage()
const panelsRef = ref([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15])

function handleClose(name) {
  const {value: panels} = panelsRef;
  if (panels.length === 1) {
    message.error("最后一个了");
    return;
  }
  message.info(`关掉 ${name}`);
  const index = panels.findIndex((v) => name === v);
  panels.splice(index, 1);
  if (nameRef.value === name) {
    nameRef.value = panels[index];
  }
}

const panels = panelsRef

const name = nameRef
</script>

<template>
  <!--图标-->
  <div class="flex items-center ">
    <!-- 左侧图标 -->
    <div class="flex justify-center items-center h-60px w-50px">
      <n-icon class="i-ep-expand " color="rgba(87, 89, 87, 0.9)" size="25"/>
    </div>

    <div>
      <n-tabs
          v-model:value="name"


          class="w-700px "
          closable
          type="card"
          @close="handleClose"
      >
        <n-tab
            v-for="panel in panels"
            :key="panel"
            :name="panel"
            :tab="panel.toString()"
        >
          666
        </n-tab>
      </n-tabs>
    </div>

  </div>

</template>

<style scoped>

</style>