(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 所有弹窗模块定义js（依赖第三方bootstrap modal插件）
     */
    angular.module('eolinker.modal', ['ui.bootstrap.modal'])

        .directive('eoModal', [function () {
            return {
                restrict: 'AE',
                templateUrl: 'app/modal/index.html',
                controller: eoModalController
            }
        }])
    eoModalController.$inject = ['$scope', '$uibModal', '$rootScope']

    function eoModalController($scope, $uibModal, $rootScope) {
        //弹窗引用
        $rootScope.Gateway_CopyApiModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'Gateway_CopyApiModal',
                controller: 'Gateway_CopyApiModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.Gateway_GpeditApiPluginModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'Gateway_GpeditApiPluginModal',
                controller: 'Gateway_GpeditApiPluginModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.Gateway_ServiceModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'Gateway_ServiceModal',
                controller: 'Gateway_ServiceModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.MixInputModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'MixInputModal',
                controller: 'MixInputModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        /**
         * @description 接口网关模块定义
         */
        $rootScope.Gateway_NodeCheckErrorReportModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'Gateway_NodeCheckErrorReportModal',
                controller: 'Gateway_NodeCheckErrorReportModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.Gateway_ChangePasswordModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'Gateway_ChangePasswordModal',
                controller: 'Gateway_ChangePasswordModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.GatewayGpeditDefaultModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'GatewayGpeditDefaultModal',
                controller: 'GatewayGpeditDefaultModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }

        $rootScope.GatewayClusterModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'GatewayClusterModal',
                controller: 'GatewayClusterModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        //公用相关定义
        $rootScope.ExportModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'ExportModal',
                controller: 'ExportModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.CommonChangePasswordModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'CommonChangePasswordModal',
                controller: 'CommonChangePasswordModalCtrl',
                resolve: {
                    input: function() {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.ImportModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'ImportModal',
                controller: 'ImportModalCtrl',
                resolve: {
                    input: function() {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.Common_SingleInputModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'Common_SingleInputModal',
                controller: 'Common_SingleInputModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.Common_LoginModal = function openModal(callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'Common_LoginModal',
                controller: 'Common_LoginModalCtrl'
            });
            modalInstance.result.then(callback);
            return modalInstance;
        }
        $rootScope.RequestParamDetailModal = function openModal(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'RequestParamDetailModal',
                controller: 'RequestParamDetailModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.RequestParamEditModal = function openModal(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'RequestParamEditModal',
                controller: 'RequestParamEditModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.ResponseParamEditModal = function openModal(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'ResponseParamEditModal',
                controller: 'ResponseParamEditModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.ResponseParamDetailModal = function openModal(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'ResponseParamDetailModal',
                controller: 'ResponseParamDetailModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.ExpressionBuilderModal = function openModel(data, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'ExpressionBuilderModal',
                controller: 'ExpressionBuilderModalCtrl',
                resolve: {
                    data: function () {
                        return data;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.InfoModal = function openModel(info, type, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'InfoModal',
                controller: 'InfoModalCtrl',
                displayClass: 'modal-info-display',
                resolve: {
                    info: function () {
                        return info;
                    },
                    type: function () {
                        return type;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.EnsureModal = function openModel(title, necessity, info, input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'EnsureModal',
                controller: 'EnsureModalCtrl',
                resolve: {
                    title: function () {
                        return title;
                    },
                    necessity: function () {
                        return necessity;
                    },
                    info: function () {
                        return info;
                    },
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.GroupModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'GroupModal',
                controller: 'GroupModalCtrl',
                resolve: {
                    input: function() {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }
        $rootScope.SelectVisualGroupModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'SelectVisualGroupModal',
                controller: 'SelectVisualGroupModalCtrl',
                resolve: {
                    input: function () {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }

        $rootScope.SingleSelectModal = function openModel(input, callback) {
            var modalInstance = $uibModal.open({
                animation: true,
                templateUrl: 'SingleSelectModal',
                controller: 'SingleSelectModalCtrl',
                resolve: {
                    input: function() {
                        return input;
                    }
                }
            });
            modalInstance.result.then(callback);
        }





























































    }

})();