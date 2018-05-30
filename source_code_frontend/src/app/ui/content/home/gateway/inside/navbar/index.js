(function() {
    /*
     * author：riverLethe
     * 网关内页页内包模块navbar相关js
     */
    angular.module('goku')
        .component('gatewayNavbar', {
            template: '<sidebar-common-component shrink-object="$ctrl.shrinkObject" main-object="$ctrl.data.component.sidebarCommonObject.mainObject" ></sidebar-common-component>',
            controller: indexController,
            bindings: {
                shrinkObject: '<'
            }
        })

    indexController.$inject = ['$scope', '$state'];

    function indexController($scope, $state) {
        var vm = this;
        vm.data = {
            component: {
                sidebarCommonObject: {}
            },
            info: {
                menu: [{ base: '/overview', name: '网关概况', sref: 'home.gateway.inside.overview',icon:'icon-tongjibaobiao',  power: -1 },
                    { base: '/api/', name: '接口列表', sref: 'home.gateway.inside.api', icon: 'icon-api', childSref: 'home.gateway.inside.api.list', params: { groupID: null }, power: -1 },
                    { base: '/backend', name: '后端管理', sref: 'home.gateway.inside.backend',  icon: 'icon-icocode',  power: -1 },
                    { 
                        base: '/gpedit', 
                        name: '权限控制', 
                        sref: 'home.gateway.inside.gpedit', 
                        childSref: 'home.gateway.inside.gpedit.default', 
                        childName:'策略组列表',
                        icon: 'icon-jiqirendaan', 
                        power: -1, 
                        status:'un-spreed',
                        childList: [
                            { name: '策略组列表', sref: 'home.gateway.inside.gpedit',childSref: 'home.gateway.inside.gpedit.default', parentName:'权限控制' }
                        ]
                    },{
                        name: '全局-黑白名单',
                        sref: 'home.gateway.inside.env',
                        icon: 'icon-waibuhuanjing',
                        power: -1
                    }
                ]
            },
            fun: {
                menu: null, //菜单功能函数
                shrink: null //收缩功能函数
            }
        }
        vm.data.fun.shrink = function() {
            $scope.$emit('$Home_ShrinkSidebar', { shrink: vm.shrinkObject.isShrink });
        }
        vm.$onInit = function() {
            vm.data.component.sidebarCommonObject = {
                mainObject: {
                    baseInfo: {
                        menu: vm.data.info.menu,
                        navigation: [{ name: '网关管理', sref: 'home.gateway.default' }, { name: $state.params.gatewayName }],
                    },
                    baseFun: {
                        shrink: vm.data.fun.shrink,
                    }
                }
            }
        }
    }

})();
