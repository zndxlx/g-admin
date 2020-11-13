const path = require('path')

function resolve(dir) {
    return path.join(__dirname, dir)
}

module.exports = {
    devServer: {
        //port: 8080,
        open: true,
        overlay: {
            warnings: false,
            errors: true
        },
        proxy: {
            // 把key的路径代理到target位置
            // detail: https://cli.vuejs.org/config/#devserver-proxy
            [process.env.VUE_APP_BASE_API]: { //需要代理的路径   例如 '/api'
                target: `http://127.0.0.1:8888`, //代理到 目标路径
                changeOrigin: true,
                pathRewrite: { // 修改路径数据
                    ['^' + process.env.VUE_APP_BASE_API]: '' // 举例 '^/api:""' 把路径中的/api字符串删除
                }
            }
        },
    },


    //webpack配置
    configureWebpack: {
        //警告 webpack 的性能提示
        performance: {
            hints: 'warning',
            //入口起点的最大体积
            maxEntrypointSize: 50000000,
            //生成文件的最大体积
            maxAssetSize: 30000000,
            //只给出 js 文件的性能提示
            assetFilter: function (assetFilename) {
                return assetFilename.endsWith('.js');
            }
        }
    }
}