'use strict';

var gutil = require('gulp-util');// node require 动态加载gulp-util 

exports.paths = {//node export/module.exports(m.e优先级最高) 将函数、变量等导出，以使其它JavaScript脚本通过require()函数引入并使用
  src:  'src',
  dist: 'dist',
  tmp:  '.tmp',
  e2e:  'test_e2e',
  env:{

  }
};

exports.modules={
  ConstantModuleName:'goku',
  templateModuleName:'goku'
}

/**
 * [依赖配置]
 */
exports.vendor = {
  // 程序启动依赖模块
  base: {
    source: require('../vendor.base.json'),
    dest: 'src/app',
    name: 'vendor'
  },
  
  // 按需加载模块
  app: {
    source: require('../vendor.json'),
    dest: 'src/vendor'
  }
};

/**
 *  错误处理
 */
exports.errorHandler = function() {
  return function (err) {
    gutil.beep();/*# 发出滴声提示*/
    gutil.log(err.toString());/*# 输出错误信息*/
  }
};

