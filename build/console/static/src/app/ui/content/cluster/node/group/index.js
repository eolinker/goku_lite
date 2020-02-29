(function () {
    /*
     * author：广州银云信息科技有限公司
     */
    angular.module('eolinker')
        .component('clusterNodeGroup', {
            template: '<group-default-common-component authority-object="$ctrl.service.authority.permission.default.nodeManagement"  fun-object="$ctrl.component.groupCommonObject.funObject"  request-object="$ctrl.component.groupCommonObject.requestObject" main-object="$ctrl.component.groupCommonObject.mainObject"></group-default-common-component>',
            controller: indexController,
            bindings:{
                groupArr:'='
            }
        })

    indexController.$inject = ['GatewayResource', '$state','Authority_CommonService'];

    function indexController(GatewayResource, $state,Authority_CommonService) {
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
            cluster:$state.params.cluster,
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
        vm.$onInit = function () {
            vm.component.groupCommonObject = {
                requestObject: {
                    resource: GatewayResource.NodeGroup,
                    baseRequest: {
                        cluster:vm.ajaxRequest.cluster
                    }
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
                            vm.groupArr=groupList||[];
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
                                message: '删除分组后，该分组内的网关节点也将被删除，该操作无法撤销，确认删除？'
                            }
                        }
                    }]
                }
            }
        }
    }

})();