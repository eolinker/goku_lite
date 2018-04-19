(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 侧边栏组件
     */
    angular.module('goku')
        .component('sidebarCommonComponent', {
            templateUrl: 'app/component/common/sidebar/index.html',
            controller: indexController,
            bindings: {
                shrinkObject: '<',
                mainObject: '<',
                powerObject: '<',
                pluginList: '<'
            }
        })

    indexController.$inject = ['$scope', '$state', 'NavbarService'];

    function indexController($scope, $state, NavbarService) {
        var vm = this;
        vm.data = {
            info: {
                current: null
            },
            service: {
                default: NavbarService
            },
            fun: {
                shrink: null
            }
        }
        var fun = {};
        /**
         * 更新title
         * @param {object} item 选中列表项 
         */
        fun.setTitle = function (item) {
            if (item.childName) {
                $scope.$emit('$WindowTitleSet', {
                    list: [item.name, item.childName]
                });
            } else if (item.parentName) {
                $scope.$emit('$WindowTitleSet', {
                    list: [item.parentName, item.name]
                });
            } else {
                $scope.$emit('$WindowTitleSet', {
                    list: [item.name]
                });
            }
        }
        vm.data.fun.initMenu = function (arg) {
            if ($state.current.name.indexOf(arg.item.sref) > -1) {
                vm.data.info.current = arg.item;
                if (arg.item.childList) {
                    vm.data.service.default.info.navigation = {
                        query: vm.mainObject.baseInfo.navigation || [{
                            name: arg.item.name
                        }]
                    }
                    for (var $index = 0; $index < arg.item.childList.length; $index++) {
                        var val = arg.item.childList[$index];
                        if ($state.current.name.indexOf(val.sref) > -1) {
                            vm.data.service.default.info.navigation.current = val.name;
                            fun.setTitle(val);
                            break;
                        }
                    }
                } else {
                    vm.data.service.default.info.navigation = {
                        query: vm.mainObject.baseInfo.navigation || null,
                        current: arg.item.name
                    }
                    fun.setTitle(arg.item);
                }
            }
        }
        vm.data.fun.shrink = function () {
            vm.shrinkObject.isShrink = !vm.shrinkObject.isShrink;
            if (vm.mainObject.baseFun && vm.mainObject.baseFun.shrink) {
                vm.mainObject.baseFun.shrink();
            }
        }
        vm.data.fun.menu = function (arg, status) {
            if (arg.item.disable) return;
            if (!arg.item.href) {
                arg.item.back = false;
                vm.data.info.current = arg.item;
                if (arg.item.childList) {
                    switch (arg.item.status) {
                        case 'un-spreed':
                            {
                                break;
                            }
                        default:
                            {
                                vm.shrinkObject.isShrink = false;
                                break;
                            }
                    }
                    vm.data.service.default.info.navigation = {
                        query: vm.mainObject.baseInfo.navigation || [{
                            name: arg.item.name
                        }],
                        current: arg.item.childList[0].name
                    }
                } else {
                    switch (status) {
                        case 'child':
                            {
                                vm.data.service.default.info.navigation.current = arg.item.name;
                                break;
                            }
                        default:
                            {
                                vm.data.service.default.info.navigation = {
                                    query: vm.mainObject.baseInfo.navigation || null,
                                    current: arg.item.name
                                }
                                break
                            }
                    }
                }
            }
            if (arg.item.childSref) {
                if (arg.item.otherChildSref && $state.params.companyHashKey) {
                    $state.go(arg.item.otherChildSref, arg.item.otherParams);
                } else {
                    $state.go(arg.item.childSref, arg.item.params);
                }
            } else if (arg.item.sref) {
                $state.go(arg.item.sref, arg.params);
            } else {
                window.open(arg.item.href);
            }
            fun.setTitle(arg.item);

        }
    }
})();