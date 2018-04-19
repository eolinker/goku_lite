'use strict';

var gulp = require('gulp'),
        config = require('./config'),
        path = require('path'),
        browserSync = require('browser-sync'),
        proxyMiddleware = require('http-proxy-middleware'),
        browserSyncSpa = require('browser-sync-spa'),
        gulpSequence = require('gulp-sequence');



gulp.task('watch', ['inject', 'vendor'], function () {
    //监控index.html,和bower.json文件
    gulp.watch([path.join(config.paths.src, '/*.html'), 'bower.json', 'vendor.base.json', 'vendor.json'], ['inject']);
    //监控CSS文件
    gulp.watch([path.join(config.paths.src, '/app/**/*.scss')], function (event) {
        if (event.type === 'changed') {
            gulp.start('styles:compass');
        } else {
            gulp.start('inject');
        }
    });
    //监控插件SCSS文件
    gulp.watch([path.join(config.paths.src,'/plug/**/*.scss')], function (event) {
        if (event.type === 'changed') {
            gulp.start('plug:compass');
        } 
    });
    //监控JS文件
    gulp.watch([path.join(config.paths.src, '/app/**/*.js')], function (event) {
        if (event.type === 'changed') {
            gulp.start('jshint');
        } else {
            gulp.start('inject');
        }
    });
    //监控html文件
    gulp.watch([
        path.join(config.paths.src, '/app/**/*.html')
    ], function (event) {
        browserSync.reload(event.path);
    });

});


function browserSyncInit(baseDir, open, port) {
    
    var onProxyRes = function (proxyRes, req, res) {
        // 重写set-cookie位置
        if (proxyRes.headers['set-cookie']) {
            proxyRes.headers['set-cookie'][0] = proxyRes.headers['set-cookie'][0].replace('/', '')
        }
    }
    
    browserSync.use(browserSyncSpa({
        selector: '[ng-app]'
    }));
    browserSync.init({
        startPath: '/',
        port: port || 3000,
        open: open || false,//决定Browsersync启动时自动打开的网址。默认为“本地”  false://停止自动打开浏览器
        server: {
            baseDir: baseDir,
            routes: {
                "/bower_components": "bower_components"
            },

            //使用代理
            
            middleware: [
                proxyMiddleware(['/Web'], {onProxyRes: onProxyRes, target: 'http://127.0.0.1:8080', changeOrigin: true,secure: false})
                // ,proxyMiddleware(['/automated'], {onProxyRes: onProxyRes, target: 'http://localhost:11205', changeOrigin: true})
            ]
        }
    });
}

exports.browserSyncInit = browserSyncInit;

gulp.task('serve', ['dev-config', 'watch'], function () {
    browserSyncInit([path.join(config.paths.tmp, '/serve'), config.paths.src], true);
});
gulp.task('serve:dist', ['build'], function () {
    browserSyncInit(config.paths.dist, true);
});