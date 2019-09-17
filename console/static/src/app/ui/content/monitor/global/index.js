(function () {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .component('panel', {
            templateUrl: 'app/ui/content/monitor/global/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource','$state', 'CODE', '$rootScope', 'uibDateParser'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, uibDateParser) {

        var vm = this,
        privateFun = {}
        vm.data = {
            cluster:'',
            tabSummaryList: [{
                name: '今天',
                active: 0
            }, {
                name: '近3天',
                active: 1
            }, {
                name: '近7天',
                active: 2
            }, {
                active: 3,
                type: 'html'
            }],
            granularityList: [{
                name: '小时',
                active: 1
            }, {
                name: '天',
                active: 0
            }]
        }
        vm.fun = {};
        vm.ajaxRequest={
            projectHashKey: $state.params.projectHashKey,
            table: {
                beginTime: null,
                endTime: null,
                period: 0
            }
        }
        vm.ajaxResponse={
            monitorInfo: null,
            redisArr:[]
        };
        vm.directive = {
            tableTimeObject: {
                show: false,
                maxDate: new Date(),
                maxMode: 'month',
                request:{}
            }
        }
        vm.component = {
            overviewObject: {},
            listDefaultCommonObject: null
        }
        privateFun.filterTime = function (mark) {
            var template = {
                startTime: uibDateParser.filter(vm.directive[mark + 'TimeObject'].request.startTime, 'yyyy-M!-dd'),
                endTime: uibDateParser.filter(vm.directive[mark + 'TimeObject'].request.endTime, 'yyyy-M!-dd')
            }
            if (!template.startTime) {
                $rootScope.InfoModal('请选择开始日期', 'error');
            } else if (!template.endTime) {
                $rootScope.InfoModal('请选择结束日期', 'error');
            } else {

                vm.ajaxRequest[mark].period = 3;
                vm.directive[mark + 'TimeObject'].show = false;
                if (template.startTime > template.endTime) {
                    template.templateTime = template.startTime;
                    template.startTime = template.endTime;
                    template.endTime = template.templateTime;
                }
                vm.ajaxRequest[mark].beginTime = template.startTime;
                vm.ajaxRequest[mark].endTime = template.endTime;
                privateFun.initTable();
            }
        }
        vm.fun.tableFilterTime = function (arg) {
            if (arg) arg.$event.stopPropagation();
            privateFun.filterTime('table')
        }
        vm.fun.changeMenu = function (mark, arg) {
            if (arg.item.active == 3) {
                vm.directive.tableTimeObject.show = true;
                arg.$event.stopPropagation();
                return;
            } else {
                vm.directive.tableTimeObject.request = {};
                vm.ajaxRequest.table.beginTime = null;
            }
            vm.ajaxRequest.table.period = arg.item.active;
            privateFun.initTable();
        }
        privateFun.initTable=function(){
            var template = {
                promise: null,
                request: {
                    beginTime: vm.ajaxRequest.table.beginTime,
                    endTime: vm.ajaxRequest.table.endTime,
                    period: vm.ajaxRequest.table.period,
                    cluster:vm.data.cluster
                }
            }
            template.promise = GatewayResource.Monitor.Info(template.request).$promise;
            template.promise.then(function (response) {
                vm.ajaxResponse.monitorInfo = response || {};
            })
            return template.promise;
        }
        privateFun.refresh=function(){
            var tmpPromise=GatewayResource.Monitor.Refresh().$promise;
            tmpPromise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            $rootScope.InfoModal('立即刷新成功!', 'success', function () {
                                $scope.$emit('$TransferStation', {
                                    state: '$Init_LoadingCommonComponent'
                                });
                            });
                            break;
                        }
                }
            })
            return tmpPromise;
        }
        privateFun.initComponent = function () {
            vm.component.overviewObject = {
                setting: {
                    title: '基本信息',
                    showOperate:true
                }
            }
        }
        vm.fun.init = function (arg) {
            arg=arg||{
                type:'default'
            }
            switch(arg.type){
                case 'default':{
                    if(vm.data.cluster){
                        return privateFun.initTable();
                    }else{
                        let tmpPromise=privateFun.initCluster();
                        tmpPromise.finally(()=>{
                            privateFun.initTable();
                        })
                        return tmpPromise;
                    }
                }
                case 'refresh':{
                    return privateFun.refresh();
                }
                case 'cluster':{
                    return privateFun.initTable();
                }
            }
            
        }
        vm.fun.refresh = function () {
            $scope.$emit('$TransferStation', {
                state: '$Init_LoadingCommonComponent',
                data:{
                    type:'refresh',
                    tips:'刷新'
                }
            });
        }
        vm.fun.changeClutser=()=>{
            $scope.$emit('$TransferStation', {
                state: '$Init_LoadingCommonComponent',
                data:{
                    type:'cluster'
                }
            });
        }
        privateFun.initCluster=()=>{
            let tmpPromise=GatewayResource.Cluster.SimpleQuery().$promise;
            tmpPromise.then((response)=>{
                vm.ajaxResponse.clusterArr=[{
                    title:"所有集群",
                    name:null
                }].concat(response.clusters||[]);
                vm.data.cluster=vm.ajaxResponse.clusterArr[0].name;
            })
            return tmpPromise;
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['监控面板']
            });
            privateFun.initComponent();
        }
    }
})();