(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页外包模块相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside', {
                    url: '/inside?gatewayName?gatewayAlias',
                    template: '<gateway></gateway>'
                });
        }])
        .component('gateway', {
            templateUrl: 'app/ui/content/home/gateway/inside/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', '$state',  'Cache_CommonService','GroupService'];

    function indexController($scope, $state, Cache_CommonService,GroupService) {

        var vm = this;
        vm.data = {
            info: {
                shrinkObject: {}
            }
        }
        vm.service = {
            cache: Cache_CommonService,
            group:GroupService
        }
        vm.$onInit=function() {
            vm.service.cache.clear('backend');
            vm.service.group.clear();
        }
    }
})();
