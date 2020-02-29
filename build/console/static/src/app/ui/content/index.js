(function () {
    //author：广州银云信息科技有限公司
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function ($stateProvider, helper) {
            $stateProvider
                .state('home', {
                    url: '/home',
                    template: '<home class="home-div"></home>',
                    containerRouter:true,
                    resolve: helper.resolveFor( 'DATEPICKER')
                })
        }])
        .component('home', {
            templateUrl: 'app/ui/content/index.html',
            controller: indexController,
        })
    indexController.$inject = ['$scope','GroupService', '$state','Sidebar_CommonService'];

    function indexController($scope,GroupService, $state,Sidebar_CommonService) {
        var vm = this;
        vm.data = {
            info: {
                shrinkObject: {},
                sidebarShow: null
            },
            fun: {
                $Home_ShrinkSidebar: null,
                init: null //初始化功能函数
            }
        }
        vm.service={
            group:GroupService,
            sidebar: Sidebar_CommonService
        }
        vm.data.fun.init = function (arg) {
            if (!/inside/.test($state.current.name)) {
                vm.data.info.sidebarShow = true;
            } else {
                vm.data.info.sidebarShow = false;
            }
        }
        vm.data.fun.init({
            key: window.location.href
        });
        $scope.$on('$locationChangeSuccess', function () {
            if (!/inside/.test($state.current.name.toLowerCase())) {
                vm.data.info.sidebarShow = true;
            } else {
                vm.data.info.sidebarShow = false;
            }
        })
    }
})();