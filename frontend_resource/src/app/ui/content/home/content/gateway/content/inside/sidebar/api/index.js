(function() {
    /*
     * author：riverLethe
     * 网关内页页内包模块sidebar相关js
     */
    angular.module('goku')
        .component('gatewayApiSidebar', {
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/sidebar/api/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', 'GroupService','Group_GatewayCommonService'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, GroupService,Group_GatewayCommonService) {
        var vm = this;
        vm.data = {
            service: {
                defaultCommon: Group_GatewayCommonService
            },
            static: {
                query: [{ groupID: -1, groupName: "所有接口" }]
            },
            component: {
                groupCommonObject: {}
            },
            info: {
                sidebarShow: null
            },
            interaction: {
                request: {
                    gatewayHashKey: $state.params.gatewayHashKey,
                    groupID: $state.params.groupID || -1,
                    childGroupID: $state.params.childGroupID,
                    apiID: $state.params.apiID
                },
                response: {
                    query: null
                }
            },
            fun: {
                more: null, //更多功能函数
                init: null, //初始化功能函数
                click: {
                    parent: null, //父分组单击事件
                    child: null //子分组单击事件
                },
                edit: {
                    parent: null, //父分组编辑事件
                    child: null //子分组编辑事件
                },
                delete: {
                    parent: null, //父分组删除事件
                    child: null //子分组删除事件
                }
            }
        }

        vm.data.fun.init = function() {
            var template = {
                request: {
                    gatewayHashKey: vm.data.interaction.request.gatewayHashKey,
                    groupID: vm.data.interaction.request.groupID,
                    childGroupID: vm.data.interaction.request.childGroupID,
                    apiID: vm.data.interaction.request.apiID
                }
            }
            if ($state.current.name.indexOf('edit') > -1) {
                vm.data.info.sidebarShow = false;
            } else {
                vm.data.info.sidebarShow = true;
            }
            GatewayResource.ApiGroup.Query(template.request).$promise.then(function(response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.data.interaction.response.query = response.groupList||[];
                            if ($state.current.name.indexOf('edit') > -1) {
                                GroupService.set(response.groupList, true);
                            } else {
                                GroupService.set(response.groupList);
                            }
                        }
                }
            })
        }
        vm.data.fun.init();
        vm.data.fun.click.child = function(arg) {
            vm.data.interaction.request.childGroupID = arg.item.groupID;
            $state.go('home.gateway.inside.api.list', { groupID: vm.data.interaction.request.groupID, childGroupID: arg.item.groupID, apiID: null, search: null });
        }
        vm.data.fun.click.parent = function(arg) {
            vm.data.interaction.request.groupID = arg.item.groupID || -1;
            vm.data.interaction.request.childGroupID = null;
            arg.item.isSpreed = true;
            $state.go('home.gateway.inside.api.list', { 'groupID': arg.item.groupID, childGroupID: null});
        }
        vm.data.fun.edit.parent = function(arg) {
            arg = arg || {};
            var template = {
                options: {
                    callback: vm.data.fun.init,
                    resource: GatewayResource.ApiGroup,
                    originGroupQuery: vm.data.interaction.response.query,
                    status: 'parent-' + (arg.isEdit ? 'edit' : 'add'),
                    baseRequest: {
                        gatewayHashKey: vm.data.interaction.request.gatewayHashKey
                    }
                }
            }
            vm.data.service.defaultCommon.fun.operate('edit', arg, template.options);
        }
        vm.data.fun.edit.child = function(arg) {
            arg.item = arg.childItem || {};
            var template = {
                options: {
                    callback: vm.data.fun.init,
                    resource: GatewayResource.ApiGroup,
                    originGroupQuery: vm.data.interaction.response.query,
                    status: 'child-' + (arg.isEdit ? 'edit' : 'add'),
                    baseRequest: {
                        gatewayHashKey: vm.data.interaction.request.gatewayHashKey
                    }
                }
            }
            arg.item.$index = arg.$outerIndex + 1;
            vm.data.service.defaultCommon.fun.operate('edit', arg, template.options);
        }
        vm.data.fun.delete.child = function(arg) {
            arg = arg || {};
            var template = {
                modal: {
                    title: '删除分组',
                    message: '此操作无法恢复，确认删除？'
                }
            }
            $rootScope.EnsureModal(template.modal.title, false, template.modal.message, {}, function(callback) {
                if (callback) {
                    GatewayResource.ApiGroup.Delete({ gatewayHashKey: vm.data.interaction.request.gatewayHashKey, groupID: arg.childItem.groupID }).$promise.then(function(response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    arg.item.childGroupList.splice(arg.$index, 1);
                                    $rootScope.InfoModal('分组删除成功', 'success');
                                    if (vm.data.interaction.request.childGroupID == arg.childItem.groupID) {
                                        vm.data.fun.click.parent({ item: arg.item });
                                    }
                                    break;
                                }
                        }
                    })
                }
            });
        }
        vm.data.fun.delete.parent = function(arg) {
            arg = arg || {};
            var template = {
                modal: {
                    title: '删除分组',
                    message: '此操作无法恢复，确认删除？'
                }
            }
            $rootScope.EnsureModal(template.modal.title, false, template.modal.message, {}, function(callback) {
                if (callback) {
                    GatewayResource.ApiGroup.Delete({ gatewayHashKey: vm.data.interaction.request.gatewayHashKey, groupID: arg.item.groupID }).$promise.then(function(response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    vm.data.interaction.response.query.splice(arg.$index, 1);
                                    $rootScope.InfoModal('分组删除成功', 'success');
                                    if (vm.data.interaction.response.query.length > 1) {
                                        GroupService.set(vm.data.interaction.response.query.slice(1));
                                    } else {
                                        GroupService.set(null);
                                    }
                                    if (vm.data.interaction.request.groupID == arg.item.groupID||$state.params.groupID) {
                                        vm.data.fun.click.parent({ item: vm.data.interaction.response.query[0] });
                                    }else if($state.params.groupID==-1){
                                        vm.data.fun.click.parent({ item: {} });
                                    }

                                    break;
                                }
                        }
                    })
                }
            });
        }
        $scope.$on('$stateChangeSuccess', function() { //路由更改函数
            vm.data.interaction.request.groupID = $state.params.groupID || -1;
            vm.data.interaction.request.childGroupID = vm.data.interaction.request.childGroupID;
            if ($state.current.name.indexOf('edit') > -1) {
                vm.data.info.sidebarShow = false;
            } else {
                vm.data.info.sidebarShow = true;
            }
        })
        vm.$onInit = function() {
            vm.data.component.groupCommonObject = {
                funObject: {
                    btnGroupList: {
                        edit: {
                            key: '新建分组',
                            class: 'eo-button-success',
                            icon: 'tianjia',
                            fun: vm.data.fun.edit.parent
                        }
                    }
                },
                mainObject: {
                    level: 2,
                    baseInfo: {
                        name: 'groupName',
                        id: 'groupID',
                        childID: 'childGroupID',
                        child: 'childGroupList',
                        interaction: vm.data.interaction.request
                    },
                    staticQuery: vm.data.static.query,
                    parentFun: {
                        addChild: {
                            fun: vm.data.fun.edit.child,
                            key: '添加子分组',
                            params: { $outerIndex: null, isEdit: false },
                            class: 'add-child-btn'
                        },
                        edit: {
                            fun: vm.data.fun.edit.parent,
                            key: '修改',
                            params: { item: null, isEdit: true }
                        },
                        delete: {
                            fun: vm.data.fun.delete.parent,
                            key: '删除',
                            params: { item: null, $index: null }
                        }
                    },
                    childFun: {
                        edit: {
                            fun: vm.data.fun.edit.child,
                            key: '修改',
                            params: { childItem: null, $outerIndex: null, isEdit: true }
                        },
                        delete: {
                            fun: vm.data.fun.delete.child,
                            key: '删除',
                            params: { item: null, childItem: null, $index: null }
                        }
                    },
                    baseFun: {
                        parentClick: vm.data.fun.click.parent,
                        childClick: vm.data.fun.click.child,
                        spreed: function(arg){
                            if (arg.$event) {
                                arg.$event.stopPropagation();
                            }
                            arg.item.isSpreed = !arg.item.isSpreed;
                        }
                    }
                }
            }
        }
    }

})();
