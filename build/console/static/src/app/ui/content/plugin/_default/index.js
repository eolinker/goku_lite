(function () {
    'use strict';
    angular.module('eolinker')
        .component('pluginDefault', {
            templateUrl: 'app/ui/content/plugin/_default/index.html',
            controller: indexController
        })

    indexController.$inject = [ '$scope', 'GatewayResource', '$state', '$rootScope', 'CODE', 'Authority_CommonService'];

    function indexController( $scope, GatewayResource, $state, $rootScope, CODE, Authority_CommonService) {
        var vm = this;
        vm.data = {
            batch: {
                address: []
            }
        }
        vm.ajaxResponse={}
        vm.service = {
            authority: Authority_CommonService
        }
        vm.ajaxRequest={
            pluginName: [],
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        }
        vm.component = {
            menuObject: {},
            listRequireObject: null
        }
        vm.fun={};
        var privateFun = {},
            data = {
                checking: false
            }
        vm.fun.init = function () {
            var tmpAjaxRequest={
                condition:vm.component.menuObject.active.condition
            }
            if(vm.ajaxRequest.keyword){
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_Plugin = GatewayResource.Plugin.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_Plugin.$promise.then(function (response) {
                vm.ajaxResponse.query = response.pluginList || [];
                vm.ajaxResponse.query.map((val)=>{
                    val.isAlreadyStart=val.pluginStatus===1;
                })
            })
            return $rootScope.global.ajax.Query_Plugin.$promise;
        }
        privateFun.conditionFilter = function (arg) {
            if (vm.component.menuObject.active.condition == arg.item.value) return;
            vm.component.menuObject.active.condition = arg.item.value;
            vm.ajaxRequest.ids = null;
            $scope.$broadcast('$Init_LoadingCommonComponent');
        }
        privateFun.search = function (arg) {
            window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
            $state.reload($state.current.name);
        }
        privateFun.check = function (arg) {
            if (data.checking) {
                $rootScope.InfoModal('当前已有插件正在进行检测，请稍后再试', 'error');
                return;
            }
            var template = {
                request: {
                    pluginName: arg.item.pluginName
                },
                pluginStatus:arg.item.pluginStatus
            }
            data.checking=true;
            vm.ajaxResponse.query[arg.$index].pluginStatus=2;
            GatewayResource.Plugin.Check(template.request).$promise.then(function (response) {
                data.checking=false;
                vm.ajaxResponse.query[arg.$index].pluginStatus=template.pluginStatus;
                vm.ajaxResponse.query[arg.$index].isAlreadyStart=template.pluginStatus;
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            $rootScope.InfoModal('检测在全部节点内检测成功，插件可正常使用', 'success');
                            vm.ajaxResponse.query[arg.$index].isCheck=true;
                            break;
                        }
                        case '210000':{
                            $rootScope.Gateway_NodeCheckErrorReportModal({
                                query:response.errNodeList,
                                pluginName:arg.item.pluginName
                            }, 'success');
                            vm.ajaxResponse.query[arg.$index].isCheck=false;
                            break;
                        }
                }
            })
        }
        privateFun.batchOperate = function (status) {
            var template = {
                tip: '',
                resource: null,
                pluginStatus: 0
            }
            switch (status) {
                case 'start':
                    {
                        template.tip = '开启';
                        template.resource = GatewayResource.Plugin.BatchStart;
                        template.pluginStatus = 1;
                        break;
                    }
                case 'stop':
                    {
                        template.tip = '关闭';
                        template.resource = GatewayResource.Plugin.BatchStop;
                        break;
                    }
            }
            template.resource({
                pluginNameList: vm.data.batch.query.join(',')
            }).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            
                            $rootScope.InfoModal('批量' + template.tip + '插件成功！', 'success');
                            $state.reload('home.plugin.default');
                            break;
                        }
                }
            })
        }
        privateFun.operate = function (arg) {
            GatewayResource.Plugin[arg.status]({
                pluginName: arg.item.pluginName
            }).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.ajaxResponse.query[arg.$index].pluginStatus = arg.status == 'Start' ? 1 : 0;
                            vm.ajaxResponse.query[arg.$index].isAlreadyStart=vm.ajaxResponse.query[arg.$index].pluginStatus;
                            vm.ajaxResponse.query[arg.$index].isCheck = 0;
                            $rootScope.InfoModal((arg.status == 'Start' ? '开启' : '关闭') + arg.item.pluginName + '成功！', 'success');
                            break;
                        }
                }
            })
        }
        privateFun.edit = function (arg) {
            var template = {
                uri: {
                    status: arg.status
                }
            }
            switch (arg.status) {
                case 'edit':
                    {
                        template.uri.pluginName = arg.item.pluginName;
                        break;
                    }
            }
            $state.go('home.plugin.operate', template.uri);
        }
        privateFun.delete = function (arg) {
            var template = {
                request: {
                    pluginName: arg.item.pluginName
                }
            }
            $rootScope.EnsureModal('删除' + arg.item.pluginName, null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.Plugin.Delete(template.request).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    vm.ajaxResponse.query.splice(arg.$index, 1);
                                    $rootScope.InfoModal(arg.item.pluginName + '删除成功', 'success');
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
                list: ['插件管理']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'pluginName',
                    default: [{
                        key: '名称',
                        html: '{{item.pluginName}}'
                    }, {
                        key: '描述',
                        html: '{{item.pluginDesc}}'
                    }
                    // , {
                    //     key: '类别',
                    //     html: '<span ng-switch-when=false>自定义插件</span><tip-directive ng-if="!item.official" input="如需让自定义的插件生效，必须先重启/重载网关"></tip-directive><span ng-switch-when=true>官方插件</span>',
                    //     switch: 'official',
                    //     keyStyle: {
                    //         'width': '120px'
                    //     }
                    // }
                    , {
                        key: '插件类型',
                        html: '<span ng-switch-when="0">全局</span><span ng-switch-when="1">策略</span><span ng-switch-when="2">API</span>',
                        switch: 'pluginType',
                        class:'w_100'
                    }
                    // , {
                    //     key: '版本号',
                    //     html: '{{item.version}}',
                    //     keyStyle: {
                    //         'width': '90px'
                    //     }
                    // }
                    , {
                        key: '优先级（0-3000）',
                        html: '{{item.pluginPriority}}',
                        class:'w_150'
                    }, {
                        key: '状态',
                        html: '<span class="eo-status-tips" ng-switch-when=0>关闭</span><span class="eo-status-success" ng-switch-when=1>开启</span><span class="eo-status-default" ng-switch-when=2>检测中</span><span class="eo-status-default iconfont icon-jiazai_shuaxin" ng-switch-when=2></span>',
                        switch: 'pluginStatus',
                        class:'w_80'
                    }],
                    operate: {
                        funArr: [{
                                key: '配置',
                                itemExpression:'ng-if="item.pluginType===0"',
                                fun: privateFun.edit,
                                params: {
                                    status: 'edit'
                                }
                            }, {
                                key: '关闭',
                                fun: privateFun.operate,
                                itemExpression:`ng-if="item.isAlreadyStart"`,
                                params: {
                                    status: 'Stop'
                                }
                            }, {
                                key: '开启',
                                fun: privateFun.operate,
                                itemExpression:'ng-if="!item.isAlreadyStart" ng-disabled="!item.isCheck"',
                                params: {
                                    status: 'Start'
                                }
                            },
                            {
                                key: '<span>检测</span><tip-directive input="开启插件前，需先检测插件是否可用，如放置路径是否正确等"></tip-directive>',
                                itemExpression:`ng-disabled="item.isAlreadyStart"`,
                                fun: privateFun.check
                            },

                            {
                                key: '删除',
                                fun: privateFun.delete
                            }
                        ],
                        power: -1,
                        class:'w_250'
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
                            name: '全局插件',
                            value: 1,
                            active: 1,
                            fun: {
                                default: privateFun.conditionFilter
                            }
                        }, {
                            name: '策略插件',
                            value: 2,
                            active: 2,
                            fun: {
                                default: privateFun.conditionFilter
                            }
                        }, {
                            name: '接口插件',
                            value: 3,
                            active: 3,
                            fun: {
                                default: privateFun.conditionFilter
                            }
                        }]
    
                    }
                ],
                batchList: [{
                    type: 'btn',
                    disabledPoint: 'isBatchSelected',
                    class: 'pull-left',
                    btnList: [{
                        name: '批量开启',
                        show: true,
                        disabled: 0,
                        fun: {
                            default: privateFun.batchOperate,
                            params: '"start"'
                        }
                    }, {
                        name: '批量关闭',
                        show: true,
                        disabled: 0,
                        fun: {
                            default: privateFun.batchOperate,
                            params: '"stop"'
                        }
                    }]
                }],
                setting: {
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    class: "common-menu-fixed-seperate common-menu-lg",
                    titleAuthority: 'showTitle',
                    title: '扩展插件'
                },
                active:{
                    condition:0
                }
            }
        }
    }
})();