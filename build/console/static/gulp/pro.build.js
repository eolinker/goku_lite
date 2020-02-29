'use strict';
/**
 * @name 编译
 * @author 广州银云信息科技有限公司
 */
/**
 * @version 4.0
 * @description 去掉图片压缩&字体压缩
 */
var gulp = require('gulp');
var path = require('path');
var config = require('./config');
var _ = require('lodash');
var $ = require('gulp-load-plugins')({
    pattern: ['gulp-*', 'main-bower-files', 'uglify-save-license', 'del', 'imagemin-pngquant']
});
var sass = require('gulp-sass');
var replace = require('gulp-replace');
/**
 * [生成Html模版文件]
 */

gulp.task('build:partials', function () {
    return gulp.src([
            path.join(config.paths.src, '/app/**/*.html')
        ])
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('templateCacheHtml.js', {
            module: config.modules.templateModuleName,
            root: 'app'
        }))
        .pipe(gulp.dest(config.paths.tmp + '/partials/'));
});

gulp.task('build:partials:_default', function () {
    return gulp.src([
            path.join(config.paths.src, '/module/_default/**/*.html')
        ])
        .pipe($.minifyHtml({
            empty: true,
            spare: true,
            quotes: true
        }))
        .pipe($.angularTemplatecache('_default_templateCacheHtml.js', {
            module: config.modules.templateModuleName,
            root: 'module/_default'
        }))
        .pipe(gulp.dest(config.paths.tmp + '/partials/'));
});
/**
 * [Html,Js,Css压缩合并]
 */
gulp.task('build:html', ['build:inject', 'build:partials','build:partials:_default'], function () {
    var partialsInjectFile = gulp.src([path.join(config.paths.tmp, '/partials/templateCacheHtml.js'),path.join(config.paths.tmp, '/partials/_default_templateCacheHtml.js')], {
        read: false
    });
    var partialsInjectOptions = {
        starttag: '<!-- inject:partials -->',
        ignorePath: path.join(config.paths.tmp, '/partials'),
        addRootSlash: false
    };

    var htmlFilter = $.filter('*.html', {
        restore: true
    });
    var jsFilter = $.filter('**/*.js', {
        restore: true
    });
    var cssFilter = $.filter('**/*.css', {
        restore: true
    });

    return gulp.src(path.join(config.paths.tmp, '/serve/*.html'))
        //error 
        .pipe($.plumber(config.errorHandler()))
        //inject template
        .pipe($.inject(partialsInjectFile, partialsInjectOptions))
        //js
        .pipe($.useref()) //合并和压缩
        .pipe(jsFilter)
        //修复HTML图片地址
        .pipe($.replace('app/assets/', 'assets/'))
        .pipe($.stripDebug())
        .pipe($.uglify())
        .pipe(jsFilter.restore)
        //css 
        .pipe(cssFilter)
        //修复HTML图片地址
        .pipe($.replace('app/assets/', 'assets/'))
        .pipe($.replace('(assets/', '(../assets/'))
        .pipe($.autoprefixer({
            browsers: ['last 20 versions'],
            cascade: false
        }))
        .pipe($.csso())
        .pipe(cssFilter.restore)
        //md5后缀
        .pipe($.if('*.css', $.rev()))
        .pipe($.if('*.js', $.rev()))
        //替换md5后缀的文件名
        .pipe($.revReplace())
        //html处理
        .pipe(htmlFilter)
        // .pipe($.replace('<base href="/">', '<base href="/eolinker/">'))
        // .pipe($.minifyHtml({
        //     empty: true,
        //     spare: true,
        //     quotes: true,
        //     conditionals: true
        // }))
        .pipe(htmlFilter.restore)
        .pipe(gulp.dest(path.join(config.paths.dist, '/')))
        .pipe($.size({
            title: path.join(config.paths.dist, '/'),
            showFiles: true
        }));

});


gulp.task('build', $.sequence('prod-config', ['clean', 'clear_inject_sass', 'build:html'],'other:vendor', 'other:libs', 'other:assets'));//,'translate'
gulp.task('build:e2e', $.sequence('test-config', ['clean', 'build:html'],'other:vendor', 'other:libs', 'other:assets'));