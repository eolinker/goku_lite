(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 网关设置，基础配置相关js
     */
    angular.module('eolinker')
        .component('settingBasic', {
            templateUrl: 'app/ui/content/setting/basic/index.html',
            controller: indexController
        })

    indexController.$inject = ['$rootScope', 'CODE', 'GatewayResource','$scope','Authority_CommonService'];

    function indexController($rootScope, CODE, GatewayResource,$scope,Authority_CommonService) {
        var vm = this;
        vm.data = {
            status: {
                config: {
                    isEdit: false
                }
            },
            userEmail: [{
                value: ''
            }]
        }
        vm.interaction = {
            response: {
                gatewayConfig: {}
            }
        }
        vm.fun = {};
        vm.constant = {
            monitorUpdatePeriodQuery: [{
                key: '30秒',
                value: 30
            }, {
                key: '60秒',
                value: 60
            }, {
                key: '180秒',
                value: 180
            }]
        }
        var data = {
            template: {
                gatewayConfig: {}
            }
        }
        vm.service={
            authority:Authority_CommonService
        }
        vm.fun.editBasicInfo = function () {
            vm.data.status.config.submitted = true;
            if ($scope.BasicForm.$invalid) return;
            var template = {
                request: {
                    successCode: vm.interaction.response.gatewayConfig.successCode,
                    monitorUpdatePeriod: vm.interaction.response.gatewayConfig.monitorUpdatePeriod,
                    nodeUpdatePeriod: vm.interaction.response.gatewayConfig.nodeUpdatePeriod
                }
            }
            GatewayResource.Config.BaseEdit(template.request).$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            $rootScope.InfoModal('修改成功！', 'success');
                            vm.data.status.config.isEdit = false;
                            angular.copy(vm.interaction.response.gatewayConfig, data.template.gatewayConfig);
                            break;
                        }
                }
            })
        }
        vm.fun.cancleBasicInfo = function () {
            vm.data.status.config.isEdit = false;
            vm.interaction.response.gatewayConfig = angular.copy(data.template.gatewayConfig);
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['基本设置','网关设置']
            });
            GatewayResource.Config.BaseInfo().$promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.interaction.response.gatewayConfig = {
                                successCode: response.gatewayConfig.successCode,
                                monitorUpdatePeriod: response.gatewayConfig.monitorUpdatePeriod,
                                nodeUpdatePeriod: response.gatewayConfig.nodeUpdatePeriod
                            }
                            angular.copy(vm.interaction.response.gatewayConfig, data.template.gatewayConfig);
                            break;
                        }
                }
            })
        };
    }
})();