(function() {
    'use strict';
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway', {
                    url: '/gateway',
                    template: '<home-gateway></home-gateway>'
                });
        }])
        .component('homeGateway', {
            templateUrl: 'app/ui/content/home/content/gateway/index.html',
            controller: homeGatewayController
        })

    homeGatewayController.$inject = [];

    function homeGatewayController() {
        var vm = this;
    }
})();
