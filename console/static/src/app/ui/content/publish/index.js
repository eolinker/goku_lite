(function () {
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.publish', {
                    url: '/publish',
                    template: '<publish></publish>'
                })
        }])
        .component('publish', {
            templateUrl: 'app/ui/content/publish/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', '$rootScope', 'CODE', 'Authority_CommonService'];

    function indexController($scope, GatewayResource, $state, $rootScope, CODE, Authority_CommonService) {
        var vm = this;
        vm.data = {
            batch: {},
            alreadyHadPublishConf:false //是否已经存在已发布的配置
        }
        vm.ajaxRequest = {
            versionID: [],
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        };
        vm.ajaxResponse = {};
        vm.fun = {};
        vm.service = {
            authority: Authority_CommonService
        }
        vm.component = {
            menuObject: null
        }
        var privateFun = {},cache={
            publishQuery:[]
        };

        vm.fun.init = function () {
            let tmpAjaxRequest = {};
            if (vm.ajaxRequest.keyword) {
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            vm.data.alreadyHadPublishConf=false;
            $rootScope.global.ajax.Query_Version = GatewayResource.Version.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_Version.$promise.then(function (response) {
                vm.ajaxResponse.query = response.configList || [];
                if(vm.ajaxResponse.query.length>0&&vm.ajaxResponse.query[0].publishStatus===1){
                    vm.data.alreadyHadPublishConf=true;
                }
            })
            return $rootScope.global.ajax.Query_Version.$promise;
        }
        privateFun.generateConf = function () {
            let tmpModal = {
                title: "生成配置",
                resource: GatewayResource.Version.Add,
                btnObject:{
                    text:"保存"
                },
                tmpBtnObj:{
                    ajaxRequest:{
                        publish:1
                    },
                    text:"保存并发布"
                },
                textArray: [{
                    type: 'input',
                    title: "名称",
                    key:"name",
                    placeholder:"配置名称",
                    required:true
                },{
                    type: 'input',
                    title: '版本号',
                    key:"version",
                    placeholder:"配置版本号,如1.0"
                },{
                    type: 'input',
                    title: '备注',
                    key:"remark",
                    placeholder:"配置备注"
                }]
            }
            $rootScope.MixInputModal(tmpModal,(callback)=>{
                if(callback){
                    vm.fun.init();
                    $rootScope.InfoModal(`${tmpModal.title}成功`,'success');
                }
            })
        }
        privateFun.publish = function (inputArg) {
            let tmpModal = {
                    title: "配置管理"
                },tmpAjaxRequest={
                    versionID:inputArg.item.versionID
                }
            $rootScope.EnsureModal(tmpModal.title, null, '配置发布立即生效，确定对各节点配置管理？', {
                btnType:2,
                btnMessage:"确定"
            }, function (callback) {
                if (callback) {
                    GatewayResource.Version.Publish(tmpAjaxRequest).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                vm.fun.init();
                                $rootScope.InfoModal(tmpModal.title + '成功', 'success');
                                break;
                            }
                        }
                    })
                }
            });
        }
        privateFun.delete = function (inputArg) {
            let tmpAjaxRequest = {
                    ids: inputArg.status == 'batch' ? vm.data.batch.query : [inputArg.item.versionID]
                },
                tmpModal = {
                    title: "删除配置"
                };
            $rootScope.EnsureModal(tmpModal.title, null, '确认删除配置？', {}, function (callback) {
                if (callback) {
                    tmpAjaxRequest.ids=JSON.stringify(tmpAjaxRequest.ids);
                    GatewayResource.Version.Delete(tmpAjaxRequest).$promise.then(function (response) {
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
                                $rootScope.InfoModal(tmpModal.title + '成功', 'success');
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
            vm.ajaxResponse.query=cache.publishQuery.concat(vm.ajaxResponse.query);
        };
        privateFun.search = function (arg) {
            window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
            $state.reload($state.current.name);
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['配置列表']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'versionID',
                    resource: GatewayResource.Version,
                    default: [{
                        key: '状态',
                        html: `<span ng-class="{'c999':!item.publishStatus,'eo-status-success':item.publishStatus}">{{item.publishStatus?"已发布":"未发布"}}</span>`,
                        class: 'w_80'
                    },{
                            key: '名称',
                            html: '{{item.name}}',
                            contentClass: 'item-title-li'
                        },{
                            key: '版本号',
                            html: '{{item.version}}',
                            contentClass: 'item-title-li'
                        },{
                            key: '备注',
                            html: '{{item.remark}}',
                            contentClass: 'item-title-li'
                        },
                        {
                            key: '保存时间',
                            html: '{{item.createTime}}',
                            class: 'w_150'
                        },
                        {
                            key: '发布时间',
                            html: '{{item.publishTime}}',
                            class: 'w_150',
                            isUnneccessary: true
                        }
                    ],
                    operate: {
                        funArr: [{
                                key: '发布',
                                fun: privateFun.publish,
                                itemExpression:`ng-disabled="item.publishStatus"`
                            },
                            {
                                key: '删除',
                                fun: privateFun.delete,
                                itemExpression:`ng-disabled="item.publishStatus"`,
                                params: {
                                    status: 'single'
                                }
                            }
                        ],
                        class: 'w_150'
                    }
                },
                setting: {
                    batch: true,
                    warning: '尚未生成任何配置',
                    defaultFoot: true,
                    unhover:true
                }
            }
            vm.component.menuObject = {
                list: [{
                        type: 'btn',
                        authority: 'edit',
                        class: 'pull-left',
                        btnList: [{
                            name: '生成配置',
                            icon: 'jiahao',
                            class: 'eo_theme_btn_success block-btn',
                            fun: {
                                default: privateFun.generateConf,
                                params: {
                                    status: 'add'
                                }
                            }
                        }]
                    }, {
                        type: 'search',
                        class: 'pull-right mr15',
                        keyword: vm.ajaxRequest.keyword,
                        fun: privateFun.search,
                        placeholder: "输入搜索内容"
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
                baseFun: {
                    batchDefault: () => {
                        if(vm.data.alreadyHadPublishConf){
                            cache.publishQuery=[vm.ajaxResponse.query.shift()]
                        }
                    }
                },
                setting: {
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    class: "common-menu-fixed-seperate common-menu-lg",
                    title: "配置管理",
                    secondTitle:"网关的配置内容支持版本管理，可以对配置进行发布和回滚<br/><br/>注意：配置发布后会立即生效，请谨慎操作",
                    titleAuthority: "showTitle"
                }
            }
        }
    }
})();