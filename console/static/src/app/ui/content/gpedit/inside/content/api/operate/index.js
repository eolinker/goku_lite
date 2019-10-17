(function () {
    'use strict';
    angular.module('eolinker')
        .component('gpeditInsideApiOperate', {
            templateUrl: 'app/ui/content/gpedit/inside/content/api/operate/index.html',
            controller: indexController
        })

    indexController.$inject = ['CODE', '$state', '$rootScope', 'GatewayResource','Group_MultistageService'];

    function indexController(CODE, $state, $rootScope, GatewayResource,Group_MultistageService) {
        var vm = this;
        vm.data = {
            batch:{
                isOperating:true,
                query:[],
                indexAddress:{}
            },
            pagination: {
                maxSize: 10,
                pageSize: 50,
                page: 0,
                msgCount: 0
            }
        }
        vm.fun={};
        vm.ajaxResponse={}
        vm.component = {
            groupCommonObject: {
                firstGroup: {},
                secondGroup: {}
            }
        };
        var ajaxRequest={
            strategyID: $state.params.strategyID,
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        },data = {
            projectObject: {},
            positionObject: {
                strategyID: $state.params.strategyID,
                projectID: null,
                groupID: -1
            }
        },privateFun={},service={
            groupCommon: Group_MultistageService
        },
        groupInfo = {
            project: null,
            apiGroup: null,
        };
        privateFun.scrollLoading = function () {
            let tmpFlag = {
                hasItem: vm.ajaxResponse.query && vm.ajaxResponse.query.length !== 0,
                hasNextPage: vm.data.pagination.page < (vm.data.pagination.msgCount / vm.data.pagination.pageSize),
                isQuerying: data.isQuerying
            }
            if (tmpFlag.hasItem && tmpFlag.hasNextPage && !tmpFlag.isQuerying) {
                data.isQuerying=true;
                privateFun.getApiList('preload',vm.data.pagination.page+1);
            }
        }
        vm.fun.init = function () {
            var tmpAjaxRequest={
                strategyID: ajaxRequest.strategyID
            }
            $rootScope.global.ajax.ProjectQuery_Common = GatewayResource.Project.QueryAndGroup(tmpAjaxRequest);
            $rootScope.global.ajax.ProjectQuery_Common.$promise.then(function (response) {
                vm.ajaxResponse.firstGroupList = response.projectList || [];
                    if (vm.ajaxResponse.firstGroupList.length > 0) {
                        var groupObj = service.groupCommon.sort.init(vm.ajaxResponse.firstGroupList[0])
                        vm.ajaxResponse.seondGroupList = groupObj.groupList;
                        groupInfo.secondGroup = groupObj.groupInfo;
                        vm.component.groupCommonObject.firstGroup.mainObject.baseInfo.resetFlag = !vm.component.groupCommonObject.firstGroup.mainObject.baseInfo.resetFlag;
                        vm.component.groupCommonObject.secondGroup.mainObject.baseInfo.resetFlag = !vm.component.groupCommonObject.secondGroup.mainObject.baseInfo.resetFlag;
                        data.positionObject.projectID = vm.ajaxResponse.firstGroupList[0].projectID;
                        privateFun.getApiList("reset");
                    } else {
                        $rootScope.InfoModal("目前暂无可以被引用的API项目", 'error');
                        vm.fun.back();
                    }

            })
            return $rootScope.global.ajax.ProjectQuery_Common.$promise;
        }
        privateFun.search = function (arg) {
            ajaxRequest.keyword=arg.item.keyword;
            window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
            privateFun.getApiList("reset");
        }
        privateFun.selectAll = function (type) {
            console.log(type,vm.data.allQueryID)
            switch (type) {
                case 'selectAll': {
                    vm.data.allQueryID.map(function (val, key) {
                        console.log(val)
                        if (!vm.data.batch.indexAddress[val]) {
                            vm.data.batch.query.push(val);
                            vm.data.batch.indexAddress[val] = key + 1;
                        }
                    })
                    vm.data.batch.selectAll = true;
                    break;
                }
                case 'selectView': {
                    vm.ajaxResponse.query.map(function (val, key) {
                        if (!vm.data.batch.indexAddress[val.apiID]) {
                            vm.data.batch.query.push(val.apiID);
                            vm.data.batch.indexAddress[val.apiID] = key + 1;
                        }
                    })
                    let benginIndex = vm.ajaxResponse.query.length;
                    for (var i = benginIndex; i < vm.data.allQueryID.length; i++) {
                        if (vm.data.batch.indexAddress[vm.data.allQueryID[i]]) {
                            vm.data.batch.query.splice(vm.data.batch.query.indexOf(vm.data.allQueryID[i]), 1);
                            delete vm.data.batch.indexAddress[vm.data.allQueryID[i]];
                        }
                    }
                    vm.data.batch.selectAll = true;
                    break;
                }
                case 'cancelAll': {
                    vm.data.allQueryID.map(function (val, key) {
                        vm.data.batch.query.splice(vm.data.batch.query.indexOf(val), 1);
                        delete vm.data.batch.indexAddress[val];
                    })
                    vm.data.batch.selectAll = false;
                    break;
                }
            }
        }
        vm.fun.confirm = function () {
            if(vm.data.batch.query.length==0){
                $rootScope.InfoModal('尚未勾选任何API','error');
                return;
            }
            GatewayResource.StrategyApi.Add({
                strategyID: ajaxRequest.strategyID,
                apiID: vm.data.batch.query.join(',')
            }, function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            $rootScope.InfoModal('批量新增成功', 'success');
                            vm.fun.back();
                            break;
                        }
                }
            })
        }
        vm.fun.back = function () {
            $state.go('home.gpedit.inside.api.default');
        }
        privateFun.click = function (arg) {
            data.positionObject.groupID = arg.item.groupID;
            service.groupCommon.generalFun.initGroupStatus({
                currentGroupID: arg.item.groupID,
                groupInfo: groupInfo.secondGroup,
                list: vm.ajaxResponse.seondGroupList,
            });
            privateFun.getApiList("reset");
        }
        privateFun.batchBtnClickFun = () => {
            var tmpAjaxRequest={
                strategyID: ajaxRequest.strategyID,
                projectID: data.positionObject.projectID,
                groupID:data.positionObject.groupID
            },tmpPromise=null;
            if(ajaxRequest.keyword){
                tmpAjaxRequest.keyword = ajaxRequest.keyword;
            }
            tmpPromise = GatewayResource.StrategyApi.UnassignIDQuery(tmpAjaxRequest).$promise;
            tmpPromise.then(function (response) {
                vm.data.allQueryID = response.apiIDList||[];
            })
        }
        privateFun.checkIsSelectAll=()=>{
            vm.data.batch.selectAll = false;
                let returnFlag = false;
                for (var i = 0; i < vm.ajaxResponse.query.length; i++) {
                    if (vm.data.batch.query.indexOf(vm.ajaxResponse.query[i].apiID) === -1) {
                        returnFlag = true;
                        break;
                    }
                }
                if (vm.ajaxResponse.query.length && !returnFlag) {
                    vm.data.batch.selectAll = true;
                }
        }
        privateFun.getApiList = function (inputOpr,inputPage,inputPageSize) {
            var tmpPromise,tmpAjaxRequest={
                strategyID: ajaxRequest.strategyID,
                projectID: data.positionObject.projectID,
                pageSize: inputPageSize||vm.data.pagination.pageSize,
                groupID:data.positionObject.groupID
            },tmpResponseArr=[];
            switch(inputOpr){
                case 'reset':{
                    vm.ajaxResponse.query=[];
                    tmpAjaxRequest.page=1;
                    privateFun.batchBtnClickFun();
                    break;
                }
                default:{
                    tmpResponseArr=vm.ajaxResponse.query;
                    tmpAjaxRequest.page=inputPage;
                    break;
                }
            }
            if(ajaxRequest.keyword){
                tmpAjaxRequest.keyword = ajaxRequest.keyword;
            }
            tmpPromise = GatewayResource.StrategyApi.All(tmpAjaxRequest);
            tmpPromise.$promise.then(function (response) {
                data.isQuerying=false;
                data.projectObject[data.positionObject.projectID] = vm.ajaxResponse.query = tmpResponseArr.concat(response.apiList || []);
                if(inputOpr==='reset'){
                    document.getElementsByClassName('tbody_container_ldcc')[0].scrollTop = 0;
                    privateFun.checkIsSelectAll();
                }
                vm.data.pagination.msgCount = response.page.totalNum||0;
                if(!inputPageSize)vm.data.pagination.page=tmpAjaxRequest.page;
            })
            return tmpPromise.$promise;
        }
        vm.$onInit = function () {
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'apiID',
                    default: [{
                        key: 'APIs',
                        html: `<span class="mr5 plr3 ptb2 fs12" ng-class="{'eo-label-warning':item.apiType,'eo-label-success':!item.apiType}">{{item.apiType?"编排":"普通"}}</span><span>{{item.apiName}}</span>`,
                    }, {
                        key: '请求方式',
                        html: '{{item.requestMethod}}'
                    }, {
                        key: '请求URL',
                        html: '{{item.requestURL}}'
                    },
                    {
                        key: '负载后端',
                        html: '{{item.target}}'
                    }, {
                        key: '转发方式',
                        html: '{{item.isFollow?item.requestMethod:item.targetMethod}}',
                        draggableCacheMark: 'targetMethod'
                    }, {
                        key: '转发URL',
                        html: '{{item.targetURL}}'
                    }]
                },
                baseFun:{
                    selectAll:privateFun.selectAll,
                    scrollLoading: privateFun.scrollLoading
                },
                setting: {
                    batchInitStatus:"open",
                    page: true,
                    scroll: true,
                    scrollRemainRatio: 7,
                    isFixedHeight: true,
                    isWantToWatchSelectAll:true,
                    isAcrossToSelect:true,
                    batch: true,
                    batchInitFun: privateFun.resetBatchInfo,
                    titleAuthority: 'showTitle',
                    warning: '尚未存在任何接口'
                }
            }
            vm.component.menuDefaultObject = {
                list: [{
                    type: 'btn',
                    class: 'pull-left menu-ml15',
                    btnList: [{
                        name: '返回列表',
                        icon: 'chexiao',
                        fun: {
                            default: vm.fun.back
                        }
                    }, {
                        name: '保存',
                        class: 'eo_theme_btn_success block-btn',
                        fun: {
                            disabled: 1,
                            default: vm.fun.confirm
                        }
                    }]
                }],
                setting: {
                    class: "common-menu-fixed-seperate"
                }
            }
            vm.component.menuObject = {
                list: [{
                    type: 'search',
                    class: 'pull-right',
                    keyword: ajaxRequest.keyword,
                    fun: privateFun.search,
                    placeholder:"输入搜索内容"
                }],
                setting: {}
            }
            vm.component.groupCommonObject.firstGroup = {
                funObject: {
                    baseFun: {
                        click: function (arg) {
                            data.positionObject.projectID = arg.item.projectID;
                            data.positionObject.groupID = -1;
                            var groupObj = service.groupCommon.sort.init(arg.item);
                            vm.ajaxResponse.seondGroupList = groupObj.groupList;
                            vm.component.groupCommonObject.secondGroup.mainObject.baseInfo.resetFlag = !vm.component.groupCommonObject.secondGroup.mainObject.baseInfo.resetFlag;
                            groupInfo.secondGroup = groupObj.groupInfo;
                            privateFun.getApiList("reset");
                        }
                    }
                },
                mainObject: {
                    baseInfo: {
                        status: 'cancelRequest',
                        name: 'projectName',
                        id: 'projectID',
                        current: data.positionObject,
                        title: "项目",
                        disabledShrink: true
                    }
                }
            }
            vm.component.groupCommonObject.secondGroup = {
                funObject: {
                    baseFun: {
                        click: privateFun.click
                    }
                },
                mainObject: {
                    baseInfo: {
                        status: 'cancelRequest',
                        name: 'groupName',
                        id: 'groupID',
                        current: data.positionObject,
                        disabledShrink: true
                    },
                    staticQuery: [{
                        groupID: -1,
                        groupName: "所有接口",
                        icon: 'caidan'
                    },{
                        groupID: 0,
                        groupName: "未分组",
                        icon: 'caidan'
                    }]
                }
            }
        }
    }
})();