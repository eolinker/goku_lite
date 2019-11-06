(function () {
    /*
     * author：广州银云信息科技有限公司
     * 全局sidebar指令相关js
     */
    angular.module('eolinker')
        .component('eoSidebar', {
            template: '<sidebar-common-component power-object="$ctrl.service.authority.permission.default" shrink-object="$ctrl.shrinkObject" main-object="$ctrl.data.component.sidebarCommonObject.mainObject"></sidebar-common-component>',
            controller: indexController,
            bindings: {
                shrinkObject: '<'
            }
        })

    indexController.$inject = ['$scope', 'Authority_CommonService'];

    function indexController($scope, Authority_CommonService) {

        var vm = this;
        vm.data = {
            component: {
                sidebarCommonObject: {}
            },
            info: {
                current: null,
                menu: [{
                        name: '首页',
                        sref: 'home.panel',
                        icon: 'icon-shouye_o',
                        power: -1,
                        status:'un-spreed'
                    },
                    {
                        name: '网关节点',
                        sref: 'home.cluster',
                        childSref: 'home.cluster.default',
                        icon: 'icon-yingyongAPP_o',
                        power: -1
                    }, {
                        name: '注册方式与负载',
                        sref: 'home.balance',
                        icon: 'icon-renwuzhongxin_o',
                        childSref: 'home.balance.service',
                        childSref: 'home.balance.list.default',
                        power: -1,
                        childList: [{
                                name: '服务注册方式',
                                sref: 'home.balance.service'
                            },
                            {
                                name: '负载配置',
                                sref: 'home.balance.list',
                                childSref:'home.balance.list.default',
                                params:{
                                    cluster:'default'
                                }
                            }
                        ]
                    }, {
                        name: '接口管理',
                        sref: 'home.project',
                        icon: 'icon-lianjie_o',
                        childSref: 'home.project.default',
                        power: -1
                    }, {
                        name: '访问策略',
                        sref: 'home.gpedit',
                        icon: 'icon-yuechi_o',
                        power: -1,
                        childSref: 'home.gpedit.default'
                    }, {
                        name: 'API监控设置',
                        sref: 'home.monitor',
                        icon: 'icon-yuechi_o',
                        power: -1
                    },
                    {
                        name: '扩展插件',
                        sref: 'home.plugin',
                        childSref: 'home.plugin.default',
                        icon: 'icon-cengji_o',
                        power: -1,
                    },
                    {
                        name: '网关设置',
                        sref:'home.setting',
                        icon: 'icon-quanjushezhi_o',
                        power: -1,
                        status:'un-spreed',
                        childList: [
                            {
                                name: '日志设置',
                                sref: 'home.setting.log'
                            }
                        ]
                    }
                    , {
                        name: '配置管理',
                        sref: 'home.publish',
                        icon: 'icon-liangliangduibi_o',
                        power: -1,
                    }
                ]
            }
        }
        vm.service={
            authority:Authority_CommonService
        }
        vm.$onInit = function () {
            vm.shrinkObject.isShrink = false;
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