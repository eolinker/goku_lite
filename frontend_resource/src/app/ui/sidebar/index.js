(function() {
    /*
     * author：广州银云信息科技有限公司
     * 全局sidebar指令相关js
     */
    angular.module('goku')
        .component('eoSidebar', {
            templateUrl: 'app/ui/sidebar/index.html',
            controller: indexController,
            bindings: {
                shrinkObject: '<'
            }
        })

    indexController.$inject = ['$scope', '$state','NavbarService'];

    function indexController($scope, $state, NavbarService) {

        var vm = this;
        vm.data = {
            component: {
                sidebarCommonObject: {}
            },
            service: {
                default:NavbarService
            },
            info: {
                current: null,
                menu: [{
                        name: '接口网关',
                        sref: 'home.gateway.default',
                        icon: 'icon-hexin',
                        power: -1
                    }
                ]
            },
            fun:{
                $Sidebar_ResetCurrent:null
            }
        }
        vm.data.fun.$Sidebar_ResetCurrent = function(_default) {
            vm.data.info.current = vm.data.info.menu[0];
            vm.data.service.default.info.navigation = {
                current: vm.data.info.menu[0].name
            }
            vm.data.service.pro.fun.init();
        }
        $scope.$on('$Sidebar_ResetCurrent', vm.data.fun.$Sidebar_ResetCurrent);
        vm.$onInit=function(){
            vm.shrinkObject.isShrink=false;
            vm.data.component.sidebarCommonObject = {
                mainObject: {
                    baseInfo: {
                        menu: vm.data.info.menu
                    }
                }
            }
        }
    }

})();