<script setup>

</script>

<template>

  <router-link class="h-60px flex items-center justify-center " to="/">
    <div class="flex">
      <img alt="Logo" class="h-50" src="@/assets/images/arlog.webp">
    </div>

    <h2
        class="ml-10 max-w-140 flex-shrink-0 text-20  font-bold text-[#8ba923db] "
    >
      Arlog Admin
    </h2>

  </router-link>

</template>

<style scoped>
.color-primary, .text-primary {
  --un-text-opacity: 1;
  color: rgba(var(49, 108, 114), var(--un-text-opacity)) /* rgba(var(--primary-color)) */;
}
</style>