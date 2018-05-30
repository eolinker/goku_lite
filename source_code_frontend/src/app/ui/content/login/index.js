(function() {
    'use strict';
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('index', {
                    url: '/?redirectUri',
                    auth: true,
                    template: '<login></login>'
                });
        }])
        .component('login', {
            templateUrl: 'app/ui/content/login/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'CommonResource', '$state', 'CODE','$rootScope', 'NavbarService', '$cookies'];

    function indexController($scope, CommonResource, $state, CODE, $rootScope, NavbarService, $cookies) {
        var vm = this;
        vm.data = {
            service: {
                nav: NavbarService
            },
            info: {
                submitted: false,
                password: {
                    isShow: false
                },
                isRemember: false
            },
            interaction: {
                request: {
                    loginName: '',
                    loginPassword: '',
                }
            },
            fun: {
                enterConsole: null, //进入控制台功能函数
                confirm: null, //确认登录功能函数
                changeView: null //密码是否显示功能函数
            }
        }

        vm.data.fun.enterConsole = function(arg) {
            var template = {
                storage: JSON.parse(window.localStorage['VERSIONINFO'] || '{}')
            }
            $state.go('home.gateway.default');
        }
        vm.data.fun.confirm = function() {
            var template = {
                href: decodeURIComponent(window.location.href).split("redirectUri=")[1],
                time: new Date(),
                cookieConfig: {
                    path: '/',
                    domain: window.location.hostname
                },
                request: {
                    loginName: vm.data.interaction.request.loginName,
                    loginPassword: CryptoJS.MD5(vm.data.interaction.request.loginPassword).toString()
                }
            }
            if ($scope.loginForm.$valid) {
                vm.data.info.submitted = false;
                CommonResource.Guest.Login(template.request).$promise.then(function(response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.data.fun.enterConsole({ loginName: template.request.loginName });
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
                vm.data.info.submitted = true;
            }
        }
        vm.data.fun.changeView = function() {
            if (vm.data.interaction.request.loginPassword) {
                vm.data.info.password.isShow = !vm.data.info.password.isShow;
            }
        }
        vm.$onInit = function() {
            vm.data.service.nav.fun.$router();
        }
    }
})();
