(function () {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.cluster', {
                    url: '/cluster',
                    template: '<div ui-view></div>',
                    containerRouter: true
                })
                .state('home.cluster.default', {
                    url: '/',
                    template: '<cluster-default></cluster-default>'
                })
                .state('home.cluster.node', {
                    url: '/node?clusterName',
                    template: '<cluster-node></cluster-node>',
                    needExtraNav:true
                })
                .state('home.cluster.node.default', {
                    url: '/:cluster?groupID',
                    template: '<cluster-node-default class="mt50 dp_b" group-arr="$ctrl.groupArr"></cluster-node-default>',
                    needExtraNav:true
                })
        }])
})();