'use strict';

var gulp = require('gulp'),
  path = require('path'),
  config = require('./config'),
  _ = require('lodash'),
  $ = require('gulp-load-plugins')({
    pattern: ['gulp-*', 'event-stream', 'main-bower-files', 'uglify-save-license', 'del']
  }),
  browserSync = require('browser-sync'),
  reload = browserSync.reload; //实时刷新
const babel = require('gulp-babel');
/**
 * [代码质量管理]
 */

gulp.task('jshint', function () {
  return gulp.src([
      path.join(config.paths.src, 'module/**/*.js'),
      path.join(config.paths.src, 'app/**/*.js'),
      // path.join(config.paths.src, '.tmp/**/*.js'),
      path.join('!' + config.paths.src, 'module/**/module.js')
    ])
    .pipe($.plumber(config.errorHandler()))
    .pipe($.jshint())
    // .pipe(reload({
    //   stream: true
    // }))
    .pipe($.size());
});

/**
 * [编译之前将scss注入index.scss]
 */

gulp.task('inject_sass', function () {

  /**
   * @description module inject
   */
  var injectFiles = gulp.src([
    path.join(config.paths.src, 'module/**/*.scss'),
    path.join('!' + config.paths.src, 'module/**/import.scss')
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

  /**
   * @description common inject
   */
  var tmpInjectCommonFiles = gulp.src([
    path.join(config.paths.src, 'app/**/*.scss'),
    path.join('!' + config.paths.src, 'app/*.scss'),
    path.join('!' + config.paths.src, 'app/**/import.scss'),
    path.join('!' + config.paths.src, 'app/scss/common/animation.scss')
  ], {
    read: false
  });

  var tmpInjectCommonOptions = {
    transform: function (filePath) {
      filePath = filePath.replace(config.paths.src + '/app/', './');
      return '@import "' + filePath + '";';
    },
    starttag: '//common-scss-injector',
    endtag: '//common-scss-end',
    addRootSlash: false
  };
  return gulp.src([
      path.join(config.paths.src, 'app/index.scss'),
      path.join(config.paths.src, 'module/index.scss')
    ])
    .pipe($.inject(injectFiles, injectOptions))
    .pipe($.inject(tmpInjectCommonFiles, tmpInjectCommonOptions))
    .pipe(gulp.dest(path.join(config.paths.src, 'app/')))
});
gulp.task('clean:tmp', function () {
  $.del([path.join(config.paths.tmp, '/')]);
});

gulp.task('clean', function () {
  $.del([path.join(config.paths.dist, '/'), path.join(config.paths.tmp, '/')]);
});

/**
 * [SASS预编译模块,依赖compass模块编译]
 */

gulp.task('styles:compass', ['inject_sass'], function () {
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

gulp.task('styles:theme:compass', function () {
  return gulp.src(path.join(config.paths.src, '/theme/*.scss'))
    .pipe($.plumber(config.errorHandler()))
    .pipe($.compass({
      config_file: path.join(__dirname, '/../config.rb'),
      css: path.join(config.paths.tmp, '/serve/app/theme/'),
      sass: path.join(config.paths.src, '/theme/')
    }))
    .pipe(gulp.dest(path.join(config.paths.tmp, '/serve/app/theme/')))
    //css改变时无刷新改变页面
    .pipe(reload({
      stream: true
    }));
});



/**
 * [Html中的CSS以及JS注入]
 */

gulp.task('inject', ['build:tmpHtml:app','build:tmpHtml:module','jshint', 'styles:compass','vendor:base'], function () {
  var injectStyles = gulp.src([
    path.join(config.paths.tmp, '/serve/app/**/*.css'),
    path.join(config.paths.tmp, '/serve/module/**/*.css'),
    '!' + path.join(config.paths.tmp, '/serve/app/theme/*.css')
  ], {
    read: false
  });

  var injectScripts = gulp.src([
      path.join(config.paths.src, '/module/**/*.js'),
      path.join(config.paths.src, '/app/**/*.js'),
      path.join(config.paths.tmp, '/serve/**/*TmpHtml.js'),
      // path.join('!' + config.paths.src, '/.tmp/vendor.js'),
      path.join('!' + config.paths.src, '/app/vendor.js'),
      path.join('!' + config.paths.src, 'module/**/module.js')
    ])
    .pipe(babel({
      "presets": ["stage-3", "es2015"]
    }))
    .pipe($.angularFilesort());

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