module.exports = {
  // 配置项目启动端口及自动打开浏览器
  devServer: {
    port: 3333,
    open: true
  },
  // 配置scss全局变量
  css: {
    loaderOptions: {
      sass: {
        // 新版本sass-loader， 将data改成prependData进行配置
        // data: `@import "@/assets/scss/_variable.scss";`
        prependData: `@import "@/assets/scss/_variable.scss";`
      }
    }
  }
};
