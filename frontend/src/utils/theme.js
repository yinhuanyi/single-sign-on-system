/**
 * @Author: Robby
 * @Date: 2022/3/4
 * @Filename: theme.js
 * @Function:
 **/

import color from 'css-color-function'
import rgbHex from 'rgb-hex'
import formula from '@/constant/formula.json'
import axios from 'axios'

// 把生产的样式表，写入到style中
export const writeNewStyle = (newStyle) => {
  // 在html的header部分，将 <style>标签，内部是样式
  const style = document.createElement('style')
  style.innerText = newStyle
  document.head.appendChild(style)
}

// 基于主题颜色，生成样式表
export const generateNewStyle = async (primaryColor) => {
  // 1：根据主色值，生成最新的样式表
  const colors = generateColors(primaryColor)
  // 2：将element-plus的默认样式表，将色值打标记
  let cssText = await getOriginalStyle()
  // console.log(cssText)
  // 3：遍历生成的样式表，在 CSS 的原样式中进行全局替换
  Object.keys(colors).forEach((key) => {
    cssText = cssText.replace(
      new RegExp('(:|\\s+)' + key, 'g'),
      '$1' + colors[key]
    )
  })
  return cssText
}

// 生成16进制的颜色对象
export const generateColors = (primary) => {
  if (!primary) return
  const colors = {
    primary
  }
  // 遍历formula对象
  Object.keys(formula).forEach((key) => {
    const value = formula[key].replace(/primary/g, primary)
    colors[key] = '#' + rgbHex(color.convert(value))
  })
  return colors
}

// 获取默认样式表
const getOriginalStyle = async () => {
  // 获取当前element-plus版本号
  const version = require('element-plus/package.json').version
  // 基于当前项目element-plus版本号，获取样式表
  const url = `https://unpkg.com/element-plus@${version}/dist/index.css`
  // 获取样式表的数据，这里axios返回的是一个promise对象
  const { data } = await axios(url)
  // 把获取到的数据筛选为原样式模板
  return getStyleTemplate(data)
}

// 将element-plus原生的样式表
const getStyleTemplate = (data) => {
  // element-plus 默认色值
  const colorMap = {
    '#3a8ee6': 'shade-1',
    '#409eff': 'primary',
    '#53a8ff': 'light-1',
    '#66b1ff': 'light-2',
    '#79bbff': 'light-3',
    '#8cc5ff': 'light-4',
    '#a0cfff': 'light-5',
    '#b3d8ff': 'light-6',
    '#c6e2ff': 'light-7',
    '#d9ecff': 'light-8',
    '#ecf5ff': 'light-9'
  }
  // 根据默认色值为要替换的色值打上标记
  Object.keys(colorMap).forEach((key) => {
    const value = colorMap[key]
    data = data.replace(new RegExp(key, 'ig'), value)
  })
  return data
}
