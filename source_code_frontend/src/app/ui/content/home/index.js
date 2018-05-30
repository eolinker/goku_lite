(function() {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('goku')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('home', {
                    url: '/home',
                    template: '<home></home>'
                });
        }])
        .component('home', {
            templateUrl: 'app/ui/content/home/index.html',
            controller: indexController,
        })
    indexController.$inject = ['$scope', '$state', 'HTML_LAZYLOAD'];

    function indexController($scope, $state, HTML_LAZYLOAD) {
        var vm = this;
        vm.data = {
            constant: {
                lazyload: HTML_LAZYLOAD[0]
            },
            info: {
                shrinkObject: {},
                sidebarShow: null
            },
            fun: {
                $Home_ShrinkSidebar: null,
                init: null //初始化功能函数
            }
        }
        vm.data.fun.init = function(arg) {
            if (!/inside/.test(arg.key.toLowerCase())) {
                vm.data.info.sidebarShow = true;
            } else {
                vm.data.info.sidebarShow = false;
            }
        }
        vm.data.fun.$Home_ShrinkSidebar = function(_default, arg) {
            vm.data.info.shrinkObject.isShrink = arg.shrink;
        }
        vm.data.fun.init({ key: window.location.href });
        $scope.$on('$locationChangeSuccess', function() {
            if (!/inside/.test($state.current.name.toLowerCase())) {
                vm.data.info.sidebarShow = true;
            } else {
                vm.data.info.sidebarShow = false;
            }
        })
        $scope.$on('$Home_ShrinkSidebar', vm.data.fun.$Home_ShrinkSidebar);
    }
})();
