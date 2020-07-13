import Cookie from "js-cookie";
export default {
  state: {
    token: ""
  },
  mutations: {
    // 保存Token
    setToken(state, val) {
      state.token = val;
      Cookie.set("token", val);
    },
    // 清除Token
    clearToken(state) {
      state.token = "";
      Cookie.remove("token");
    },
    // 获取Token
    getToken(state) {
      state.token = Cookie.get("token");
    }
  },
  actions: {}
};
