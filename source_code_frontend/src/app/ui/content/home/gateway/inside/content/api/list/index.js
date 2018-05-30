(function () {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api列表相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.api.list', {
                    url: '/list?groupID',
                    template: '<gateway-api-list></gateway-api-list>'
                });
        }])
        .component('gatewayApiList', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/api/list/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', 'PATH_INFO', 'GroupService', 'Cache_CommonService', 'Communicate_CommonService'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, PATH_INFO, GroupService, Cache_CommonService, Communicate_CommonService) {
        var vm = this;
        vm.data = {
            info: {},
            interaction: {
                request: {
                    gatewayAlias: $state.params.gatewayAlias,
                    groupID: $state.params.groupID || -1,
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
            menuObject: {},
            listRequireObject: {
                mainObject: {
                    baseInfo: {
                        class: 'hover-tr'
                    },
                    tdList: [],
                    fun: {}
                }
            }
        };
        vm.service = {
            cache: Cache_CommonService,
            communicate: Communicate_CommonService
        }
        vm.data.fun.filter = function (arg) {
            if (!vm.data.interaction.request.tips) return arg;
            if (arg.apiName.toLowerCase().indexOf(vm.data.interaction.request.tips.toLowerCase()) > -1 || arg.requestURL.toLowerCase().indexOf(vm.data.interaction.request.tips.toLowerCase()) > -1) return arg;
        }
        vm.data.fun.init = function (arg) {
            var template = {
                promise: null,
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    groupID: vm.data.interaction.request.groupID,
                    tips: vm.data.interaction.request.tips
                }
            }

            if (template.request.groupID == -1) {
                template.promise = GatewayResource.Api.All(template.request).$promise;
                template.promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.data.interaction.response.query = response.apiList || [];
                                vm.data.info.gatewayUrl = PATH_INFO.INHERIT_HOST + response.gatewayInfo.gatewayPort + '/' + vm.data.interaction.request.gatewayAlias;
                                break;
                            }
                    }
                })
            } else {
                template.promise = GatewayResource.Api.Query(template.request).$promise;
                template.promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.data.interaction.response.query = response.apiList || [];
                                vm.data.info.gatewayUrl = PATH_INFO.INHERIT_HOST + response.gatewayInfo.gatewayPort + '/' + vm.data.interaction.request.gatewayAlias;
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
        vm.data.fun.edit = function (arg) {
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
                    gatewayAlias: vm.data.interaction.request.gatewayAlias
                },
                uri: {
                    groupID: vm.data.interaction.request.groupID,
                    apiID: arg.item ? arg.item.apiID : null,
                    status: arg.item ? 'edit' : 'add'
                }
            }
            if ((!template.cache.group) || (template.cache.group.length == 0)) {
                $rootScope.InfoModal('请先建立分组！', 'error');
            } else {
                if (!arg.item) {
                    if (template.cache.backend) {
                        $state.go('home.gateway.inside.api.edit', template.uri);
                    } else {
                        GatewayResource.Backend.Query(template.request).$promise.then(function (response) {
                            vm.service.cache.set(response.backendList, 'backend');
                            $state.go('home.gateway.inside.api.edit', template.uri);
                        })
                    }
                } else {
                    $state.go('home.gateway.inside.api.edit', template.uri)
                }
            }
        }
        vm.data.fun.delete = function (arg) {
            arg = arg || {};
            arg.$event.stopPropagation();
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    apiID: arg.item.apiID
                }
            }
            $rootScope.EnsureModal('删除Api', false, '确认删除', {}, function (callback) {
                if (callback) {
                    GatewayResource.Api.Delete(template.request).$promise.then(function (response) {
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
        vm.data.fun.enter = function (arg) {
            arg = arg || {};
            var template = {
                uri: {
                    groupID: vm.data.interaction.request.groupID,
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
                type: 'customized-html',
                html: '<span>接口名称&nbsp;</span><span class="count-span">{{\'[\'+($ctrl.list.length||0)+\']\'}}</span>',
                key: 'apiName',
                style: {
                    'width': '300px'
                }
            }, {
                name: '接口URI',
                keyType: 'customized-html',
                keyHtml: '<span>{{($ctrl.otherObject.GATEWAY_URL_SHOULD_BE_CONTACT?$ctrl.otherObject.gatewayUrl:"")+($ctrl.otherObject.GPEDIT_ID||"")+item.requestURL}}<span>'
            }, {
                name: '操作',
                keyType: 'btn',
                style: {
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