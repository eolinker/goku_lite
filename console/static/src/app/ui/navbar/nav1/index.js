(function () {
    /*
     * author：广州银云信息科技有限公司
     *home/project
     * navbar指令相关js
     */
    angular.module('eolinker')
        .component('eoNavbar1', {
            templateUrl: 'app/ui/navbar/nav1/index.html',
            controller: indexController
        })
    indexController.$inject = ['CODE', '$state', 'CommonResource','Sidebar_CommonService', 'Authority_CommonService','$rootScope','NavbarService'];

    function indexController(CODE, $state, CommonResource,Sidebar_CommonService, Authority_CommonService,$rootScope,NavbarService) {

        var vm = this,
            fun = {},
            service = {
                authority: Authority_CommonService
            };
            vm.service = {
                navbar: NavbarService,
                sidebar:Sidebar_CommonService
            }
        vm.data = {
            showLogout: /home/.test($state.current.name),
            userInfo:null
        }
        vm.fun = {};
        vm.fun.changePassword = function() {
            $rootScope.Gateway_ChangePasswordModal();
        }
        vm.fun.logout = function () {
            CommonResource.User.LoginOut()
                .$promise.then(function (response) {
                    switch (response.statusCode) {
                        case CODE.COMMON.SUCCESS:
                            {
                                vm.service.navbar.info.navigation={};
                                $state.go('index');
                                break;
                            }
                    }
                })
        }
        fun.init = ((function () {
            if(vm.data.showLogout){
                $rootScope.global.ajax.Info_User=CommonResource.User.Info();
                $rootScope.global.ajax.Info_User.$promise.then(function (response) {
                    vm.service.navbar.userInfo=vm.data.userInfo=response.userInfo;
                    service.authority.fun.setPermission('default', response.userInfo||{userType:0});
                })
            }
        }))();
    }
})();