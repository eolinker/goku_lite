(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块backend相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.accessControl', {
                    url: '/accessControl',
                    template: '<gateway-control></gateway-control>'
                });
        }])
        .component('gatewayControl', {
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/content/accessControl/index.html',
            controller: gatewayControlController
        })

    gatewayControlController.$inject = ['$scope'];

    function gatewayControlController($scope) {
        var vm = this;
        
    }
})();
