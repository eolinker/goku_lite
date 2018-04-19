(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块backend相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.accessControl.ipFirewall', {
                    url: '/ipFirewall',
                    template: '<gateway-control-list></gateway-control-list>'
                });
        }])
        .component('gatewayControlList', {
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/content/accessControl/content/ipFirewall/index.html',
            controller: gatewayControlListController
        })

    gatewayControlListController.$inject = ['$scope', '$rootScope', '$state', 'CODE', 'GatewayResource', '$timeout'];

    function gatewayControlListController($scope, $rootScope, $state, CODE, GatewayResource, $timeout) {
        var vm = this;
        vm.data = {
            info: {
                menu: [{ key: '不设置', keyID: 0 }, { key: '黑名单', keyID: 1 }, { key: '白名单', keyID: 2 }],
                template: {
                    preMenu: '0',
                    request:{
                        chooseType:'0',
                        blackList:'',
                        whiteList:''
                    }
                }
            },
            interaction: {
                request: {
                    gatewayHashKey: $state.params.gatewayHashKey,
                    ipList: ''
                },
                response: {
                    ipInfo: {
                        chooseType: '0',
                        blackList: '',
                        whiteList: ''
                    }
                }
            },
            fun: {
                cancle:null,//取消功能函数
                changeSwitch: null, //单选按钮change功能函数
                edit: null, //编辑功能函数
                init: null //初始化功能函数
            }
        }
        vm.data.fun.edit = function() {
            var template = {
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey,
                    chooseType: vm.data.info.template.request.chooseType,
                    ipList: vm.data.interaction.request.ipList
                }
            }
            GatewayResource.Ip.Edit(template.request).$promise.then(function(response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                    case '170004':
                        {
                            $rootScope.InfoModal('保存黑白名单成功！', 'success');
                            vm.data.info.isEdit = false;
                            break;
                        }
                    default:
                        {
                            $rootScope.InfoModal('保存黑白名单操作失败！', 'error');
                            console.log('操作失败!');
                        }
                }
            })
        }
        vm.data.fun.cancle = function() {
            vm.data.info.isEdit = false;
            angular.copy(vm.data.interaction.response.ipInfo,vm.data.info.template.request);
            vm.data.info.template.preMenu='0';
            vm.data.fun.changeSwitch();
        }
        vm.data.fun.changeSwitch = function() {
            switch (vm.data.info.template.preMenu) {
                case '1':
                    {
                        vm.data.info.template.request.blackList = vm.data.interaction.request.ipList;
                        break;
                    }
                case '2':
                    {
                        vm.data.info.template.request.whiteList = vm.data.interaction.request.ipList;
                        break;
                    }
            }
            vm.data.info.template.preMenu = vm.data.info.template.request.chooseType;
            switch (vm.data.info.template.request.chooseType) {
                case '0':
                    {
                        vm.data.interaction.request.ipList = '';
                        break;
                    }
                case '1':
                    {
                        vm.data.interaction.request.ipList = vm.data.info.template.request.blackList;
                        break;
                    }
                case '2':
                    {
                        vm.data.interaction.request.ipList = vm.data.info.template.request.whiteList;
                        break;
                    }
            }
        }
        vm.data.fun.init = (function() {
            var template = {
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey
                }
            }
            GatewayResource.Ip.Info(template.request).$promise.then(function(response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            response.ipInfo.chooseType=response.ipInfo.chooseType.toString();
                            vm.data.interaction.response.ipInfo = response.ipInfo;
                            angular.copy(response.ipInfo,vm.data.info.template.request);
                            vm.data.fun.changeSwitch();
                            break;
                        }
                }
            })
        })();
    }
})();
