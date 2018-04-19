'use strict';

var gulp = require('gulp');
var fs = require('fs');

fs.readdirSync('./gulp').forEach(function (file) {
	if((/\.(js|coffee)$/i).test(file)){
		require('./gulp/' + file);
	}
});

gulp.task('default', ['clean'], function () {
  gulp.start('build');
});