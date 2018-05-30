(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关管理二级菜单网关列表页指令
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.default', {
                    url: '/',
                    template: '<home-gateway-default></home-gateway-default>'
                });
        }])
        .component('homeGatewayDefault', {
            templateUrl: 'app/ui/content/home/gateway/default/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', '$timeout', '$filter'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, $timeout, $filter) {
        var vm = this;
        vm.data = {
            interaction: {
                response: {
                    query: null
                }
            },
            fun: {
                enter: null, //进入内页功能函数
                edit: null, //编辑功能函数
                delete: null, //删除功能函数
                init: null //初始化功能函数
            }
        }
        vm.component = {
            listDefaultCommonObject: null,
            menuObject: {
                list: null
            }
        }
        vm.data.fun.init = function() {
            var template = {
                promise: null
            }
            template.promise = GatewayResource.Gateway.Query(template.request).$promise;
            template.promise.then(function(response) {
                vm.data.interaction.response.query = response.gatewayList || [];
            })
            return template.promise;
        }
        vm.data.fun.edit = function(arg) {
            arg = arg || {};
            var template = {
                modal: {
                    title: arg.item ? '修改网关' : '新增网关',
                    status:arg.item?'edit':'add',
                    request:arg.item
                },
                response: null
            }
            $rootScope.Gateway_DefaultModal(template.modal, function(callback) {
                if (callback) {
                    $rootScope.InfoModal(template.modal.title + '成功', 'success');
                    if (arg.item) {
                        arg.item.gatewayName = callback.gatewayName;
                        arg.item.gatewayAlias = callback.gatewayAlias;
                        arg.item.updateTime = $filter('currentTimeFilter')();
                        vm.data.interaction.response.query.splice(arg.$index,1,arg.item);
                    } else {
                        template.response = {
                            gatewayAlias: callback.gatewayAlias,
                            gatewayName: callback.gatewayName,
                            updateTime: $filter('currentTimeFilter')()
                        };
                        vm.data.interaction.response.query.splice(0,0,template.response);
                    }
                }
            });
        }
        vm.data.fun.delete = function(arg) {
            arg = arg || {};
            var template = {
                request: {
                    gatewayAlias: arg.item.gatewayAlias
                },
                GPEDIT_COMPONENT_TABLE:JSON.parse(window.localStorage['GPEDIT_COMPONENT_TABLE'] || '{}')
            }
            $rootScope.EnsureModal('删除网关', true, '确认删除？', {}, function(callback) {
                if (callback) {
                    GatewayResource.Gateway.Delete(template.request).$promise.then(function(response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    window.localStorage.setItem('GPEDIT_COMPONENT_TABLE', JSON.stringify(template.GPEDIT_COMPONENT_TABLE, function(key, val) {
                                        if (key === arg.item.gatewayAlias) {
                                            return undefined;
                                        }
                                        return val;
                                    }));
                                    vm.data.interaction.response.query.splice(arg.$index, 1);
                                    $rootScope.InfoModal('网关删除成功', 'success');
                                    break;
                                }
                        }
                    })
                }
            });
        }
        vm.data.fun.enter = function(arg) {
            $state.go('home.gateway.inside.overview', { gatewayName: arg.item.gatewayName, gatewayAlias: arg.item.gatewayAlias });
        }

        vm.$onInit = function() {
            vm.component.listDefaultCommonObject = {
                mainObject: {
                    item: {
                        default: [
                            { key: '网关名称', html: '{{item.gatewayName}}', keyStyle: { 'min-width': '400px' } },
                            { key: '状态', html: '{{item.gatewayStatus=="off"?\'下线\':\'启用\'}}', keyStyle: { 'width': '160px' } },
                            { key: '网关最后改动时间', keyStyle: { 'width': '280px' }, html: '{{item.updateTime}}', contentClass: 'unnecessary-td' }
                        ],
                        fun: {
                            array: [
                                { icon: 'bianji', key: '修改', fun: vm.data.fun.edit,show:-1 },
                                { icon: 'shanchu', key: '删除', fun: vm.data.fun.delete,show:-1 }
                            ],
                            keyStyle: { 'width': '250px' },
                            power:1
                        }
                    },
                    baseInfo: {
                        colspan: 5,
                        warning:'尚未新建任何接口网关'
                    },
                    baseFun: {
                        click: vm.data.fun.enter
                    }
                }
            }
            vm.component.menuObject.list = [{
                type: 'btn',
                class: 'btn-group-li pull-left',
                btnList: [{
                    name: '新建网关',
                    icon: 'tianjia',
                    class: 'eo-button-success',
                    fun: {
                        default: vm.data.fun.edit,
                        params: ''
                    }
                }]
            }]
        }
    }
})();
