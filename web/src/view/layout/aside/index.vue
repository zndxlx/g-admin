<template>
  <div class="aside">
    <div class="tilte">
      <img alt class="logoimg" src="~@/assets/nav_logo.png" />
      <h2 class="tit-text" v-if="!isCollapse">Gin-Vue-Admin</h2>
    </div>

    <el-scrollbar
      style="height: calc(100vh - 64px)"
      wrap-style="overflow-x:hidden;"
    >
      <transition
        :duration="{ enter: 800, leave: 100 }"
        mode="out-in"
        name="el-fade-in-linear"
      >
        <el-menu
          :collapse="isCollapse"
          :collapse-transition="true"
          :default-active="active"
          @select="selectMenuItem"
          active-text-color="#fff"
          class="el-menu-vertical"
          text-color="rgb(191, 203, 217)"
          background-color="#001529"
          unique-opened
        >
          <template v-for="item in asyncRouters[0].children">
            <aside-component
              :key="item.name"
              :routerInfo="item"
              v-if="!item.hidden"
            />
          </template>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters, mapMutations } from "vuex";
import AsideComponent from "@/view/layout/aside/asideComponent";
export default {
  name: "Aside",
  components: {
    AsideComponent,
  },
  data() {
    return {
      //isSider: true,
      isCollapse: false,
    };
  },
  computed: {
    ...mapGetters("router", ["asyncRouters"]),
  },
  methods: {
    selectMenuItem(index, _, ele) {
      const query = {};
      const params = {};
      //console.log(index);
      ele.route.parameters &&
        ele.route.parameters.map((item) => {
          if (item.type == "query") {
            query[item.key] = item.value;
          } else {
            params[item.key] = item.value;
          }
        });
      if (index === this.$route.name) return;
      if (index.indexOf("http://") > -1 || index.indexOf("https://") > -1) {
        window.open(index);
      } else {
        this.$router.push({ name: index, query, params });
      }
    },
  },
  created() {
    this.$bus.on("collapse", (item) => {
      this.isCollapse = item;
    });
  },
  mounted() {},
};
</script>

<style  lang="scss">
.aside {
  overflow: hidden;
  .tilte {
    background: #001529;
    min-height: 64px;
    line-height: 64px;
    background: #002140;
    text-align: center;
    .logoimg {
      width: 30px;
      height: 30px;
      vertical-align: middle;
      background: #fff;
      border-radius: 50%;
      padding: 3px;
    }
    .tit-text {
      display: inline-block;
      color: #fff;
      font-weight: 600;
      font-size: 20px;
      vertical-align: middle;
    }
  }
  .el-scrollbar {
    .el-scrollbar__wrap {
      overflow-x: hidden;
    }
  }
  .el-menu-vertical {
    height: calc(100vh - 64px) !important;
    visibility: auto;
    &:not(.el-menu--collapse) {
      width: 220px;
    }
  }
  .el-menu--collapse {
    width: 54px;
    li {
      .el-tooltip,
      .el-submenu__title {
        padding: 0px 15px !important;
      }
    }
  }
}
</style>