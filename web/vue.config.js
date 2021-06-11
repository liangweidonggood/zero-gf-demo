module.exports = {
    devServer: {
        /* 自动打开浏览器 */
        open: true,
        /* 设置为0.0.0.0则所有的地址均能访问 */
        host: '0.0.0.0',
        port: 8000,
        https: false,
        hotOnly: false,
        /* 使用代理 */
        proxy: {
            '/api': {
                /* 目标代理服务器地址 */
                target: 'http://localhost:8888/',
                /* 允许跨域 */
                changeOrigin: true,
                pathRewrite: {
                    '^/api': '' //路径重写
                }
            },
        },
    },
}
