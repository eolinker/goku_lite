(function () {
    'use strict';
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('arrange', {
                    url: '/arrange',
                    auth: true,
                    template: '<arrange></arrange>'
                });
        }])
        .component('arrange', {
            templateUrl: 'app/ui/content/arrange/index.html',
            controller: indexController
        })

    indexController.$inject = ['$cookies', '$scope', 'CommonResource','CODE','$rootScope'];

    function indexController($cookies, $scope, CommonResource, CODE,$rootScope) {
        var vm = this;
        vm.data = {
            info: {
                conn: {
                    db: null,
                    redis: null
                },
                status: 'set-first', //安装状态 'set'需安装，'failure'安装失败 'success' 安装成功
            },
            interaction: {
                request: {
                    userName: '',
                    userPassword: '',
                    gatewayConfPath: '',
                    port: ''
                }
            },
            fun: {
                check: {
                    env: null,
                    database: null
                },
                confirm: null,
            }
        }
        vm.data.fun.load = function (arg) {
            return arg.promise;
        }
        vm.data.fun.confirm = function () {
            var template = {
                request: null,
                promise: null
            }
            template.request = {
                userName: vm.data.interaction.request.userName,
                userPassword: CryptoJS.MD5(vm.data.interaction.request.userPassword).toString(),
                gatewayConfPath:vm.data.interaction.request.gatewayConfPath,
                port: vm.data.interaction.request.port
            };
            template.promise = CommonResource.Install.Post(template.request).$promise;
            template.promise.then(function (response) {
                switch (response.statusCode) {
                    case CODE.COMMON.SUCCESS:
                        {
                            vm.data.info.status = 'success';
                            break;
                        }
                    default:
                        {
                            vm.data.info.status = 'failure';
                            break;
                        }
                }
            })
            $scope.$broadcast('$LoadingInit', {
                promise: template.promise
            });
            return template.promise;
        }
    }
})();