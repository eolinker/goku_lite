(function() {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gpedit', {
                    url: '/gpedit',
                    template: '<div ui-view></div>',
                    containerRouter:true
                })
                .state('home.gpedit.default', {
                    url: '/overview',
                    template: '<gpedit-overview></gpedit-overview>'
                })
                .state('home.gpedit.common', {
                    url: '/list',
                    template: '<gpedit-common></gpedit-common>'
                })
                .state('home.gpedit.common.list', {
                    url: '/?groupID',
                    template: '<gpedit-default class="mt50 dp_b"></gpedit-default>'
                })
                .state('home.gpedit.inside', {
                    url: '/inside/:groupType?strategyID?groupID?strategyName',
                    template: '<gepdit-inside></gepdit-inside>',
                    containerRouter:true
                })
                .state('home.gpedit.inside.setting', {
                    url: '/setting',
                    template: '<gpedit-setting></gpedit-setting>'
                })
                .state('home.gpedit.inside.api', {
                    url: '/api',
                    template: '<div ui-view></div>',
                    containerRouter:true
                })
                .state('home.gpedit.inside.api.default', {
                    url: '/',
                    template: '<gpedit-inside-api-default></gpedit-inside-api-default>'
                })
                .state('home.gpedit.inside.api.operate', {
                    url: '/operate',
                    template: '<gpedit-inside-api-operate></gpedit-inside-api-operate>'
                })
                .state('home.gpedit.inside.api.plugin', {
                    url: '/plugin/:apiID',
                    template: '<gpedit-inside-api-plugin></gpedit-inside-api-plugin>',
                    resolve: helper.resolveFor( 'ACE_EDITOR')
                })
                .state('home.gpedit.inside.auth', {
                    url: '/auth',
                    template: '<gpedit-inside-auth></gpedit-inside-auth>'
                })
                .state('home.gpedit.inside.plugin', {
                    url: '/plugin',
                    template: '<div ui-view></div>',
                    containerRouter:true
                })
                .state('home.gpedit.inside.plugin.gpedit', {
                    url: '/gpedit',
                    template: '<gpedit-inside-plugin-gpedit></gpedit-inside-plugin-gpedit>',
                    sidebarIndex:1
                })
                .state('home.gpedit.inside.plugin.operate', {
                    url: '/operate/:status?pluginName?chineseName',
                    template: '<gpedit-inside-plugin-operate opr-target="gpedit"></gpedit-inside-plugin-operate>',
                    resolve: helper.resolveFor( 'ACE_EDITOR')
                });
        }])
})();
