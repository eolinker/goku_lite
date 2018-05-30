(function() {
    /*
     * author：riverLethe
     * 网关内页页内包模块sidebar相关js
     */
    angular.module('goku')
        .component('gatewayApiSidebar', {
            template: '<div ng-if="$ctrl.data.info.sidebarShow"><group-common-component authority-object="{edit:true}" fun-object="$ctrl.data.component.groupCommonObject.funObject" main-object="$ctrl.data.component.groupCommonObject.mainObject" list="$ctrl.data.interaction.response.query"></group-common-component></div>',
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
                    gatewayAlias: $state.params.gatewayAlias,
                    groupID: $state.params.groupID || -1,
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
                },
                edit: {
                    parent: null, //父分组编辑事件
                },
                delete: {
                    parent: null, //父分组删除事件
                }
            }
        }

        vm.data.fun.init = function() {
            var template = {
                request: {
                    gatewayAlias: vm.data.interaction.request.gatewayAlias,
                    groupID: vm.data.interaction.request.groupID,
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
        vm.data.fun.click.parent = function(arg) {
            vm.data.interaction.request.groupID = arg.item.groupID || -1;
            arg.item.isSpreed = true;
            $state.go('home.gateway.inside.api.list', { 'groupID': arg.item.groupID});
        }
        vm.data.fun.edit.parent = function(arg) {
            arg = arg || {};
            var template = {
                options: {
                    callback: vm.data.fun.init,
                    resource: GatewayResource.ApiGroup,
                    status: 'parent-' + (arg.isEdit ? 'edit' : 'add'),
                    baseRequest: {
                        gatewayAlias: vm.data.interaction.request.gatewayAlias
                    }
                }
            }
            vm.data.service.defaultCommon.fun.operate('edit', arg, template.options);
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
                    GatewayResource.ApiGroup.Delete({ gatewayAlias: vm.data.interaction.request.gatewayAlias, groupID: arg.item.groupID }).$promise.then(function(response) {
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
                    baseInfo: {
                        name: 'groupName',
                        id: 'groupID',
                        interaction: vm.data.interaction.request
                    },
                    staticQuery: vm.data.static.query,
                    parentFun: {
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
                    baseFun: {
                        parentClick: vm.data.fun.click.parent
                    }
                }
            }
        }
    }

})();
