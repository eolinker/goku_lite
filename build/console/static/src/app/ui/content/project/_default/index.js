(function () {
    'use strict';
    angular.module('eolinker')
        .component('projectDefault', {
            templateUrl: 'app/ui/content/project/_default/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', '$rootScope', 'CODE', 'Authority_CommonService'];

    function indexController($scope, GatewayResource, $state, $rootScope, CODE, Authority_CommonService) {
        var vm = this;
        vm.data = {
            batch: {}
        }
        vm.ajaxRequest = {
            projectID: [],
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
        var privateFun = {};

        vm.fun.init = function () {
            let tmpAjaxRequest={};
            if(vm.ajaxRequest.keyword){
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_Project = GatewayResource.Project.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_Project.$promise.then(function (response) {
                vm.ajaxResponse.query = response.projectList || [];
            })
            return $rootScope.global.ajax.Query_Project.$promise;
        }
        privateFun.import = function () {
            var tmpModal={
                title: '导入项目',
                resource: GatewayResource.ImportAms.Project
            };
            $rootScope.ImportModal(tmpModal, function (callback) {
                if (callback) {
                    $rootScope.InfoModal(tmpModal.title + '成功', 'success');
                    $scope.$broadcast('$Init_LoadingCommonComponent');
                }
            });
        }
        privateFun.edit = function (inputArg) {
            let tmpAjaxRequest={},tmpModal={
                title: (inputArg.status == 'edit' ? '修改' : '新增') + '项目',
                placeholder: '项目名称',
                text: inputArg.item.projectName || ''
            };
            $rootScope.Common_SingleInputModal(tmpModal, function (callback) {
                if (callback) {
                    tmpAjaxRequest.projectName = callback.text;
                    switch (inputArg.status) {
                        case 'add': {
                            GatewayResource.Project.Add(tmpAjaxRequest)
                                .$promise.then(function (response) {
                                    switch (response.statusCode) {
                                        case CODE.COMMON.SUCCESS: {
                                            $rootScope.InfoModal(tmpModal.title + '成功', 'success');
                                            vm.fun.init();
                                            break;
                                        }
                                    }
                                })
                            break;
                        }
                        case 'edit': {
                            tmpAjaxRequest.projectID = inputArg.item.projectID;
                            GatewayResource.Project.Edit(tmpAjaxRequest)
                                .$promise.then(function (response) {
                                    switch (response.statusCode) {
                                        case CODE.COMMON.SUCCESS: {
                                            $rootScope.InfoModal(tmpModal.title + '成功', 'success');
                                            vm.fun.init();
                                            break;
                                        }
                                    }
                                })
                            break;
                        }
                    }
                }
            })
        }
        privateFun.delete = function (inputArg) {
            let tmpAjaxRequest={
                projectIDList: inputArg.status == 'batch' ? vm.data.batch.query.join(',') : inputArg.item.projectID
            },tmpModal={
                title: inputArg.status == 'batch' ? '批量删除项目' : ('删除项目-' + inputArg.item.projectName)
            }
            $rootScope.EnsureModal(tmpModal.title, null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.Project.Delete(tmpAjaxRequest).$promise.then(function (response) {
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
        privateFun.click = function (arg) {
            $state.go('home.project.api.default', {
                projectID: arg.item.projectID,
                projectName: arg.item.projectName
            });
        }
        privateFun.resetBatchInfo = function () {
            vm.data.batch.isOperating = false;
            vm.data.batch.selectAll = false;
            vm.data.batch.query = [];
            vm.data.batch.groupQuery = [];
            vm.data.batch.indexAddress = {};
            vm.data.batchOprArchive = [];
        };
        privateFun.search = function (arg) {
            window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
            $state.reload($state.current.name);
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['项目列表']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'projectID',
                    resource: GatewayResource.Project,
                    default: [{
                            key: '名称',
                            html: '{{item.projectName}}',
                            contentClass: 'item-title-li'
                        },
                        {
                            key: '更新时间',
                            html: '{{item.updateTime}}',
                            class:'w_250',
                            isUnneccessary: true
                        }
                    ],
                    operate: {
                        funArr: [
                            {
                                key: '编辑',
                                fun: privateFun.edit,
                                params: {
                                    status: 'edit',
                                }
                            },
                            {
                                key: '删除',
                                fun: privateFun.delete,
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
                    warning: '尚未新建任何项目',
                    defaultFoot:true
                },
                baseFun: {
                    click: privateFun.click,
                    batchSelectAll: (inputOpr, inputItem) => {
                        switch (inputOpr) {
                            case 'select': {
                                switch (inputItem.isArchive) {
                                    case 1: {
                                        let tmpIndex = vm.data.batchOprArchive.indexOf(inputItem.projectHashKey);
                                        if (tmpIndex === -1) {
                                            vm.data.batchOprArchive.push(inputItem.projectHashKey)
                                        }
                                        break;
                                    }
                                }
                                break;
                            }
                            case 'cancel': {
                                vm.data.batchOprArchive = [];
                                break;
                            }
                        }
                    },
                    batchItemClick: (inputItem) => {
                        switch (inputItem.isArchive) {
                            case 1: {
                                let tmpIndex = vm.data.batchOprArchive.indexOf(inputItem.projectHashKey);
                                if (tmpIndex === -1) {
                                    vm.data.batchOprArchive.push(inputItem.projectHashKey)
                                } else {
                                    vm.data.batchOprArchive.splice(tmpIndex, 1);
                                }
                                break;
                            }
                        }
                    }
                }
            }
            vm.component.menuObject = {
                list: [{
                        type: 'btn',
                        authority: 'edit',
                        class: 'pull-left',
                        btnList: [{
                            name: '新建项目',
                            icon: 'jiahao',
                            class: 'eo_theme_btn_success block-btn',
                            fun: {
                                default: privateFun.edit,
                                params: {
                                    status:'add'
                                }
                            }
                        }]
                    },
                    {
                        type: 'btn',
                        authority: 'edit',
                        class: 'pull-left',
                        btnList: [{
                            name: '导入',
                            icon: 'yunshangchuan',
                            fun: {
                                default: privateFun.import
                            }
                        }]
                    },{
                        type: 'search',
                        class: 'pull-right mr15',
                        keyword: vm.ajaxRequest.keyword,
                        fun: privateFun.search,
                        placeholder:"输入项目名称"
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
                    title:"接口管理",
                    titleAuthority:"showTitle"
                }
            }
        }
    }
})();