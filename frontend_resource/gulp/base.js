'use strict';

var gulp = require('gulp'),
    path = require('path'),
    fs = require('fs'),
    config = require('./config'),
    _ = require('lodash'),
    $ = require('gulp-load-plugins')({
      pattern: ['gulp-*', 'event-stream', 'main-bower-files', 'uglify-save-license', 'del']
    }),
    browserSync = require('browser-sync'),
    gulpsync    = $.sync(gulp),
    reload      = browserSync.reload;//实时刷新


gulp.task('dev-config',function () {
  return gulp.src('app.conf.json')
        .pipe($.ngConfig(config.modules.ConstantModuleName,{
          environment: 'development',
          createModule: false,
          wrap: true
        }))
        .pipe(gulp.dest(path.join(config.paths.src,'/app')))//join() 方法用于把数组中的所有元素放入一个字符串。
});
gulp.task('prod-config',function () {
  return gulp.src('app.conf.json')
        .pipe($.ngConfig(config.modules.ConstantModuleName,{
          environment: 'production',
          createModule: false,
          wrap: true//生成闭包
        }))
        .pipe(gulp.dest(path.join(config.paths.src,'/app')))
});
gulp.task('test-config',function () {
  return gulp.src('app.conf.json')
        .pipe($.ngConfig(config.modules.ConstantModuleName,{
          environment: 'test',
          createModule: false,
          wrap: true//生成闭包
        }))
        .pipe(gulp.dest(path.join(config.paths.src,'/app')))
});


/**
 * [代码质量管理]
 */

gulp.task('jshint',function () {
  return gulp.src(path.join(config.paths.src,'app/**/*.js'))
  .pipe($.plumber(config.errorHandler()))
  .pipe($.jshint())
  .pipe(reload({ stream: true }))
  .pipe($.size());
});


/**
 * [清理DIST,TEMP文件夹]
 */

gulp.task('clean', function () {
  $.del([path.join(config.paths.dist, '/'), path.join(config.paths.tmp, '/')]);
});


/**
 * [编译之前将scss注入index.scss]
 */

gulp.task('inject_sass',function () {
  var injectFiles = gulp.src([
      path.join(config.paths.src,'app/**/*.scss'),
      path.join('!'+ config.paths.src, 'app/index.scss')
    ],{read:false});

  var injectOptions = {
    transform: function(filePath) {
      filePath = filePath.replace(config.paths.src + '/app/', '');
      return '@import "' + filePath + '";';
    },
    starttag: '// injector',
    endtag: '// endinjector',
    addRootSlash: false
  };
  return gulp.src(path.join(config.paths.src,'app/index.scss'))
          .pipe($.inject(injectFiles,injectOptions))
          .pipe(gulp.dest(path.join(config.paths.src,'app/')))
});


gulp.task('clean', function () {
  $.del([path.join(config.paths.dist, '/'), path.join(config.paths.tmp, '/')]);
});

/**
 * [SASS预编译模块,依赖compass模块编译]
 */

gulp.task('styles:compass',['inject_sass'],function () {
  return gulp.src(path.join(config.paths.src,'app/index.scss'))
    .pipe($.plumber(config.errorHandler()))
    .pipe($.compass({
      config_file: path.join(__dirname, '/../config.rb'),
        css: path.join(config.paths.tmp, '/serve/app/'),
        sass: path.join(config.paths.src, '/app/'),
    }))
    //sprite图片路径修复
    .pipe($.replace('../../../src/assets/images/', '../assets/images/'))
    .pipe(gulp.dest(path.join(config.paths.tmp,'/serve/app/')))
    //css改变时无刷新改变页面
    .pipe(reload({ stream: true }));
});



/**
 * [Html中的CSS以及JS注入]
 */

gulp.task('inject', ['jshint', 'styles:compass','vendor:base'], function () {
  var injectStyles = gulp.src([
    path.join(config.paths.tmp, '/serve/app/**/*.css')
  ], { read: false });

  var injectScripts = gulp.src([
    path.join(config.paths.src, '/app/**/*.js'),
    path.join('!' +config.paths.src, '/app/vendor.js'),
  ]).pipe($.angularFilesort());

  var injectOptions = {
    ignorePath: [config.paths.src, path.join(config.paths.tmp, '/serve')],
    addRootSlash: false
  };

  return gulp.src(path.join(config.paths.src, '/*.html'))
    .pipe($.plumber(config.errorHandler()))
    .pipe($.inject($.eventStream.merge(
      injectStyles,
      injectScripts
    ),injectOptions))
    .pipe(gulp.dest(path.join(config.paths.tmp, '/serve')));

});

gulp.task('vendor', gulpsync.sync(['vendor:base']) );
// gulp.task('vendor', gulpsync.sync(['vendor:base', 'vendor:app']) );


/**
 * [复制依赖文件]
 */

gulp.task('vendor:base', function() {
    var jsFilter = $.filter('**/*.js',{restore: true}),
        cssFilter = $.filter('**/*.css',{restore: true});
    return gulp.src(config.vendor.base.source)
        .pipe($.expectFile(config.vendor.base.source))
        .pipe(jsFilter)
        .pipe($.concat(config.vendor.base.name+'.js'))
        .pipe(jsFilter.restore)
        .pipe(cssFilter)
        .pipe($.concat(config.vendor.base.name+'.scss'))
        .pipe(cssFilter.restore)
        .pipe(gulp.dest(config.vendor.base.dest))
        ;
});

gulp.task('vendor:app', function() {

  var jsFilter = $.filter('*.js',{restore: true}),
      cssFilter = $.filter('*.css',{restore: true});

  return gulp.src(config.vendor.app.source, {base: 'bower_components'})
      .pipe($.expectFile(config.vendor.app.source))
      .pipe(jsFilter)
      .pipe(jsFilter.restore)
      .pipe(cssFilter)
      .pipe(cssFilter.restore)
      .pipe(gulp.dest(config.vendor.app.dest) );

});