export default {
  state: {
    stockData: ""
  },
  mutations: {
    // 保存Token
    setStockData(state, val) {
      state.stockData = val;
    },
    // 清除Token
    clearStockData(state) {
      state.stockData = "";
    }
  },
  actions: {}
};
