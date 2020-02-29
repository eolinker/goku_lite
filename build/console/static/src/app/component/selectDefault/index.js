(function () {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 人员下拉菜单
     */

    angular.module('eolinker')
        .component('selectDefaultCommonComponent', {
            templateUrl: 'app/component/selectDefault/index.html',
            bindings: {
                input: '<',
                output: '=',
                required: '@',
                multiple: '@',
                modelKey: '@',
                inputChangeFun: '&',
                disabled: '<',
                mainObject:'<'
            },
            controller: indexController
        });

    indexController.$inject = ['$rootScope', '$scope', '$element'];

    function indexController($rootScope, $scope, $element) {
        var vm = this;
        vm.data = {
            text: '',
            query: null,
            searchInputElem: null,
            inputElem: $element[0].getElementsByClassName('input-text'),
            q: ''
        }
        vm.fun = {};
        var fun = {},
            data = {
                originalElemCount: 0
            };
        vm.fun.inputMousedown = function ($event) {
            if ($event) $event.stopPropagation();
            vm.data.currentElementCount = data.originalElemCount - 1;
            fun.clearText();
        }
        vm.fun.searchChange = function () {
            var tmpQuery = angular.copy(vm.input.query);
            vm.data.currentElementCount = data.originalElemCount;
            if (!vm.data.q) {
                vm.data.query = tmpQuery;
                return;
            }
            vm.data.query = tmpQuery.filter(function (val, key) {
                if (((val[vm.input.key] || '').toLowerCase()).indexOf((vm.data.q || '').toLowerCase()) > -1) {
                    return val;
                } else {
                    return undefined;
                }
            })
        }
        vm.fun.divFocus = function () {
            vm.fun.inputMousedown();
            vm.data.inputElem[0].focus();
        }
        vm.fun.keydown = function (_default) {
            if (!vm.data.hasOwnProperty('currentElementCount')) {
                vm.data.currentElementCount = data.originalElemCount - 1;
            }
            switch (_default.keyCode) {
                case 38: // up
                    {
                        vm.data.currentElementCount = vm.data.currentElementCount <= data.originalElemCount ? ((vm.data.query || []).length || 1) - 1 : (vm.data.currentElementCount - 1);
                        if (vm.data.currentElementCount == data.originalElemCount) {
                            if (vm.data.searchInputElem) {
                                vm.data.searchFocusStatus = true;
                                vm.data.searchInputElem[0].click();
                                vm.data.searchInputElem[0].focus();
                                // return;
                            }
                        } else if (vm.data.currentElementCount == 4) {
                            vm.data.inputElem[0].focus();
                        }
                        $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        break;
                    }
                case 40: // down
                    {
                        _default.preventDefault();
                        vm.data.currentElementCount++;
                        if (vm.data.currentElementCount == (vm.data.query || []).length) {
                            vm.data.currentElementCount = data.originalElemCount;
                        }
                        if (vm.data.currentElementCount == data.originalElemCount) {
                            if (vm.data.searchInputElem) {
                                vm.data.searchFocusStatus = true;
                                vm.data.searchInputElem[0].click();
                                vm.data.searchInputElem[0].focus();
                                // return;
                            }
                        } else if (vm.data.currentElementCount == 0) {
                            vm.data.inputElem[0].focus();
                        }

                        $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        break;
                    }
                case 13:
                    { //enter
                        _default.preventDefault();
                        if (vm.data.currentElementCount >= 0) {
                            fun.select(vm.data.query[vm.data.currentElementCount], vm.data.currentElementCount);
                            if (vm.multiple!=='true') {
                                vm.data.inputElem[0].blur();
                            }
                            $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        }
                        return false;
                    }
            }
        }
        fun.clearText = function () {
            vm.data.q = '';
            vm.data.query = vm.input.query;
        }
        fun.select = function (arg, index) {
            vm.output = vm.output || {};
            let tmpOriginalValue=vm.output[vm.modelKey];
            if (vm.multiple==='true') {
                vm.output[vm.modelKey] = vm.output[vm.modelKey] || {};
                if (arg[vm.input.value] in vm.output[vm.modelKey]) {
                    if (vm.required && !(Object.keys(vm.output[vm.modelKey]).length > 1)) return;
                    delete vm.output[vm.modelKey][arg[vm.input.value]];
                    vm.data.text = '';
                    angular.forEach(vm.output[vm.modelKey], function (val, key) {
                        vm.data.text = vm.data.text + (vm.data.text ? ',' : '') + val;
                    })
                } else {
                    vm.output[vm.modelKey][arg[vm.input.value]] = arg[vm.input.key];
                    vm.data.text = vm.data.text + (vm.data.text ? ',' : '') + arg[vm.input.key];
                }
            } else {
                vm.data.text = arg[vm.input.key];
                vm.output[vm.modelKey] = arg[vm.input.value];
            }
            /**
             * @desc 当存在值变化&外部change函数，触发以下事件
             */
            if (tmpOriginalValue!==vm.output[vm.modelKey]&&vm.inputChangeFun) {
               vm.inputChangeFun();
            }
        }
        vm.fun.clear = function ($event) {
            $event.stopPropagation();
            if (vm.multiple==='true') {
                vm.data.text = '';
                vm.output[vm.modelKey] = {};
            } else {
                vm.data.text = '';
                vm.output = vm.output || {};
                vm.output[vm.modelKey] = null;
            }

            if (vm.inputChangeFun) vm.inputChangeFun();
        }
        vm.fun.domClick=(inputEvent)=>{
            inputEvent.stopPropagation();
        }
        vm.fun.listMouseDown = function ($event) {
            $event.stopPropagation();
            var template = {};
            try {
                template.point = $event.target.classList[0];
            } catch (e) {
                template.point = 'default';
            }
            switch (template.point) {
                case 'select-btn-item':
                case 'sd_check_box':
                    {
                        if (vm.multiple==='true') {
                            $event.preventDefault();
                            vm.data.containerFocus = true;
                        }
                        template.index = $event.target.getAttribute('eo-attr-index');
                        fun.select(Object.assign({}, vm.data.query[template.index]), template.index);
                        break;
                    }
            }
        }
        vm.fun.searchActiveStatus = function (inputFocusStatus) {
            vm.data.searchFocusStatus = inputFocusStatus;
        }
        fun.initial = function () {
            vm.data.text = "";
            if (vm.input.initialData === undefined) return;
            if (vm.multiple==='true') {
                angular.forEach(vm.input.initialData, function (val, key) {
                    if (key) {
                        vm.data.text += (vm.data.text ? ',' : '') + val;
                        vm.output[vm.modelKey][key] = val;
                    }
                })
            } else {
                for (var key in vm.data.query) {
                    var val = vm.data.query[key];
                    if (val[vm.input.value] == vm.input.initialData) {
                        vm.data.text = val[vm.input.key];
                        vm.output = vm.output || {};
                        vm.output[vm.modelKey] = val[vm.input.value];
                        break;
                    }
                }
            }
        }
        $rootScope.global.$watch.push($scope.$watch('$ctrl.input.query+$ctrl.input.initialData', function () {
            if (!vm.input.query) return;
            vm.data.query = vm.input.query;
            if (vm.data.query.length >= 5) {
                data.originalElemCount = -1;
                vm.data.searchInputElem = $element[0].getElementsByClassName('input-search');
            } else {
                data.originalElemCount = 0;
            }
            fun.initial();
        }));
        vm.$onInit = function () {
            vm.modelKey = vm.modelKey || 'value';
            $element.bind('keydown', vm.fun.keydown);
        }
    }
})();