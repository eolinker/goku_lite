var gulp = require('gulp'),
    argv = require('minimist'),
    path = require('path'),
    config = require('./config'),
    $ = require('gulp-load-plugins')({
        pattern: ['gulp-*', 'event-stream', 'main-bower-files', 'uglify-save-license', 'del']
    }),
    gulpsync = $.sync(gulp),
    sass = require('gulp-sass'),
    replace = require('gulp-replace');
let eoSetPointSrc = () => {
    let pointLangIsEn = argv(process.argv.slice(2)).en;
    if(pointLangIsEn){
        config.paths.src = '.tmp/enApp';
        config.vendor.base.dest='.tmp/enApp/app'
    }
}
gulp.task('dev-config', function () {
    eoSetPointSrc();
    return gulp.src('app.conf.json')
        .pipe($.ngConfig(config.modules.ConstantModuleName, {
            environment: 'development',
            createModule: false,
            wrap: true
        }))
        .pipe(gulp.dest(path.join(config.paths.src, '/app'))) //join() 方法用于把数组中的所有元素放入一个字符串。
});
gulp.task('prod-config', function () {
    eoSetPointSrc();
    return gulp.src('app.conf.json')
        .pipe($.ngConfig(config.modules.ConstantModuleName, {
            environment: 'production',
            createModule: false,
            wrap: true //生成闭包
        }))
        .pipe(gulp.dest(path.join(config.paths.src, '/app')))
});
gulp.task('test-config', function () {
    eoSetPointSrc();
    return gulp.src('app.conf.json')
        .pipe($.ngConfig(config.modules.ConstantModuleName, {
            environment: 'test',
            createModule: false,
            wrap: true //生成闭包
        }))
        .pipe(gulp.dest(path.join(config.paths.src, '/app')))
});
gulp.task('clean:dist', function () {
    $.del([path.join(config.paths.dist, '/')]);
});

/**
 * [清理DIST,TEMP文件夹]
 */

gulp.task('clean', function () {
    $.del([path.join(config.paths.dist, '/'), path.join(config.paths.tmp, '/')]);
});

gulp.task('vendor', gulpsync.sync(['vendor:base']));

/**
 * [复制依赖文件]
 */

gulp.task('vendor:base', function () {
    var jsFilter = $.filter('**/*.js', {
            restore: true
        }),
        cssFilter = $.filter('**/*.css', {
            restore: true
        });
    return gulp.src(config.vendor.base.source)
        .pipe($.expectFile(config.vendor.base.source))
        .pipe(jsFilter)
        .pipe($.concat(config.vendor.base.name + '.js'))
        .pipe(jsFilter.restore)
        .pipe(cssFilter)
        .pipe($.concat(config.vendor.base.name + '.scss'))
        .pipe(cssFilter.restore)
        .pipe(gulp.dest(config.vendor.base.dest));
});

gulp.task('vendor:app', function () {

    var jsFilter = $.filter('*.js', {
            restore: true
        }),
        cssFilter = $.filter('*.css', {
            restore: true
        });

    return gulp.src(config.vendor.app.source, {
            base: 'bower_components'
        })
        .pipe($.expectFile(config.vendor.app.source))
        .pipe(jsFilter)
        .pipe(jsFilter.restore)
        .pipe(cssFilter)
        .pipe(cssFilter.restore)
        .pipe(gulp.dest(config.vendor.app.dest));

});

/**
 * [图片压缩]
 */
gulp.task('images', function () {
    return gulp.src([
            path.join(config.paths.src, '/assets/images/**/*'),
            path.join('!' + config.paths.src, '/assets/images/sprite/**/*')
        ])
        .pipe($.imagemin({
            progressive: true,
            svgoPlugins: [{
                removeViewBox: false
            }],
            use: [$.imageminPngquant()]
        }))
        .pipe(gulp.dest(path.join(config.paths.dist, '/assets/images')));
});

gulp.task('fonts', function () {

    return gulp.src(config.vendor.base.source, {
            base: 'bower_components'
        })
        .pipe($.filter('**/*.{eot,svg,ttf,woff,woff2}'))
        .pipe($.flatten())
        .pipe(gulp.dest(path.join(config.paths.dist, '/fonts/')));
});

/**
 * [复制文件] 前端依赖库以及静态文件
 */
gulp.task('other:vendor', function () {
    return gulp.src([
            path.join(config.paths.src, '/vendor/**/*'),
        ])
        .pipe($.filter(function (file) {
            return file.stat.isFile();
        }))
        .pipe(gulp.dest(path.join(config.paths.dist, '/vendor')));
});
gulp.task('other:libs:js', function () {
    return gulp.src([
            path.join(config.paths.src, '/libs/**/*.js')
        ])
        .pipe($.filter(function (file) {
            return file.stat.isFile();
        }))
        // .pipe($.babel())
        .pipe($.stripDebug())
        .pipe($.uglify())
        .pipe(gulp.dest(path.join(config.paths.dist, '/libs')));
});
gulp.task('other:libs', ['other:libs:js'], function () {
    return gulp.src([
            path.join(config.paths.src, '/libs/**/*'),
            path.join('!' + config.paths.src, '/libs/**/*.js')
        ])
        .pipe($.filter(function (file) {
            return file.stat.isFile();
        }))
        .pipe(gulp.dest(path.join(config.paths.dist, '/libs')));
});
gulp.task('other:assets', function () {
    return gulp.src([
            path.join(config.paths.src, '/app/assets/**/*')
        ])
        .pipe($.filter(function (file) {
            return file.stat.isFile();
        }))
        .pipe(gulp.dest(path.join(config.paths.dist, '/assets')));
});


gulp.task('version:scar', function () {
    gulp.src([
            path.join(config.paths.src, '../version/*.md'),
        ])
        .pipe(replace(/(position\s)([^\u4e00-\u9fa5].*)/g, function (match) {
            return match.replace(/position\s/, "position " + config.version.scar).replace(/\\/g, "/");
        }))
        .pipe(gulp.dest(path.join(config.paths.version, '/file')));
});
gulp.task('version:lethe', function () {
    gulp.src([
            path.join(config.paths.src, '../version/*.md'),
        ])
        .pipe(replace(/(position\s)([^\u4e00-\u9fa5].*)/g, function (match) {
            return match.replace(/position\s/, "position " + config.version.lethe).replace(/\\/g, "/");
        }))
        .pipe(gulp.dest(path.join(config.paths.version, '/file')));
});