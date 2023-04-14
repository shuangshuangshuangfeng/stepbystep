## echart 用法



### 1. React Echart

**使用：**

```js
npm install --save echarts-for-react


```



**React demo**

```js
import React from 'react';
import ReactECharts from 'echarts-for-react';


<ReactECharts
  option={this.getOption()} // option 相关字段在官方给的配置文档
  notMerge={true} // 后一个option是否与前一个option merge， 这个要选true
  lazyUpdate={true}
  theme={"theme_name"}
  onChartReady={this.onChartReadyCallback}
  onEvents={EventsDict}
  opts={}
/>


```



**echart官方文档：**

[Documentation - Apache ECharts](https://echarts.apache.org/zh/option.html#color)

**echart-for-react官方文档：**

https://www.npmjs.com/package/echarts-for-react


