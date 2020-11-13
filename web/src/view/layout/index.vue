<template>
  <el-container class="layout-cont">
    <el-aside>
      <Aside></Aside>
    </el-aside>
    <el-container>
      <el-header><Header></Header></el-header>
      <el-main>
        <router-view class="admin-box"></router-view>
      </el-main>
      <el-footer>
        <Footer></Footer>
      </el-footer>
    </el-container>
  </el-container>
</template>

<script>
import Footer from "@/view/layout/footer/footer";
import Aside from "@/view/layout/aside";
import Header from "@/view/layout/header/header";
export default {
  name: "Layout",
  components: {
    Aside,
    Footer,
    Header,
  },
  data() {
    return {};
  },
  computed: {},
  methods: {},
  created() {},
  mounted() {},
};
</script>

<style scoped lang="scss">
.layout-cont {
  margin-bottom: 40px;
  .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    line-height: 200px;
    overflow: hidden;
    width: auto !important;
  }
  .el-container {
    .el-header {
      line-height: 60px;
      color: #606266;
      background-color: #fff;
    }
    .el-footer {
      background-color: #B3C0D1;
      color: #333;
      line-height: 60px;
    }

    .el-main {
      background-color: #E9EEF3;
      color: #333;
      // text-align: center;
      // line-height: 160px;
    }
    .admin-box {
      background-color: #fff;
      padding: 24px;
      border-radius: 2px;
    }
  }
}
</style>