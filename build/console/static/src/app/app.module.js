(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 全局定义app模块js
     */
    angular.module('eolinker', [
        //thrid part
        'ui.router',
        'oc.lazyLoad',
        'ngResource',
        'angular-md5',
        //custom part
        'eolinker.constant',
        'eolinker.resource',
        'eolinker.modal',
        'eolinker.filter',
        'eolinker.directive',
        'eolinker.service'
    ])

    .config(AppConfig)

    .run(AppRun);


    AppConfig.$inject = ['$qProvider','$controllerProvider', '$compileProvider', '$filterProvider', '$provide', '$logProvider', '$stateProvider', '$urlRouterProvider', '$locationProvider', '$httpProvider', 'isDebug'];


    function AppConfig($qProvider,$controllerProvider, $compileProvider, $filterProvider, $provide, $logProvider, $stateProvider, $urlRouterProvider, $locationProvider, $httpProvider, IsDebug) {

        var data = {
            fun: {
                init: null, //初始化功能函数
                param: null, //解析请求参数格式功能函数
            }
        }
        data.fun.param = function(arg) {
            var query = '',
                name, value, fullSubName, subName, subValue, innerObj, i;

            for (name in arg.object) {
                value = arg.object[name];

                if (value instanceof Array) {
                    for (i = 0; i < value.length; ++i) {
                        subValue = value[i];
                        fullSubName = name + '[' + i + ']';
                        innerObj = {};
                        innerObj[fullSubName] = subValue;
                        query += data.fun.param({ object: innerObj }) + '&';
                    }
                } else if (value instanceof Object) {
                    for (subName in value) {
                        subValue = value[subName];
                        fullSubName = name + '[' + subName + ']';
                        innerObj = {};
                        innerObj[fullSubName] = subValue;
                        query += data.fun.param({ object: innerObj }) + '&';
                    }
                } else if (value !== undefined && value !== null)
                    query += encodeURIComponent(name) + '=' + encodeURIComponent(value) + '&';
            }

            return query.length ? query.substr(0, query.length - 1) : query;
        };
        data.fun.init = (function() {
            $httpProvider.defaults.headers.put['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';
            $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';
            // Override $http service's default transformRequest
            $httpProvider.defaults.transformRequest = [function(callback) {
                return angular.isObject(callback) && String(callback) !== '[object File]' ? data.fun.param({ object: callback }) : callback;
            }];
            // Enable log
            $logProvider.debugEnabled(IsDebug);
            $urlRouterProvider.otherwise('/');
            $qProvider.errorOnUnhandledRejections(false);//注销$q.reject抛出错误配置
        })();
    }

    AppRun.$inject = ['$rootScope', '$state', '$stateParams', '$window', '$templateCache', '$http'];


    function AppRun($rootScope, $state, $stateParams, $window, $templateCache, $http) {

        $rootScope.$state = $state;
        $rootScope.$stateParams = $stateParams;
        window.$eo=window.$eo||{
            directive:{}
        };
    }
})();
