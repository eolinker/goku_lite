(function () {
    'use strict';
    angular.module('eolinker')
        .component('apiDefault', {
            templateUrl: 'app/ui/content/project/api/_default/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'CommonResource','GatewayResource', '$state', '$rootScope', 'CODE', 'GroupService', 'Authority_CommonService'];

    function indexController($scope,CommonResource, GatewayResource, $state, $rootScope, CODE, GroupService, Authority_CommonService) {
        var vm = this;
        vm.data = {
            batch: {},
            pagination: {
                maxSize: 10,
                pageSize: 50,
                page: 0,
                msgCount: 0
            }
        }
        vm.fun = {};
        vm.ajaxRequest = {
            projectID: $state.params.projectID,
            groupID: $state.params.groupID || -1,
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        };
        vm.ajaxResponse = {
            query: null
        }
        vm.service = {
            authority: Authority_CommonService
        }
        vm.component = {
            listRequireObject: {},
            menuObject: {}
        }
        var privateFun = {},data={};
        vm.fun.init = function () {
            return privateFun.getQuery('reset');
        }
        privateFun.batchBtnClickFun = () => {
            if (vm.data.pagination.page === Math.ceil(vm.data.pagination.msgCount / vm.data.pagination.pageSize)) return;
            var tmpAjaxRequest={
                projectID: vm.ajaxRequest.projectID,
                groupID: vm.ajaxRequest.groupID,
                condition: vm.component.menuObject.active.condition
            },tmpPromise=null;
            if (vm.ajaxRequest.keyword) {
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            } 
            tmpPromise = GatewayResource.Api.IDQuery(tmpAjaxRequest).$promise;
            tmpPromise.then(function (response) {
                vm.data.allQueryID = response.apiIDList||[];
            })
        }
        privateFun.getQuery = function (inputOpr,inputPage,inputPageSize) {
            var tmpAjaxRequest={
                projectID: vm.ajaxRequest.projectID,
                groupID: vm.ajaxRequest.groupID,
                condition: vm.component.menuObject.active.condition,
                pageSize: inputPageSize||vm.data.pagination.pageSize
            },tmpResponseArr=[];
            switch(inputOpr){
                case 'reset':{
                    tmpAjaxRequest.page=1;
                    break;
                }
                default:{
                    tmpResponseArr=vm.ajaxResponse.query;
                    tmpAjaxRequest.page=inputPage;
                    break;
                }
            }
            if (vm.ajaxRequest.keyword) {
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_Api = GatewayResource.Api.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_Api.$promise.then(function (response) {
                if(inputOpr==='reset'){
                    document.getElementsByClassName('tbody_container_ldcc')[0].scrollTop = 0;
                }
                data.isQuerying=false;
                vm.ajaxResponse.query = tmpResponseArr.concat(response.apiList || []);
                vm.data.pagination.msgCount = response.page.totalNum||0;
                if(!inputPageSize)vm.data.pagination.page=tmpAjaxRequest.page;
            })
            return $rootScope.global.ajax.Query_Api.$promise;
        }
        privateFun.scrollLoading = function () {
            let tmpFlag = {
                hasItem: vm.ajaxResponse.query && vm.ajaxResponse.query.length !== 0,
                hasNextPage: vm.data.pagination.page < (vm.data.pagination.msgCount / vm.data.pagination.pageSize),
                isQuerying: data.isQuerying
            }
            
            if (tmpFlag.hasItem && tmpFlag.hasNextPage && !tmpFlag.isQuerying) {
                data.isQuerying=true;
                privateFun.getQuery('preload',vm.data.pagination.page+1);
            }
        }
        privateFun.search = function (arg) {
            window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
            $state.reload($state.current.name);
        }
        privateFun.import = function () {
            var template = {
                modal: {
                    title: '导入接口',
                    request: {
                        projectID: vm.ajaxRequest.projectID,
                        groupID: vm.ajaxRequest.groupID==-1?0:vm.ajaxRequest.groupID
                    },
                    resource: GatewayResource.ImportAms.Api
                }
            }
            $rootScope.ImportModal(template.modal, function (callback) {
                if (callback) {
                    $rootScope.InfoModal(template.modal.title + '成功', 'success');
                    $scope.$broadcast('$Init_LoadingCommonComponent');
                }
            });
        }
        privateFun.edit = function (arg) {
            var template = {
                uri: {
                    status: arg.status,
                    groupID: vm.ajaxRequest.groupID
                }
            }
            switch (arg.status) {
                case 'edit': {
                    template.uri.apiID = arg.item.apiID;
                    break;
                }
            }
            $state.go('home.project.api.operate', template.uri);
        }
        privateFun.getDefaultBalance=()=>{
           let tmpPromise= GatewayResource.Balance.SimpleQuery().$promise;
           tmpPromise.then((response)=>{
                vm.ajaxResponse.balanceList=response.balanceNames;
            })
            return tmpPromise;
        }
        privateFun.delete = function (status, arg) {
            var template = {
                request: {
                    projectID: vm.ajaxRequest.projectID,
                    apiIDList: status == 'batch' ? vm.data.batch.query.join(',') : arg.item.apiID
                }
            }
            $rootScope.EnsureModal('删除API', null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.Api.Delete(template.request).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                switch (status) {
                                    case 'batch': {
                                        privateFun.resetBatchInfo();
                                        privateFun.getQuery("reset");
                                        break;
                                    }
                                    case 'single': {
                                        vm.ajaxResponse.query.splice(arg.$index, 1);
                                        if(vm.data.pagination.msgCount>vm.data.pagination.page*vm.data.pagination.pageSize)privateFun.getQuery('preload',vm.data.pagination.page*vm.data.pagination.pageSize,1)
                                        vm.data.pagination.msgCount--;
                                        break;
                                    }
                                }
                                $rootScope.InfoModal('API删除成功', 'success');
                                break;
                            }
                        }
                    })
                }
            });
        }
        privateFun.batchMoveGroup = function () {
            var template = {
                modal: {
                    list: [{groupID:0,groupName:"未分组"}].concat(angular.copy(GroupService.get())),
                    title: '批量编辑接口分组',
                    current: vm.ajaxRequest
                },
                request: {
                    projectID: vm.ajaxRequest.projectID,
                    apiIDList: vm.data.batch.query.join(','),
                    groupID: ''
                }
            }
            $rootScope.SelectVisualGroupModal(template.modal, function (callback) {
                if (callback) {
                    template.request.groupID = callback.groupID;
                    GatewayResource.Api.ChangeGroup(template.request).$promise
                        .then(function (response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS: {
                                    privateFun.resetBatchInfo();
                                    privateFun.getQuery("reset");
                                    $rootScope.InfoModal('Api批量编辑分组成功', 'success');
                                    break;
                                }
                            }
                        })
                }
            })
        }
        privateFun.copy = (inputArg) => {
            let tmpFunCopy=(inputArr)=>{
                let tmpItem = angular.copy(inputArg.item),
                tmpModal = {
                    apiName: "复制-" + tmpItem.apiName,
                    requestURL: tmpItem.requestURL,
                    balanceName: tmpItem.target,
                    targetURL: tmpItem.targetURL,
                    targetMethod: tmpItem.isFollow?'-1':tmpItem.targetMethod.toLowerCase(),
                    protocol: tmpItem.protocol||"http",
                    apiID:tmpItem.apiID,
                    groupID:tmpItem.groupID||"0",
                    projectID:vm.ajaxRequest.projectID,
                    balanceList:inputArr,
                    requestMethodList: [{
                        key: 'POST',
                        value: 'post'
                    }, {
                        key: 'GET',
                        value: 'get'
                    }, {
                        key: 'PUT',
                        value: 'put'
                    }, {
                        key: 'DELETE',
                        value: 'delete'
                    }, {
                        key: 'HEAD',
                        value: 'head'
                    }, {
                        key: 'OPTIONS',
                        value: 'options'
                    }, {
                        key: 'PATCH',
                        value: 'patch'
                    }]
                };
            let tmpRequestMethod = tmpItem.requestMethod.split(',');
            for (var key in tmpRequestMethod) {
                var val = tmpRequestMethod[key];
                switch (val.toLowerCase()) {
                    case 'post': {
                        tmpModal.requestMethodList[0].checkbox = true;
                        break;
                    }
                    case 'get': {
                        tmpModal.requestMethodList[1].checkbox = true;
                        break;
                    }
                    case 'put': {
                        tmpModal.requestMethodList[2].checkbox = true;
                        break;
                    }
                    case 'delete': {
                        tmpModal.requestMethodList[3].checkbox = true;
                        break;
                    }
                    case 'head': {
                        tmpModal.requestMethodList[4].checkbox = true;
                        break;
                    }
                    case 'options': {
                        tmpModal.requestMethodList[5].checkbox = true;
                        break;
                    }
                    case 'batch': {
                        tmpModal.requestMethodList[6].checkbox = true;
                        break;
                    }
                }
            }
            $rootScope.Gateway_CopyApiModal(tmpModal,(callback)=>{
                if(callback){
                    $rootScope.InfoModal('Api复制成功', 'success');
                    privateFun.getQuery("reset");
                }
            });
            }
            if(vm.ajaxResponse.balanceList){
                tmpFunCopy(vm.ajaxResponse.balanceList)
            }else{
                let tmpPromise=privateFun.getDefaultBalance();
                tmpPromise.finally(()=>{
                    tmpFunCopy(vm.ajaxResponse.balanceList)
                })
            }
        }
        privateFun.resetBatchInfo = function () {
            vm.data.batch.isOperating = false;
            vm.data.batch.selectAll = false;
            vm.data.batch.query = [];
            vm.data.batch.indexAddress = {};
        };
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['APIs列表']
            });

            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'apiID',
                    default: [{
                        key: 'APIs',
                        html: `<span class="mr5 plr3 ptb2 fs12" ng-class="{'eo-label-warning':item.apiType,'eo-label-success':!item.apiType}">{{item.apiType?"编排":"普通"}}</span><span>{{item.apiName}}</span>`,
                        draggableCacheMark: 'name'
                    }, {
                        key: '请求方式',
                        html: '{{item.requestMethod}}',
                        draggableCacheMark: 'requestMethod'
                    },{
                        key: '请求URL',
                        html: '{{item.requestURL}}',
                        draggableCacheMark: 'url'
                    }, {
                        key: '负载后端',
                        html: '{{item.target}}',
                        draggableCacheMark: 'target'
                    }, {
                        key: '转发方式',
                        html: '{{item.isFollow?item.requestMethod:item.targetMethod}}',
                        draggableCacheMark: 'targetMethod'
                    }, {
                        key: '转发URL',
                        html: '{{item.targetURL}}',
                        draggableCacheMark: 'targetURL'
                    }, {
                        key: '分组',
                        html: '{{item.groupName}}',
                        draggableCacheMark: 'group'
                    }, {
                        key: '更新时间',
                        html: '{{item.updateTime}}',
                        draggableCacheMark: 'updateTime'
                    }],
                    operate: {
                        funArr: [{
                            key: '复制',
                            fun: privateFun.copy
                        }, {
                            key: '修改',
                            fun: privateFun.edit,
                            params: {
                                status: 'edit'
                            }
                        }, {
                            key: '删除',
                            fun: privateFun.delete,
                            params: '"single",arg'
                        }],
                        class: 'w_150'
                    }
                },
                setting: {
                    page: true,
                    scroll: true,
                    scrollRemainRatio: 7,
                    isFixedHeight: true,
                    draggable: true,
                    dragCacheVar: 'AMS_API_LIST_DRAG_VAR',
                    dragCacheObj: {
                        name: '250px',
                        url: '250px',
                        target: '150px',
                        targetURL: '150px',
                        group: '150px',
                        manager: '150px',
                        updater: '150px',
                        updateTime: '150px',
                        requestMethod: '150px',
                        targetMethod: '150px'
                    },
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    titleAuthority: 'showTitle',
                    unhover: true,
                    warning: '尚未新建任何接口'
                },
                baseFun:{
                    scrollLoading: privateFun.scrollLoading
                }
            }

            vm.component.menuObject = {
                list: [{
                    type: 'more-btn',
                    authority: 'edit',
                    icon: 'jiahao',
                    name: '新建',
                    class: 'pull-left',
                    btnClass: 'eo_theme_btn_success',
                    btnList: [{
                        name: '普通接口',
                        fun: {
                            default: privateFun.edit,
                            params: {
                                status: 'add-common'
                            }
                        }
                    }, {
                        name: '链式调用(服务编排)',
                        fun: {
                            default: privateFun.edit,
                            params: {
                                status: 'add-link'
                            }
                        }
                    }, {
                        name: '导入接口',
                        fun: {
                            default: privateFun.import
                        }
                    }]
                }, {
                    type: 'search',
                    class: 'pull-right',
                    keyword: vm.ajaxRequest.keyword,
                    fun: privateFun.search,
                    placeholder: "输入搜索内容"
                }],
                batchList: [{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '修改分组',
                        fun: {
                            default: privateFun.batchMoveGroup
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
                    title: "接口列表"
                },
                active: {
                    condition: 0
                },
                baseFun:{
                    batchDefault: privateFun.batchBtnClickFun
                }
            };
        }
    }
})();