#goku-gateway ce
##author 广州银云信息科技有限公司

* 开发目录结构要求：采取navbar、footer、header、content、sidebar树状结构
* js命名采取小驼峰模式（eg:eoNavbar）
* css class命名采取-拼接模式（eg:eo-navbar）id命名设计js需标明（eg:eo-navbar-js）
* 常量命名采用全大写及下划线拼接方式（eg:HTTP_CONSTANT）
* 任何服务及过滤器均采用开头大写小驼峰模式（eg:AuthInterceptor）
* 所有文件命名均采用小驼峰模式
* controller中不能设计dom操作
* 涉及交互尽量采用service

# 部署说明


## 运行

首先进入要存放文件的文件夹路径（建议存放目录不要有中文字符）

运行`npm install`安装运行依赖

运行`bower install`安装前端依赖

安装完毕后可运行

`gulp serve` --开发模式

`gulp build` --编译模式（将项目文件输出为上线文件）

`gulp serve:dist` --调试上线模式

### Gulp

目录结构

--gulp

* --base.js    配置基本项目依赖
* --build.js   生成上线版本脚本
* --config.js  全局配置文件
* --server.js  代理服务器/服务器启动脚本

详细配置见文件注释

## 框架目录

主要功能目录

* gulp 脚本管理
* src  项目源文件
* app.conf.json  Angular全部变量配置文件,通过[dev-config/prod-config任务编译]
* config.rb compass配置文件
* vendor.base.json 前端启动依赖文件(打包会随源文件一同压缩)
* vendor.json  前端依赖库文件(通过Lazyload模块加载)
* package.json 编译模块依赖文件以及项目配置--新增模块请注意加上 `npm install --save 新安装模块`
* bower.json 前端依赖库json文件 `bower install --save 新安装模块`


## 源文件

### app目录

主要文件

* assets 存放静态文件
* config 全局配置文件,包括路由配置模块`routes`,全局定义模块`core`,以及按需加载模块`lazyload`
* component 组建文件
* directive 指令模块,页面所有的指令文件写在这里,模块位置为`eo-help(项目名).directive`
* service 服务模块,页面所有的服务文件写在这里,模块位置为`eo-help(项目名).service`
* filter 过滤器模块,页面所有的过滤器文件写在这里,模块位置为`eo-help(项目名).filter`
* constant 存放常量文件
* resource Api配置模块,全局的Api配置位置(目前Api数量较少,不考虑多文件,后期可能会改为多文件)
* app.module.js 全局模块依赖声明模块,如无需全局依赖更改,不要随意改动该文件内容.
* app.conf.js 由app.conf.json编译而来的全局变量文件,配置当前开发模式DEV/PRODUCTION
* vendor.js 前端依赖js库文件,随index.html注入文档
* vendor.scss 前端依赖scss库文件,通过在index.scss中引入
* index.scss 全局的样式文件.
* ui view视图

