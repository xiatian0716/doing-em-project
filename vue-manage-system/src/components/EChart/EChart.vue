<template>
  <div style="height: 100%" ref="echart">
    echart
  </div>
</template>

<script>
import echarts from "echarts";
export default {
  // 数据部分(由父组件传进来的)
  props: {
    chartData: {
      type: Object,
      default() {
        return {
          xData: [], // x轴
          series: [] // y轴
        };
      }
    },
    // 是否有坐标轴
    isAxisChart: {
      type: Boolean,
      default: true
    }
  },
  computed: {
    // 计算最终使用的option
    options() {
      return this.isAxisChart ? this.axisOption : this.normalOption;
    },
    // 计算属性isCollapse
    isCollapse() {
      return this.$store.state.tab.isCollapse;
    }
  },
  data() {
    return {
      // 容器(存放echart对象)
      echart: null,
      // 有坐标轴的配置
      axisOption: {
        legend: {
          textStyle: {
            color: "#333"
          }
        },
        grid: {
          left: "20%"
        },
        tooltip: {
          trigger: "axis"
        },
        xAxis: {
          type: "category",
          data: [], // 数据
          axisLine: {
            lineStyle: {
              color: "#17b3a3"
            }
          },
          axisLabel: {
            color: "#333"
          }
        },
        yAxis: [
          {
            type: "value",
            axisLine: {
              lineStyle: {
                color: "#17b3a3"
              }
            }
          }
        ],
        color: [
          "#2ec7c9",
          "#b6a2de",
          "#5ab1ef",
          "#ffb980",
          "#d87a80",
          "#8d98b3",
          "#e5cf0d",
          "#97b552",
          "#95706d",
          "#dc69aa",
          "#07a2a4",
          "#9a7fd1",
          "#588dd5",
          "#f5994e",
          "#c05050",
          "#59678c",
          "#c9ab00",
          "#7eb00a",
          "#6f5553",
          "#c14089"
        ],
        series: [] // 数据
      },
      // 没坐标轴的配置
      normalOption: {
        tooltip: {
          trigger: "item"
        },
        color: [
          "#0f78f4",
          "#dd536b",
          "#9462e5",
          "#a6a6a6",
          "#e1bb22",
          "#39c362",
          "#3ed1cf"
        ],
        series: [] // 数据
      }
    };
  },
  methods: {
    // initChart(初始化方法)
    initChart() {
      // 初始化图标时先渲染数据
      this.initChartData();
      // 容器存在
      if (this.echart) {
        this.echart.setOption(this.options);
        // 容器不存在
      } else {
        this.echart = echarts.init(this.$refs.echart);
        this.echart.setOption(this.options);
      }
    },
    // 配置图表数据方法
    initChartData() {
      // 有坐标轴的数据
      if (this.isAxisChart) {
        this.axisOption.xAxis.data = this.chartData.xData;
        this.axisOption.series = this.chartData.series;
        // 没有坐标轴的数据
      } else {
        this.normalOption.series = this.chartData.series;
      }
    },
    // 自适应-对图表尺寸进行重新计算
    resizeChart() {
      this.echart ? this.echart.resize() : "";
    }
  },
  // 什么时候去初始化图表呢？
  // chartData发生变化时
  // 通过监听的方式
  // 对象(深度监听)
  watch: {
    chartData: {
      handler: function() {
        this.initChart();
      },
      deep: true
    },
    // 自适应-折叠菜单时
    isCollapse() {
      setTimeout(() => {
        this.resizeChart();
      }, 300);
    }
  },
  // 自适应-窗口大小-钩子
  mounted() {
    // 注册resize监听事件
    window.addEventListener("resize", this.resizeChart);
  },
  destroyed() {
    // 销毁resize监听事件-防止内存溢出(组件销毁时销毁)
    window.removeEventListener("resize", this.resizeChart);
  }
};
</script>

<style lang="scss" scoped></style>
