(function() {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.project', {
                    url: '/project',
                    template: '<div ui-view></div>',
                    containerRouter:true
                })
                .state('home.project.default', {
                    url: '/',
                    template: '<project-default></project-default>'
                })
                .state('home.project.api', {
                    url: '/api?projectID?projectName',
                    template: '<api></api>',
                    containerRouter:true
                })
                .state('home.project.api.default', {
                    url: '/?groupID',
                    template: '<api-default class="mt50 dp_b"></api-default>',
                    needExtraNav:true
                })
                .state('home.project.api.operate', {
                    url: '/operate/:status?apiID?groupID',
                    template: '<api-operate></api-operate>',
                    needExtraNav:true
                });
        }])
})();
