(function() {
    'use strict';
    /**
     * [懒加载相关常量]
     * @author [广州银云信息科技有限公司]
     * @description [按需加载文件存放常量集]
     */
    angular
        .module('goku.constant')
        /**
         * [路由预加载]
         * @type {constant}
         */
        .constant('APP_REQUIRES', {
            // jQuery based and standalone scripts
            SCRIPTS: {},
            // Angular based script (use the right module name)
            MODULES: [
                // options {serie: true,insertBefore: '#load_styles_before'}
                {
                    name: 'CLIPBOARD',
                    files: ['vendor/clipboard/dist/clipboard.min.js']
                }
            ]
        })
        /**
         * [html 懒加载]
         * @type {constant}
         */
        .constant('HTML_LAZYLOAD', [{
            name: 'PAGINATION',
            files: [
                "libs/pagination/pagination.js"
            ]
        }]);

})();
