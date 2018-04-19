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

    indexController.$inject = ['$cookies', '$scope', 'CommonResource', 'md5', 'CODE','$rootScope'];

    function indexController($cookies, $scope, CommonResource, md5, CODE,$rootScope) {
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
                    mysqlHost: '',
                    mysqlDBName: '',
                    mysqlUserName: '',
                    mysqlPassword: '',
                    userName: '',
                    userPassword: '',
                    redisHost: '',
                    redisDB: '',
                    redisPassword: '',
                    gatewayPort: ''
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
        vm.data.fun.check.redis = function () {
            var template = {
                request: {
                    redisHost: vm.data.interaction.request.redisHost,
                    redisDB: vm.data.interaction.request.redisDB,
                    redisPassword: vm.data.interaction.request.redisPassword
                },
                promise: null
            }
            template.promise = CommonResource.Install.CheckRedis(template.request).$promise;
            template.promise.then(function (response) {
                vm.data.info.conn.redis = response.statusCode == CODE.COMMON.SUCCESS ? 1 : 0;
            })
            return template.promise;
        }
        vm.data.fun.check.database = function () {
            var template = {
                request: {
                    mysqlHost: vm.data.interaction.request.mysqlHost,
                    mysqlDBName: vm.data.interaction.request.mysqlDBName,
                    mysqlUserName: vm.data.interaction.request.mysqlUserName,
                    mysqlPassword: vm.data.interaction.request.mysqlPassword
                },
                promise: null
            }
            template.promise = CommonResource.Install.CheckDatabase(template.request).$promise;
            template.promise.then(function (response) {
                vm.data.info.conn.db = response.statusCode == CODE.COMMON.SUCCESS ? 1 : 0;
            })
            return template.promise;
        }
        vm.data.fun.confirm = function (status) {
            var template = {
                request: null,
                promise: null
            }
            switch (status) {
                case 'set-admin':
                    {
                        template.request = {
                            userName: vm.data.interaction.request.userName,
                            userPassword: md5.createHash(vm.data.interaction.request.userPassword)
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
                        break;
                    }
                case 'set-config':
                    {
                        template.request = {
                            redisHost: vm.data.interaction.request.redisHost,
                            redisDB: vm.data.interaction.request.redisDB,
                            redisPassword: vm.data.interaction.request.redisPassword,
                            mysqlHost: vm.data.interaction.request.mysqlHost,
                            mysqlDBName: vm.data.interaction.request.mysqlDBName,
                            mysqlUserName: vm.data.interaction.request.mysqlUserName,
                            mysqlPassword: vm.data.interaction.request.mysqlPassword,
                            gatewayPort: vm.data.interaction.request.gatewayPort
                        };
                        template.promise = CommonResource.Install.PostConfig(template.request).$promise;
                        template.promise.then(function (response) {
                            switch (response.statusCode) {
                                case CODE.COMMON.SUCCESS:
                                    {
                                        vm.data.info.status = 'set-second';
                                        break;
                                    }
                                default:
                                    {
                                        $rootScope.InfoModal('配置生成失败！','error');
                                        break;
                                    }
                            }
                        })
                        break;
                    }
            }


            $scope.$broadcast('$LoadingInit', {
                promise: template.promise
            });
            return template.promise;
        }
    }
})();