(function () {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.monitor', {
                    url: '/monitor',
                    template: '<div ui-view></div>'
                })
                .state('home.monitor.global', {
                    url: '/',
                    template: '<panel></panel>'
                })
        }])
})();