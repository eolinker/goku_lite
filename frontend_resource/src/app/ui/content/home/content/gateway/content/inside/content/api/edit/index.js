(function() {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api编辑相关js
     */
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home.gateway.inside.api.edit', {
                    url: '/operate/:status?apiID?groupID?childGroupID',
                    template: '<gateway-api-edit></gateway-api-edit>'
                });
        }])
        .component('gatewayApiEdit', {
            templateUrl: 'app/ui/content/home/content/gateway/content/inside/content/api/edit/index.html',
            controller: gatewayApiEditController
        })

    gatewayApiEditController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', 'GroupService', 'Cache_CommonService', '$filter'];

    function gatewayApiEditController($scope, GatewayResource, $state, CODE, $rootScope, GroupService, Cache_CommonService, $filter) {
        var vm = this;
        var code = CODE.COMMON.SUCCESS;
        vm.data = {
            info: {
                input: {
                    disable: false,
                    submited: false
                },
                group: {
                    parent: [],
                    child: []
                },
                reset: {
                    gatewayHashKey: $state.params.gatewayHashKey,
                    groupID: $state.params.groupID || -1,
                    childGroupID: $state.params.childGroupID,
                    apiID: $state.params.apiID,
                    status: $state.params.status,
                    backendItem: {},
                    backendQuery: [{ backendName: '无', backendID: -1 }],
                    type: {
                        gateway: true,
                        backend: true
                    },
                    check: {
                        gatewayRequestParam: false,
                        constantResultParam: false
                    }
                }
            },
            interaction: {
                response: {
                    apiInfo: {
                        gatewayHashKey: $state.params.gatewayHashKey,
                        groupID: $state.params.groupID,
                        childGroupID: $state.params.childGroupID,
                        apiID: $state.params.apiID,
                        apiName: '',
                        gatewayProtocol: '0',
                        gatewayRequestType: '0',
                        gatewayRequestPath: '',
                        backendProtocol: '0',
                        backendRequestType: '0',
                        backendID: '',
                        backendURI: '',
                        backendRequestPath: '',
                        isRequestBody: false,
                        gatewayRequestBodyNote: '',
                        gatewayRequestParam: [],
                        constantResultParam: []
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
                    method: null, //改变请求方式 boolean:true gatewayRequestType false:backendRequestType
                    required:null,//接受参数映射表是否必填
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
                filterArray:null,//过滤数组功能函数
            }
        }
        vm.component = {
            menuObject: {
                list: []
            }
        };
        vm.service={
            cache:Cache_CommonService
        }
        vm.data.assistantFun.init = function() {
            var apiGroup = GroupService.get();
            vm.data.info.group.parent = apiGroup;
            if (vm.data.interaction.response.apiInfo.groupID > 0) {
                for (var i = 0; i < vm.data.info.group.parent.length; i++) {
                    var val = vm.data.info.group.parent[i];
                    if (val.groupID == vm.data.interaction.response.apiInfo.groupID) {
                        vm.data.info.group.child = [{ groupID: -1, groupName: '可选[二级菜单]' }].concat(val.childGroupList);
                        break;
                    }
                }
            } else {
                vm.data.info.group.child = [{ groupID: -1, groupName: '可选[二级菜单]' }].concat(vm.data.info.group.parent[0].childGroupList);
            }
            if (vm.data.interaction.response.apiInfo.apiID || vm.data.interaction.response.apiInfo.groupID > 0) {
                vm.data.interaction.response.apiInfo.groupID = parseInt(vm.data.interaction.response.apiInfo.groupID);
                if (vm.data.interaction.response.apiInfo.childGroupID) {
                    vm.data.interaction.response.apiInfo.childGroupID = parseInt(vm.data.interaction.response.apiInfo.childGroupID);
                } else {
                    vm.data.interaction.response.apiInfo.childGroupID = -1;
                }
            } else {
                vm.data.interaction.response.apiInfo.groupID = vm.data.info.group.parent[0].groupID;
                vm.data.interaction.response.apiInfo.childGroupID = -1;
            }
        }

        vm.data.fun.init=function() {
            var template={
                response:null
            }
            if (vm.data.info.reset.apiID) {
                GatewayResource.Api.Detail({
                    apiID: vm.data.info.reset.apiID,
                    groupID: vm.data.info.reset.childGroupID || vm.data.info.reset.groupID,
                    gatewayHashKey: vm.data.info.reset.gatewayHashKey
                }).$promise.then(function(data) {
                    if (code == data.statusCode) {
                        template.response=vm.data.interaction.response.apiInfo = data.apiInfo.apiJson;
                        vm.data.interaction.response.apiInfo.gatewayRequestParam = template.response.requestParams||[];
                        vm.data.interaction.response.apiInfo.constantResultParam = template.response.constantParams||[];
                        vm.data.interaction.response.apiInfo.backendRequestPath = template.response.backendPath;
                        vm.data.interaction.response.apiInfo.gatewayProtocol=(vm.data.interaction.response.apiInfo.gatewayProtocol||0).toString();
                        vm.data.interaction.response.apiInfo.gatewayRequestType=(vm.data.interaction.response.apiInfo.gatewayRequestType||0).toString();
                        vm.data.interaction.response.apiInfo.backendProtocol=(vm.data.interaction.response.apiInfo.backendProtocol||0).toString();
                        vm.data.interaction.response.apiInfo.backendRequestType=(vm.data.interaction.response.apiInfo.backendRequestType||0).toString();
                        vm.data.interaction.response.apiInfo.isRequestBody = vm.data.interaction.response.apiInfo.isRequestBody == '0' ? false : true;
                        $scope.$emit('$windowTitle', { apiName: (vm.data.info.reset.status == 'copy' ? '[另存为]' : '[修改]') + vm.data.interaction.response.apiInfo.apiName });
                        if (!!vm.data.interaction.response.apiInfo.parentGroupID) {
                            vm.data.interaction.response.apiInfo.childGroupID = data.apiInfo.groupID;
                            vm.data.interaction.response.apiInfo.groupID = vm.data.interaction.response.apiInfo.parentGroupID;
                        } else {
                            vm.data.interaction.response.apiInfo.childGroupID = -1;
                        }
                        for (var i = 0; i < vm.data.info.reset.backendQuery.length; i++) {
                            if (vm.data.info.reset.backendQuery[i].backendID == vm.data.interaction.response.apiInfo.backendID) {
                                vm.data.info.reset.backendItem = vm.data.info.reset.backendQuery[i];
                                break;
                            }
                        }
                        angular.forEach(vm.data.interaction.response.apiInfo.gatewayRequestParam, function(val, key) {
                            val.checkbox = val.isNotNull == '1' ? true : false;
                        })
                        vm.data.assistantFun.init();
                    }
                });
            } else {
                vm.data.assistantFun.init();
                vm.data.info.reset.backendItem = vm.data.info.reset.backendQuery[0];
                $scope.$emit('$windowTitle', { apiName: '[新增接口]' });
            }
        }
        vm.data.assistantFun.checkInitStatus = function() {
            if (!!GroupService.get()) {
                vm.data.info.reset.backendQuery = vm.service.cache.get('backend');
                if (!vm.data.info.reset.backendQuery) {
                    GatewayResource.Backend.Query({ gatewayHashKey: vm.data.info.reset.gatewayHashKey }).$promise.then(function(data) {
                        if (code == data.statusCode) {
                            vm.data.info.reset.backendQuery = data.backendList;
                            vm.service.cache.set(data.backendList,'backend')
                            vm.data.fun.init();
                        } else {
                            $state.go('home.gateway.inside.api.list', { 'groupID': vm.data.info.reset.groupID, 'childGroupID': vm.data.info.reset.childGroupID });
                        }
                    })
                } else {
                    vm.data.fun.init();
                }
            }
        }
        vm.data.assistantFun.checkInitStatus();
        vm.data.fun.change.group = function() {
            for (var i = 0; i < vm.data.info.group.parent.length; i++) {
                var val = vm.data.info.group.parent[i];
                if (val.groupID == vm.data.interaction.response.apiInfo.groupID) {
                    vm.data.info.group.child = [{ groupID: -1, groupName: '可选[二级菜单]' }].concat(val.childGroupList);
                    vm.data.interaction.response.apiInfo.childGroupID = -1;
                    break;
                }
            }
        }
        vm.data.fun.requestParamList.add = function() {
            var info = {
                gatewayParamPosition: '0',
                isNotNull: '1',
                checkbox: true,
                paramType: '0',
                gatewayParamKey: '',
                backendParamPosition: '0',
                backendParamKey: ''
            }
            vm.data.interaction.response.apiInfo.gatewayRequestParam.push(info);
        }
        vm.data.fun.change.required = function(query) {
            query.isNotNull = query.checkbox ? '1' : '0';
        }
        vm.data.fun.requestParamList.delete = function(arg) {
            vm.data.interaction.response.apiInfo.gatewayRequestParam.splice(arg.$index, 1);
        }
        vm.data.fun.resultParamList.add = function() {
            var info = {
                checkbox: true,
                paramPosition: '0',
                backendParamKey: '',
                paramValue: '',
                paramName: ''
            }
            vm.data.interaction.response.apiInfo.constantResultParam.push(info);
        }
        vm.data.fun.resultParamList.delete = function(arg) {
            vm.data.interaction.response.apiInfo.constantResultParam.splice(arg.$index, 1);
        }
        vm.data.fun.back = function() {
            if (vm.data.info.reset.apiID) {
                $state.go('home.gateway.inside.api.simple', { 'groupID': vm.data.info.reset.groupID, 'childGroupID': vm.data.info.reset.childGroupID, 'apiID': vm.data.info.reset.apiID });
            } else {
                $state.go('home.gateway.inside.api.list', { 'groupID': vm.data.info.reset.groupID, 'childGroupID': vm.data.info.reset.childGroupID });
            }
        }
        vm.data.fun.change.method = function(boolean) { 
            if (boolean) {
                switch (vm.data.interaction.response.apiInfo.gatewayRequestType) {
                    case '0':
                    case '2':
                    case '6':
                        {
                            vm.data.info.reset.type.gateway = true;
                            break;
                        }
                    default:
                        {
                            vm.data.info.reset.type.gateway = false;
                            break;
                        }
                }
            } else {
                switch (vm.data.interaction.response.apiInfo.backendRequestType) {
                    case '0':
                    case '2':
                    case '6':
                        {
                            vm.data.info.reset.type.backend = true;
                            break;
                        }
                    default:
                        {
                            vm.data.info.reset.type.backend = false;
                            break;
                        }
                }
            }
        }
        vm.data.fun.filter = function(e) {
            if (((vm.data.interaction.response.apiInfo.isRequestBody || !vm.data.info.reset.type.gateway) && (e.gatewayParamPosition == '1' || e.paramPosition == '1')) || (!vm.data.info.reset.type.backend && e.backendParamPosition == '1')) {
                return false;
            } else {
                return true;
            }
        }
        vm.data.assistantFun.filterArray = function(array, which) {
            var array = angular.copy(array);
            for (var i = 0; i < array.length; i++) {
                var val = array[i];
                if ((val[which] == '1' && (vm.data.interaction.response.apiInfo.isRequestBody || !vm.data.info.reset.type.gateway)) || (!vm.data.info.reset.type.backend && val.backendParamPosition == '1')) {
                    array.splice(i, 1);
                    i--;
                }
            }
            return array;
        }
        vm.data.assistantFun.confirm = function() {
            var info = {
                gatewayHashKey: vm.data.info.reset.gatewayHashKey,
                groupID: vm.data.interaction.response.apiInfo.childGroupID > 0 ? vm.data.interaction.response.apiInfo.childGroupID : vm.data.interaction.response.apiInfo.groupID,
                apiID: vm.data.info.reset.apiID,
                apiName: vm.data.interaction.response.apiInfo.apiName,
                gatewayProtocol: vm.data.interaction.response.apiInfo.gatewayProtocol,
                gatewayRequestType: vm.data.interaction.response.apiInfo.gatewayRequestType,
                gatewayRequestPath: /^\//.test(vm.data.interaction.response.apiInfo.gatewayRequestPath) ? vm.data.interaction.response.apiInfo.gatewayRequestPath : ('/' + vm.data.interaction.response.apiInfo.gatewayRequestPath),
                backendProtocol: vm.data.interaction.response.apiInfo.backendProtocol,
                backendRequestType: vm.data.interaction.response.apiInfo.backendRequestType,
                backendID: vm.data.info.reset.backendItem.backendID,
                backendURI: vm.data.info.reset.backendItem.backendURI,
                backendRequestPath: /^\//.test(vm.data.interaction.response.apiInfo.backendRequestPath) ? vm.data.interaction.response.apiInfo.backendRequestPath : ('/' + vm.data.interaction.response.apiInfo.backendRequestPath),
                isRequestBody: vm.data.interaction.response.apiInfo.isRequestBody && vm.data.info.reset.type.gateway ? '1' : '0',
                gatewayRequestBodyNote: vm.data.interaction.response.apiInfo.gatewayRequestBodyNote,
                gatewayRequestParam: '',
                constantResultParam: ''
            }
            var template = {
                gatewayRequestParam: vm.data.interaction.response.apiInfo.gatewayRequestParam,
                constantResultParam: vm.data.interaction.response.apiInfo.constantResultParam
            };
            var i = 0,
                j = 0;
            vm.check = {
                constantResultParam: false,
                gatewayRequestParam: false
            }
            for (i = template.gatewayRequestParam.length - 1; i >= 0; i--) { //请求参数映射表
                if (!template.gatewayRequestParam[i].gatewayParamKey) {
                    if (!template.gatewayRequestParam[i].backendParamKey) {
                        vm.data.interaction.response.apiInfo.gatewayRequestParam.splice(i, 1);
                    } else {
                        vm.check.gatewayRequestParam = true;
                    }
                }
            }
            for (i = template.constantResultParam.length - 1; i >= 0; i--) { //常量参数
                if (!template.constantResultParam[i].backendParamKey) {
                    if ((!template.constantResultParam[i].paramValue) && (!template.constantResultParam[i].paramName)) {
                        vm.data.interaction.response.apiInfo.constantResultParam.splice(i, 1);
                    } else {
                        vm.check.constantResultParam = true;
                    }
                }
            }
            info.gatewayRequestParam = JSON.stringify(vm.data.assistantFun.filterArray(template.gatewayRequestParam, 'gatewayParamPosition'), function(key, val) {
                if (key === "$$hashKey" || key === "checkbox") {
                    return undefined;
                }
                return val;
            });
            info.constantResultParam = JSON.stringify(vm.data.assistantFun.filterArray(template.constantResultParam, 'paramPosition'), function(key, val) {
                if (key === "$$hashKey" || key === "checkbox") {
                    return undefined;
                }
                return val;
            });
            return info;
        }
        vm.data.fun.load = function(arg) {
            $scope.$emit('$TransferStation', { state: '$LoadingInit', data: arg });
        }
        vm.data.fun.requestProcessing = function(arg) { //arg status:（0：继续添加 1：快速保存，2：编辑（修改/新增））
            var template = {
                request: vm.data.assistantFun.confirm(),
                promise: null
            }
            if ($scope.editForm.$valid && (!vm.check.constantResultParam) && (!vm.check.gatewayRequestParam)) {
                vm.data.info.input.disable = true;
                switch (arg.status) {
                    case 0:
                        {
                            template.promise = vm.data.assistantFun.keep({ request: template.request });
                            break;
                        }
                    case 1:
                        {
                            template.promise = vm.data.assistantFun.edit({ request: template.request });
                            break;
                        }
                }
            } else {
                $rootScope.InfoModal('Api编辑失败，请检查信息是否填写完整！', 'error');
                vm.data.info.input.submited = true;
            }
            return template.promise;
        }
        vm.data.assistantFun.keep = function(arg) {
            var template = {
                promise: null
            }
            template.promise = GatewayResource.Api.Add(arg.request).$promise;
            template.promise.then(function(data) {
                vm.data.info.input.disable = false;
                if (data.statusCode == code) {
                    $rootScope.InfoModal('Api添加成功', 'success');
                    vm.data.interaction.response.apiInfo = {
                        gatewayHashKey: $state.params.gatewayHashKey,
                        groupID: vm.data.info.reset.groupID == '-1' ? vm.data.info.group.parent[0].groupID : parseInt(vm.data.info.reset.groupID),
                        apiID: $state.params.apiID,
                        apiName: '',
                        gatewayProtocol: '0',
                        gatewayRequestType: '0',
                        gatewayRequestPath: '',
                        backendProtocol: '0',
                        backendRequestType: '0',
                        backendID: '',
                        backendURI: '',
                        backendRequestPath: '',
                        isRequestBody: false,
                        gatewayRequestBodyNote: '',
                        gatewayRequestParam: [],
                        constantResultParam: []
                    };
                    vm.data.info.reset.backendItem = vm.data.info.reset.backendQuery[0];
                    if (vm.data.info.reset.groupID > 0) {
                        for (var i = 0; i < vm.data.info.group.parent.length; i++) {
                            var val = vm.data.info.group.parent[i];
                            if (val.groupID == vm.data.info.reset.groupID) {
                                vm.data.info.group.child = [{ groupID: -1, groupName: '可选[二级菜单]' }].concat(val.childGroupList);
                                break;
                            }
                        }
                    } else {
                        vm.data.info.group.child = [{ groupID: -1, groupName: '可选[二级菜单]' }].concat(vm.data.info.group.parent[0].childGroupList);
                    }
                    if (vm.data.info.reset.childGroupID) {
                        vm.data.interaction.response.apiInfo.childGroupID = parseInt(vm.data.info.reset.childGroupID);
                    } else {
                        vm.data.interaction.response.apiInfo.childGroupID = -1;
                    }
                    vm.data.info.input.submited = false;
                    window.scrollTo(0, 0);
                }
            })
            return template.promise;
        }
        vm.data.assistantFun.edit = function(arg) {
            var template = {
                promise: null
            }
            if (vm.data.info.reset.status=='edit') {
                template.promise = GatewayResource.Api.Edit(arg.request).$promise;
                template.promise.then(function(data) {
                    vm.data.info.input.disable = false;
                    if (data.statusCode == code) {
                        $state.go('home.gateway.inside.api.simple', { 'groupID': vm.data.info.reset.groupID, 'childGroupID': vm.data.info.reset.childGroupID, 'apiID': vm.data.info.reset.apiID });
                        $rootScope.InfoModal('Api修改成功', 'success');
                    }
                })
            } else {
                template.promise = GatewayResource.Api.Add(arg.request).$promise;
                template.promise.then(function(data) {
                    vm.data.info.input.disable = false;
                    if (data.statusCode == code) {
                        $state.go('home.gateway.inside.api.simple', { 'groupID': vm.data.info.reset.groupID, 'childGroupID': vm.data.info.reset.childGroupID, 'apiID': data.apiID });
                        $rootScope.InfoModal('Api添加成功', 'success');
                    }
                })
            }
            return template.promise;
        }

        $scope.$on('$SidebarFinish', function() {
            GatewayResource.Backend.Query({ gatewayHashKey: vm.data.info.reset.gatewayHashKey }).$promise.then(function(data) {
                if (code == data.statusCode) {
                    vm.data.info.reset.backendQuery = data.backendList;
                    vm.service.cache.set(data.backendList,'backend');
                    vm.data.fun.init();
                } else {
                    $state.go('home.gateway.inside.api.list', { 'groupID': vm.data.info.reset.groupID, 'childGroupID': vm.data.info.reset.childGroupID });
                }
            })
        })
        vm.$onInit = function() {
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
                                params: { status: 0 }
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
                class: 'btn-group-li pull-right',
                btnList: template.array.concat([{
                    name: '保存',
                    class: 'eo-button-success',
                    fun: {
                        disabled: 1,
                        default: vm.data.fun.requestProcessing,
                        params: { status: 1 }
                    }
                }])
            }];
        }
    }
})();
