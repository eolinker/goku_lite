(function () {
    /*
     * author：广州银云信息科技有限公司
     */
    angular.module('eolinker')
        .component('gpeditGroup', {
            template: '<group-default-common-component authority-object="$ctrl.service.authority.permission.default.strategyManagement"  fun-object="$ctrl.component.groupCommonObject.funObject"  request-object="$ctrl.component.groupCommonObject.requestObject" main-object="$ctrl.component.groupCommonObject.mainObject"></group-default-common-component>',
            controller: indexController
        })

        indexController.$inject = ['GatewayResource','Cache_CommonService', '$state','Authority_CommonService'];

        function indexController(GatewayResource,Cache_CommonService, $state,Authority_CommonService) {
            var vm = this;
            vm.data = {
                static: {
                    query: [{
                        groupID: -1,
                        groupName: "所有分组",
                        icon: 'caidan'
                    },{
                        groupID: 0,
                        groupName: "未分组",
                        icon: 'caidan'
                    }]
                }
            }
            vm.ajaxRequest={
                groupID: $state.params.groupID||-1
            }
            vm.ajaxResponse={
                query: []
            }
            vm.service = {
                authority: Authority_CommonService
            }
            vm.component = {
                groupCommonObject: {}
            }
            let service = {
                cache: Cache_CommonService
            };
            vm.$onInit = function () {
                vm.component.groupCommonObject = {
                    requestObject: {
                        resource: GatewayResource.StrategyGroup,
                        baseRequest: {}
                    },
                    funObject: {
                        btnGroupList: [{
                            type: 'btn',
                            authority: 'edit',
                            icon: 'jiahao',
                            name: '新建分组',
                            funName: 'edit',
                            class: 'eo_theme_btn_success'
                        }],
                        callback: {
                            querySuccess: function (groupList) {
                                service.cache.set(groupList,'gpeditGroup');
                            }
                        }
                    },
                    mainObject: {
                        baseInfo: {
                            name: 'groupName',
                            id: 'groupID',
                            current: vm.ajaxRequest
                        },
                        staticQuery: vm.data.static.query,
                        itemFun: [{
                            funName: 'edit',
                            key: '编辑',
                            authority: 'edit',
                            params: '"edit",arg'
                        }, {
                            funName: 'delete',
                            key: '删除',
                            authority: 'edit',
                            params: {
                                modal: {
                                    title: '删除分组',
                                    message: '删除分组后，该分组内的策略也将被删除，该操作无法撤销，确认删除？'
                                }
                            }
                        }]
                    }
                }
            }
        }

})();