(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 环境管理相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.env', {
                    url: '/env',
                    template: '<gateway-env></gateway-env>'
                });
        }])
        .component('gatewayEnv', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/env/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', '$rootScope', '$state', 'CODE', 'GatewayResource'];

    function indexController($scope, $rootScope, $state, CODE, GatewayResource) {
        var vm = this;
        vm.data = {
            info: {
                menu: [{
                    key: '不设置',
                    keyID: 'none'
                }, {
                    key: '黑名单',
                    keyID: 'black'
                }, {
                    key: '白名单',
                    keyID: 'white'
                }],
                template: {
                    request: {
                        ipLimitType: 'none',
                        ipBlackList: '',
                        ipWhiteList: ''
                    }
                }
            },
            interaction: {
                request: {
                    strategyID: $state.params.strategyID,
                    gatewayAlias: $state.params.gatewayAlias
                },
                response: {
                    ipList: {
                        ipLimitType: 'none',
                        ipBlackList: '',
                        ipWhiteList: ''
                    }
                }
            },
            fun: {
                cancle: null, //取消功能函数
                edit: null, //编辑功能函数
                init: null //初始化功能函数
            }
        }
        vm.component = {
            menuObject: {
                show: {
                    status: {
                        isEdit: false
                    }
                },
                list: null
            }
        }
        vm.data.fun.edit = function () {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    ipLimitType: vm.data.interaction.response.ipList.ipLimitType,
                    ipBlackList: vm.data.interaction.response.ipList.ipBlackList,
                    ipWhiteList: vm.data.interaction.response.ipList.ipWhiteList
                }
            }
            GatewayResource.Ip.GlobalEdit(template.request).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                    case '170004':
                        {
                            $rootScope.InfoModal('保存黑白名单成功！', 'success');
                            vm.data.info.status.ip.isEdit = false;
                            break;
                        }
                    default:
                        {
                            $rootScope.InfoModal('保存黑白名单操作失败！', 'error');
                        }
                }
            })
        }
        vm.data.fun.cancle = function () {
            vm.data.info.status.ip.isEdit = false;
            vm.data.interaction.response.ipList = angular.copy(vm.data.info.template.request);
        }
        vm.data.fun.init = (function () {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias
                }
            }
            GatewayResource.Ip.GlobalInfo(template.request).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.data.interaction.response.ipList = response.ipList;
                            vm.data.info.template.request = angular.copy(response.ipList);
                            break;
                        }
                }
            })
        })();
    }
})();