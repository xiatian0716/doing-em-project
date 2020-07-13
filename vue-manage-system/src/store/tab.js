import Cookie from "js-cookie";
export default {
  // state
  state: {
    isCollapse: false,
    menu: [],
    // 当前目录
    currentMenu: {},
    // tab列表
    tabsList: [
      {
        path: "/",
        name: "home",
        label: "首页",
        icon: "home"
      }
    ]
  },
  // 方法
  mutations: {
    // 获取菜单-登录
    setMenu(state, val) {
      state.menu = val;
      Cookie.set("menu", JSON.stringify(val));
    },
    // 清除菜单-登出
    clearMenu(state) {
      state.menu = [];
      Cookie.remove("menu");
    },
    // 动态添加到路由
    addMenu(state, router) {
      // 从Cookie里面获取
      if (!Cookie.get("menu")) {
        return;
      }
      let menu = JSON.parse(Cookie.get("menu"));
      // 将menu添加到state
      state.menu = menu;
      let currentMenu = [
        {
          path: "/",
          component: () => import(`@/views/Main`),
          // 动态添加的路由都放到children里
          children: []
        }
      ];
      // 循环操作来拼接路径
      menu.forEach(item => {
        if (item.children) {
          item.children = item.children.map(item => {
            item.component = () => import(`@/views/${item.url}`);
            return item;
          });
          currentMenu[0].children.push(...item.children);
        } else {
          item.component = () => import(`@/views/${item.url}`);
          currentMenu[0].children.push(item);
        }
      });
      // 动态添加路由
      router.addRoutes(currentMenu);
    },
    getMenu(state) {
      if (Cookie.get("tagList")) {
        let tagList = JSON.parse(Cookie.get("tagList"));
        state.tabsList = tagList;
      }
    },

    // 点击后将val传进来
    selectMenu(state, val) {
      if (val.name !== "home") {
        state.currentMenu = val;
        // 判断是否存在？然后添加
        let result = state.tabsList.findIndex(item => item.name === val.name);
        result === -1 ? state.tabsList.push(val) : "";
      } else {
        state.currentMenu = null;
      }
    },
    closeTab(state, val) {
      // 找到并删除
      let result = state.tabsList.findIndex(item => item.name === val.name);
      state.tabsList.splice(result, 1);
    },
    collapseMenu(state) {
      state.isCollapse = !state.isCollapse;
    }
  },
  actions: {}
};
