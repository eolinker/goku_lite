(function() {
    /*
     * author：riverLethe
     * 网关内页页内包模块overview相关js
     */
    'use strict';
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.overview', {
                    url: '/overview',
                    template: '<gateway-overview></gateway-overview>'
                });
        }])
        .component('gatewayOverview', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/overview/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope','PATH_INFO'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope,PATH_INFO) {

        var vm = this;
        vm.data = {
            info: {
                hash: ''
            },
            interaction: {
                request: {
                    gatewayAlias: $state.params.gatewayAlias
                },
                response: {
                    gatewayInfo: {}
                }
            },
            fun: {
                init: null, //初始化功能函数
                edit: null //编辑网关信息功能函数
            }
        }

        vm.data.fun.init = (function() {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias
                }
            }
            GatewayResource.Gateway.Info(template.request).$promise
                .then(function(response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.data.interaction.response.gatewayInfo = response.gatewayInfo||{};
                                vm.data.info.hash=PATH_INFO.INHERIT_HOST+response.gatewayInfo.gatewayPort+'/';
                            }
                    }
                })
        })();
    }
})();
