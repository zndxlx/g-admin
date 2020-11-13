<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "app",
  components: {},
  created() {
    console.log("VUE_APP_BASE_API:", process.env.VUE_APP_BASE_API);
  },
};
</script>

<style lang="scss">
// 引入初始化样式
@import "@/style/main.scss";
@import "@/style/base.scss";
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
}
</style>
