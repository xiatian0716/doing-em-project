import Vue from "vue";
import Vuex from "vuex";

import tab from "./tab";
import user from "./user";
import esTest from "./esTest";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    tab,
    user,
    esTest
  }
});
