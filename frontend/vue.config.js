/**
 * @Author: Robby
 * @Date: 2021/12/6
 * @Filename: vue.config.js
 * @Function:
 **/

// 下面三行代码是为svg图标加载提供服务的，与其他的webpack配置无关
const path = require('path')
function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  // 对开发服务器的配置，配置了devServer后需要重启服务
  // devServer: {
  //   // 配置反向代理
  //   proxy: {
  //     // 当地址中有/api的时候会触发代理机制
  //     '/api': {
  //       // 要代理的服务器地址  这里不用写 api
  //       target: 'https://api.imooc-admin.lgdsunday.club/',
  //       changeOrigin: true // 是否跨域
  //     }
  //   }
  // },
  devServer: {
    // 配置反向代理
    proxy: {
      // 当地址中有/api的时候会触发代理机制
      '/v1': {
        // 要代理的服务器地址
        target: 'http://localhost:8888/',
        changeOrigin: true // 是否跨域
      }
    }
  },
  chainWebpack(config) {
    // 设置 svg-sprite-loader
    config.module.rule('svg').exclude.add(resolve('src/icons')).end()
    config.module
      .rule('icons')
      .test(/\.svg$/)
      .include.add(resolve('src/icons'))
      .end()
      .use('svg-sprite-loader')
      .loader('svg-sprite-loader')
      .options({
        symbolId: 'icon-[name]'
      })
      .end()
    // config.resolve.alias.set('@', resolve('src'))
  }
}
