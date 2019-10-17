(function () {
    'use strict';
    angular.module('eolinker')
        .component('gpeditInsideApiDefault', {
            templateUrl: 'app/ui/content/gpedit/inside/content/api/_default/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope','CommonResource', 'GatewayResource', '$state', '$rootScope', 'CODE', 'Authority_CommonService'];

    function indexController($scope,CommonResource, GatewayResource, $state, $rootScope, CODE, Authority_CommonService) {
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
            strategyID: $state.params.strategyID,
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        };
        vm.ajaxResponse = {
            query: null,
            balanceList:null
        }
        vm.component = {
            listDefaultCommonObject: null,
            menuObject: {}
        }
        vm.service = {
            authority: Authority_CommonService
        }
        var privateFun = {},data={};
        privateFun.batchBtnClickFun = () => {
            if (vm.data.pagination.page === Math.ceil(vm.data.pagination.msgCount / vm.data.pagination.pageSize)) return;
            var tmpAjaxRequest={
                strategyID: vm.ajaxRequest.strategyID,
                condition:vm.component.menuObject.active.condition
            },tmpPromise=null;
            if (vm.ajaxRequest.keyword) {
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            } 
            tmpPromise = GatewayResource.StrategyApi.IDQuery(tmpAjaxRequest).$promise;
            tmpPromise.then(function (response) {
                vm.data.allQueryID = response.apiIDList||[];
            })
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
        vm.fun.init = function () {
            return privateFun.getQuery('reset');
        }
        privateFun.getQuery = function (inputOpr,inputPage,inputPageSize) {
            var tmpAjaxRequest={
                strategyID: vm.ajaxRequest.strategyID,
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
            if(vm.ajaxRequest.ids){
                tmpAjaxRequest.ids=JSON.stringify(vm.ajaxRequest.ids);
            }
            if(vm.ajaxRequest.keyword){
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_StrategyApi = GatewayResource.StrategyApi.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_StrategyApi.$promise.then(function (response) {
                if(inputOpr==='reset'){
                    document.getElementsByClassName('tbody_container_ldcc')[0].scrollTop = 0;
                }
                data.isQuerying=false;
                vm.ajaxResponse.query = tmpResponseArr.concat(response.apiList || []);
                vm.data.pagination.msgCount = response.page.totalNum||0;
                if(!inputPageSize)vm.data.pagination.page=tmpAjaxRequest.page;
            })
            return $rootScope.global.ajax.Query_StrategyApi.$promise;
        }
        privateFun.edit = function (arg) {
            var tmpUri = {
                status: arg.status
            }
            switch (arg.status) {
                case 'edit': {
                    tmpUri.apiID = arg.item.apiID;
                    break;
                }
            }
            $state.go('home.gpedit.inside.api.operate', tmpUri);
        }
        privateFun.delete = function (status, arg) {
            var tmpAjaxRequest = {
                    strategyID: vm.ajaxRequest.strategyID,
                    apiIDList: status == 'batch' ? vm.data.batch.query.join(',') : arg.item.apiID
                },
                tmpModal = {
                    title: status == 'batch' ? '批量删除绑定API' : ('删除绑定API-' + arg.item.apiName)
                }
            $rootScope.EnsureModal(tmpModal.title, null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.StrategyApi.Delete(tmpAjaxRequest).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                switch (status) {
                                    case 'batch': {
                                        privateFun.resetBatchInfo();
                                        $rootScope.InfoModal(tmpModal.title + '成功', 'success');
                                        privateFun.getQuery("reset");
                                        break;
                                    }
                                    case 'single': {
                                        vm.ajaxResponse.query.splice(arg.$index, 1);
                                        if(vm.data.pagination.msgCount>vm.data.pagination.page*vm.data.pagination.pageSize)privateFun.getQuery('preload',vm.data.pagination.page*vm.data.pagination.pageSize,1)
                                        vm.data.pagination.msgCount--;
                                        $rootScope.InfoModal(tmpModal.title + '成功', 'success');
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
        privateFun.getDefaultBalance=()=>{
            let tmpPromise=GatewayResource.Balance.SimpleQuery().$promise;
            tmpPromise.then((response)=>{
                vm.ajaxResponse.balanceList=response.balanceNames;
            })
            return tmpPromise;
        }
        privateFun.setGpeditBalance = function (inputArg) {
            let tmpFunModal=()=>{
                let tmpSpecialObj={
                    itemArr:[]
                };
                switch(inputArg.status){
                    case "batch":{
                        tmpSpecialObj.apiIDs=vm.data.batch.query;
                        break;
                    }
                    case "single":{
                        tmpSpecialObj.itemArr=[{
                            type: 'input',
                            title: '[原]API负载后端',
                            value: inputArg.item['target'],
                            disabled:true
                        }];
                        tmpSpecialObj.apiIDs=[inputArg.item.apiID];
                        break;
                    }
                }
                let tmpModal = {
                    title: '设置策略负载',
                    resource: GatewayResource.StrategyApi.ChangeTarget,
                    textArray: tmpSpecialObj.itemArr.concat([{
                        type: 'html',
                        title: '[新]策略负载后端',
                        key:'target',
                        value: inputArg.item['rewriteTarget'],
                        query:vm.ajaxResponse.balanceList,
                        autoCompleteObj:{pattern:'[\\w\\._\\/\\-\\:]+'},
                        html:`<auto-complete-component ng-class="{'eo-had-input-error':(ConfirmForm.value.$invalid)&&data.submitted}" model="item" key-name='value' array="item.query" placeholder="支持字母、数字、_、.、：、/" main-object="item.autoCompleteObj"> </auto-complete-component>`
                    }]),
                    request:{
                        strategyID: vm.ajaxRequest.strategyID,
                        apiIDs:JSON.stringify(tmpSpecialObj.apiIDs)
                    }
                }
                $rootScope.MixInputModal(tmpModal,(callback)=>{
                    if(callback){
                        if(inputArg.status==="batch"){
                            privateFun.resetBatchInfo();
                            privateFun.getQuery("reset");
                        }else{
                            vm.ajaxResponse.query[inputArg.$index].rewriteTarget=callback.target;
                        }
                        $rootScope.InfoModal('设置策略负载成功','success');
                    }
                })
            }
            if(vm.ajaxResponse.balanceList){
                tmpFunModal();
            }else{
                let tmpPromise=privateFun.getDefaultBalance();
                tmpPromise.finally(()=>{
                    tmpFunModal();
                })
            }
        }
        privateFun.resetBatchInfo = function () {
            vm.data.batch.isOperating = false;
            vm.data.batch.selectAll = false;
            vm.data.batch.query = [];
            vm.data.batch.indexAddress = {};
        };
        privateFun.setGpeditPlugin=(inputArg)=>{
            $state.go('home.gpedit.inside.api.plugin',{
                apiID:inputArg.item.apiID
            })
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['API接口', '策略']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey:'apiID',
                    default: [{
                            key: 'APIs',
                            html: `<span class="mr5 plr3 ptb2 fs12" ng-class="{'eo-label-warning':item.apiType,'eo-label-success':!item.apiType}">{{item.apiType?"编排":"普通"}}</span><span>{{item.apiName}}</span>`,
                            draggableCacheMark: 'name'
                        },{
                            key: '请求方式',
                            html: '{{item.requestMethod}}',
                            draggableCacheMark: 'requestMethod'
                        },
                        {
                            key: '请求URL',
                            html: '{{item.requestURL}}',
                            draggableCacheMark: 'url'
                        }, {
                            key: '转发方式',
                            html: '{{item.isFollow?item.requestMethod:item.targetMethod}}',
                            draggableCacheMark: 'targetMethod'
                        }, {
                            key: '转发URL',
                            html: '{{item.targetURL}}',
                            draggableCacheMark: 'targetURL',
                            draggableCacheMark: 'targetURL'
                        },
                        {
                            key: '<span>API负载后端</span><span class="iconfont icon-guanyu api_balance_tip fwn"></span>',
                            html: '{{item.target}}',
                            draggableCacheMark: 'target'
                        },
                        {
                            key: '<span>策略负载后端</span><span class="iconfont icon-guanyu gpedit_balance_tip fwn"></span>',
                            html: '{{item.rewriteTarget}}',
                            draggableCacheMark: 'rewriteTarget'
                        }, {
                            key: '更新时间',
                            html: '{{item.updateTime}}',
                            draggableCacheMark: 'updateTime'
                        }
                    ],
                    operate: {
                        funArr: [{
                            key: '设置插件',
                            show: false,
                            fun: privateFun.setGpeditPlugin
                        },{
                                key: '设置策略负载',
                                show: false,
                                fun: privateFun.setGpeditBalance,
                                params: {
                                    status:"single"
                                }
                            },
                            {
                                key: '删除',
                                show: false,
                                fun: privateFun.delete,
                                params: '"single",arg'
                            }
                        ],
                        class:'w_200'
                    }
                },
                setting: {
                    draggable: true,
                    dragCacheVar: 'GEPDIT_API_LIST_DRAG_VAR',
                    dragCacheObj: {
                        name: '250px',
                        url: '250px',
                        target: '150px',
                        targetURL: '150px',
                        updateTime: '150px',
                        requestMethod: '150px',
                        targetMethod: '150px',
                        rewriteTarget:'150px'
                    },
                    page: true,
                    scroll: true,
                    scrollRemainRatio: 7,
                    isFixedHeight: true,
                    unhover: true,
                    batch:true
                },
                baseFun:{
                    scrollLoading: privateFun.scrollLoading
                }
            }
            vm.component.menuObject = {
                list: [{
                        type: 'btn',
                        authority: 'edit',
                        class: 'pull-left',
                        btnList: [{
                            name: '绑定接口',
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
                        placeholder:"输入搜索内容"
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
                    },{
                        name: '设置策略负载',
                        fun: {
                            default: privateFun.setGpeditBalance,
                            params: {
                                status:"batch"
                            }
                        }
                    }]
                }],
                setting: {
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    titleAuthority: 'showTitle',
                    title: 'API列表'
                },
                active:{
                    condition:0
                },
                baseFun:{
                    batchDefault: privateFun.batchBtnClickFun
                }
            }
        }
    }
})();