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
  gulpsync = $.sync(gulp),
  reload = browserSync.reload; //实时刷新

const babel = require('gulp-babel');

/**
 * [编译之前将scss注入index.scss]
 */

gulp.task('clear_inject_sass', function () {
  var injectFiles = gulp.src([
    path.join(config.paths.src, 'module/_default/**/*.scss')
  ], {
    read: false
  });

  var injectOptions = {
    transform: function (filePath) {
      filePath = filePath.replace(config.paths.src + '/module/', '../module/');
      return '@import "' + filePath + '";';
    },
    starttag: '//injector',
    endtag: '//endinjector',
    addRootSlash: false
  };
  return gulp.src([path.join(config.paths.src, 'app/index.scss')])
    .pipe($.inject(injectFiles, injectOptions))
    .pipe(gulp.dest(path.join(config.paths.src, 'app/')))
});
/**
 * [代码质量管理]
 */

gulp.task('build:jshint', function () {
  return gulp.src([
    path.join(config.paths.src, '/module/_default/**/*.js'),
    path.join(config.paths.src, 'app/**/*.js')
  ])
    .pipe($.plumber(config.errorHandler()))
    .pipe($.jshint())
    .pipe(reload({
      stream: true
    }))
    .pipe($.size())
    
});


/**
 * [SASS预编译模块,依赖compass模块编译]
 */

gulp.task('build:styles:compass', function () {
  return gulp.src(path.join(config.paths.src, 'app/index.scss'))
    .pipe($.plumber(config.errorHandler()))
    .pipe($.compass({
      config_file: path.join(__dirname, '/../config.rb'),
      css: path.join(config.paths.tmp, '/serve/app/'),
      sass: path.join(config.paths.src, '/app/'),
    }))
    //sprite图片路径修复
    .pipe($.replace('../../../src/assets/images/', '../assets/images/'))
    .pipe(gulp.dest(path.join(config.paths.tmp, '/serve/app/')))
    //css改变时无刷新改变页面
    .pipe(reload({
      stream: true
    }));
});

gulp.task('build:styles:theme:compass', function () {
  return gulp.src(path.join(config.paths.src, '/theme/*.scss'))
    .pipe($.plumber(config.errorHandler()))
    .pipe($.compass({
      config_file: path.join(__dirname, '/../config.rb'),
      css: path.join(config.paths.tmp, '/serve/app/theme/'),
      sass: path.join(config.paths.src, '/theme/')
    }))
    .pipe(gulp.dest(path.join(config.paths.dist, '/theme/')))
});

/**
 * [生成Html模版文件]
 */
gulp.task('build:tmpHtml:app', function () {
  return gulp.src([
          path.join(config.paths.src, '/app/**/*.tmp.html')
      ])
      .pipe($.minifyHtml({
          empty: true,
          spare: true,
          quotes: true
      }))
      .pipe($.angularTemplatecache('eoAppTmpHtml.js', {
          module: config.modules.templateModuleName,
          root: 'app'
      }))
      .pipe(gulp.dest(config.paths.tmp + '//serve/'));
});
gulp.task('build:tmpHtml:module', function () {
  return gulp.src([
          path.join(config.paths.src, '/module/**/*.tmp.html')
      ])
      .pipe($.minifyHtml({
          empty: true,
          spare: true,
          quotes: true
      }))
      .pipe($.angularTemplatecache('eoModuleTmpHtml.js', {
          module: config.modules.templateModuleName,
          root: 'module'
      }))
      .pipe(gulp.dest(config.paths.tmp + '//serve/'));
});
/**
 * [Html中的CSS以及JS注入]
 */

gulp.task('build:inject', ['build:jshint', 'build:styles:compass', 'vendor:base'], function () {
  var injectStyles = gulp.src([
    path.join(config.paths.tmp, '/serve/app/**/*.css'),
    path.join(config.paths.src, 'module/_default/**/*.scss'),
    '!' + path.join(config.paths.tmp, '/serve/app/theme/*.css')
  ], {
    read: false
  });

  var injectScripts = gulp.src([
    path.join(config.paths.src, '/app/**/*.js'),
    path.join(config.paths.src, '/module/_default/**/*.js'),
    path.join('!' + config.paths.src, '/app/vendor.js'),
  ])
  .pipe(babel({
    "presets": ["stage-3", "es2015"]
  }))
  .pipe($.angularFilesort())
  .pipe(gulp.dest(config.paths.tmp + '/serve/app/'))
  

  var injectOptions = {
    ignorePath: [config.paths.src, path.join(config.paths.tmp, '/serve')],
    addRootSlash: false
  };

  return gulp.src(path.join(config.paths.src, '/*.html'))
    .pipe($.plumber(config.errorHandler()))
    .pipe($.inject($.eventStream.merge(
      injectStyles,
      injectScripts
    ), injectOptions))
    .pipe(gulp.dest(path.join(config.paths.tmp, '/serve')));

});