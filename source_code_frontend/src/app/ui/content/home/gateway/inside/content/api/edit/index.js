(function () {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api编辑相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.api.edit', {
                    url: '/operate/:status?apiID?groupID',
                    template: '<gateway-api-edit></gateway-api-edit>'
                });
        }])
        .component('gatewayApiEdit', {
            templateUrl: 'app/ui/content/home/gateway/inside/content/api/edit/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', 'GroupService', 'Cache_CommonService', '$filter'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, GroupService, Cache_CommonService, $filter) {
        var vm = this;
        vm.data = {
            info: {
                requestMethodList: [{
                    key: 'POST',
                    value:'post'
                }, {
                    key: 'GET',
                    value:'get'
                }, {
                    key: 'PUT',
                    value:'put'
                }, {
                    key: 'DELETE',
                    value:'delete'
                }, {
                    key: 'HEAD',
                    value:'head'
                }, {
                    key: 'OPTIONS',
                    value:'options'
                }, {
                    key: 'PATCH',
                    value:'patch'
                }],
                input: {
                    disable: false,
                    submited: false
                },
                group: {
                    parent: []
                },
                reset: {
                    gatewayAlias: $state.params.gatewayAlias,
                    groupID: $state.params.groupID || -1,
                    apiID: $state.params.apiID,
                    status: $state.params.status,
                    backendItem: {},
                    backendQuery: [{
                        backendName: '无',
                        backendID: -1
                    }],
                    type: {
                        gateway: true,
                        backend: true
                    },
                    check: {
                        proxyParams: false,
                        constantParams: false
                    }
                }
            },
            interaction: {
                response: {
                    apiInfo: {
                        gatewayAlias: $state.params.gatewayAlias,
                        groupID: $state.params.groupID,
                        apiID: $state.params.apiID,
                        apiName: '',
                        requestMethod: '',
                        requestURL: '',
                        backendProtocol: '0',
                        proxyMethod: '-1',
                        backendID: -1,
                        backendPath: '',
                        proxyURL: '',
                        isRaw: false,
                        proxyParams: [],
                        constantParams: []
                    }
                }
            },
            fun: {
                init: null, //初始化功能函数
                load: null, //编辑相关系列按钮功能函数
                requestProcessing: null, //发送存储请求时预处理功能函数
                filter: null, //过滤分组
                change: {
                    group: null, //更改父分组
                    required: null, //接受参数映射表是否必填
                },
                requestParamList: {
                    add: null, //添加请求参数值功能函数
                    delete: null, //删除请求参数值功能函数
                },
                resultParamList: {
                    add: null, //添加返回参数值功能函数
                    delete: null, //删除返回参数值功能函数
                },
                back: null, //返回功能函数
            },
            assistantFun: {
                init: null, //辅助初始化功能函数
                confirm: null, //辅助确认功能函数
                keep: null, //辅助继续添加功能函数
                checkInitStatus: null, //辅助检测准备状态功能函数
                edit: null, //编辑功能函数
                filterArray: null, //过滤数组功能函数
            }
        }
        vm.component = {
            menuObject: {
                list: []
            }
        };
        vm.service = {
            cache: Cache_CommonService
        }
        vm.data.assistantFun.init = function () {
            var apiGroup = GroupService.get();
            vm.data.info.group.parent = apiGroup;
            if (vm.data.interaction.response.apiInfo.apiID || vm.data.interaction.response.apiInfo.groupID > 0) {
                vm.data.interaction.response.apiInfo.groupID = parseInt(vm.data.interaction.response.apiInfo.groupID);
            } else {
                vm.data.interaction.response.apiInfo.groupID = vm.data.info.group.parent[0].groupID;
            }
        }

        vm.data.fun.init = function () {
            var template = {
                requestMethod:[]
            }
            if (vm.data.info.reset.apiID) {
                GatewayResource.Api.Detail({
                    apiID: vm.data.info.reset.apiID,
                    groupID: vm.data.info.reset.groupID,
                    gatewayAlias: vm.data.info.reset.gatewayAlias
                }).$promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.data.interaction.response.apiInfo = response.apiInfo;
                                vm.data.interaction.response.apiInfo.proxyParams = response.apiInfo.proxyParams || [];
                                vm.data.interaction.response.apiInfo.constantParams = response.apiInfo.constantParams || [];
                                vm.data.interaction.response.apiInfo.proxyMethod=response.apiInfo.follow?'-1':vm.data.interaction.response.apiInfo.proxyMethod;
                                template.requestMethod=response.apiInfo.requestMethod.split(',');
                                for(var key in template.requestMethod){
                                    var val=template.requestMethod[key];
                                    switch(val){
                                        case 'post':{
                                            vm.data.info.requestMethodList[0].checkbox=true;
                                            break;
                                        }
                                        case 'get':{
                                            vm.data.info.requestMethodList[1].checkbox=true;
                                            break;
                                        }
                                        case 'put':{
                                            vm.data.info.requestMethodList[2].checkbox=true;
                                            break;
                                        }
                                        case 'delete':{
                                            vm.data.info.requestMethodList[3].checkbox=true;
                                            break;
                                        }
                                        case 'head':{
                                            vm.data.info.requestMethodList[4].checkbox=true;
                                            break;
                                        }
                                        case 'options':{
                                            vm.data.info.requestMethodList[5].checkbox=true;
                                            break;
                                        }
                                        case 'batch':{
                                            vm.data.info.requestMethodList[6].checkbox=true;
                                            break;
                                        }
                                    }
                                }
                                $scope.$emit('$windowTitle', {
                                    apiName: (vm.data.info.reset.status == 'copy' ? '[另存为]' : '[修改]') + vm.data.interaction.response.apiInfo.apiName
                                });
                                for (var i = 0; i < vm.data.info.reset.backendQuery.length; i++) {
                                    if (vm.data.info.reset.backendQuery[i].backendID == vm.data.interaction.response.apiInfo.backendID) {
                                        vm.data.info.reset.backendItem = vm.data.info.reset.backendQuery[i];
                                        break;
                                    }
                                }
                                vm.data.assistantFun.init();
                                break;
                            }
                    }
                });
            } else {
                vm.data.assistantFun.init();
                vm.data.info.reset.backendItem = vm.data.info.reset.backendQuery[0];
                $scope.$emit('$windowTitle', {
                    apiName: '[新增接口]'
                });
            }
        }
        vm.data.assistantFun.checkInitStatus = function () {
            if (!!GroupService.get()) {
                vm.data.info.reset.backendQuery = vm.service.cache.get('backend');
                if (!vm.data.info.reset.backendQuery) {
                    GatewayResource.Backend.Query({
                        gatewayAlias: vm.data.info.reset.gatewayAlias
                    }).$promise.then(function (response) {
                        vm.data.info.reset.backendQuery = [{
                            backendName: '无',
                            backendID:-1
                        }].concat(response.backendList||[]);
                        vm.service.cache.set(response.backendList, 'backend')
                        vm.data.fun.init();
                    })
                } else {
                    vm.data.info.reset.backendQuery = [{
                        backendName: '无',
                        backendID:-1
                    }].concat(vm.data.info.reset.backendQuery);
                    vm.data.fun.init();
                }
            }
        }
        vm.data.assistantFun.checkInitStatus();
        vm.data.fun.requestParamList.add = function () {
            var info = {
                keyPosition: 'header',
                notEmpty: true,
                key: '',
                proxyKeyPosition: 'header',
                proxyKey: ''
            }
            vm.data.interaction.response.apiInfo.proxyParams.push(info);
        }
        vm.data.fun.requestParamList.delete = function (arg) {
            vm.data.interaction.response.apiInfo.proxyParams.splice(arg.$index, 1);
        }
        vm.data.fun.resultParamList.add = function () {
            var info = {
                position: 'header',
                key: '',
                value: ''
            }
            vm.data.interaction.response.apiInfo.constantParams.push(info);
        }
        vm.data.fun.resultParamList.delete = function (arg) {
            vm.data.interaction.response.apiInfo.constantParams.splice(arg.$index, 1);
        }
        vm.data.fun.back = function () {
            if (vm.data.info.reset.apiID) {
                $state.go('home.gateway.inside.api.simple', {
                    'groupID': vm.data.info.reset.groupID,
                    'apiID': vm.data.info.reset.apiID
                });
            } else {
                $state.go('home.gateway.inside.api.list', {
                    'groupID': vm.data.info.reset.groupID
                });
            }
        }
        vm.data.fun.filter = function (e) {
            if (vm.data.interaction.response.apiInfo.isRaw && (e.keyPosition == 'body' || e.position == 'body')) {
                return false;
            } else {
                return true;
            }
        }
        vm.data.assistantFun.filterArray = function (array, which) {
            var array = angular.copy(array);
            for (var i = 0; i < array.length; i++) {
                var val = array[i];
                if (val[which] == '1' && (vm.data.interaction.response.apiInfo.isRaw)) {
                    array.splice(i, 1);
                    i--;
                }
            }
            return array;
        }
        vm.data.assistantFun.confirm = function () {
            var info = {
                gatewayAlias: vm.data.info.reset.gatewayAlias,
                groupID: vm.data.interaction.response.apiInfo.groupID,
                apiID: vm.data.info.reset.apiID,
                apiName: vm.data.interaction.response.apiInfo.apiName,
                requestMethod: [],
                requestURL: /^\//.test(vm.data.interaction.response.apiInfo.requestURL) ? vm.data.interaction.response.apiInfo.requestURL : ('/' + vm.data.interaction.response.apiInfo.requestURL),
                proxyMethod: vm.data.interaction.response.apiInfo.proxyMethod=='-1'?null:vm.data.interaction.response.apiInfo.proxyMethod,
                backendID: vm.data.info.reset.backendItem.backendID==-1?null:vm.data.info.reset.backendItem.backendID,
                backendPath: vm.data.info.reset.backendItem.backendPath,
                proxyURL: /^\//.test(vm.data.interaction.response.apiInfo.proxyURL) ? vm.data.interaction.response.apiInfo.proxyURL : ('/' + vm.data.interaction.response.apiInfo.proxyURL),
                isRaw: vm.data.interaction.response.apiInfo.isRaw||false,
                proxyParams: '',
                constantParams: '',
                follow:vm.data.interaction.response.apiInfo.proxyMethod=='-1'?true:false
            }
            var template = {
                proxyParams: vm.data.interaction.response.apiInfo.proxyParams,
                constantParams: vm.data.interaction.response.apiInfo.constantParams
            };
            var i = 0,
                j = 0;
            vm.check = {
                constantParams: false,
                proxyParams: false,
                requestMethod:false
            }
            for(var key in vm.data.info.requestMethodList){
                if(vm.data.info.requestMethodList[key].checkbox){
                    info.requestMethod.push(vm.data.info.requestMethodList[key].value);
                }
            }
            if(info.requestMethod.length>0){
                info.requestMethod=info.requestMethod.join(',');
            }else{
                vm.check.requestMethod=true;
            }
            for (i = template.proxyParams.length - 1; i >= 0; i--) { //请求参数映射表
                if (!template.proxyParams[i].key) {
                    if (!template.proxyParams[i].proxyKey) {
                        vm.data.interaction.response.apiInfo.proxyParams.splice(i, 1);
                    } else {
                        vm.check.proxyParams = true;
                    }
                }
            }
            for (i = template.constantParams.length - 1; i >= 0; i--) { //常量参数
                if (!template.constantParams[i].key) {
                    if (!template.constantParams[i].value) {
                        vm.data.interaction.response.apiInfo.constantParams.splice(i, 1);
                    } else {
                        vm.check.constantParams = true;
                    }
                }
            }
            info.proxyParams = JSON.stringify(vm.data.assistantFun.filterArray(template.proxyParams, 'keyPosition'), function (key, val) {
                if (key === "$$hashKey") {
                    return undefined;
                }
                return val;
            });
            info.constantParams = JSON.stringify(vm.data.assistantFun.filterArray(template.constantParams, 'position'), function (key, val) {
                if (key === "$$hashKey") {
                    return undefined;
                }
                return val;
            });
            return info;
        }
        vm.data.fun.load = function (arg) {
            $scope.$emit('$TransferStation', {
                state: '$LoadingInit',
                data: arg
            });
        }
        vm.data.fun.requestProcessing = function (arg) { //arg status:（0：继续添加 1：快速保存，2：编辑（修改/新增））
            var template = {
                request: vm.data.assistantFun.confirm(),
                promise: null
            }
            if ($scope.editForm.$valid && (!vm.check.constantParams) && (!vm.check.proxyParams)&&(!vm.check.requestMethod)) {
                vm.data.info.input.disable = true;
                switch (arg.status) {
                    case 0:
                        {
                            template.promise = vm.data.assistantFun.keep({
                                request: template.request
                            });
                            break;
                        }
                    case 1:
                        {
                            template.promise = vm.data.assistantFun.edit({
                                request: template.request
                            });
                            break;
                        }
                }
            } else {
                $rootScope.InfoModal('Api编辑失败，请检查信息是否填写完整！', 'error');
                vm.data.info.input.submited = true;
            }
            return template.promise;
        }
        vm.data.assistantFun.keep = function (arg) {
            var template = {
                promise: null
            }
            template.promise = GatewayResource.Api.Add(arg.request).$promise;
            template.promise.then(function (response) {
                vm.data.info.input.disable = false;
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            $rootScope.InfoModal('Api添加成功', 'success');
                            vm.data.interaction.response.apiInfo = {
                                gatewayAlias: $state.params.gatewayAlias,
                                groupID: vm.data.info.reset.groupID == '-1' ? vm.data.info.group.parent[0].groupID : parseInt(vm.data.info.reset.groupID),
                                apiID: $state.params.apiID,
                                apiName: '',
                                requestMethod: '',
                                requestURL: '',
                                proxyMethod: '-1',
                                backendID: -1,
                                backendPath: '',
                                proxyURL: '',
                                isRaw: false,
                                proxyParams: [],
                                constantParams: []
                            };
                            vm.data.info.reset.backendItem = vm.data.info.reset.backendQuery[0];
                            vm.data.info.input.submited = false;
                            window.scrollTo(0, 0);
                            break;
                        }
                    case '190000':
                        {
                            $rootScope.InfoModal('Api添加失败，接口地址请求路径重复！', 'error');
                            break;
                        }
                    default:
                        {
                            $rootScope.InfoModal('Api添加失败!', 'error');
                            break;
                        }
                }
            })
            return template.promise;
        }
        vm.data.assistantFun.edit = function (arg) {
            var template = {
                promise: null
            }
            if (vm.data.info.reset.status == 'edit') {
                template.promise = GatewayResource.Api.Edit(arg.request).$promise;
                template.promise.then(function (response) {
                    vm.data.info.input.disable = false;
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                $state.go('home.gateway.inside.api.simple', {
                                    'groupID': vm.data.info.reset.groupID,
                                    'apiID': vm.data.info.reset.apiID
                                });
                                $rootScope.InfoModal('Api修改成功', 'success');
                                break;
                            }
                    }
                })
            } else {
                template.promise = GatewayResource.Api.Add(arg.request).$promise;
                template.promise.then(function (response) {
                    vm.data.info.input.disable = false;
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                $state.go('home.gateway.inside.api.simple', {
                                    'groupID': vm.data.info.reset.groupID,
                                    'apiID': response.apiID
                                });
                                $rootScope.InfoModal('Api添加成功', 'success');
                                break;
                            }
                        case '190000':
                            {
                                $rootScope.InfoModal('Api添加失败，接口地址请求路径重复！', 'error');
                                break;
                            }
                        default:
                            {
                                $rootScope.InfoModal('Api添加失败!', 'error');
                                break;
                            }
                    }
                })
            }
            return template.promise;
        }

        $scope.$on('$SidebarFinish', function () {
            GatewayResource.Backend.Query({
                gatewayAlias: vm.data.info.reset.gatewayAlias
            }).$promise.then(function (response) {
                vm.data.info.reset.backendQuery = [{
                    backendName: '无',
                    backendID:-1
                }].concat(response.backendList||[]);
                vm.service.cache.set(response.backendList, 'backend')
                vm.data.fun.init();
            })
        })
        vm.$onInit = function () {
            var template = {
                array: []
            }
            switch (vm.data.info.reset.status) {
                case 'add':
                    {
                        template.array = [{
                            name: '继续添加',
                            class: 'eo-button-success',
                            fun: {
                                disabled: 1,
                                default: vm.data.fun.requestProcessing,
                                params: {
                                    status: 0
                                }
                            }
                        }]
                        break;
                    }
            }
            vm.component.menuObject.list = [{
                type: 'btn',
                class: 'btn-group-li pull-left',
                btnList: [{
                    name: vm.data.info.reset.status == 'add' ? '返回列表' : '返回详情',
                    icon: 'xiangzuo',
                    fun: {
                        default: vm.data.fun.back
                    }
                }]
            }, {
                type: 'btn',
                class: 'btn-group-li',
                btnList: template.array.concat([{
                    name: '保存',
                    class: 'eo-button-info',
                    fun: {
                        disabled: 1,
                        default: vm.data.fun.requestProcessing,
                        params: {
                            status: 1
                        }
                    }
                }])
            }];
        }
    }
})();