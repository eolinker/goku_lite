(function() {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.plugin', {
                    url: '/plugin',
                    template: '<div ui-view></div>',
                    containerRouter:true
                })
                .state('home.plugin.default', {
                    url: '/',
                    template: '<plugin-default></plugin-default>'
                })
                .state('home.plugin.operate', {
                    url: '/operate/:status?pluginName',
                    template: '<plugin-operate></plugin-operate>',
                    resolve: helper.resolveFor( 'ACE_EDITOR')
                });
        }])
})();
