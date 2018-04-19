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
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/content/api/index.html',
            controller: gatewayApiController
        })

    gatewayApiController.$inject = [];

    function gatewayApiController() {
        var vm = this;
        
    }
})();
