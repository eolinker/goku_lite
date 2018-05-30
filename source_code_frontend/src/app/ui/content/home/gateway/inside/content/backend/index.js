(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块backend相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.backend', {
                    url: '/backend',
                    template: '<gateway-backend></gateway-backend>'
                });
        }])
        .component('gatewayBackend', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/backend/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', '$timeout', 'Cache_CommonService'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, $timeout, Cache_CommonService) {
        var vm = this;
        var code = CODE.COMMON.SUCCESS;
        vm.data = {
            interaction: {
                request: {
                    gatewayAlias: $state.params.gatewayAlias
                },
                response: {
                    query: []
                }
            },
            fun: {
                edit: null, //编辑功能函数
                delete: null, //删除功能函数
                init: null //初始化功能函数
            }
        }
        vm.service={
            cache:Cache_CommonService
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
                cache: vm.service.cache.get('backend'),
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias
                }
            }
            if (template.cache) {
                vm.data.interaction.response.query = template.cache;
            } else {
                template.promise = GatewayResource.Backend.Query(template.request).$promise;
                template.promise.then(function(data) {
                    if (code == data.statusCode) {
                        vm.data.interaction.response.query = data.backendList;
                    }
                })
            }
            return template.promise;
        }
        vm.data.fun.edit = function(arg) {
            arg = arg || {};
            var template = {
                modal: {
                    title: arg.item ? '修改后端服务' : '新增后端服务'
                }
            }
            $rootScope.GatewayBackendModal(template.modal.title, arg.item, function(callback) {
                if (callback) {
                    callback.gatewayAlias = vm.data.interaction.request.gatewayAlias;
                    if (arg.item) {
                        GatewayResource.Backend.Edit(callback).$promise
                            .then(function(response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS:
                                        {
                                            $rootScope.InfoModal(template.modal.title + '成功', 'success');
                                            arg.item.backendName = callback.backendName;
                                            arg.item.backendPath = callback.backendPath;
                                            vm.data.interaction.response.query.splice(arg.$index,1,arg.item);
                                            break;
                                        }
                                }
                            })
                    } else {
                        GatewayResource.Backend.Add(callback).$promise
                            .then(function(response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS:
                                        {
                                            $rootScope.InfoModal(template.modal.title + '成功', 'success');
                                            callback.backendID = response.backendID;
                                            vm.data.interaction.response.query.splice(0, 0, callback);
                                            break;
                                        }
                                }
                            })
                    }
                }
            });
        }
        vm.data.fun.delete = function(arg) {
            arg = arg || {};
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    backendID: arg.item.backendID
                }
            }
            $rootScope.EnsureModal('删除后端服务', false, '确认删除', {}, function(callback) {
                if (callback) {
                    GatewayResource.Backend.Delete(template.request).$promise
                        .then(function(response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS:
                                    {
                                        vm.data.interaction.response.query.splice(arg.$index, 1);
                                        $rootScope.InfoModal('后端服务删除成功', 'success');
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
            $scope.$emit('$windowTitle', { groupName: '后端管理' });
            vm.component.listDefaultCommonObject = {
                mainObject: {
                    item: {
                        default: [
                            { key: '后端名称', html: '{{item.backendName}}', keyStyle: { 'width': '200px' } },
                            { key: '后端域名/IP地址', html: '{{item.backendPath}}', keyStyle: { 'min-width': '400px' } },
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
                        warning:'尚未新建任何后端服务'
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
                    name: '添加后端',
                    icon: 'tianjia',
                    class: 'eo-button-success',
                    fun: {
                        default: vm.data.fun.edit,
                        params: ''
                    }
                }]
            }]
        }
        $scope.$on('$destroy', function() {
            if (vm.data.interaction.response.query.length > 0) {
                vm.service.cache.set(vm.data.interaction.response.query,'backend');
            } else {
                vm.service.cache.clear('backend');
            }
        });
    }
})();
