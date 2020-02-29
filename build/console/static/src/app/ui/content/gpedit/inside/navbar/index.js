(function () {
    //author：广州银云信息科技有限公司
    angular.module('eolinker')
        .component('gpeditNavbar', {
            templateUrl: 'app/ui/content/gpedit/inside/navbar/index.html',
            bindings: {
                shrinkObject: '<'
            },
            controller: indexController
        })

    indexController.$inject = ['$scope', '$state'];

    function indexController($scope, $state) {
        var vm = this;
        vm.data = {
            component: {
                sidebarCommonObject: {}
            },
            info: {
                menu: [{
                        base: '/api/',
                        name: 'API列表',
                        sref: 'home.gpedit.inside.api',
                        childSref:'home.gpedit.inside.api.default',
                        icon: 'icon-api',
                        power: -1
                    },
                    {
                        base: '/plugin/',
                        name: '策略插件',
                        sref: 'home.gpedit.inside.plugin',
                        childSref:'home.gpedit.inside.plugin.gpedit',
                        icon: 'icon-cengji_o',
                        power: -1
                    },
                    {
                        base: '/auth',
                        name: '鉴权方式',
                        sref: 'home.gpedit.inside.auth',
                        icon: 'icon-dunpaibaoxianrenzheng_o',
                        power: -1
                    },
                    {
                        base: '/setting',
                        name: '策略管理',
                        sref: 'home.gpedit.inside.setting',
                        icon: 'icon-quanjushezhi_o',
                        power: -1
                    }
                ]
            }
        }
        var fun={};
        fun.shrink = function () {
            $scope.$emit('$Home_ShrinkSidebar', {
                shrink: vm.shrinkObject.isShrink
            });
        }
        vm.$onInit = function() {
            let tmpRouterObj={
                common:"home.gpedit.common.list",
                open:"home.gpedit.default"
            }
            vm.data.component.sidebarCommonObject = {
                mainObject: {
                    baseInfo: {
                        staticTop:true,
                        menu: vm.data.info.menu,
                        navigation: [{
                            name: '访问策略',
                            sref: "home.gpedit.default"
                        }, {
                            name: $state.params.strategyName
                        }],
                        staticQuery: [{
                            name: '返回',
                            sref: tmpRouterObj[$state.params.groupType],
                            icon: 'icon-chexiao',
                            params:{
                                'groupID':$state.params.groupID
                            }
                        }]
                    },
                    baseFun: {
                        shrink: fun.shrink
                    }
                }
            }
        }
    }

})();