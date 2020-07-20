import Mock from "mockjs";

export default {
  getStockData: () => {
    return {
      code: 20000,
      data: {
        menu: [
          {
            path: "/",
            name: "home",
            label: "首页",
            icon: "s-home",
            url: "Home/Home"
          },
          {
            path: "/video",
            name: "video",
            label: "视频管理页",
            icon: "video-play",
            url: "VideoManage/VideoManage"
          }
        ],
        token: Mock.Random.guid(),
        message: "获取成功"
      }
    };
  }
};
