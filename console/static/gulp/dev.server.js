'use strict';

var gulp = require('gulp'),
    config = require('./config'),
    path = require('path'),
    browserSync = require('browser-sync'),
    proxyMiddleware = require('http-proxy-middleware'),
    browserSyncSpa = require('browser-sync-spa');

gulp.task('watch', ['inject', 'vendor'], function () {
    //监控index.html,和bower.json文件
    // gulp.watch([path.join(config.paths.src, '/*.html'), 'bower.json', 'vendor.base.json', 'vendor.json'], ['inject']);
    //监控CSS文件
    gulp.watch([
        path.join(config.paths.src, '/app/**/*.scss'),
        path.join(config.paths.src, '/module/**/*.scss')
    ], function (event) {
        if (event.type === 'changed') {
            gulp.start('styles:compass');
        } else {
            gulp.start('inject');
        }
    });
    gulp.watch([
        path.join(config.paths.src, '/theme/**/*.scss')
    ], function (event) {
        if (event.type === 'changed') {
            gulp.start('styles:theme:compass');
            browserSync.reload(event.path);
        }
    });
    //监控JS文件
    gulp.watch([
        path.join(config.paths.src, '/module/**/*.js'),
        path.join(config.paths.src, '/app/**/*.js'),
        path.join('!' + config.paths.src, 'module/**/module.js')
    ], function (event) {
        if (event.type === 'changed') {
            gulp.start('jshint');
            browserSync.reload(event.path);

        } else {
            gulp.start('inject');
        }
    });
    //监控html文件
    gulp.watch([
        path.join(config.paths.src, '/app/**/*.html'),
        path.join(config.paths.src, '/module/**/*.html'),
        path.join('!' + config.paths.src, '/**/*.tmp.html')
    ], function (event) {
        browserSync.reload(event.path);
    });
    gulp.watch([
        path.join(config.paths.src, '/app/**/*.tmp.html'),
    ], function (event) {
        gulp.start('build:tmpHtml:app');
        browserSync.reload(event.path);
    });
    gulp.watch([
        path.join(config.paths.src, '/module/**/*.tmp.html'),
    ], function (event) {
        gulp.start('build:tmpHtml:module');
        browserSync.reload(event.path);
    });
});




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
        // https: true,
        open: open || false, //决定Browsersync启动时自动打开的网址。默认为“本地”  false://停止自动打开浏览器
        server: {
            baseDir: baseDir,
            routes: {
                "/bower_components": "bower_components"
            },
            //使用代理
            middleware: [
                proxyMiddleware(['/config','/cluster','/balance', '/apis', '/strategy', '/node', '/plugin', '/auth', '/project', '/guest', '/import', '/gateway', '/user', '/message','/account', '/permission', '/monitor'], {
                    onProxyRes: onProxyRes,
                    target: 'http://127.0.0.1',
                    changeOrigin: true,
                    secure: false
                })
            ]
        }
    });
}

exports.browserSyncInit = browserSyncInit;

gulp.task('serve', ['clean:tmp', 'dev-config', 'watch'], function () {
    browserSyncInit([path.join(config.paths.tmp, '/serve'), config.paths.src], true);
});