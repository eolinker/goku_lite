(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块backend相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.accessControl.requestRateLimit', {
                    url: '/requestRateLimit',
                    template: '<gateway-control-rate></gateway-control-rate>'
                });
        }])
        .component('gatewayControlRate', {
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/content/accessControl/content/requestRateLimit/index.html',
            controller: gatewayControlRateController
        })

    gatewayControlRateController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', '$timeout'];

    function gatewayControlRateController($scope, GatewayResource, $state, CODE, $rootScope, $timeout) {
        var vm = this;
        vm.data = {
            interaction: {
                request: {
                    gatewayHashKey: $state.params.gatewayHashKey
                },
                response: {
                    query: []
                }
            },
            fun: {
                delete: null, //删除请求频率限制功能函数
                edit: null, //编辑请求频率限制功能函数
                init: null //初始化功能函数
            }
        }
        vm.component = {
            listDefaultCommonObject: null,
            menuObject: {
                list: null
            }
        }
        vm.data.fun.init = function(arg) {
            var template = {
                promise: null,
                request: { gatewayHashKey: vm.data.interaction.request.gatewayHashKey }
            }
            $scope.$emit('$windowTitle', { groupName: '访问控制' });
            template.promise = GatewayResource.Frequency.Query(template.request).$promise;
            template.promise.then(function(response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.data.interaction.response.query = response.frequencyList;
                            break;
                        }
                }
            })
            return template.promise;
        }
        vm.data.fun.edit = function(arg) {
            arg = arg || {};
            var template = {
                modal: {
                    title: arg.item ? '修改请求频率限制' : '新增请求频率限制',
                    array: [{ groupName: '1秒' }, { groupName: '1分钟' }],
                    item: arg.item ? { groupName: arg.item.count, $index: arg.item.intervalType } : null
                },
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey,
                    intervalType: '0',
                    count: '0'
                }
            }
            $rootScope.GatewayRateLimitModal(template.modal.title, template.modal.item, '请求频率限制', template.modal.array, function(callback) {
                if (callback) {
                    template.request.intervalType = callback.$index;
                    template.request.count = callback.groupName;
                    if (arg.item) {
                        GatewayResource.Frequency.Edit(template.request)
                            .$promise.then(function(response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS:
                                        {
                                            $rootScope.InfoModal('修改请求频率限制成功！', 'success');
                                            vm.data.fun.init();
                                            break;
                                        }
                                }
                            })
                    } else {
                        GatewayResource.Frequency.Add(template.request)
                            .$promise.then(function(response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS:
                                        {
                                            $rootScope.InfoModal('新增请求频率限制成功！', 'success');
                                            vm.data.fun.init();
                                            break;
                                        }
                                    case '180003':
                                        {
                                            $rootScope.InfoModal('请求频率限制，同一个单位时间只能新增一次！', 'error');
                                            break;
                                        }
                                }
                            })
                    }
                }
            });
        }
        vm.data.fun.delete = function(arg) {
            $rootScope.EnsureModal('删除请求频率限制', false, '确认删除', {}, function(data) {
                if (data) {
                    GatewayResource.Frequency.Delete({ gatewayHashKey: vm.data.interaction.request.gatewayHashKey, intervalType: arg.item.intervalType }).$promise
                        .then(function(response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS:
                                    {
                                        vm.data.interaction.response.query.splice(arg.$index, 1);
                                        $rootScope.InfoModal('请求频率限制删除成功', 'success');
                                        break;
                                    }
                                default:
                                    {
                                        $rootScope.InfoModal('删除失败，请稍候再试或到论坛提交bug', 'error');
                                        break;
                                    }
                            }
                        })
                }
            });
        }
        vm.$onInit = function() {
            $scope.$emit('$windowTitle', { groupName: '访问控制' });
            vm.component.listDefaultCommonObject = {
                mainObject: {
                    item: {
                        default: [
                            { key: '单位时间', html: '{{item.intervalType==0?\'1秒\':\'1分钟\'}}', keyStyle: { 'width': '200px' } },
                            { key: '限制请求次数', html: '{{item.count}}', keyStyle: { 'min-width': '400px' } },
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
                        colspan: 3,
                        warning:'尚未新建任何限制'
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
                    name: '添加限制',
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