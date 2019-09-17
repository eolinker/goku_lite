(function () {
    'use strict';
    angular.module('eolinker')
        .component('balanceList', {
            templateUrl: 'app/ui/content/balance/list/index.html',
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
        vm.ajaxRequest={
            balanceName: [],
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        };
        vm.ajaxResponse={
            query:null
        }
        vm.service = {
            authority: Authority_CommonService
        }
        vm.component = {
            menuObject: {},
            listRequireObject: null
        }
        vm.fun={};
        var privateFun = {},
            ajaxResponse = {};

        vm.fun.init = function () {
            let tmpAjaxRequest={};
            if(vm.ajaxRequest.keyword){
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_Balance = GatewayResource.Balance.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_Balance.$promise.then(function (response) {
                vm.ajaxResponse.query = response.balanceList || [];
            })
            return $rootScope.global.ajax.Query_Balance.$promise;
        }
        privateFun.search = function (arg) {
            window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
            $state.reload($state.current.name);
        }
        privateFun.resetBatchInfo = function () {
            vm.data.batch.isOperating = false;
            vm.data.batch.selectAll = false;
            vm.data.batch.query = [];
            vm.data.batch.indexAddress = {};
        };
        privateFun.initRegistry = () => {
            if ($rootScope.global.ajax.Query_ServiceDiscovery) $rootScope.global.ajax.Query_ServiceDiscovery.$cancelRequest();
            $rootScope.global.ajax.Query_ServiceDiscovery = GatewayResource.ServiceDiscovery.SimpleQuery();
            $rootScope.global.ajax.Query_ServiceDiscovery.$promise.then(function (response) {
                ajaxResponse.registry = response.data.list;
            })
            return $rootScope.global.ajax.Query_ServiceDiscovery.$promise;
        }
        privateFun.edit = function (arg) {
            var tmpUri = {
                status: arg.status
            }
            switch (arg.status) {
                case 'edit': {
                    tmpUri.balanceName = arg.item.balanceName;
                    $state.go('home.balance.list.operate', tmpUri);
                    break;
                }
                default: {
                    let tmpPromise = privateFun.initRegistry();
                    tmpPromise.finally(() => {
                        if (ajaxResponse.registry.length === 0) {
                            $rootScope.InfoModal('请先新建服务注册方式', 'error');
                        } else {
                            $state.go('home.balance.list.operate', tmpUri);
                        }
                    })
                    break;
                }
            }
        }
        privateFun.delete = function (inputArg) {
            var template = {
                request: {
                    balanceNames: inputArg.status == 'batch' ? vm.data.batch.query.join(',') : inputArg.item.balanceName
                },
                modal: {
                    title: inputArg.status == 'batch' ? '批量删除负载' : ('删除负载-' + inputArg.item.balanceName)
                },
                loop: {
                    num: 0
                }
            }
            $rootScope.EnsureModal(template.modal.title, null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.Balance.Delete(template.request).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                switch (inputArg.status) {
                                    case 'batch': {
                                        privateFun.resetBatchInfo();
                                        vm.fun.init();
                                        break;
                                    }
                                    case 'single': {
                                        vm.ajaxResponse.query.splice(inputArg.$index, 1);
                                        break;
                                    }
                                }
                                $rootScope.InfoModal(template.modal.title + '成功', 'success');
                                break;
                            }
                        }
                    })
                }
            });
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['负载管理']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'balanceName',
                    default: [{
                        key: '名称',
                        html: '{{item.balanceName}}'
                    }, {
                        key: '服务注册方式',
                        html: '{{item.serviceName}}'
                    }, {
                        key: '更新时间',
                        html: '{{item.updateTime}}',
                        class:'w_200'
                    }],
                    operate: {
                        funArr: [{
                            key: '修改',
                            fun: privateFun.edit,
                            params: {
                                status: 'edit'
                            }
                        }, {
                            key: '删除',
                            fun: privateFun.delete,
                            params: {
                                status:"single"
                            }
                        }],
                        class:'w_150'
                    }
                },
                setting: {
                    batch:true,
                    unhover: true,
                    warning: '尚未添加负载',
                    defaultFoot:true
                }
            }
            vm.component.menuObject = {
                list: [{
                        type: 'btn',
                        authority: 'edit',
                        class: 'pull-left',
                        btnList: [{
                            name: '添加负载',
                            icon: 'jiahao',
                            class: 'eo_theme_btn_success block-btn',
                            fun: {
                                default: privateFun.edit,
                                params: {
                                    status:'add'
                                }
                            }
                        }]
                    },{
                        type: 'search',
                        class: 'pull-right',
                        keyword: vm.ajaxRequest.keyword,
                        fun: privateFun.search,
                        placeholder:"输入负载名称或服务注册方式"
                    }
                ],
                batchList: [{
                    type: 'btn',
                    disabledPoint: 'isBatchSelected',
                    class: 'pull-left',
                    btnList: [{
                        name: '删除',
                        icon: 'shanchu',
                        disabled: false,
                        fun: {
                            default: privateFun.delete,
                            params: {
                                status: 'batch'
                            }
                        }
                    }]
                }],
                setting: {
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    class: "common-menu-fixed-seperate common-menu-lg",
                    titleAuthority: 'showTitle',
                    title: '负载配置',
                    secondTitle:"配置API的转发目标服务器（负载后端），创建之后可以设置为 API 的转发地址 / 负载后端（Target / Upstream）"
                }
            }
        }
    }
})();