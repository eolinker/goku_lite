(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * ip黑白名单相关js
     */
    angular.module('goku')
        .component('gatewayGpeditIp', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/authority/gpedit/ip/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', '$rootScope', '$state', 'CODE', 'GatewayResource', '$timeout'];

    function indexController($scope, $rootScope, $state, CODE, GatewayResource, $timeout) {
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
                show:{
                    status:{
                        isEdit:false
                    }
                },
                list: null
            }
        }
        vm.data.fun.edit = function () {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    strategyID: vm.data.interaction.request.strategyID,
                    ipLimitType: vm.data.interaction.response.ipList.ipLimitType,
                    ipBlackList: vm.data.interaction.response.ipList.ipBlackList,
                    ipWhiteList:vm.data.interaction.response.ipList.ipWhiteList
                }
            }
            GatewayResource.Ip.Edit(template.request).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                    case '170004':
                        {
                            $rootScope.InfoModal('保存黑白名单成功！', 'success');
                            vm.component.menuObject.show.status.isEdit = false;
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
        vm.data.fun.cancle = function () {
            vm.component.menuObject.show.status.isEdit = false;
            vm.data.interaction.response.ipList=angular.copy(vm.data.info.template.request);
        }
        vm.data.fun.init = (function () {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    strategyID: vm.data.interaction.request.strategyID
                }
            }
            GatewayResource.Ip.Info(template.request).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.data.interaction.response.ipList = response.ipList;
                            vm.data.info.template.request=angular.copy(response.ipList);
                            break;
                        }
                }
            })
        })();

        vm.$onInit = function () {
            vm.component.menuObject.list = [{
                type: 'btn',
                class: 'btn-group-li pull-left',
                showVariable: 'isEdit',
                showPoint: 'status',
                btnList: [{
                    name: '策略组列表',
                    icon: 'xiangzuo',
                    show: -1,
                    fun: {
                        default: function () {
                            $state.go('home.gateway.inside.gpedit.default', {
                                strategyID: null
                            });
                        }
                    }
                }, {
                    name: '编辑',
                    class: 'eo-button-success',
                    show: false,
                    fun: {
                        default: function(){
                            vm.component.menuObject.show.status.isEdit=true;
                        }
                    }
                }, {
                    name: '保存',
                    class: 'eo-button-success',
                    show: true,
                    fun: {
                        default: vm.data.fun.edit
                    }
                }, {
                    name: '取消',
                    class: 'eo-button-default',
                    show: true,
                    fun: {
                        default: vm.data.fun.cancle
                    }
                }]
            }]
        }
    }
})();