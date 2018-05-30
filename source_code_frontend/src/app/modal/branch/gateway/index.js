(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 专业版本专用弹窗controller js
     */
    angular.module('goku.modal')

        .directive('eoGatewayModal', [function () {
            return {
                restrict: 'AE',
                templateUrl: 'app/modal/branch/gateway/index.html'
            }
        }])

        .controller('Gateway_DefaultModalCtrl', Gateway_DefaultModalCtrl)

        .controller('GatewayBackendModalCtrl', GatewayBackendModalCtrl)

        .controller('GatewayRateLimitModalCtrl', GatewayRateLimitModalCtrl)

    Gateway_DefaultModalCtrl.$inject = ['$scope', '$uibModalInstance', '$filter', 'GatewayResource', 'CODE', 'input'];

    function Gateway_DefaultModalCtrl($scope, $uibModalInstance, $filter, GatewayResource, CODE, input) {
        $scope.interaction = {
            request: {
                gatewayName: ''
            }
        }
        $scope.data = {
            title: input.title,
            status: input.status
        }
        $scope.fun = {};
        var fun = {
            init: null
        }
        var data = {
            aliasAjax: null
        }

        /**
         * 初始化信息
         */
        fun.init = (function () {
            switch (input.status) {
                case 'edit':
                    {
                        $scope.interaction.request = angular.copy(input.request);
                        $scope.interaction.request.oldGatewayAlias = input.request.gatewayAlias;
                        break;
                    }
            }
            $scope.$watch('interaction.request.gatewayAlias', function () {
                fun.checkAlias();
            }, true);
        })()

        $scope.fun.random = function () {
            $scope.interaction.request.gatewayAlias = $filter('tokenFilter')($scope.interaction.request.gatewayName);
        }
        fun.checkAlias = function () {
            if (!$scope.interaction.request.gatewayAlias) return;
            var template = {
                request: {
                    gatewayAlias: $scope.interaction.request.gatewayAlias
                }
            }
            if (data.aliasAjax) {
                data.aliasAjax.$cancelRequest();
            }
            data.aliasAjax = GatewayResource.Gateway.CheckAlias(template.request);
            data.aliasAjax.$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            $scope.ConfirmForm.alias.$invalid = true;
                            break;
                        }
                }
            })
        }
        /**
         * 确认保存网关
         */
        $scope.fun.confirm = function () {
            if ($scope.ConfirmForm.$valid) {
                switch (input.status) {
                    case 'add':
                        {
                            GatewayResource.Gateway.Add($scope.interaction.request).$promise.then(function (response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS:
                                        {
                                            $uibModalInstance.close($scope.interaction.request);
                                            break;
                                        }
                                    default:
                                        {
                                            $scope.submited = true;
                                            break;
                                        }
                                }
                            });
                            break;
                        }
                    default:
                        {
                            GatewayResource.Gateway.Edit($scope.interaction.request).$promise.then(function (response) {
                                switch (response.statusCode) {
                                    case CODE.COMMON.SUCCESS:
                                        {
                                            $uibModalInstance.close($scope.interaction.request);
                                            break;
                                        }
                                    default:
                                        {
                                            $scope.submited = true;
                                            break;
                                        }
                                }
                            });
                            break;
                        }
                }
            } else {
                $scope.submited = true;
            }
        };

        /**
         * 取消编辑
         */
        $scope.fun.cancel = function () {
            $uibModalInstance.close(false);
        };
    }

    GatewayBackendModalCtrl.$inject = ['$scope', '$uibModalInstance', '$timeout', 'CODE', 'title', 'info'];

    function GatewayBackendModalCtrl($scope, $uibModalInstance, $timeout, CODE, title, info) {
        var code = CODE.COMMON.SUCCESS;
        var vm = this;
        $scope.title = title;
        $scope.info = {
            backendName: '',
            backendPath: '',
            isAdd: true
        }

        function init() {
            if (info) {
                $scope.info = {
                    backendID: info.backendID,
                    backendName: info.backendName,
                    backendPath: info.backendPath,
                    isAdd: false
                }
            }
        }
        init();
        $scope.ok = function () {
            if ($scope.ConfirmForm.$valid) {
                $uibModalInstance.close($scope.info);
            } else {
                $scope.submited = true;
            }
        };

        $scope.cancel = function () {
            //$uibModalInstance.dismiss(false);
            $uibModalInstance.close(false);
        };
    }

    GatewayRateLimitModalCtrl.$inject = ['$scope', '$uibModalInstance', 'input'];

    function GatewayRateLimitModalCtrl($scope, $uibModalInstance, input) {
        var vm = this;
        $scope.data = {
            constant: {
                viewArray: [{
                    name: '允许访问',
                    id: true
                }, {
                    name: '禁止访问',
                    id: false
                }],
                intervalArray: [{
                    name: '1秒',
                    id: 'sec'
                }, {
                    name: '1分钟',
                    id: 'min'
                }, {
                    name: '1小时',
                    id: 'hour'
                }, {
                    name: '1天',
                    id: 'day'
                }],
                timeArray: [{
                    id: 0
                }, {
                    id: 1
                }, {
                    id: 2
                }, {
                    id: 3
                }, {
                    id: 4
                }, {
                    id: 5
                }, {
                    id: 6
                }, {
                    id: 7
                }, {
                    id: 8
                }, {
                    id: 9
                }, {
                    id: 10
                }, {
                    id: 11
                }, {
                    id: 12
                }, {
                    id: 13
                }, {
                    id: 14
                }, {
                    id: 15
                }, {
                    id: 16
                }, {
                    id: 17
                }, {
                    id: 18
                }, {
                    id: 19
                }, {
                    id: 20
                }, {
                    id: 21
                }, {
                    id: 22
                }, {
                    id: 23
                }, {
                    id: 24
                }]
            }
        };
        $scope.input = angular.copy(input);
        $scope.fun = {};
        $scope.fun.confirm = function () {
            if ($scope.ConfirmForm.$valid) {
                if (!$scope.input.request.allow) {
                    try {
                        delete $scope.input.request['period'];
                        delete $scope.input.request['limit'];
                    } catch (e) {
                        console.error(e)
                    }
                }
                $uibModalInstance.close($scope.input.request);
            }
        };

        $scope.fun.cancel = function () {
            //$uibModalInstance.dismiss(false);
            $uibModalInstance.close(false);
        };
    }


})();