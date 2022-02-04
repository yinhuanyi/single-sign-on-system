/**
 * @Author: Robby
 * @Date: 2021/12/6
 * @Filename: index.js
 * @Function: 导入所有的svg图标，完成svgIcon组件的全局注册
 **/
import SvgIcon from '@/components/SvgIcon'

// 此时返回了一个require函数，可以接收一个request参数
// const svgRequire = require.context('./svg', false, '/.svg$/')
const svgRequire = require.context('./svg', false, /\.svg$/)

// 可以通过svgRequire.keys()获得目录下所有的图标，遍历图标，把图标作为request的参数导入，这样就完成了注册
svgRequire.keys().forEach((svgicon) => svgRequire(svgicon))

// 这里是全局组成组件
export default (app) => {
  app.component('svg-icon', SvgIcon)
}
