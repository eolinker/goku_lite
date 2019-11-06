(function () {
    'use strict';
    angular.module('eolinker')
        .component('clusterNodeDefault', {
            templateUrl: 'app/ui/content/cluster/node/_default/index.html',
            bindings: {
                groupArr: '<'
            },
            controller: indexController
        })
    indexController.$inject = ['$scope', 'GatewayResource', '$state', '$rootScope', 'CODE', 'Authority_CommonService'];

    function indexController($scope, GatewayResource, $state, $rootScope, CODE, Authority_CommonService) {
        var vm = this;
        vm.data = {
            batch: {
                address: []
            }
        }
        vm.fun = {};
        vm.ajaxRequest = {
            nodeID: [],
            groupID: $state.params.groupID || -1,
            cluster: $state.params.cluster,
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        }
        vm.ajaxResponse = {
            query: null
        }
        vm.component = {
            listDefaultCommonObject: {},
            menuObject: {}
        }
        vm.service = {
            authority: Authority_CommonService
        }
        var privateFun = {},
            CONST = {
                GROUP_ARR: [{
                    groupID: 0,
                    groupDepth: 1,
                    groupName: '未分组'
                }]
            };
            privateFun.search = function (arg) {
                window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
                $state.reload($state.current.name);
            }
        privateFun.initQueryAjax = function () {
            let tmpAjaxRequest={
                groupID: vm.ajaxRequest.groupID,
                cluster: vm.ajaxRequest.cluster
            };
            if(vm.ajaxRequest.keyword){
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_Node = GatewayResource.Node.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_Node.$promise.then(function (response) {
                vm.ajaxResponse.query = response.nodeList || [];
            })
            return $rootScope.global.ajax.Query_Node.$promise;
        }
        vm.fun.init = function () {
            return privateFun.initQueryAjax();
        }
        privateFun.changeGroup = function () {
            var template = {
                modal: {
                    title: '批量修改节点分组',
                    query: CONST.GROUP_ARR.concat(vm.groupArr),
                    position: {
                        key: 'groupName'
                    }
                },
                request: {
                    nodeIDList: vm.data.batch.query.join(','),
                    groupID: '',
                    cluster: vm.ajaxRequest.cluster
                }
            }
            $rootScope.SingleSelectModal(template.modal, function (callback) {
                if (callback) {
                    template.request.groupID = template.modal.query[callback.$index].groupID;
                    GatewayResource.Node.ChangeGroup(template.request).$promise
                        .then(function (response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS: {
                                    privateFun.resetBatchInfo();
                                    $rootScope.InfoModal('节点批量修改分组成功', 'success');
                                    vm.fun.init();

                                    break;
                                }
                            }
                        })
                }
            });
        }
        privateFun.edit = function (arg) {
            let tmpModal={
                title: arg.status == 'edit' ? '修改节点' : '新建节点',
                group: CONST.GROUP_ARR.concat(vm.groupArr),
                item: arg.status == 'edit' ? arg.item : {}
            }
            var template = {
                request: {}
            }
            tmpModal.item.groupID = arg.item.groupID||parseInt(vm.ajaxRequest.groupID);
            if(tmpModal.item.groupID===-1)tmpModal.item.groupID=0;
            $rootScope.GatewayClusterModal(tmpModal, function (callback) {
                if (callback) {
                    template.request = {
                        groupID: callback.groupID,
                        listenAddress: callback.listenAddress,
                        nodeName: callback.nodeName,
                        adminAddress: callback.adminAddress,
                        gatewayPath: callback.gatewayPath,
                        cluster: vm.ajaxRequest.cluster
                    }
                    switch (arg.status) {
                        case 'add': {
                            GatewayResource.Node.Add(template.request).$promise.then(function (response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS: {
                                        vm.fun.init();
                                        $rootScope.InfoModal('新增节点成功!', 'success')
                                        break;
                                    }
                                }
                            })
                            break;
                        }
                        default: {
                            template.request.nodeID = arg.item.nodeID;
                            GatewayResource.Node.Edit(template.request).$promise.then(function (response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS: {
                                        vm.fun.init();
                                        $rootScope.InfoModal('修改节点成功!', 'success')
                                        break;
                                    }
                                }
                            })
                            break;
                        }
                    }
                }
            });
        }
        privateFun.delete = function (status, arg) {
            var template = {
                request: {
                    nodeIDList: status == 'batch' ? vm.data.batch.query.join(',') : arg.item.nodeID,
                    cluster: vm.ajaxRequest.cluster
                }
            }
            $rootScope.EnsureModal('删除节点', null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.Node.Delete(template.request).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                switch (status) {
                                    case 'batch': {
                                        privateFun.resetBatchInfo();
                                        vm.fun.init();
                                        break;
                                    }
                                    case 'single': {
                                        vm.ajaxResponse.query.splice(arg.$index, 1);
                                        break;
                                    }
                                }
                                $rootScope.InfoModal('节点删除成功', 'success');
                                break;
                            }
                        }
                    })
                }
            });
        }
        privateFun.resetBatchInfo = function () {
            vm.data.batch.isOperating = false;
            vm.data.batch.selectAll = false;
            vm.data.batch.query = [];
            vm.data.batch.indexAddress = {};
        };
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['节点管理']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'nodeID',
                    default: [{
                        key: '名称',
                        html: '{{item.nodeName}}',
                        draggableCacheMark: 'name'
                    }, {
                        key: 'Node Key',
                        html: '{{item.nodeKey}}',
                        draggableCacheMark: 'key'
                    }, {
                        key: '监听地址',
                        html: '{{item.listenAddress}}',
                        draggableCacheMark: 'listenAddress'
                    }, {
                        key: '管理地址',
                        html: '{{item.adminAddress}}',
                        draggableCacheMark: 'adminAddress'
                    }, {
                        key: '状态',
                        html: `<span class="eo-status-warning" ng-if="item.nodeStatus=='0'">未运行</span><span class="eo-status-danger" ng-if="item.nodeStatus=='2'">异常</span><span class="eo-status-success" ng-if="item.nodeStatus=='1'">运行中</span>`,
                        draggableCacheMark: 'status'
                    }, {
                        key: '分组',
                        html: '{{item.groupName}}',
                        draggableCacheMark: 'group'
                    }, {
                        key: '版本',
                        html: '{{item.version}}',
                        draggableCacheMark: 'version'
                    }, {
                        key: '更新时间',
                        html: '{{item.updateTime}}',
                        draggableCacheMark: 'time'
                    }],
                    operate: {
                        funArr: [{
                            type: 'html',
                            html: `<button class="eo-operate-btn" ng-click="$ctrl.mainObject.baseFun.edit({status:'edit',item:item})">修改</button>`,
                        }, {
                            key: '删除',
                            fun: privateFun.delete,
                            params: '"single",arg'
                        }],
                        class: 'w_150'
                    }
                },
                baseFun: {
                    edit: privateFun.edit
                },
                setting: {
                    isFixedHeight: true,
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    titleAuthority: 'showTitle',
                    unhover: true,
                    warning: '尚未新建任何节点',
                    fixFoot:true,
                    draggable: true,
                    dragCacheVar: 'NODE_LIST_DRAG_VAR',
                    dragCacheObj: {
                        name: '250px',
                        key: '250px',
                        listenAddress: '150px',
                        adminAddress: '150px',
                        status: '150px',
                        group: '150px',
                        version: '150px',
                        time:'180px'
                    }
                }
            }

            vm.component.menuObject = {
                list: [{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    authority: 'edit',
                    btnList: [{
                        name: '新建节点',
                        icon: 'jiahao',
                        class: 'eo_theme_btn_success block-btn',
                        fun: {
                            default: privateFun.edit,
                            params: {
                                status: 'add'
                            }
                        }
                    }]
                },{
                    type: 'search',
                    class: 'pull-right',
                    keyword: vm.ajaxRequest.keyword,
                    fun: privateFun.search,
                    placeholder:"输入节点名称或IP信息"
                }],
                batchList: [{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '修改分组',
                        fun: {
                            default: privateFun.changeGroup
                        }
                    }, {
                        name: '批量删除',
                        fun: {
                            default: privateFun.delete,
                            params: '"batch",arg'
                        }
                    }]
                }],
                setting: {
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    titleAuthority: 'showTitle',
                    title: "节点列表"
                }
            };
        }
    }
})();