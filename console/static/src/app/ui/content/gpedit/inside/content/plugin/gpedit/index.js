(function () {
    /**
     * @name 策略组插件
     * @author 广州银云信息科技有限公司
     */
    /**
     * @version 1.4
     * @description 添加批量操作
     */
    'use strict';
    angular.module('eolinker')
        .component('gpeditInsidePluginGpedit', {
            templateUrl: 'app/ui/content/gpedit/inside/content/plugin/gpedit/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', '$rootScope', 'CODE', 'Authority_CommonService'];

    function indexController($scope, GatewayResource, $state, $rootScope, CODE, Authority_CommonService) {
        var vm = this;
        vm.data = {
            batch: {}
        }
        vm.fun={};
        vm.ajaxRequest={
            strategyID: $state.params.strategyID,
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        };
        vm.ajaxResponse={
            query: null
        }
        vm.component = {
            listDefaultCommonObject: null,
            menuObject: {
                show: {
                    batch: {
                        disable: false
                    }
                }
            }
        }
        vm.service = {
            authority: Authority_CommonService
        }
        var privateFun = {};
        privateFun.search = function (arg) {
            window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
            $state.reload($state.current.name);
        }
        vm.fun.init = function () {
            var tmpAjaxRequest={
                strategyID: vm.ajaxRequest.strategyID,
                condition:vm.component.menuObject.active.condition
            }
            if(vm.ajaxRequest.keyword){
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_PluginStrategy = GatewayResource.PluginStrategy.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_PluginStrategy.$promise.then(function (response) {
                vm.ajaxResponse.query = response.strategyPluginList || [];
            })
            return $rootScope.global.ajax.Query_PluginStrategy.$promise;
        }
        privateFun.edit = function (arg) {
            var template = {
                uri: {
                    status: arg.status,
                    from: 'gpedit'
                }
            }
            switch (arg.status) {
                case 'edit':
                    {
                        template.uri.pluginName = arg.item.pluginName;
                        template.uri.chineseName=arg.item.pluginDesc;
                        break;
                    }
            }
            $state.go('home.gpedit.inside.plugin.operate', template.uri);
        }
        privateFun.batchOperate = function (arg) {
            var template = {
                promise: null,
                title: vm.data.batch.isOperating ? '批量' : '',
                loop: {
                    num: 0
                },
                fun: null
            }
            switch (arg.status) {
                case 'Start':
                    {
                        template.title += '启用';
                        break;
                    }
                case 'Stop':
                    {
                        template.title += '停用';
                        break;
                    }
            }
            template.fun = function () {
                GatewayResource.PluginStrategy[arg.status]({
                    strategyID: vm.ajaxRequest.strategyID,
                    connIDList: vm.data.batch.isOperating ? vm.data.batch.query.join(',') : arg.item.connID
                }).$promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                privateFun.resetBatchInfo();
                                $rootScope.InfoModal(template.title + '成功', 'success');
                                vm.fun.init();
                                break;
                            }
                    }
                })
            }
            template.fun();
        }
        privateFun.delete = function (status,arg) {
            var template = {
                modal:{
                    title:status == 'batch' ?'批量删除插件':('删除插件-' + arg.item.pluginName)
                },
                request: {
                    strategyID: vm.ajaxRequest.strategyID,
                    connIDList: status == 'batch' ? vm.data.batch.query.join(',') : arg.item.connID
                },
                loop:{
                    num:0
                }
            }
            $rootScope.EnsureModal(template.modal.title, null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.PluginStrategy.Delete(template.request).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    switch (status) {
                                        case 'batch':
                                            {
                                                privateFun.resetBatchInfo();
                                                $rootScope.InfoModal(template.modal.title + '成功', 'success');
                                                vm.fun.init();
                                                break;
                                            }
                                        case 'single':
                                            {
                                                $rootScope.InfoModal(template.modal.title + '成功', 'success');
                                                vm.ajaxResponse.query.splice(arg.$index, 1);
                                                break;
                                            }
                                    }
                                    
                                    break;
                                }
                        }
                    })
                }
            });
        }
        privateFun.operate = function (arg) {
            var template = {
                promise: null
            }
            GatewayResource.PluginStrategy[arg.status]({
                strategyID: vm.ajaxRequest.strategyID,
                connIDList: arg.item.connID
            }).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.fun.init();
                            $rootScope.InfoModal((arg.status == 'Start' ? '启用' : '停用') + arg.item.pluginName + '成功！', 'success');
                            break;
                        }
                }
            })
        }
        privateFun.conditionFilter = function (arg) {
            if (vm.component.menuObject.active.condition == arg.item.value) return;
            vm.component.menuObject.active.condition = arg.item.value;
            vm.ajaxRequest.ids = null;
            $scope.$broadcast('$Init_LoadingCommonComponent');
        }
        privateFun.resetBatchInfo = function () {
            vm.data.batch.isOperating = false;
            vm.data.batch.selectAll = false;
            vm.data.batch.query = [];
            vm.data.batch.indexAddress = {};
        };
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['策略插件列表', '策略']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey:'connID',
                    default: [{
                            key: '插件名称',
                            html: '{{item.pluginName}}'
                        },
                        {
                            key: '描述',
                            html: '{{item.pluginDesc}}'
                        },
                        {
                            key: '状态',
                            html: '<div ng-switch="item.pluginStatus"><span class="eo-status-tips" ng-switch-when=-1>扩展插件处尚未开启该插件</span><span class="eo-status-warning" ng-switch-when=0>停用</span><span class="eo-status-success" ng-switch-when=1>启用</span></div>'
                        },, {
                            key: '更新时间',
                            html: '{{item.updateTime}}'
                        }
                    ],
                    operate: {
                        funArr: [{
                            key: '停用',
                            fun: privateFun.operate,
                            itemExpression:'ng-if="item.pluginStatus===1"',
                            params: {
                                status: 'Stop'
                            }
                        },
                        {
                            key: '启用',
                            fun: privateFun.operate,
                            itemExpression:'ng-if="item.pluginStatus!==1" ng-disabled="item.pluginStatus===-1"',
                            params: {
                                status: 'Start'
                            }
                        },
                        {
                            key: '修改',
                            fun: privateFun.edit,
                            params: {
                                status: 'edit'
                            }
                        },
                        {
                            key: '删除',
                            fun: privateFun.delete,
                            params: '"single",arg'
                        }
                    ],
                        class:'w_200'
                    }
                },
                setting: {
                    unhover: true,
                    batch:true,
                    defaultFoot:true
                }
            }
            vm.component.menuObject = {
                list: [{
                        type: 'btn',
                        authority: 'edit',
                        class: 'pull-left',
                        btnList: [{
                            name: '添加插件',
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
                        placeholder:"输入插件名称或描述"
                    },{
                        type: 'fun-list',
                        class: 'fun-list-li  pull-right mr15',
                        name: '筛选',
                        icon: 'shaixuan',
                        activePoint: 'condition',
                        funList: [{
                            name: '无',
                            value: 0,
                            active: 0,
                            fun: {
                                default: privateFun.conditionFilter
                            }
                        },{
                            name: '启用',
                            value: 2,
                            active: 2,
                            fun: {
                                default: privateFun.conditionFilter
                            }
                        }, {
                            name: '停用',
                            value: 1,
                            active: 1,
                            fun: {
                                default: privateFun.conditionFilter
                            }
                        }]
    
                    }],
                batchList: [{
                    type: 'btn',
                    class: 'pull-left',
                    btnList: [{
                        name: '删除',
                        fun: {
                            default: privateFun.delete,
                            params: '"batch"'
                        }
                    }, {
                        name: '启用',
                        fun: {
                            default: privateFun.batchOperate,
                            params: {
                                status: 'Start'
                            }
                        }
                    }, {
                        name: '停用',
                        fun: {
                            default: privateFun.batchOperate,
                            params: {
                                status: 'Stop'
                            }
                        }
                    }]
                }],
                setting: {
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    titleAuthority: 'showTitle',
                    title: '策略插件'
                },
                active:{
                    condition:0
                }
            }
        }
    }
})();