(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api详情相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.api.detail', {
                    url: '/detail?apiID?groupID',
                    template: '<gateway-api-detail></gateway-api-detail>'
                });
        }])
        .component('gatewayApiDetail', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/api/detail/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', '$timeout', 'GatewayResource', '$state', 'CODE', '$rootScope', 'Cache_CommonService','PATH_INFO','Communicate_CommonService'];

    function indexController($scope, $timeout, GatewayResource, $state, CODE, $rootScope, Cache_CommonService,PATH_INFO,Communicate_CommonService) {
        var vm = this;
        vm.data = {
            info:{},
            interaction: {
                response: {
                    apiInfo: {}
                },
                request: {
                    apiID: $state.params.apiID,
                    gatewayAlias: $state.params.gatewayAlias,
                    groupID: $state.params.groupID
                }
            },
            fun: {
                init: null //初始化功能函数
            }
        }
        vm.service={
            cache:Cache_CommonService,
            communicate:Communicate_CommonService
        }
        vm.component = {
            menuObject: {
                list: null
            }
        }
        var data={
            fun:{}
        }
        data.fun.menu = function(status, uri, cache) {
            var template = {
                uri: {
                    groupID: uri.groupID,
                    apiID: uri.apiID,
                    status:status
                }
            }
            switch (status) {
                case 'list':
                    {
                        $state.go('home.gateway.inside.api.list', template.uri);
                        break;
                    }
                case 'simple':
                    {
                        $state.go('home.gateway.inside.api.simple', template.uri);
                        break;
                    }
                case 'edit':
                    {
                        $state.go('home.gateway.inside.api.edit', template.uri);
                        break;
                    }
                case 'copy':
                    {
                        $state.go('home.gateway.inside.api.edit', template.uri);
                        break;
                    }
            }
        }
        vm.data.fun.init = function(arg) {
            var template = {
                promise: null,
                cache: vm.service.cache.get('apiInfo'),
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    groupID:vm.data.interaction.request.groupID,
                    apiID: vm.data.interaction.request.apiID
                }
            }
            if (template.cache) {
                vm.data.interaction.response.apiInfo = template.cache;
                vm.data.info.gatewayUrl=PATH_INFO.INHERIT_HOST+vm.service.cache.get('gatewayInfo').gatewayPort+'/'+vm.data.interaction.request.gatewayAlias;
            } else {
                template.promise = GatewayResource.Api.Detail(template.request).$promise;
                template.promise.then(function(response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                response.apiInfo.requestMethod=response.apiInfo.requestMethod.split(',');
                                vm.data.interaction.response.apiInfo = response.apiInfo;
                                vm.data.info.gatewayUrl=PATH_INFO.INHERIT_HOST+response.apiInfo.gatewayInfo.gatewayPort+'/'+vm.data.interaction.request.gatewayAlias;
                                break;
                            }
                    }
                })
            }
            return template.promise;
        }
        vm.data.fun.delete = function(arg) {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    apiID: vm.data.interaction.request.apiID
                },
                uri: {
                    groupID: vm.data.interaction.request.groupID
                }
            }
            $rootScope.EnsureModal('删除Api', false, '确认删除', {}, function(callback) {
                if (callback) {
                    GatewayResource.Api.Delete(template.request).$promise.then(function(response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    $state.go('home.gateway.inside.api.list', template.uri);
                                    $rootScope.InfoModal('Api删除成功', 'success');
                                    break;
                                }
                            default:
                                {
                                    $rootScope.InfoModal('删除失败，请稍候再试或到论坛提交bug', 'error');
                                }
                        }
                    })
                }
            });
        }
        $scope.$on('$stateChangeStart', function() { //路由更改函数
            vm.service.cache.clear('apiInfo')
            vm.service.cache.clear('gatewayInfo')
        })
        vm.$onInit = function() {
            var template = {
                array: []
            }
            template.array = [{
                type: 'tabs',
                class: 'menu-li pull-left',
                tabList: [{
                    name: '修改',
                    fun: {
                        default: data.fun.menu,
                        params: '\'edit\',' + JSON.stringify(vm.data.interaction.request)
                    }
                }]
            }, {
                type: 'fun-list',
                class: 'btn-li sort-btn-li pull-left',
                name: '操作',
                icon: 'caidan',
                funList: [{
                    name: '另存为(复制)',
                    icon: 'renwuguanli',
                    fun: {
                        default: data.fun.menu,
                        params: '\'copy\',' + JSON.stringify(vm.data.interaction.request)
                    }
                }, {
                    name: '删除',
                    icon: 'shanchu',
                    fun: {
                        default: vm.data.fun.delete
                    }
                }]
            }]
            vm.component.menuObject.list = [{
                type: 'btn',
                class: 'margin-left-li btn-group-li pull-left',
                btnList: [{
                    name: '接口列表',
                    icon: 'xiangzuo',
                    fun: {
                        default: data.fun.menu,
                        params: '\'list\',' + JSON.stringify(vm.data.interaction.request)
                    }
                }]
            }, {
                type: 'tabs',
                class: 'menu-li pull-left first-menu-li',
                tabList: [{
                    name: '简略',
                    fun: {
                        default: data.fun.menu,
                        params: '\'simple\',' + JSON.stringify(vm.data.interaction.request) 
                    }
                }, {
                    name: '详情',
                    class: 'elem-active'
                }]
            }, {
                type: 'divide',
                class: 'divide-li pull-left'
            }].concat(template.array);
        }
    }
})();
