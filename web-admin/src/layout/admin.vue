<script setup>

</script>

<template>
  <div class="wh-full flex">
    <aside
        :class="collapsed ? 'w-64' : 'w-220'"
        border-r="1px solid light_border dark:dark_border"
        class="flex-col flex-shrink-0 transition-width-300"
    >
      <h1>66666666</h1>
    </aside>

    <article class="w-0 flex-col flex-1">
      <h1>66666666</h1>
      <router-view></router-view>
    </article>
  </div>


</template>

<style scoped>
.collapsed {
  width: 64px;
}
</style>