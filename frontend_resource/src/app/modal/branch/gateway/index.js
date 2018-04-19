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

        .controller('Gateway_ChangePasswordModalCtrl', Gateway_ChangePasswordModalCtrl)

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
                CommonResource.User.Password(template.request).$promise
                    .then(function (response) {
                        switch (response.statusCode) {
                            case CODE.COMMON.SUCCESS:
                            case CODE.USER.UNCHANGE:
                                {
                                    $rootScope.InfoModal('修改成功', 'success');
                                    $uibModalInstance.close(true);
                                    break;
                                }
                            case CODE.USER.ERROR:
                                {
                                    $rootScope.InfoModal('旧密码错误', 'error');
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

    Gateway_DefaultModalCtrl.$inject = ['$scope', '$uibModalInstance', 'GatewayResource', 'CODE', 'input'];

    function Gateway_DefaultModalCtrl($scope, $uibModalInstance, GatewayResource, CODE, input) {
        $scope.interaction = {
            request: {
                gatewayName: '',
                gatewayDesc: ''
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

        /**
         * 初始化信息
         */
        fun.init = (function () {
            switch (input.status) {
                case 'edit':
                    {
                        $scope.interaction.request = angular.copy(input.request);
                        break;
                    }
            }
        })()

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
                                            $scope.interaction.request.gatewayHashKey = response.gatewayHashKey;
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
            backendURI: '',
            isAdd: true
        }

        function init() {
            if (info) {
                $scope.info = {
                    backendID: info.backendID,
                    backendName: info.backendName,
                    backendURI: info.backendURI,
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

    GatewayRateLimitModalCtrl.$inject = ['$scope', '$uibModalInstance', '$timeout', 'CODE', 'title', 'info', 'secondTitle', 'query'];

    function GatewayRateLimitModalCtrl($scope, $uibModalInstance, $timeout, CODE, title, info, secondTitle, query) {
        var vm = this;
        $scope.title = title;
        $scope.secondTitle = secondTitle || '分组';
        $scope.required = info ? (info.required ? true : false) : false;
        $scope.info = {
            groupName: '',
            groupID: '',
            $index: '0',
            isAdd: true
        }
        $scope.params = {
            query: query,
            hadSelected: query ? true : false
        }

        function init() {
            if (info) {
                $scope.info = {
                    groupName: info.groupName,
                    groupID: info.groupID,
                    $index: info.$index ? '' + info.$index : '0',
                    isAdd: false
                }
            }
        }
        init();
        $scope.ok = function () {
            if ($scope.editGroupForm.$valid) {
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


})();