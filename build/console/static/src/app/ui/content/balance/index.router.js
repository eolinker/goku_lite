(function() {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.balance', {
                    url: '/balance',
                    template: '<div ui-view></div>',
                    containerRouter:true
                })
                .state('home.balance.service', {
                    url: '/service',
                    template: '<balance-service></balance-service>'
                })
                .state('home.balance.list', {
                    url: '/list',
                    template: '<div ui-view></div>'
                })
                .state('home.balance.list.default', {
                    url: '/',
                    template: '<balance-list group-arr="$ctrl.balanceGroupArr"></balance-list>'
                })
                .state('home.balance.list.operate', {
                    url: '/:status?balanceName',
                    template: '<balance-operate group-arr="$ctrl.balanceGroupArr"></balance-operate>',
                    resolve: helper.resolveFor( 'ACE_EDITOR')
                });
        }])
})();
