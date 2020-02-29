(function () {
    'use strict';
    /**
     * @name 侧边栏
     * @author 广州银云信息科技有限公司
     */
    angular.module('eolinker')
        .component('sidebarCommonComponent', {
            templateUrl: 'app/component/common/sidebar/index.html',
            controller: indexController,
            bindings: {
                mainObject: '<',
                powerObject: '<',
                permissionObject: '<'
            }
        })

    indexController.$inject = ['$scope', '$state', 'NavbarService', 'Sidebar_CommonService'];

    function indexController($scope, $state, NavbarService, Sidebar_CommonService) {
        var vm = this;
        vm.data = {
            current: null,
            itemSpreedObject: {}
        }
        vm.fun = {};
        vm.service = {
            default: NavbarService,
            sidebar: Sidebar_CommonService
        }
        vm.fun.initMenu = function (arg) {
            if ($state.current.name===arg.item.sref||$state.current.name.indexOf(arg.item.sref+'.') > -1) {
                vm.data.current = arg.item;
                if (arg.item.childList && arg.item.childList.length > 0) {
                    vm.data.itemSpreedObject[arg.item.sref] = 1;
                }
                if (vm.mainObject.baseInfo.unNav || arg.item.unNav) {
                    vm.service.default.info.navigation = {};
                    return;
                }
                if (arg.item.childList) {
                    vm.service.default.info.navigation = {
                        query: vm.mainObject.baseInfo.navigation || [{
                            name: arg.item.name
                        }]
                    }
                    for (var $index = 0; $index < arg.item.childList.length; $index++) {
                        var val = arg.item.childList[$index];
                        if ($state.current.name.indexOf(val.sref) > -1 || $state.current.name.indexOf(val.otherSref) > -1) {
                            vm.service.default.info.navigation.current = val.name;
                            break;
                        }
                    }
                } else {
                    vm.service.default.info.navigation.query=vm.mainObject.baseInfo.navigation || null;
                    vm.service.default.info.navigation.current=arg.item.name;
                }
            }
        }

        /**
         * @description 单击侧边栏菜单
         * @param {string} inputMark 操作类型
         * @param {object} inputOperateItem 当前操作对象
         * @param {object} inputParentOperateItem 父操作对象
         */
        vm.fun.clickMenu = function (inputMark, inputOperateItem, inputParentOperateItem) {
            if (inputOperateItem.click) {
                let tmpStopProgrammer = inputOperateItem.click({
                    item: inputOperateItem
                });
                if (tmpStopProgrammer === true) return;
            }
            if (inputOperateItem.mark == 'text') return;
            window.sessionStorage.removeItem('COMMON_SEARCH_TIP');
            if (inputOperateItem.permissionKey == 'disabled' && vm.permissionObject[inputOperateItem.permission].disabled == inputOperateItem.permissionValue) return;
            if(inputParentOperateItem){
                vm.service.default.info.navigation = {
                    query: vm.mainObject.baseInfo.navigation || [{
                        name: inputParentOperateItem.name
                    }],
                    current: inputOperateItem.name
                }
            }else if (inputOperateItem.childList && inputOperateItem.childList.length > 0) {
                if (vm.data.itemSpreedObject[inputOperateItem.sref] == 1) {
                    delete vm.data.itemSpreedObject[inputOperateItem.sref];
                } else {
                    vm.data.itemSpreedObject[inputOperateItem.sref] = 1;
                }
                return;
            }
            vm.data.current = inputParentOperateItem || inputOperateItem;
            if (vm.mainObject.baseInfo.unNav) {
                vm.service.default.info.navigation = {};
            } else {
                if (inputOperateItem.childList) {
                    vm.service.default.info.navigation = {
                        query: vm.mainObject.baseInfo.navigation || [{
                            name: inputOperateItem.name
                        }],
                        current: inputOperateItem.childList[0].name
                    }
                } else {
                    switch (inputMark) {
                        case 'child': {
                            vm.service.default.info.navigation.current = inputOperateItem.name;
                            break;
                        }
                        default: {
                            vm.service.default.info.navigation = {
                                query: vm.mainObject.baseInfo.navigation || null,
                                current: inputOperateItem.name
                            }
                            break
                        }
                    }
                }
            }
            if (inputOperateItem.childSref) {
                if (inputOperateItem.otherChildSref && $state.params.spaceKey) {
                    $state.go(inputOperateItem.otherChildSref, inputOperateItem.otherParams);
                } else {
                    $state.go(inputOperateItem.childSref, inputOperateItem.params);
                }
            } else if (inputOperateItem.sref) {
                $state.go(inputOperateItem.sref, inputOperateItem.params);
            }
        }
        $scope.$on('$stateChangeSuccess', function () {
            if (vm.data.current && $state.current.name.indexOf(vm.data.current.sref) === -1 && $state.current.sidebarIndex) {
                var sidebarIndex =0;
                for (var i = 0; i < vm.mainObject.baseInfo.menu.length; i++) {
                    let sideBarItem = vm.mainObject.baseInfo.menu[i];
                    if($state.current.name.indexOf(sideBarItem.sref)>-1){
                        sidebarIndex=i;
                        break;
                    }
                }
                vm.fun.initMenu({
                    item: vm.mainObject.baseInfo.menu[sidebarIndex]
                });

            }
            if ($state.current.navName) {
                try {
                    vm.service.default.info.navigation.current = $state.current.navName;
                } catch (e) {}
            }
            if (!$state.current.needExtraNav) {
                vm.service.default.info.navigation.extra = null;
            }
            
        })
        vm.fun.shrinkSidebar=()=>{
            vm.service.sidebar.isShrink=!vm.service.sidebar.isShrink;
            window.localStorage.setItem('EO_CONFIG_SIDEBAR_SHRINK_STATUS',vm.service.sidebar.isShrink)
        }
        $scope.$on('$HomeProjectInsideSidebarShrink', function () {
            vm.service.sidebar.isShrink = true;
            $scope.$root && $scope.$root.$$phase || $scope.$apply();
        })

        /**
         * @description 监听重置侧边栏聚焦状态广播
         * @param {int} inputIndex focus侧边序号
         */
        $scope.$on('$ResetSidebarActiveStatus_SidebarComponent', function (_default, inputIndex) {
            vm.fun.initMenu({
                item: vm.mainObject.baseInfo.menu[inputIndex]
            })
        })

    }
})();