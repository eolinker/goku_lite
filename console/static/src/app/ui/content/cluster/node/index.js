(function () {
    /*
     * author：广州银云信息科技有限公司
     */
    angular.module('eolinker')
        .component('clusterNode', {
            templateUrl: 'app/ui/content/cluster/node/index.html',
            controller: indexController
        })

    indexController.$inject = ['NavbarService','$state'];

    function indexController( NavbarService,$state) {
        var vm=this;
        var service={
            navbar: NavbarService
        }
        vm.component = {
            menuObject: {}
        }
        vm.$onInit=()=>{
            service.navbar.info.navigation.extra = $state.params.clusterName;
            vm.component.menuObject={
                list:[{
                    type: 'btn',
                    class: 'btn-group-li pull-left',
                    btnList: [{
                        name: '返回集群列表',
                        icon: 'chexiao',
                        fun: {
                            default: ()=>{
                                $state.go("home.cluster.default");
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