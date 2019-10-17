'use strict';

var gulp = require('gulp'),
    config = require('./config'),
    browserSync = require('browser-sync'),
    proxyMiddleware = require('http-proxy-middleware'),
    browserSyncSpa = require('browser-sync-spa');
function browserSyncInit(baseDir, open, port) {
    var onProxyRes = function (proxyRes, req, res) {
        // 重写set-cookie位置
        if (proxyRes.headers['set-cookie']) {
            proxyRes.headers['set-cookie'][0] = proxyRes.headers['set-cookie'][0].replace('domain=.eolinker.com', 'domain=localhost')
        }
    }
    browserSync.use(browserSyncSpa({
        selector: '[ng-app]'
    }));
    browserSync.init({
        startPath: '/', //zh-cn/
        port: port || 3000,
        https: true,
        open: open || false, //决定Browsersync启动时自动打开的网址。默认为“本地”  false://停止自动打开浏览器
        server: {
            baseDir: baseDir,
            routes: {
                "/bower_components": "bower_components"
            },
            //使用代理
            middleware: [
                proxyMiddleware(['/apis','/strategy','/node','/plugin','/auth','/project','/guest','/balance','/import','/gateway','/user','/message','/account','/permission','/monitor'], {onProxyRes: onProxyRes, target: 'http://47.95.203.198:10003', changeOrigin: true,secure: false})
            ]
        }
    });
}

exports.browserSyncInit = browserSyncInit;

gulp.task('serve:dist', ['build'], function () {
    browserSyncInit(config.paths.dist, true);
});