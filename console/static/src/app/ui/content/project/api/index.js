(function () {
    'use strict';
    angular.module('eolinker')
        .component('api', {
            templateUrl: 'app/ui/content/project/api/index.html',
            controller: indexController
        })

    indexController.$inject = ['NavbarService','Cache_CommonService','$state','$scope'];

    function indexController( NavbarService,Cache_CommonService,$state,$scope) {
        var vm=this;
        var service={
            cache:Cache_CommonService,
            navbar: NavbarService
        }
        vm.data={};
        vm.component={
            menuObject:null
        }
        let privateFun={};
        privateFun.setNavShow=()=>{
            if ($state.current.name.indexOf('api.default') > -1) {
                vm.data.navShow = true;
            } else {
                vm.data.navShow = false;
            }
        }
        $scope.$on('$stateChangeSuccess', privateFun.setNavShow)

        vm.$onInit=()=>{
            service.cache.clear('apiGroup');
            privateFun.setNavShow();
            service.navbar.info.navigation.extra = $state.params.projectName;
            vm.component.menuObject={
                list:[{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '返回项目列表',
                        icon: 'chexiao',
                        fun: {
                            default: ()=>{
                                $state.go("home.project.default");
                            }
                        }
                    }]
                }],
                setting:{
                    class:'common-menu-fixed-seperate'
                }
            };
        }
        
    }
})();