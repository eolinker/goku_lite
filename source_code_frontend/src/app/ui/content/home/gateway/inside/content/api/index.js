(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.api', {
                    url: '/api',
                    template: '<gateway-api></gateway-api>'
                });
        }])
        .component('gatewayApi', {
            template: '<gateway-api-sidebar></gateway-api-sidebar><gpedit-gateway-component ng-show="$ctrl.data.status!=-1"></gpedit-gateway-component><div ui-view></div>',
            controller: indexController
        })

    indexController.$inject = ['$scope', '$state'];

    function indexController($scope, $state) {
        var vm = this;
        vm.data = {
            status:0
        }
        var assistantFun={}
        assistantFun.init=function(){
            switch ($state.current.name) {
                case 'home.gateway.inside.api.edit':
                    {
                        vm.data.status = -1;
                        break;
                    }
                default:
                    {
                        vm.data.status=0;
                        break;
                    }
            }
        }
        $scope.$on('$stateChangeSuccess', function() { //路由转换函数，检测是否该显示环境变量
            assistantFun.init();
        });
        vm.$onInit = (function() {
            assistantFun.init();
        })()
    }
})();
