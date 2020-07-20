<template>
  <div style="padding: 20px">{{ menu }}}</div>
</template>

<script>
export default {
  data() {
    return {};
  },
  // 计算属性
  computed: {
    menu() {
      return this.$store.state.esTest.stockData;
    }
  },
  mounted: function() {
    // 登录成功时触发操作
    this.$http.get("/home/getStockData").then(res => {
      res = res.data;
      console.log(JSON.stringify(res));
      if (res.code === 20000) {
        this.$store.commit("setStockData", res.data.menu);
      } else {
        this.$message.warning(res.data.message);
      }
    });
  }
};
</script>

<style lang="scss" scoped></style>
