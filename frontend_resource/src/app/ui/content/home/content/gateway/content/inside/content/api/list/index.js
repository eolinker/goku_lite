(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api列表相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.api.list', {
                    url: '/list?groupID?childGroupID',
                    template: '<gateway-api-list></gateway-api-list>'
                });
        }])
        .component('gatewayApiList', {
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/content/api/list/index.html',
            controller: gatewayApiListController
        })

    gatewayApiListController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', '$timeout', 'GroupService', 'Cache_CommonService'];

    function gatewayApiListController($scope, GatewayResource, $state, CODE, $rootScope, $timeout, GroupService, Cache_CommonService) {
        var vm = this;
        vm.data = {
            interaction: {
                request: {
                    gatewayHashKey: $state.params.gatewayHashKey,
                    groupID: $state.params.groupID || -1,
                    childGroupID: $state.params.childGroupID,
                    tips: null
                },
                response: {
                    query: []
                }
            },
            fun: {
                simple: null, //进入简略功能函数
                delete: null, //删除功能函数
                edit: null, //编辑功能函数
                search: null, //搜索功能函数
                init: null //初始化功能函数
            }
        }
        vm.component = {
            menuObject: {
                show: {
                    batch: {
                        disable: false
                    }
                },
                active: {
                    sort: JSON.parse(window.localStorage['PROJECT_SORTTYPE'] || '{"orderBy":3,"asc":0}'),
                    condition: 0,
                    more: parseInt(window.localStorage['PROJECT_MORETYPE']) || 1
                }
            },
            listRequireObject: {
                mainObject: {
                    baseInfo: {
                        active: 'disable',
                        class: 'hover-tr'
                    },
                    tdList: [],
                    fun: {}
                }
            }
        };
        vm.service={
            cache:Cache_CommonService
        }
        vm.data.fun.filter = function (arg) {
            if (!vm.data.interaction.request.tips) return arg;
            if (arg.apiName.toLowerCase().indexOf(vm.data.interaction.request.tips.toLowerCase()) > -1 || arg.gatewayRequestURI.toLowerCase().indexOf(vm.data.interaction.request.tips.toLowerCase()) > -1) return arg;
        }
        vm.data.fun.init = function(arg) {
            var template = {
                promise: null,
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey,
                    groupID: vm.data.interaction.request.childGroupID || vm.data.interaction.request.groupID,
                    tips: vm.data.interaction.request.tips
                }
            }
            
            if (template.request.groupID == -1) {
                template.promise = GatewayResource.Api.All(template.request).$promise;
                template.promise.then(function(response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.data.interaction.response.query = response.apiList||[];
                                break;
                            }
                    }
                })
            } else {
                template.promise = GatewayResource.Api.Query(template.request).$promise;
                template.promise.then(function(response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.data.interaction.response.query = response.apiList||[];
                                break;
                            }
                    }
                })
            }
            return template.promise;
        }
        vm.data.fun.search = function (q) {
            vm.data.interaction.request.tips = q;
        }
        vm.data.fun.edit = function(arg) {
            arg = arg || {}
            if (arg.$event) {
                arg.$event.stopPropagation();
            }
            var template = {
                cache: {
                    group: GroupService.get(),
                    backend: vm.service.cache.get('backend')
                },
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey
                },
                uri: {
                    groupID: vm.data.interaction.request.groupID,
                    childGroupID: vm.data.interaction.request.childGroupID,
                    apiID: arg.item ? arg.item.apiID : null,
                    status:arg.item?'edit':'add'
                }
            }
            if ((!template.cache.group) || (template.cache.group.length == 0)) {
                $rootScope.InfoModal('请先建立分组！', 'error');
            } else {
                if (!arg.item) {
                    if (template.cache.backend) {
                        $state.go('home.gateway.inside.api.edit', template.uri);
                    } else {
                        GatewayResource.Backend.Query(template.request).$promise.then(function(response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS:
                                    {
                                        vm.service.cache.set(response.backendList,'backend');
                                        $state.go('home.gateway.inside.api.edit', template.uri);
                                        break;
                                    }
                                default:
                                    {
                                        $rootScope.InfoModal('请先建立后端服务！', 'error');
                                        break;
                                    }
                            }
                        })
                    }
                } else {
                    $state.go('home.gateway.inside.api.edit', template.uri)
                }
            }
        }
        vm.data.fun.delete = function(arg) {
            arg = arg || {};
            arg.$event.stopPropagation();
            var template = {
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey,
                    apiID: arg.item.apiID
                }
            }
            $rootScope.EnsureModal('删除Api', false, '确认删除', {}, function(callback) {
                if (callback) {
                    GatewayResource.Api.Delete(template.request).$promise.then(function(response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    vm.data.interaction.response.query.splice(arg.$index, 1);
                                    $rootScope.InfoModal('Api删除成功', 'success');
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
        vm.data.fun.enter = function(arg) {
            arg = arg || {};
            var template = {
                uri: {
                    groupID: vm.data.interaction.request.groupID,
                    childGroupID: vm.data.interaction.request.childGroupID,
                    apiID: arg.item.apiID
                }
            }
            $state.go('home.gateway.inside.api.simple', template.uri);
        }
        vm.$onInit = function () {
            var template = {
                array: []
            }
            vm.component.listRequireObject.mainObject.fun = {
                filter: vm.data.fun.filter,
                click: vm.data.fun.enter
            }
            vm.component.listRequireObject.mainObject.tdList = [{
                type:'customized-html',
                html: '<span>接口名称&nbsp;</span><span class="count-span">{{\'[\'+($ctrl.list.length||0)+\']\'}}</span>',
                key: 'apiName',
                style: {
                    'min-width': '200px'
                }
            }, {
                name: '接口URI',
                key: 'gatewayRequestURI',
                style: {
                    'min-width': '220px'
                }
            }, {
                name: '操作',
                keyType: 'btn',
                style: {
                    'min-width': '115px',
                    'width': '250px'
                },
                btnList: [{
                    name: '修改',
                    icon: 'bianji',
                    fun: {
                        default: vm.data.fun.edit
                    }
                }, {
                    name: '删除',
                    icon: 'shanchu',
                    fun: {
                        default: vm.data.fun.delete,
                        params: {
                            switch: 0
                        }
                    }
                }]
            }];
            template.array = [{
                type: 'btn',
                class: 'btn-group-li pull-left',
                btnList: [{
                    name: '添加接口',
                    icon: 'tianjia',
                    class: 'eo-button-success',
                    fun: {
                        default: vm.data.fun.edit,
                        params: ''
                    }
                }]
            }];
            vm.component.menuObject.list = [{
                type: 'divide',
                class: 'divide-li pull-left'
            }, {
                type: 'search',
                class: 'search-li pull-left',
                placeholder: '搜索接口',
                fun: vm.data.fun.search,
                tips: ''
            }];

            vm.component.menuObject.list = template.array.concat(vm.component.menuObject.list);
        }
    }
})();
