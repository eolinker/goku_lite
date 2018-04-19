(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 公用弹窗controller js
     */
    angular.module('goku.modal')

        .directive('eoCommonModal', [function () {
            return {
                restrict: 'AE',
                templateUrl: 'app/modal/branch/common/index.html'
            }
        }])

        .controller('Common_LoginModalCtrl', Common_LoginModalCtrl)

        .controller('InfoModalCtrl', InfoModalCtrl)

        .controller('EnsureModalCtrl', EnsureModalCtrl)

        .controller('GroupModalCtrl', GroupModalCtrl)

        .controller('TableModalCtrl', TableModalCtrl)

    Common_LoginModalCtrl.$inject = ['$scope', 'CODE', 'COOKIE_CONFIG', '$rootScope', '$cookies', 'md5', 'CommonResource', '$uibModalInstance'];

    function Common_LoginModalCtrl($scope, CODE, COOKIE_CONFIG, $rootScope, $cookies, md5, CommonResource, $uibModalInstance) {
        $scope.data = {
            info: {
                submitted: false
            },
            interaction: {
                request: {
                    loginCall: '',
                    loginPassword: '',
                    verifyCode: ''
                }
            },
            fun: {
                close: null, //关闭功能函数
                confirm: null, //确认功能函数
            }
        }
        var data = {
            fun: {
                init: null, //初始化功能函数
            }
        }
        $scope.data.fun.close = function () {
            $uibModalInstance.close(false);
        }
        $scope.data.fun.confirm = function () {
            var template = {
                request: {
                    loginCall: $scope.data.interaction.request.loginCall,
                    loginPassword: md5.createHash($scope.data.interaction.request.loginPassword),
                    verifyCode: md5.createHash((new Date()).toUTCString())
                }
            }

            if ($scope.confirmForm.$valid) {
                $scope.data.info.submitted = false;
                $cookies.put("verifyCode", template.request.verifyCode, COOKIE_CONFIG);
                $rootScope.global.ajax.Login_Guest = CommonResource.Guest.Login(template.request);
                $rootScope.global.ajax.Login_Guest.$promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                $uibModalInstance.close({
                                    loginCall: template.request.loginCall
                                });
                                break;
                            }
                        default:
                            {
                                $rootScope.InfoModal('登录失败,请检查密码是否正确！', 'error');
                                break;
                            }
                    }
                })
            } else {
                $scope.data.info.submitted = true;
            }
        }
    }

    EnsureModalCtrl.$inject = ['$scope', '$uibModalInstance', 'title', 'necessity', 'info', 'input'];

    function EnsureModalCtrl($scope, $uibModalInstance, title, necessity, info, input) {

        $scope.title = title;
        $scope.necessity = necessity;
        $scope.info = {
            message: info || '确认删除？',
            btnType: input.btnType || 0, //0：warning 1：info,2:success,
            btnMessage: input.btnMessage || '删除',
            btnGroup: input.btnGroup || []
        }
        var data = {
            fun: {
                init: null
            }
        }
        $scope.data = {
            input: {}
        }

        /**
         * 初始化
         */
        data.fun.init = (function () {
            angular.copy(input, $scope.data.input);
        })()

        $scope.ok = function () {
            if ($scope.sureForm.$valid || !$scope.necessity) {
                $uibModalInstance.close(true);
            } else {
                $scope.submited = true;
            }
        };

        $scope.cancel = function () {
            //$uibModalInstance.dismiss(false);
            $uibModalInstance.close(false);
        };

    }

    InfoModalCtrl.$inject = ['$scope', '$uibModalInstance', '$timeout', 'info', 'type'];

    function InfoModalCtrl($scope, $uibModalInstance, $timeout, info, type) {

        $scope.type = type || 'info';
        $scope.info = info;
        var timer = $timeout(function () {
            $uibModalInstance.close(true);
        }, 1500, true);
        $scope.$on('$destroy', function () {
            if (timer) {
                $timeout.cancel(timer);
            }
        });
    }
    

    GroupModalCtrl.$inject = ['$scope', '$uibModalInstance', '$timeout', 'CODE', 'title', 'info', 'secondTitle', 'query'];

    function GroupModalCtrl($scope, $uibModalInstance, $timeout, CODE, title, info, secondTitle, query) {
        var code = CODE.COMMON.SUCCESS;
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
            query: [{
                groupName: '--不设置父分组--',
                groupID: '0'
            }].concat(query),
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

    TableModalCtrl.$inject = ['$scope', '$uibModalInstance', '$timeout', 'CODE', 'title', 'info', 'databaseHashKey'];

    function TableModalCtrl($scope, $uibModalInstance, $timeout,CODE, title, info, databaseHashKey) {
        var code = CODE.COMMON.SUCCESS;
        var vm = this;
        $scope.title = title;
        $scope.info = {
            databaseHashKey: databaseHashKey,
            tableID: '',
            tableName: '',
            tableDesc: '',
            isAdd: true
        }

        function init() {
            if (info) {
                $scope.info = {
                    databaseHashKey: databaseHashKey,
                    tableID: info.tableID,
                    tableName: info.tableName,
                    tableDesc: info.tableDesc,
                    isAdd: false
                }
            }
        }
        init();
        $scope.ok = function () {
            if ($scope.editTableForm.$valid) {
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