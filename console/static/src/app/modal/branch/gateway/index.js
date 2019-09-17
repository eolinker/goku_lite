(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 专业版本专用弹窗controller js
     */
    angular.module('eolinker.modal')

        .directive('eoGatewayModal', [function () {
            return {
                restrict: 'AE',
                templateUrl: 'app/modal/branch/gateway/index.html'
            }
        }])
        .controller('Gateway_ServiceModalCtrl', Gateway_ServiceModalCtrl)
        .controller('Gateway_NodeCheckErrorReportModalCtrl', Gateway_NodeCheckErrorReportModalCtrl)

        .controller('GatewayClusterModalCtrl', GatewayClusterModalCtrl)
        .controller('GatewayGpeditDefaultModalCtrl', GatewayGpeditDefaultModalCtrl)
        .controller('Gateway_ChangePasswordModalCtrl', Gateway_ChangePasswordModalCtrl)
        .controller('Gateway_GpeditApiPluginModalCtrl', Gateway_GpeditApiPluginModalCtrl)
        .controller('Gateway_CopyApiModalCtrl', Gateway_CopyApiModalCtrl)
    Gateway_CopyApiModalCtrl.$inject = ['$rootScope','GatewayResource', '$scope','CODE', '$uibModalInstance', 'input'];

    function Gateway_CopyApiModalCtrl($rootScope,GatewayResource, $scope,CODE, $uibModalInstance, input) {
        $scope.fun = {};
        $scope.input = input;
        $scope.CONST = {
            PROTOCOL_ARR: [{
                key: 'HTTP',
                value: 'http'
            }, {
                key: 'HTTPS',
                value: 'https'
            }]
        }
        $scope.data={};
        $scope.component = {
            pluginOprObj: {}
        }
        let privateFun={};
        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
        privateFun.edit = function (arg) {
            var tmpPromise = null;
            tmpPromise = GatewayResource.Api.Copy(arg.request).$promise;
                tmpPromise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS: {
                            $uibModalInstance.close(true);
                            break;
                        }
                        case '190005': {
                            $scope.data.requestMethod = true;
                            $scope.ConfirmForm.requestURL.$invalid = true;
                            break;
                        }
                    }
                })
            return tmpPromise;
        }
        privateFun.confirm = function () {
            $scope.data.requestMethod = false;
            var tmpOutput = {
                projectID: $scope.input.projectID,
                groupID: $scope.input.groupID,
                apiName: $scope.input.apiName,
                requestMethod: [],
                requestURL: $scope.input.requestURL,
                balanceName: $scope.input.balanceName,
                targetURL: $scope.input.targetURL,
                targetMethod: $scope.input.targetMethod,
                protocol:$scope.input.protocol,
                apiID:$scope.input.apiID
            }
            for (var key in $scope.input.requestMethodList) {
                if ($scope.input.requestMethodList[key].checkbox) {
                    tmpOutput.requestMethod.push($scope.input.requestMethodList[key].value);
                }
            }
            if (tmpOutput.requestMethod.length > 0) {
                tmpOutput.requestMethod = tmpOutput.requestMethod.join(',');
            } else {
                $scope.data.requestMethod = true;
            }
            if ($scope.input.targetMethod == '-1') {
                tmpOutput.isFollow = true;
                delete tmpOutput.targetMethod;
            }
            return tmpOutput;
        }
        $scope.fun.confirm = function () {
            var tmpAjaxRequest = privateFun.confirm(),
                tmpPromise = null;
            $scope.data.submitted = true;
            if ($scope.ConfirmForm.$valid && !$scope.data.requestMethod) {
                tmpPromise = privateFun.edit({
                    request: tmpAjaxRequest
                });
            } else {
                $rootScope.InfoModal('API复制失败，请检查信息是否填写完整！', 'error');

            }
            return tmpPromise;
        }
    }
    Gateway_GpeditApiPluginModalCtrl.$inject = ['$rootScope', '$scope', '$uibModalInstance', 'input'];

    function Gateway_GpeditApiPluginModalCtrl($rootScope, $scope, $uibModalInstance, input) {
        $scope.fun = {};
        $scope.input = input;
        $scope.component = {
            pluginOprObj: {}
        }
        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
        $scope.fun.confirm = () => {
            $scope.component.pluginOprObj.trigger = "save";
        }
        $rootScope.global.$watch.push($scope.$watch('component.pluginOprObj.accept', () => {
            switch ($scope.component.pluginOprObj.accept) {
                case "saveSuccess": {
                    $uibModalInstance.close(true);
                    break;
                }
            }
        }, true));
    }
    Gateway_ServiceModalCtrl.$inject = ['$scope', 'CODE', '$rootScope', 'Cache_CommonService', 'GatewayResource', '$uibModalInstance', 'input'];

    function Gateway_ServiceModalCtrl($scope, CODE, $rootScope, Cache_CommonService, GatewayResource, $uibModalInstance, input) {
        $scope.input = {
            title: input.title,
            opr: input.opr
        };
        $scope.data = {};
        $scope.CONST = {
            SERVICE_TYPE_ARR: [{
                key: "静态服务",
                value: "static",
                tip: "IP地址在网关直接配置"
            }, {
                key: "服务发现",
                value: "discovery",
                tip: "IP地址在网关通过服务发现的方式配置"
            }]
        }
        $scope.ajaxResponse = angular.copy(input.ajaxResponse);
        $scope.component = {
            listBlockObj: {
                setting: {
                    munalAddRow: true
                },
                tdList: [{
                    type: 'text',
                    thKey: '范围',
                    modelKey: 'title',
                    class: 'w_150'
                }, {
                    type: 'input',
                    thKey: '接入地址',
                    modelKey: 'value'
                }]
            }
        }
        $scope.fun = {};
        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
        $scope.fun.confirm = () => {
            $scope.data.submitted = true;
            if ($scope.ConfirmForm.$invalid) {
                $rootScope.InfoModal('请检查信息是否填写完整', 'error');
                return;
            }
            let tmpAjaxRequest = {
                name: $scope.ajaxResponse.serviceData.name
            }
            if ($scope.ajaxResponse.serviceData.type === "static") {
                tmpAjaxRequest.driver = "static";
            } else {
                tmpAjaxRequest.driver = $scope.ajaxResponse.serviceData.driver;
                tmpAjaxRequest.config = $scope.ajaxResponse.serviceData.config,
                    tmpAjaxRequest.clusterConfig = {};
                $scope.ajaxResponse.clusterQuery.map((val) => {
                    tmpAjaxRequest.clusterConfig[val.name] = val.value;
                })
                tmpAjaxRequest.clusterConfig = JSON.stringify(tmpAjaxRequest.clusterConfig);
            }
            let tmpPromise = GatewayResource.ServiceDiscovery[input.opr](tmpAjaxRequest).$promise;
            tmpPromise.then((response) => {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS: {
                        $rootScope.InfoModal(`${input.title}成功`, 'success');
                        $uibModalInstance.close(true);
                        break;
                    }
                }
            })
            return tmpPromise;
        }
        let privateFun = {},
            service = {
                cache: Cache_CommonService
            };
        /**
         * @desc 请求获取驱动类型列表
         */
        privateFun.ajaxDriverQuery = () => {
            let tmpCache = service.cache.get('EO_GOKU_DRIVER_LIST');
            if (tmpCache) {
                $scope.ajaxResponse.driverQuery = tmpCache;
            } else {
                GatewayResource.ServiceDiscovery.DriverQuery().$promise.then((response) => {
                    $scope.ajaxResponse.driverQuery = response.data || [];
                    service.cache.set($scope.ajaxResponse.driverQuery, 'EO_GOKU_DRIVER_LIST');
                })
            }
        }

        function main() {
            privateFun.ajaxDriverQuery();
        }
        main();
    }

    Gateway_NodeCheckErrorReportModalCtrl.$inject = ['$scope', '$uibModalInstance', 'input'];

    function Gateway_NodeCheckErrorReportModalCtrl($scope, $uibModalInstance, input) {
        $scope.input = input;
        $scope.fun = {};
        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
    }

    Gateway_ChangePasswordModalCtrl.$inject = ['md5', '$scope', '$uibModalInstance', '$rootScope', 'CommonResource', 'CODE', 'input'];

    function Gateway_ChangePasswordModalCtrl(md5, $scope, $uibModalInstance, $rootScope, CommonResource, CODE, input) {
        $scope.data = {
            input: {},
            interaction: {
                request: {
                    oldPassword: '',
                    newPassword: ''
                }
            },
            fun: {
                confirm: null, //确认修改功能函数
            }
        }
        $scope.data.fun.confirm = function () {
            var template = {
                request: {
                    oldPassword: md5.createHash($scope.data.interaction.request.oldPassword),
                    newPassword: md5.createHash($scope.data.interaction.request.newPassword)
                }
            }
            if ($scope.editForm.$valid) {
                CommonResource.User.ChangePassword(template.request).$promise
                    .then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS: {
                                $rootScope.InfoModal('修改成功', 'success');
                                $uibModalInstance.close(true);
                                break;
                            }
                        }
                    })
            }
        }
        $scope.cancel = function () {
            $uibModalInstance.close(false);
        };
    }
    GatewayGpeditDefaultModalCtrl.$inject = ['$scope', '$uibModalInstance', 'input'];

    function GatewayGpeditDefaultModalCtrl($scope, $uibModalInstance, input) {
        $scope.data = {
            title: input.title,
            group: input.group
        };
        $scope.output = angular.copy(input.item);
        $scope.fun = {};
        $scope.fun.confirm = function () {
            if ($scope.ConfirmForm.$valid) {
                $uibModalInstance.close($scope.output);
            }
        };

        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
    }

    GatewayClusterModalCtrl.$inject = ['$scope', '$uibModalInstance', 'input'];

    function GatewayClusterModalCtrl($scope, $uibModalInstance, input) {
        $scope.data = {
            title: input.title,
            group: input.group,
            checkStatus: 'init'
        };
        $scope.output = angular.copy(input.item);
        $scope.fun = {};
        $scope.fun.confirm = function () {
            if ($scope.ConfirmForm.$valid) {
                $uibModalInstance.close($scope.output);
            }
        };
        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
    }


})();