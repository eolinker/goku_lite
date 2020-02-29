(function () {
    'use strict';
    /**
     * 负载
     */
    angular.module('eolinker')
        .component('balanceOperate', {
            templateUrl: 'app/ui/content/balance/operate/index.html',
            controller: indexController,
            bindings:{
                groupArr:'<'
            }
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope) {
        var vm = this;
        vm.data = {
            status: $state.params.status,
            balanceNamePattern:'[\\w\\._\\/\\-\\:]+'
        }
        vm.fun = {};
        vm.ajaxResponse = {
            balanceInfo: {
                balanceName: $state.params.balanceName || '',
                serviceType:'static'
            },
            serviceQuery:[{
                registryName:'无',
                serviceDiscoveryID:0
            }]
        }
        vm.component = {
            menuObject: {
                list: []
            },
            listBlockObj:{}
        };
        vm.CONST = {
            SERVICE_TYPE_ARR: [{
                key: "静态服务",
                value: "static"
            }, {
                key: "服务发现",
                value: "discovery"
            }]
        }
        var privateFun = {},cache={
            staticServiceQuery:[],
            discoveryServiceQuery:[],
            serviceIDObj:{}
        };
        vm.fun.switchMenu=(inputMenu,inputUnNeedToSetServiceName)=>{
            vm.ajaxResponse.balanceInfo.serviceType=inputMenu;
            vm.ajaxResponse.serviceQuery=cache[inputMenu+'ServiceQuery'];
            if(!inputUnNeedToSetServiceName)vm.ajaxResponse.balanceInfo.serviceName=cache.serviceIDObj[inputMenu]||vm.ajaxResponse.serviceQuery[0].name;
        }
        vm.fun.init = function () {
            if(vm.data.status==='add')return;
            var tmpAjaxRequest={
                balanceName: vm.ajaxResponse.balanceInfo.balanceName
            }
            GatewayResource.Balance.Info(tmpAjaxRequest).$promise.then(function (response) {
                response.balanceInfo.staticCluster=response.balanceInfo.staticCluster||{};
                vm.ajaxResponse.balanceInfo = response.balanceInfo;
                $rootScope.global.ajax.Query_ServiceDiscovery.$promise.finally(()=>{
                    vm.fun.switchMenu(vm.ajaxResponse.balanceInfo.serviceType,true);
                })
                
                $rootScope.global.ajax.SimpleQuery_Cluster.$promise.finally(()=>{
                    vm.ajaxResponse.clusterList.map((val,key)=>{
                        if(vm.ajaxResponse.balanceInfo.staticCluster.hasOwnProperty(val.name)){
                            vm.ajaxResponse.clusterList[key].value=response.balanceInfo.staticCluster[val.name];
                        }
                    })
                })
                
            });


        };
        vm.fun.back = function () {
            $state.go('home.balance.list.default');
        }
        privateFun.initRegistry=()=>{
            $rootScope.global.ajax.Query_ServiceDiscovery = GatewayResource.ServiceDiscovery.SimpleQuery();
            $rootScope.global.ajax.Query_ServiceDiscovery.$promise.then(function (response) {
                response.data.list.map((val)=>{
                    switch(val.type){
                        case 'static':{
                            cache.staticServiceQuery.push(val);
                            break;
                        }
                        default:{
                            cache.discoveryServiceQuery.push(val);
                            break;
                        }
                    }
                })
                if(cache.staticServiceQuery.length===0){
                    cache.staticServiceQuery.push({
                        name:"无"
                    });
                }
                if(cache.discoveryServiceQuery.length===0){
                    cache.discoveryServiceQuery.push({
                        name:"无"
                    });
                }
                if(vm.data.status==='add'){
                    vm.fun.switchMenu(vm.ajaxResponse.balanceInfo.serviceType);
                }
            })
            return $rootScope.global.ajax.Query_ServiceDiscovery.$promise;
        }
        privateFun.confirm = function () {
            var tmpOutput={
                balanceName: vm.ajaxResponse.balanceInfo.balanceName,
                serviceName:vm.ajaxResponse.balanceInfo.serviceName
            }
            switch(vm.ajaxResponse.balanceInfo.serviceType){
                case 'static':{
                    tmpOutput.static=vm.ajaxResponse.balanceInfo.static;
                    tmpOutput.staticCluster={};
                    for(let val of vm.ajaxResponse.clusterList){
                        tmpOutput.staticCluster[val.name]=val.value;
                    }
                    tmpOutput.staticCluster=JSON.stringify(tmpOutput.staticCluster);
                    break;
                }
                default:{
                    tmpOutput.appName=vm.ajaxResponse.balanceInfo.appName;
                    break;
                }
            }
            
            return tmpOutput;
        }
        vm.fun.load = function (arg) {
            $scope.$emit('$TransferStation', {
                state: '$Init_LoadingCommonComponent',
                data: arg
            });
        }
        vm.fun.requestProcessing = function (arg) {
            vm.data.submitted=true;
            if(vm.ajaxResponse.balanceInfo.serviceName==="无"){
                $rootScope.InfoModal('暂未发现任何服务注册方式，请先新建服务注册方式', 'error');
                return;
            }
            var tmpAjaxRequest=privateFun.confirm(),tmpPromise=null;
            if ($scope.ConfirmForm.$valid) {
                tmpPromise = privateFun.edit({
                    request: tmpAjaxRequest
                });
            } else {
                $rootScope.InfoModal('编辑失败，请检查信息是否填写完整！', 'error');
            }
            return tmpPromise;
        }
        privateFun.edit = function (arg) {
            var tmpPromise=null;
            if (vm.data.status == 'edit') {
                tmpPromise = GatewayResource.Balance.Edit(arg.request).$promise;
                tmpPromise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS: {
                            vm.fun.back();
                            $rootScope.InfoModal('修改成功', 'success');
                            break;
                        }
                    }
                })
            } else {
                let tmpAjaxRequest=angular.copy(arg.request);
                    tmpPromise = GatewayResource.Balance.Add(tmpAjaxRequest).$promise;
                    tmpPromise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                vm.fun.back();
                                $rootScope.InfoModal('添加负载成功', 'success');
                                break;
                            }
                        }
                    })
            }
            return tmpPromise;
        }
        privateFun.initBlockTable=()=>{
            vm.component.listBlockObj={
                setting:{
                    munalAddRow:true
                },
                tdList:[{
                    type:'text',
                    thKey:'范围',
                    modelKey:'title',
                    class:'w_150'
                },{
                    type:'input',
                    thKey:'静态服务地址',
                    modelKey:'value'
                }]
            }
        }
        privateFun.initCluster=()=>{
            $rootScope.global.ajax.SimpleQuery_Cluster=GatewayResource.Cluster.SimpleQuery();
            $rootScope.global.ajax.SimpleQuery_Cluster.$promise.then((response)=>{
                vm.ajaxResponse.clusterList=response.clusters||[];
            });            
        }
        vm.$onInit = function () {
            privateFun.initCluster();
            privateFun.initRegistry();
            privateFun.initBlockTable();
            vm.component.menuObject={
                list:[{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '返回列表',
                        icon: 'chexiao',
                        fun: {
                            default: vm.fun.back
                        }
                    }]
                }, {
                    type: 'btn',
                    class: 'btn-group-li',
                    btnList: [{
                        name: '保存',
                        class: 'eo_theme_btn_success block-btn',
                        fun: {
                            disabled: 1,
                            default: vm.fun.requestProcessing,
                            params: {
                                status: 1
                            }
                        }
                    }]
                }],
                setting:{
                    class:'common-menu-fixed-seperate'
                }
            };
        }
    }
})();