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
        .component('gpeditInsideApiPlugin', {
            templateUrl: 'app/ui/content/gpedit/inside/content/apiPlugin/index.html',
            controller: indexController
        })

    indexController.$inject = ['GatewayResource','$scope', '$state', '$rootScope', 'CODE', 'Authority_CommonService'];

    function indexController(GatewayResource,$scope, $state, $rootScope, CODE, Authority_CommonService) {
        var vm = this;
        vm.data = {
            batch: {}
        }
        vm.fun={};
        vm.ajaxRequest={
            strategyID: $state.params.strategyID,
            apiID:$state.params.apiID,
            keyword: ""
        };
        vm.ajaxResponse={
            query: null
        }
        vm.component = {
            listDefaultCommonObject: null,
            menuObject: {}
        }
        vm.service = {
            authority: Authority_CommonService
        }
        var privateFun = {};
        privateFun.search = function (arg) {
            vm.ajaxRequest.keyword=arg.item.keyword;
            vm.fun.init();
        }
        vm.fun.init = function () {
            var tmpAjaxRequest={
                strategyID: vm.ajaxRequest.strategyID,
                apiID:vm.ajaxRequest.apiID,
                condition:vm.component.menuObject.active.condition
            }
            if(vm.ajaxRequest.keyword){
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_GpeditPluginApi = GatewayResource.GpeditPluginApi.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_GpeditPluginApi.$promise.then(function (response) {
                vm.ajaxResponse.query = response.apiPluginList || [];
                vm.ajaxResponse.apiInfo=response.apiInfo || {
                    apiName:"未知API",
                    requestURL:"未知URL"
                };
            })
            return $rootScope.global.ajax.Query_GpeditPluginApi.$promise;
        }
        privateFun.edit = function (arg) {
            arg.item=arg.item||{};
            $rootScope.Gateway_GpeditApiPluginModal({
                oprType:arg.status,
                pluginName:arg.item.pluginName||"",
                chineseName:arg.item.pluginDesc||""
            },(callback)=>{
                if(callback){
                    vm.fun.init();
                }
            });
        }
        privateFun.batchOperate = function (arg) {
            var template = {
                promise: null,
                title: vm.data.batch.isOperating ? '批量' : '',
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
                GatewayResource.PluginApi[arg.status]({
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
                }
            }
            $rootScope.EnsureModal(template.modal.title, null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.PluginApi.Delete(template.request).$promise.then(function (response) {
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
            GatewayResource.PluginApi[arg.status]({
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
        privateFun.resetBatchInfo = function () {
            vm.data.batch.isOperating = false;
            vm.data.batch.selectAll = false;
            vm.data.batch.query = [];
            vm.data.batch.indexAddress = {};
        };
        privateFun.conditionFilter = function (arg) {
            if (vm.component.menuObject.active.condition == arg.item.value) return;
            vm.component.menuObject.active.condition = arg.item.value;
            vm.ajaxRequest.ids = null;
            $scope.$broadcast('$Init_LoadingCommonComponent');
        }
        vm.$onInit = function () {
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
                        }, {
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
            vm.component.menuDefaultObject={
                list:[{
                    type: 'btn',
                    class: 'btn-group-li',
                    btnList: [{
                        name: '返回API列表',
                        icon: 'chexiao',
                        fun: {
                            default: ()=>{
                                $state.go("home.gpedit.inside.api.default");
                            }
                        }
                    }]
                },{
                    type:"customized-html",
                    html:`<p class="fs18 mt15">{{$ctrl.otherObject.requestURL}}</p><p class="fs18 mt10">{{$ctrl.otherObject.apiName}}</p>`
                }],
                setting:{
                    class:'common-menu-fixed-seperate common-menu-lg'
                }
            };
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