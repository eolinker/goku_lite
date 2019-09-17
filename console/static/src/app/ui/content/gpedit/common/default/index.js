(function () {
    'use strict';
    angular.module('eolinker')
        .component('gpeditDefault', {
            templateUrl: 'app/ui/content/gpedit/common/default/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', '$rootScope', 'CODE', 'Cache_CommonService', 'Authority_CommonService'];

    function indexController($scope, GatewayResource, $state, $rootScope, CODE, Cache_CommonService, Authority_CommonService) {
        var vm = this;
        vm.data = {
            groupType: $state.params.groupType || null,
            batch: {},
            pagination: {
                maxSize: 10,
                pageSize: 50,
                page: 1,
                msgCount: 0
            }
        }
        vm.ajaxRequest={
            groupID: $state.params.groupID||-1,
            strategyID: [],
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        };
        vm.ajaxResponse={};
        vm.fun = {};
        vm.service = {
            authority: Authority_CommonService
        }
        vm.component = {
            listDefaultCommonObject: {},
            menuObject: {}
        }
        var service = {
            cache: Cache_CommonService
        },privateFun = {},CONST={
            GROUP_ARR:[{
                groupID: 0,
                groupName: "未分组",
                icon: 'caidan'
            }]
        },data={};
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
        vm.fun.init = function () {
            return privateFun.getQuery('reset');
        }
        privateFun.batchBtnClickFun = () => {
            if (vm.data.pagination.page === Math.ceil(vm.data.pagination.msgCount / vm.data.pagination.pageSize)) return;
            var tmpAjaxRequest={
                groupID: vm.ajaxRequest.groupID,
                condition:vm.component.menuObject.active.condition
            },tmpPromise=null;
            if (vm.ajaxRequest.keyword) {
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            } 
            tmpPromise = GatewayResource.Strategy.IDQuery(tmpAjaxRequest).$promise;
            tmpPromise.then(function (response) {
                vm.data.allQueryID = response.strategyIDList||[];
            })
        }
        privateFun.getQuery = function (inputOpr,inputPage,inputPageSize) {
            var tmpAjaxRequest={
                groupID: vm.ajaxRequest.groupID,
                condition:vm.component.menuObject.active.condition,
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
            $rootScope.global.ajax.Query_Strategy = GatewayResource.Strategy.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_Strategy.$promise.then(function (response) {
                if(inputOpr==='reset'){
                    document.getElementsByClassName('tbody_container_ldcc')[0].scrollTop = 0;
                }
                data.isQuerying=false;
                vm.ajaxResponse.query = tmpResponseArr.concat(response.strategyList || []);
                vm.data.pagination.msgCount = response.page.totalNum||0;
                if(!inputPageSize)vm.data.pagination.page=tmpAjaxRequest.page;
            })
            return $rootScope.global.ajax.Query_Strategy.$promise;
        }
        privateFun.changeGroup = function () {
            var template = {
                modal: {
                    title: '批量修改策略分组',
                    query: CONST.GROUP_ARR.concat(service.cache.get('gpeditGroup')),
                    position: {
                        key: 'groupName'
                    }
                },
                request: {
                    strategyIDList: vm.data.batch.query.join(','),
                    groupID: ''
                },
                loop: {
                    num: 0
                }
            }
            $rootScope.SingleSelectModal(template.modal, function (callback) {
                if (callback) {
                    template.request.groupID = template.modal.query[callback.$index].groupID;
                    GatewayResource.Strategy.ChangeGroup(template.request).$promise
                        .then(function (response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS:
                                    {
                                        privateFun.resetBatchInfo();
                                        $rootScope.InfoModal('策略批量修改分组成功', 'success');
                                        privateFun.getQuery("reset");
                                        break;
                                    }
                            }
                        })
                }
            });
        }
        privateFun.operate = function (arg) {
            var template = {
                promise: null,
                title: vm.data.batch.isOperating ? '批量' : '',
                loop: {
                    num: 0
                },
                request:{
                    strategyIDList:vm.data.batch.isOperating ?vm.data.batch.query.join(','):arg.item.strategyID
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
            GatewayResource.Strategy[arg.status](template.request).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            privateFun.resetBatchInfo();
                            $rootScope.InfoModal(template.title + '成功', 'success');
                            privateFun.getQuery("reset");
                            break;
                        }
                }
            });
        }
        privateFun.edit = function (arg) {
            arg.item=arg.item||{};
            let tmpTargetObj={
                "edit":{
                    title:"修改访问策略",
                    opr:'Edit'
                },
                "add":{
                    title:"新增访问策略",
                    opr:"Add"
                },
                "copy":{
                    title:"复制访问策略",
                    opr:'Copy'
                }
            }
            let tmpModal={
                title: tmpTargetObj[arg.status].title,
                group: CONST.GROUP_ARR.concat(service.cache.get('gpeditGroup')),
                item: arg.item
            },tmpAjaxRequest={};
            if(arg.status==="copy"){
                tmpModal.item.strategyName="副本-"+tmpModal.item.strategyName;
            }
            tmpModal.item.groupID = parseInt(arg.item.groupID||vm.ajaxRequest.groupID);
            if(tmpModal.item.groupID===-1){
                tmpModal.item.groupID=0;
            }
            $rootScope.GatewayGpeditDefaultModal(tmpModal, function (callback) {
                if (callback) {
                    tmpAjaxRequest = {
                        strategyName: callback.strategyName,
                        groupID: callback.groupID
                    }
                    switch (arg.status) {
                        case 'copy':
                        case 'edit':
                            {
                                tmpAjaxRequest.strategyID = arg.item.strategyID;
                            }
                        case 'add':
                            {
                                GatewayResource.Strategy[tmpTargetObj[arg.status].opr](tmpAjaxRequest).$promise.then(function (response) {
                                    switch (response.statusCode) {
                                        case CODE.COMMON.SUCCESS:
                                            {
                                                $rootScope.InfoModal(tmpTargetObj[arg.status].title+'成功!', 'success');
                                                privateFun.getQuery("reset");
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
                    strategyIDList: status == 'batch' ? vm.data.batch.query.join(',') : arg.item.strategyID,
                    groupID: vm.ajaxRequest.groupID
                },
                loop: {
                    num: 0
                }
            }
            $rootScope.EnsureModal('删除策略', null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.Strategy.Delete(template.request).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    switch (status) {
                                        case 'batch':
                                            {
                                                privateFun.resetBatchInfo();
                                                $rootScope.InfoModal('策略删除成功', 'success');
                                                privateFun.getQuery("reset");
                                                break;
                                            }
                                        case 'single':
                                            {
                                                vm.ajaxResponse.query.splice(arg.$index, 1);
                                                if(vm.data.pagination.msgCount>vm.data.pagination.page*vm.data.pagination.pageSize)privateFun.getQuery('preload',vm.data.pagination.page*vm.data.pagination.pageSize,1)
                                                vm.data.pagination.msgCount--;
                                                $rootScope.InfoModal('策略删除成功', 'success');
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
        vm.$onInit=()=>{
            $scope.$emit('$WindowTitleSet', {
                list: ['普通策略列表']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'strategyID',
                    default: [{
                        key: '名称',
                        html: '{{item.strategyName}}'
                    }, {
                        key: '策略ID',
                        html: '<span>{{item.strategyID}}</span><span class="ml10 copy_btn_gd eo_theme_btn_default" copy-common-directive copy-model="item.strategyID">复制</span>'
                    }, {
                        key: '状态',
                        html: '<span class="eo-status-warning" ng-if="item.enableStatus==\'0\'">停用</span><span class="eo-status-success" ng-if="item.enableStatus==\'1\'">启用</span>'
                    }, {
                        key: '分组',
                        html: '{{item.groupName}}'
                    }, {
                        key: '更新时间',
                        html: '{{item.updateTime}}',
                        class: "w_180"
                    }],
                    operate: {
                        funArr: [{
                            key: '复制',
                            fun: privateFun.edit,
                            params: {
                                status: 'copy'
                            }
                        },{
                            key: '修改',
                            fun: privateFun.edit,
                            params: {
                                status: 'edit'
                            }
                        }, {
                            key: '删除',
                            fun: privateFun.delete,
                            params: '"single",arg'
                        }, {
                            key: '启用',
                            itemExpression: 'ng-if="item.enableStatus===0"',
                            fun: privateFun.operate,
                            params: {
                                status: 'Start'
                            }
                        }, {
                            key: '停用',
                            itemExpression: 'ng-if="item.enableStatus===1"',
                            fun: privateFun.operate,
                            params: {
                                status: 'Stop'
                            }
                        }],
                        class: 'w_220'
                    }
                },
                baseFun: {
                    edit: privateFun.edit,
                    click: (tmpInputArg)=>{
                        $state.go('home.gpedit.inside.api.default', {
                            groupType:"common",
                            strategyID: tmpInputArg.item.strategyID,
                            strategyName: tmpInputArg.item.strategyName,
                            groupID: vm.ajaxRequest.groupID
                        });
                    },
                    scrollLoading: privateFun.scrollLoading
                },
                setting: {
                    page: true,
                    scroll: true,
                    scrollRemainRatio: 7,
                    isFixedHeight: true,
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    titleAuthority: 'showTitle',
                    warning: '尚未新建任何策略'
                }
            }
            vm.component.menuObject = {
                list: [{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    authority: 'edit',
                    btnList: [{
                        name: '新建策略',
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
                    placeholder:"输入策略名称或策略ID"
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
                    }, {
                        name: '批量开启',
                        fun: {
                            default: privateFun.operate,
                            params: {
                                status: 'Start'
                            }
                        }
                    }, {
                        name: '批量关闭',
                        fun: {
                            default: privateFun.operate,
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
                    title: "普通策略列表"
                },
                baseFun:{
                    batchDefault: privateFun.batchBtnClickFun
                },
                active:{
                    condition:0
                }
            };
        }
    }
})();