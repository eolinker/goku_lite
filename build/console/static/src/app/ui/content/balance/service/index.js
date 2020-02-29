(function () {
    'use strict';
    angular.module('eolinker')
        .component('balanceService', {
            templateUrl: 'app/ui/content/balance/service/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', '$rootScope', 'CODE', 'Authority_CommonService'];

    function indexController($scope, GatewayResource, $state, $rootScope, CODE, Authority_CommonService) {
        var vm = this;
        vm.data = {
            batch: {}
        }
        vm.ajaxRequest={
            name: [],
            keyword: window.sessionStorage.getItem('COMMON_SEARCH_TIP')
        }
        vm.ajaxResponse={
            query: []
        }
        vm.fun={};
        vm.service = {
            authority: Authority_CommonService
        }
        vm.component = {
            menuObject: {
                show: {
                    batch: {
                        disable: false
                    }
                }
            },
            listRequireObject: null
        }
        var privateFun = {},ajaxResponse={
            clusterQuery:null
        },data={
            clusterNameObj:{}
        };

        vm.fun.init = function () {
            let tmpAjaxRequest={};
            if(vm.ajaxRequest.keyword){
                tmpAjaxRequest.keyword = vm.ajaxRequest.keyword;
            }
            $rootScope.global.ajax.Query_ServiceDiscovery = GatewayResource.ServiceDiscovery.Query(tmpAjaxRequest);
            $rootScope.global.ajax.Query_ServiceDiscovery.$promise.then(function (response) {
                vm.ajaxResponse.query = response.data || [];
            })
            return $rootScope.global.ajax.Query_ServiceDiscovery.$promise;
        }
        privateFun.search = function (arg) {
            window.sessionStorage.setItem('COMMON_SEARCH_TIP', arg.item.keyword);
            $state.reload($state.current.name);
        }
        privateFun.initCluster=()=>{
            let tmpPromise=GatewayResource.Cluster.SimpleQuery().$promise;
            return tmpPromise;
            
        }
        privateFun.getServiceData=(inputServiceName)=>{
            let tmpPromise=GatewayResource.ServiceDiscovery.Info({
                name:inputServiceName
            }).$promise;
            return tmpPromise;
        }
        privateFun.edit = function (inputOp,inputArg) {
            inputArg=inputArg||{
                item:{}
            }
            let tmpConf={
                edit:{
                    title:'编辑服务注册方式',
                    resource:'Edit'
                },
                add:{
                    title:'添加服务注册方式',
                    resource:'Add'
                }
            },tmpFunParseData=(inputClusterArr)=>{
                let tmpFunShowModal=(tmpServiceInfo)=>{
                    tmpServiceInfo=Object.assign({},{
                        type:'static',
                        driver:'eureka'
                    },tmpServiceInfo);
                    let tmpModal = {
                        title: tmpConf[inputOp].title,
                        opr:tmpConf[inputOp].resource,
                        ajaxResponse:{
                            serviceData:tmpServiceInfo,
                            clusterQuery:inputClusterArr
                        }
                    }
                    $rootScope.Gateway_ServiceModal(tmpModal,(callback)=>{
                        if(callback){
                            vm.fun.init();
                        }
                    })
                }
                if(inputOp==='edit'){
                    let tmpPromise=privateFun.getServiceData(inputArg.item.name);
                    tmpPromise.then((response)=>{
                        if(inputArg.item.type==='discovery'){
                            let tmpAccessAddrs=response.data.clusterConfig||{};//最后需修改,目前先对接
                            for(let key in tmpAccessAddrs){
                                if(data.clusterNameObj.hasOwnProperty(key)){
                                    inputClusterArr[data.clusterNameObj[key]].value=tmpAccessAddrs[key];
                                }
                            }
                        }
                        tmpFunShowModal(response.data)
                    })
                }else{
                    tmpFunShowModal();
                }
                
            }
            if(ajaxResponse.clusterQuery){
                tmpFunParseData(angular.copy(ajaxResponse.clusterQuery));
            }else{
                let tmpPromise=privateFun.initCluster();
                tmpPromise.then((response)=>{
                    ajaxResponse.clusterQuery=response.clusters||[];
                    ajaxResponse.clusterQuery.map((val,key)=>{
                        data.clusterNameObj[val.name]=key;
                    })
                    tmpFunParseData(ajaxResponse.clusterQuery);
                })
            }
        }
        privateFun.delete = function (inputArg) {
            let tmpAjaxRequest={
                names: inputArg.status == 'batch' ? vm.data.batch.query.join(',') : inputArg.item.name
            },tmpModal={
                title: inputArg.status == 'batch' ? '批量删除服务注册方式' : ('删除服务注册方式-' + inputArg.item.name)
            };
            $rootScope.EnsureModal(tmpModal.title, null, '确认删除？', {}, function (callback) {
                if (callback) {
                    GatewayResource.ServiceDiscovery.Delete(tmpAjaxRequest).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                                {
                                    switch (inputArg.status) {
                                        case 'batch':
                                            {
                                                privateFun.resetBatchInfo();
                                                vm.fun.init();
                                                break;
                                            }
                                        case 'single':
                                            {
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
        };
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['服务注册方式']
            });
            vm.component.listDefaultCommonObject = {
                item: {
                    primaryKey: 'name',
                    default: [{
                        key: '服务注册方式',
                        html: '{{item.name}}'
                    },{
                        key:'服务类型',
                        html:'{{item.driver}}'
                    }, {
                        key: '更新时间',
                        html: '{{item.updateTime}}',
                        keyStyle: {
                            'width': '200px'
                        }
                    }],
                    operate: {
                        funArr: [{
                                key: '修改',
                                show: false,
                                fun: privateFun.edit,
                                params: `"edit",arg`
                            },
                            {
                                key: '删除',
                                show: false,
                                fun: privateFun.delete,
                                params: {
                                    status:'single'
                                }
                            }
                        ],
                        class:'w_200'
                    }
                },
                setting: {
                    batch: true,
                    unhover: true,
                    warning:'尚未添加任何内容',
                    defaultFoot:true
                }
            }
            vm.component.menuObject = {
                list: [{
                        type: 'btn',
                        authority: 'edit',
                        class: 'pull-left',
                        btnList: [{
                            name: '服务注册方式',
                            icon: 'jiahao',
                            class: 'eo_theme_btn_success block-btn',
                            fun: {
                                default: privateFun.edit,
                                params: '"add"'
                            }
                        }]
                    },{
                        type: 'search',
                        class: 'pull-right',
                        keyword: vm.ajaxRequest.keyword,
                        fun: privateFun.search,
                        placeholder:"输入服务注册方式名称"
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
                    titleAuthority: 'showTitle',
                    title: '服务注册方式',
                    secondTitle:"您可以通过静态或动态的方式来注册（发现）您的后端服务，创建好服务注册方式后，您可以在某个方式的基础上创建一个或多个负载（Upstream）"
                }
            }
        }
    }
})();