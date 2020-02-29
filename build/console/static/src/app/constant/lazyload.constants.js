(function () {
    'use strict';
    /**
     * [懒加载相关常量]
     * @author [广州银云信息科技有限公司]
     * @description [按需加载文件存放常量集]
     */
    angular
        .module('eolinker.constant')
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
                }, {
                    name: 'ZEPTO',
                    files: [
                        "vendor/zepto/zepto.min.js"
                    ]
                },  {
                    name: 'ACE_EDITOR_AUTOCOMPLETE',
                    files: [
                        "libs/ace-builds/src/ext-language_tools.js"
                    ]
                }, {
                    name: 'ACE_EDITOR',
                    files: [
                        "libs/ace-builds/src/ace.js"
                    ]
                },{
                    name: 'DATEPICKER',
                    files: [
                        "libs/datepicker/lib/position.js",
                        "libs/datepicker/lib/dateparser.js",
                        "libs/datepicker/lib/datepicker.js",
                        "libs/datepicker/index.js"

                    ]
                }
            ]
        });

})();