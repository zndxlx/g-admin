<template>
  <div class="footer">
    <div>
      <span>Powered by</span>
      <span>
        <a href="https://github.com/flipped-aurora/gin-vue-admin"
          >gin-vue-admin</a
        >
      </span>
      <el-divider direction="vertical"></el-divider>
      <span>Copyright</span>
      <span>
        <a href="https://github.com/flipped-aurora">flipped-aurora团队</a>
      </span>
    </div>
  </div>
</template>

<script>
export default {
  name: "footer",
  components: {},
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
.footer {
  color: #888;
  height: 30px;
  line-height: 60px;
  a {
    color: #888;
  }
  div {
    display: flex;
    justify-content: center;
    span {
      margin: 0 3px;
    }
  }
}
</style>