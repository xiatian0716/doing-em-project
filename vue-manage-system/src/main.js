import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

// 全局配置
import "@/assets/scss/reset.scss";

// 完整引入ElementUI
import ElementUI from "element-ui";
import "element-ui/lib/theme-chalk/index.css";
Vue.use(ElementUI);

// 使用axios-mockjs
import "./mock";
import http from "@/api/config";
Vue.prototype.$http = http;

Vue.config.productionTip = false;

router.beforeEach((to, from, next) => {
  // 防止刷新后vuex里丢失token
  store.commit("getToken");
  // 防止刷新后vuex里丢失标签列表tagList
  store.commit("getMenu");
  let token = store.state.user.token;
  // 过滤登录页，防止死循环
  if (!token && to.name !== "login") {
    next({ name: "login" });
  } else {
    next();
  }
});

new Vue({
  router,
  store,
  render: h => h(App),

  // 刷新时触发操作
  created() {
    store.commit("addMenu", router);
  }
}).$mount("#app");
