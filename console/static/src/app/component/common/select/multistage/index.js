(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 多级下拉菜单
     */

    angular.module('eolinker')
        .component('selectMultistageCommonComponent', {
            templateUrl: 'app/component/common/select/multistage/index.html',
            bindings: {
                input: '<',
                disabled: "<",
                output: '='
            },
            controller: indexController
        });

    indexController.$inject = ['$rootScope', '$scope', '$document', 'Group_MultistageService'];

    function indexController($rootScope, $scope, $document, Group_MultistageService) {
        var vm = this;
        vm.data = {
            wantToSelect: false,
            textList: [],
            parentNodeList: [],
            query: null
        }
        vm.fun = {};
        var groupInfo = null

        vm.fun.filter = function (arg) {
            if (!vm.data.queryHaveChild && arg.hasChild) {
                vm.data.queryHaveChild = true;
            }
            if (vm.data.q && arg[vm.input.key].indexOf(vm.data.q) == -1) {
                return false;
            }
            return arg;
        }
        var fun = {},
            groupFun = {},
            data = {
                initialObject: {
                    textList: [],
                    query: []
                }
            },
            service = {
                groupCommon: Group_MultistageService
            }
        /**
         * 返回子分组函数
         * @param {number} groupID 参数，eg:{$event:dom,item:单击所处父分组项}
         */
        groupFun.getChildGroup = function (currentGroupID) {
            var template = {
                output: []
            }
            if (groupInfo.childGroupPath[currentGroupID] && groupInfo.childGroupPath[currentGroupID].length) {
                angular.forEach(groupInfo.childGroupPath[currentGroupID], function (val, key) {
                    var group = groupInfo.groupObj[val];
                    template.output.push({
                        groupName: group[vm.input.key],
                        groupID: group[vm.input.value],
                        hasChild: groupInfo.childGroupPath[val] && groupInfo.childGroupPath[val].length ? true : false,
                    })
                })
                return template.output;
            } else {
                return [];
            }
        }

        groupFun.getBrotherGroup = function (currentGroupID) {
            var template = {
                output: []
            }
            var currentGroup = groupInfo.groupObj[currentGroupID] || {};
            currentGroup.parentGroupID = currentGroup.parentGroupID || 0;
            return groupFun.getChildGroup(currentGroup.parentGroupID);
        }
        vm.fun.wantToSelect = function ($event) {
            $event.stopPropagation();
            if (vm.disabled) return;
            vm.data.wantToSelect = !vm.data.wantToSelect;
            fun.clearText();
        }
        $document.on("click", function (_default) {
            if (vm.input.selectDeepest) return;
            vm.data.wantToSelect = false;
            $scope.$root && $scope.$root.$$phase || $scope.$apply();
        });
        fun.setText = function (newValue) {
            var text = '';
            for (var key in vm.data.textList) {
                text = text + (text ? ' / ' : '') + vm.data.textList[key][vm.input.key];
            }
            data.selectGroup = groupInfo.groupObj[newValue];
            if (vm.input.selectDeepest) {
                let hasChild = groupInfo.childGroupPath[newValue] && groupInfo.childGroupPath[newValue].length ? true : false;
                if (!hasChild) {
                    vm.output.new = {
                        text: text,
                        value: newValue
                    }
                }
            } else {
                vm.output.new = {
                    text: text,
                    value: newValue
                }
            }

        }
        fun.resetGroup = function (currentGroupID) {
            vm.data.textList = service.groupCommon.fun.getGroupPath({
                currentGroupID: currentGroupID,
                groupInfo: groupInfo
            });
            fun.setText(currentGroupID);
        }
        fun.goToParent = function () {
            var currentFirstGroup = groupInfo.groupObj[vm.data.query[0].groupID];
            if (data.selectGroup.groupID == currentFirstGroup.parentGroupID) {
                //选中分组为当前列表的父分组
                vm.data.query = groupFun.getBrotherGroup(data.selectGroup.groupID);
            } else {
                //选中分组属于当前列表
                vm.data.query = groupFun.getBrotherGroup(data.selectGroup.parentGroupID);
                fun.resetGroup(data.selectGroup.parentGroupID)
            }
            vm.data.hasParent = data.selectGroup.parentGroupID ? true : false;
        }
        fun.clickGroup = function (arg) {
            var groupID = arg[vm.input.value];
            var childGroup = groupFun.getChildGroup(groupID);
            if (childGroup.length > 0) {
                vm.data.hasParent = true;
                vm.data.query = groupFun.getChildGroup(groupID);
            } else {
                vm.data.wantToSelect = false;
            }
            fun.resetGroup(groupID)
        }
        fun.clearText = function () {
            vm.data.q = '';
        }
        vm.fun.click = function ($event) {
            $event.stopPropagation();
            var template = {};
            try {
                template.point = $event.target.classList[0];
            } catch (e) {
                template.point = 'default';
            }
            switch (template.point) {
                case 'select-multistage-btn-clear-text':
                    {
                        fun.clearText();
                        break;
                    }
                case 'select-multistage-btn-close':
                    {

                        vm.data.wantToSelect = false;;
                        break;
                    }
                case 'select-multistage-btn-back':
                    {
                        fun.goToParent();
                        break;
                    }
                case 'select-multistage-btn-item':
                    {
                        template.index = $event.target.getAttribute('eo-attr-index');
                        template.query = [];
                        vm.data.query.map(function (val, key) {
                            if (vm.fun.filter(val)) {
                                template.query.push(val);
                            }
                        })
                        fun.clickGroup(template.query[template.index]);
                        break;
                    }
                case 'select-multistage-btn-item-span':
                    {
                        template.index = $event.target.parentNode.getAttribute('eo-attr-index');
                        template.query = [];
                        vm.data.query.map(function (val, key) {
                            if (vm.fun.filter(val)) {
                                template.query.push(val);
                            }
                        })
                        fun.clickGroup(template.query[template.index]);
                        break;
                    }
            }


        }
        fun.initial = function () {
            if ((vm.data.query || []).length <= 0) return;
            groupInfo = service.groupCommon.fun.generalGroupInfo({
                list: vm.data.query
            })
            if (vm.input[vm.input.value] <= 0) {
                data.initialObject.query = groupFun.getChildGroup(0);
                data.initialObject.textList = service.groupCommon.fun.getGroupPath({
                    currentGroupID: data.initialObject.query[0][vm.input.value],
                    groupInfo: groupInfo
                });
                vm.data.hasParent = false;
                data.selectGroup = groupInfo.groupObj[data.initialObject.query[0][vm.input.value]];
            } else {
                data.initialObject.query = groupFun.getBrotherGroup(vm.input[vm.input.value]);
                data.initialObject.textList = data.initialObject.textList = service.groupCommon.fun.getGroupPath({
                    currentGroupID: vm.input[vm.input.value],
                    groupInfo: groupInfo
                });
                data.selectGroup = groupInfo.groupObj[vm.input[vm.input.value]];
                vm.data.hasParent = data.selectGroup.parentGroupID ? true : false;
            }
            fun.resetInitial();
        }
        fun.resetInitial = function () {
            vm.output = {};
            vm.data.query = angular.copy(data.initialObject.query);
            vm.data.textList = angular.copy(data.initialObject.textList);
            var text = '';
            for (var key in vm.data.textList) {
                text = text + (text ? '>' : '') + vm.data.textList[key][vm.input.key];
            }
            vm.output.new = vm.output.original = {
                text: text,
                value: vm.data.textList[vm.data.textList.length - 1][vm.input.value]
            }
        }
        data.broadcast = $scope.$on('$ResetInitial_SelectMultistageCommonComponent', fun.resetInitial);
        $rootScope.global.$watch.push($scope.$watch('$ctrl.input.query', function () {
            if (!vm.input.query) return;
            vm.data.query = vm.input.query;
            fun.initial();
        }));
        $scope.$on('$destroy', function () {
            data.broadcast();
        });
    }
})();