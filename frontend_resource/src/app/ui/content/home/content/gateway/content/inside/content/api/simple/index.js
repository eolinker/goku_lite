(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api简略相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.api.simple', {
                    url: '/simple?apiID?groupID?childGroupID',
                    template: '<gateway-api-simple></gateway-api-simple>'
                });
        }])
        .component('gatewayApiSimple', {
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/content/api/simple/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope','Cache_CommonService','PATH_INFO'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope,Cache_CommonService,PATH_INFO) {
        var vm = this;
        vm.data = {
            interaction: {
                request: {
                    apiID: $state.params.apiID,
                    gatewayHashKey: $state.params.gatewayHashKey,
                    groupID: $state.params.groupID,
                    childGroupID: $state.params.childGroupID
                },
                response: {
                    apiInfo: {}
                }
            },
            fun: {
                delete: null, //删除功能函数
                init: null //初始化功能函数
            }
        }
        vm.service={
            cache:Cache_CommonService
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
                    childGroupID: uri.childGroupID,
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
                case 'detail':
                    {
                        $state.go('home.gateway.inside.api.detail', template.uri);
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
        vm.data.fun.init = function() {
            var template = {
                promise: null,
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey,
                    groupID: vm.data.interaction.request.childGroupID || vm.data.interaction.request.groupID,
                    apiID: vm.data.interaction.request.apiID
                }
            }
            template.promise = GatewayResource.Api.Detail(template.request).$promise;
            template.promise.then(function(response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.data.interaction.response.apiInfo = response.apiInfo.apiJson;
                            vm.data.interaction.response.apiInfo.hash=PATH_INFO.INHERIT_HOST+response.apiInfo.gatewayPort+'/'+ $state.params.gatewayHashKey+vm.data.interaction.response.apiInfo.gatewayRequestPath
                            vm.service.cache.set(vm.data.interaction.response.apiInfo,'apiInfo');
                            break;
                        }
                }
            })
            return template.promise;
        }
        vm.data.fun.delete = function() {
            var template = {
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey,
                    apiID: vm.data.interaction.request.apiID
                },
                uri:{
                    groupID:vm.data.interaction.request.groupID,
                    childGroupID:vm.data.interaction.request.childGroupID
                }
            }
            $rootScope.EnsureModal('删除Api', false, '确认删除', {}, function(callback) {
                if (callback) {
                    GatewayResource.Api.Delete(template.request).$promise.then(function(response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    $state.go('home.gateway.inside.api.list',template.uri);
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
                    class: 'elem-active'
                }, {
                    name: '详情',
                    fun: {
                        default: data.fun.menu,
                        params: '\'detail\',' + JSON.stringify(vm.data.interaction.request) 
                    }
                }]
            }, {
                type: 'divide',
                class: 'divide-li pull-left'
            }].concat(template.array);
        }
    }
})();
