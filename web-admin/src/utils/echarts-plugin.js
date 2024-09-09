import * as echarts from 'echarts/core';
// 引入渲染器
import {CanvasRenderer} from 'echarts/renderers';

// 自动引入图表和组件
const chartTypes = [
    'LineChart',
    'BarChart',
    'PieChart',
    'ScatterChart',
    // 添加其他需要的图表类型
];

const components = [
    'TitleComponent',
    'TooltipComponent',
    'GridComponent',
    'LegendComponent',
    'DatasetComponent',
    'DataZoomComponent',
    // 添加其他需要的组件
];

chartTypes.forEach((chart) => {
    echarts.use(require(`echarts/charts`).default[chart]);
});

components.forEach((component) => {
    echarts.use(require(`echarts/components`).default[component]);
});

echarts.use(CanvasRenderer);

export default echarts;