(function() {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.alert', {
                    url: '/alert',
                    template: '<div ui-view></div>',
                    containerRouter:true
                })
                .state('home.alert.list', {
                    url: '/list',
                    template: '<alert-list></alert-list>'
                })
                .state('home.alert.setting', {
                    url: '/',
                    template: '<alert-setting></alert-setting>'
                });
        }])
})();
