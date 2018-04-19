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
                    template: '<gateway-overview></gateway-overview>',
                    resolve: helper.resolveFor('CLIPBOARD')
                });
        }])
        .component('gatewayOverview', {
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/content/overview/index.html',
            controller: gatewayOverviewController
        })

    gatewayOverviewController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope','PATH_INFO'];

    function gatewayOverviewController($scope, GatewayResource, $state, CODE, $rootScope,PATH_INFO) {

        var vm = this;
        vm.data = {
            info: {
                hash: ''
            },
            interaction: {
                request: {
                    gatewayHashKey: $state.params.gatewayHashKey
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


        vm.data.fun.edit = function() {
            var template={
                modal:{
                    title:'修改网关',
                    status:'edit',
                    request:vm.data.interaction.response.gatewayInfo
                }
            }
            vm.data.interaction.response.gatewayInfo.gatewayHashKey = vm.data.interaction.request.gatewayHashKey;
            $rootScope.Gateway_DefaultModal(template.modal, function(callback) {
                if (callback) {
                    vm.data.interaction.response.gatewayInfo.gatewayName = callback.gatewayName;
                    vm.data.interaction.response.gatewayInfo.gatewayDesc = callback.gatewayDesc;
                    $rootScope.InfoModal('修改网关成功', 'success');
                }
            });
        }
        vm.data.fun.init = (function() {
            var template = {
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey
                }
            }
            GatewayResource.Gateway.Info(template.request).$promise
                .then(function(response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.data.interaction.response.gatewayInfo = response.gatewayInfo||{};
                                vm.data.info.hash=PATH_INFO.INHERIT_HOST+response.gatewayInfo.gatewayPort+'/'+ $state.params.gatewayHashKey
                            }
                    }
                })
        })();
    }
})();
