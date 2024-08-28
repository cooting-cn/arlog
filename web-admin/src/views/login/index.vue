<script setup>


import {ref} from "vue";

const timestamp = ref(118313526e4)

</script>

<template>
  <n-date-picker v-model:value="timestamp" type="date"/>

  <n-space>
    <n-button>Default</n-button>
    <n-button type="tertiary">
      Tertiary
    </n-button>
    <n-button type="primary">
      Primary
    </n-button>
    <n-button type="info">
      Info
    </n-button>
    <n-button type="success">
      Success
    </n-button>
    <n-button type="warning">
      Warning
    </n-button>
    <n-button type="error">
      Error
    </n-button>
  </n-space>
  <n-space vertical>
    <n-switch>
      <template #checked>
        老板
      </template>
      <template #unchecked>
        打工人
      </template>
    </n-switch>
    <n-menu/>
  </n-space>


</template>

<style scoped>

</style>