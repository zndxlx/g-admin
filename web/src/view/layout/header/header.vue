<template>
  <div class="header">
    <div class="line1">
      <span @click="totalCollapse" class="menu-total">
        <i class="el-icon-s-unfold" v-if="isCollapse"></i>
        <i class="el-icon-s-fold" v-else></i>
      </span>
      <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
        <el-breadcrumb-item
          :key="item.path"
          v-for="item in matched.slice(1, matched.length)"
        >
          {{ item.meta.title }}
        </el-breadcrumb-item>
      </el-breadcrumb>
    </div>
  </div>
</template>

<script>
export default {
  name: "header",
  components: {},
  data() {
    return {
      isCollapse: false,
    };
  },
  computed: {
    matched() {
      return this.$route.matched;
    },
  },
  methods: {
    totalCollapse() {
      this.isCollapse = !this.isCollapse;
      this.$bus.emit("collapse", this.isCollapse);
    },
  },
  created() {},
  mounted() {},
};
</script>

<style scoped lang="scss">
.line1 {
  height: 48px;
  background-color: pink;
  line-height: 48px;
  .menu-total {
    display: inline-block;
    margin-left: -10px;
    i {
      font-size: 22px;
      font-weight: 500;
    }
  }
  .breadcrumb {
    //line-height: 48px;
    display: inline-block;
    padding: 0 24px;
  }
}
</style>