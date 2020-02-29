(function() {
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('index', {
                    url: '/',
                    auth: true,
                    template: '<login></login>'
                });
        }])
        .component('login', {
            templateUrl: 'app/ui/content/login/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope', 'CommonResource', '$state', 'CODE', 'md5', '$rootScope'];

    function indexController($scope, CommonResource, $state, CODE, md5, $rootScope) {

        var vm = this;
        vm.data = {
            submitted: false,
                password: {
                    isShow: false
                }
        }
        vm.ajaxRequest={
            loginCall: '',
            loginPassword: '',
            verifyCode: ''
        }
        vm.fun={};
        vm.fun.confirm = function() {
            var tmpAjaxRequest={
                loginCall: vm.ajaxRequest.loginCall,
                loginPassword: md5.createHash(vm.ajaxRequest.loginPassword||''),
            }
            if(tmpAjaxRequest.loginCall&&tmpAjaxRequest.loginPassword){
                CommonResource.Guest.Login(tmpAjaxRequest).$promise.then(function(response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                window.localStorage.setItem('LOGINCALL', tmpAjaxRequest.loginCall);
                                $state.go('home.panel');
                                break;
                            }
                    }
                })
            }else{
                vm.data.submitted=true;
            }
            
        }
        vm.fun.changeView = function() {
            if (vm.ajaxRequest.loginPassword) {
                vm.data.password.isShow = !vm.data.password.isShow;
            }
        }
    }
})();