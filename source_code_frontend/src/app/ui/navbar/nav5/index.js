(function() {
    /*
     * author：广州银云信息科技有限公司
     *home/project
     * navbar5指令相关js
     */
    angular.module('goku')
        .component('eoNavbar5', {
            templateUrl: 'app/ui/navbar/nav5/index.html',
            controller: navbar
        })

    navbar.$inject = ['$scope', '$rootScope', 'CommonResource', 'NavbarService', '$state'];

    function navbar($scope, $rootScope, CommonResource, NavbarService, $state) {

        var vm = this;
        vm.data = {
            service: {
                navbar: NavbarService
            },
            interaction: {
                response: {
                    companyList: null
                }
            },
            fun: {
                logout: null, //退出登录功能函数
            },
            assistantFun: {
                init: null, //辅助初始化功能函数
            }
        }
        vm.data.fun.logout = function() {
            vm.data.service.navbar.fun.logout();
        }
    }

})();
