/**=========================================================
 * Module: config.js
 * App routes and resources configuration
 =========================================================*/


(function() {
    'use strict';

    angular
        .module('eolinker')
        .config(routesConfig)
        .run(routesRun);

    routesConfig.$inject = ['$stateProvider', '$httpProvider', '$locationProvider', '$urlRouterProvider', 'RouteHelpersProvider'];

    function routesConfig($stateProvider, $httpProvider, $locationProvider, $urlRouterProvider, helper) {
        var data = {
            fun: {
                init: null //初始化功能函数
            }
        }
        data.fun.init = function() {
            $httpProvider.interceptors.push([
                '$injector',
                function($injector) {
                    return $injector.get('AuthInterceptor');
                }
            ]);
            $locationProvider.html5Mode(false).hashPrefix('');
            $urlRouterProvider.otherwise('/');
        }
        data.fun.init();
    }

    routesRun.$inject = [];

    function routesRun() {

    }
})();
