(function () {
    //author：广州银云信息科技有限公司
    angular.module('eolinker')
        .component('apiGroup', {
            template: '<group-default-common-Component ng-show="$ctrl.data.sidebarShow" authority-object="$ctrl.service.authority.permission.default.apiManagement"  fun-object="$ctrl.component.groupCommonObject.funObject" request-object="$ctrl.component.groupCommonObject.requestObject" main-object="$ctrl.component.groupCommonObject.mainObject"></group-default-common-Component>',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', '$rootScope', 'GroupService', 'Authority_CommonService', 'Cache_CommonService'];

    function indexController($scope, GatewayResource, $state, $rootScope, GroupService, Authority_CommonService, Cache_CommonService) {
        var vm = this;
        vm.data = {
            static: {
                query: [{
                    groupID: -1,
                    groupName: "所有接口",
                    icon: 'caidan'
                }, {
                    groupID: 0,
                    groupName: "未分组",
                    icon: 'caidan'
                }]
            },
            sidebarShow: null,
        }
        vm.ajaxRequest={
            projectID: $state.params.projectID,
            groupID: $state.params.groupID||-1,
            apiID: $state.params.apiID,
            orderNumber: []
        }
        vm.component = {
            groupCommonObject: {}
        }
        vm.service = {
            authority: Authority_CommonService
        }
        var fun = {};
        fun.init = function () {
            if ($state.current.name.indexOf('api.default') > -1) {
                vm.data.sidebarShow = true;
            } else {
                vm.data.sidebarShow = false;
            }
        }
        fun.init();
        fun.import = function () {
            var template = {
                modal: {
                    title: '导入分组',
                    resource: GatewayResource.ImportAms.Group,
                    request: {
                        projectID: vm.ajaxRequest.projectID
                    }
                }
            }

            $rootScope.ImportModal(template.modal, function (callback) {
                if (callback) {
                    $rootScope.InfoModal(template.modal.title + '成功', 'success');
                    vm.component.groupCommonObject.mainObject.baseInfo.resetFlag=!vm.component.groupCommonObject.mainObject.baseInfo.resetFlag;
                }
            });
        }
        $scope.$on('$stateChangeSuccess', function () { //路由更改函数
            if ($state.current.name.indexOf('api.default') > -1) {
                vm.data.sidebarShow = true;
            } else {
                vm.data.sidebarShow = false;
            }
            vm.ajaxRequest.groupID = $state.params.groupID;
        })
        vm.$onInit = function () {
            vm.component.groupCommonObject = {
                requestObject: {
                    resource: GatewayResource.ApiGroup,
                    baseRequest: {
                        projectID: vm.ajaxRequest.projectID,
                    }
                },
                funObject: {
                    unTop:true,
                    btnGroupList: [{
                        type: 'more-btn',
                        authority: 'edit',
                        icon: 'jiahao',
                        name: '新建',
                        class: 'pull-left',
                        btnClass: 'eo_theme_btn_success',
                        btnList: [{
                            funName: 'edit',
                            name: '新建分组',
                            funName: 'edit',
                            icon: 'jiahao',
                        }, {
                            name: '导入',
                            fun: fun.import
                        }]
                    }],
                    callback: {
                        querySuccess: function (groupList) {
                            if ($state.current.name.indexOf('list') > -1) {
                                GroupService.set(groupList);
                            } else {
                                GroupService.set(groupList, true);
                            }
                        }
                    },
                },
                mainObject: {
                    showRouterList: ['home.project.api.default'],
                    baseInfo: {
                        sort: true,
                        name: 'groupName',
                        id: 'groupID',
                        current: vm.ajaxRequest,
                    },
                    staticQuery: vm.data.static.query,
                    itemFun: [{
                            funName: 'edit',
                            key: '新建子分组',
                            params: '"add-child",arg',
                            authority: 'edit',
                            isStrong: 1
                        },
                        {
                            funName: 'edit',
                            key: '编辑',
                            authority: 'edit',
                            params: '"edit",arg'
                        }, {
                            funName: 'delete',
                            authority: 'edit',
                            params: {
                                modal: {
                                    title: '删除分组',
                                    message: '删除分组后，该分组内的API也将被删除，该操作无法撤销，确认删除？'
                                }
                            },
                            key: '删除'
                        }
                    ]
                }
            }
        }
    }

})();