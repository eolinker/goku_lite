(function () {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home.panel', {
                    url: '/',
                    template: '<panel></panel>'
                })
        }])
        .component('panel', {
            templateUrl: 'app/ui/content/panel/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'GatewayResource', '$state', 'CODE', '$rootScope', 'uibDateParser'];

    function indexController($scope, GatewayResource, $state, CODE, $rootScope, uibDateParser) {

        var vm = this,
            privateFun = {}
        vm.data = {}
        vm.fun = {};
        vm.ajaxRequest = {
            table: {
                beginTime: null,
                endTime: null,
                period: 0
            }
        }
        vm.ajaxResponse = {
            monitorInfo: null,
            redisArr: []
        };
        vm.directive = {
            tableTimeObject: {
                show: false,
                maxDate: new Date(),
                maxMode: 'month',
                request: {}
            }
        }
        vm.component = {
            overviewObject: {},
            listDefaultCommonObject: null
        }
        privateFun.initTable = function () {
            let tmpPromise;
            tmpPromise = GatewayResource.Monitor.Info().$promise;
            tmpPromise.then(function (response) {
                vm.ajaxResponse.monitorInfo = response || {};
            })
            return tmpPromise;
        }
        privateFun.refresh = function () {
            var tmpPromise = GatewayResource.Monitor.Refresh().$promise;
            tmpPromise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS: {
                        $rootScope.InfoModal('立即刷新成功!', 'success', function () {
                            $scope.$emit('$TransferStation', {
                                state: '$Init_LoadingCommonComponent'
                            });
                        });
                        break;
                    }
                }
            })
            return tmpPromise;
        }
        privateFun.initComponent = function () {
            vm.component.overviewObject = {
                setting: {
                    title: '基本信息',
                    showOperate: true
                }
            }
        }
        vm.fun.init = function (arg) {
            arg = arg || {
                type: 'default'
            }
            switch (arg.type) {
                default: {
                    return privateFun.initTable();;
                }
                case 'refresh': {
                    return privateFun.refresh();
                }
            }

        }
        vm.fun.refresh = function () {
            $scope.$emit('$TransferStation', {
                state: '$Init_LoadingCommonComponent',
                data: {
                    type: 'refresh',
                    tips: '刷新'
                }
            });
        }
        vm.$onInit = function () {
            $scope.$emit('$WindowTitleSet', {
                list: ['首页']
            });
            privateFun.initComponent();
        }
    }
})();