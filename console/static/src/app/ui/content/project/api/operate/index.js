(function () {
    'use strict';
    /*
     * author：riverLethe
     * 网关内页页内包模块api编辑相关js
     */
    angular.module('eolinker')
        .component('apiOperate', {
            templateUrl: 'app/ui/content/project/api/operate/index.html',
            controller: indexController
        })
        .filter('Filter_SlashSymbol', () => {
            return (inputStr) => {
                return (inputStr || '').replace(/\/{2,}/g, '/');
            }
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', 'GroupService'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, GroupService) {
        var vm = this;
        vm.data = {
            requestMethod: false,
            apiGroup: null,
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
        vm.fun = {};
        vm.ajaxRequest = {
            projectID: $state.params.projectID,
            groupID: $state.params.groupID || -1,
            apiID: $state.params.apiID
        }
        vm.ajaxResponse = {
            apiInfo: {
                apiName: '',
                requestMethod: '',
                requestURL: '/',
                balanceName: '',
                targetURL: '/',
                targetMethod: '-1',
                isFollow: true,
                stripPrefix: true,
                stripSlash:true,
                timeout: '2000',
                retryCount: '',
                alertValve: 0,
                protocol: 'http'
            },
            balanceList:[]
        }
        vm.component = {
            selectMultistageCommonComponentObject: {
                new: {},
                original: {}
            },
            menuObject: {
                list: []
            },
            balanceAutoCompleteObj:{required:true,pattern:'[\\w\\._\\/\\-\\:]+'}
        };
        vm.CONST = {
            PROTOCOL_ARR: [{
                key: 'HTTP',
                value: 'http'
            }, {
                key: 'HTTPS',
                value: 'https'
            }]
        }
        var privateFun = {};
        vm.fun.back = function () {
            $state.go('home.project.api.default', {
                'groupID': $state.params.groupID
            });
        }
        privateFun.confirm = function () {
            vm.data.requestMethod = false;
            var tmpOutput = {
                projectID: vm.ajaxRequest.projectID,
                groupID: vm.component.selectMultistageCommonComponentObject.new.value==-1?0:vm.component.selectMultistageCommonComponentObject.new.value,
                apiName: vm.ajaxResponse.apiInfo.apiName,
                requestMethod: [],
                requestURL: vm.ajaxResponse.apiInfo.requestURL,
                balanceName: vm.ajaxResponse.apiInfo.balanceName,
                targetURL: vm.ajaxResponse.apiInfo.targetURL,
                targetMethod: vm.ajaxResponse.apiInfo.targetMethod,
                stripPrefix: vm.ajaxResponse.apiInfo.stripPrefix,
                timeout: vm.ajaxResponse.apiInfo.timeout,
                retryCount: vm.ajaxResponse.apiInfo.retryCount,
                alertValve: vm.ajaxResponse.apiInfo.alertValve,
                protocol:vm.ajaxResponse.apiInfo.protocol,
                stripSlash:vm.ajaxResponse.apiInfo.stripSlash
            }
            for (var key in vm.data.requestMethodList) {
                if (vm.data.requestMethodList[key].checkbox) {
                    tmpOutput.requestMethod.push(vm.data.requestMethodList[key].value);
                }
            }
            if (tmpOutput.requestMethod.length > 0) {
                tmpOutput.requestMethod = tmpOutput.requestMethod.join(',');
            } else {
                vm.data.requestMethod = true;
            }
            switch ($state.params.status) {
                case 'edit': {
                    tmpOutput.apiID = vm.ajaxRequest.apiID
                    break;
                }
            }
            if (vm.ajaxResponse.apiInfo.targetMethod == '-1') {
                tmpOutput.isFollow = true;
                delete tmpOutput.targetMethod;
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
            var tmpAjaxRequest = privateFun.confirm(),
                tmpPromise = null;
            vm.data.submitted = true;
            if ($scope.ConfirmForm.$valid && !vm.data.requestMethod) {
                tmpPromise = privateFun.edit({
                    request: tmpAjaxRequest
                });
            } else {
                $rootScope.InfoModal('API编辑失败，请检查信息是否填写完整！', 'error');

            }
            return tmpPromise;
        }
        privateFun.edit = function (arg) {
            var tmpPromise = null;
            if ($state.params.status == 'edit') {
                tmpPromise = GatewayResource.Api.Edit(arg.request).$promise;
                tmpPromise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS: {
                            vm.fun.back();
                            $rootScope.InfoModal('API修改成功', 'success');
                            break;
                        }
                        case '190005': {
                            vm.data.requestMethod = true;
                            $scope.ConfirmForm.requestURL.$invalid = true;
                            break;
                        }
                    }
                })
            } else {
                tmpPromise = GatewayResource.Api.Add(arg.request).$promise;
                tmpPromise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS: {
                            vm.fun.back();
                            $rootScope.InfoModal('API添加成功', 'success');
                            break;
                        }
                        case '190005': {
                            vm.data.requestMethod = true;
                            $scope.ConfirmForm.requestURL.$invalid = true;
                            break;
                        }
                    }
                })
            }
            return tmpPromise;
        }
        privateFun.initApi = function () {
            var tmpAjaxRequest = {
                projectID: vm.ajaxRequest.projectID,
                apiID: vm.ajaxRequest.apiID
            }
            switch ($state.params.status) {
                case 'edit': {
                    GatewayResource.Api.Info(tmpAjaxRequest).$promise.then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                vm.ajaxResponse.apiInfo = response.apiInfo;
                                let tmpRequestMethod = response.apiInfo.requestMethod.split(',');
                                vm.ajaxResponse.apiInfo.targetMethod = response.apiInfo.isFollow ? '-1' : vm.ajaxResponse.apiInfo.targetMethod;
                                for (var key in tmpRequestMethod) {
                                    var val = tmpRequestMethod[key];
                                    switch (val.toLowerCase()) {
                                        case 'post': {
                                            vm.data.requestMethodList[0].checkbox = true;
                                            break;
                                        }
                                        case 'get': {
                                            vm.data.requestMethodList[1].checkbox = true;
                                            break;
                                        }
                                        case 'put': {
                                            vm.data.requestMethodList[2].checkbox = true;
                                            break;
                                        }
                                        case 'delete': {
                                            vm.data.requestMethodList[3].checkbox = true;
                                            break;
                                        }
                                        case 'head': {
                                            vm.data.requestMethodList[4].checkbox = true;
                                            break;
                                        }
                                        case 'options': {
                                            vm.data.requestMethodList[5].checkbox = true;
                                            break;
                                        }
                                        case 'batch': {
                                            vm.data.requestMethodList[6].checkbox = true;
                                            break;
                                        }
                                    }
                                }
                                break;
                            }
                        }
                    });
                    break;
                }
            }
        }
        privateFun.getDefaultBalance=()=>{
            GatewayResource.Balance.SimpleQuery().$promise.then((response)=>{
                vm.ajaxResponse.balanceList=response.balanceNames;
            })
        }
        vm.fun.init = (function () {
            let tmpStaticGroupArr=[{
                groupName:"未分组",
                groupID:-1
            }]
            var tmpCache = GroupService.get();
            if (tmpCache) {
                vm.data.apiGroup = tmpStaticGroupArr.concat(tmpCache);
                privateFun.initApi();
            } else {
                $rootScope.global.ajax.Query_Group.$promise.finally(function () {
                    vm.data.apiGroup = tmpStaticGroupArr.concat(GroupService.get());
                    privateFun.initApi();
                })
            }

        })()
        vm.$onInit = function () {
            privateFun.getDefaultBalance();
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