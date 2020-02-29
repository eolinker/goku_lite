(function () {
    'use strict';
    /**
     * 产品基本需求
     * （1）单击下拉按钮，下拉菜单为全部选项内容
     * （2）编辑输入框，下拉菜单为筛选后的内容
     * （3）能够通过上下箭头控制下拉菜单选中项
     */
    /**
     * @author 广州银云信息科技有限公司
     * @description 自动补全控件
     * @extends {obj} mainObject 主配置信息
     * @extends {string} placeholder 内置输入框placeholder内容[optional]
     * @extends {array} array 预设列表
     * @extends {string} model 输入框绑定对象
     * @extends {string} type 输入对象类型
     * @extends {function} inputChangeFun 输入框预设函数[optional]
     */

    angular.module('eolinker')
        .component('autoCompleteComponent', {
            templateUrl: 'app/component/autoComplete/index.html',
            controller: indexController,
            bindings: {
                mainObject:'<',
                readOnly: '<',
                placeholder: '@',
                keyName: '@',
                required: "<",
                array: '<', //自定义数组填充数组
                model: '=', //输入框绑定
                inputChangeFun: '&', //输入框值改变绑定功能函数
            }
        })

    indexController.$inject = ['$scope', '$element'];

    function indexController($scope, $element) {
        var vm = this;
        vm.data = {
            query: [],
            inputElem: $element[0].getElementsByClassName('input-text-acac'),
            inputIsFocus:false
        };
        vm.fun = {};
        var data = {
                originalElemCount: 0
            },
            privateFun = {};
        vm.fun.modelChange = function () {
            vm.data.inputIsFocus=true;
            vm.inputChangeFun();
            privateFun.clearSelectItem();
            if (vm.model[vm.keyName]) {
                vm.data.query = [];
                let tmpIndex = 0;
                angular.forEach(vm.array, function (val, key) {
                    try {
                        if (val.toLowerCase().indexOf(vm.model[vm.keyName].toLowerCase())>-1) {
                            vm.data.query.splice(tmpIndex, 0, val);
                            tmpIndex++;
                        } else if (val.toLowerCase().indexOf(vm.model[vm.keyName].toLowerCase()) > -1) {
                            vm.data.query.push(val);
                        }
                    } catch (EVAL_ERR) {
                        console.error(EVAL_ERR)
                    }
                })
                if (vm.data.query.length <= 0) {
                    vm.data.viewIsShow = false;
                }
            } else {
                vm.data.query = vm.array;
            }
        }
        vm.fun.changeSwitch = function (inputBool) {
            if (vm.readOnly) return;
            vm.data.inputIsFocus=inputBool;
            if(vm.data.inputIsFocus){
                vm.data.query = vm.array;
                
            }
            vm.data.inputElem[0].focus();
        }
        vm.fun.changeText = function (inputText) {
            vm.data.inputIsFocus=false;
            vm.model[vm.keyName] = inputText;
            vm.inputChangeFun();
        }

        /**
         * @description 重置下拉菜单选中项
         */
        privateFun.clearSelectItem = () => {
            data.originalElemCount = 0;
            vm.data.currentElementCount = data.originalElemCount - 1;
        }
        vm.fun.inputBlur = ($event) => {
            $event.stopPropagation();
            vm.data.inputIsFocus=false;
        }
        vm.fun.inputFocus = ($event) => {
            $event.stopPropagation();
            privateFun.clearSelectItem();
        }
        vm.fun.keydown = function (_default) {
            if (!vm.data.hasOwnProperty('currentElementCount')) {
                vm.data.currentElementCount = data.originalElemCount - 1;
            }
            switch (_default.keyCode) {
                case 38: // up
                    {
                        vm.data.currentElementCount = vm.data.currentElementCount <= data.originalElemCount ? ((vm.data.query || []).length || 1) - 1 : (vm.data.currentElementCount - 1);
                        $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        break;
                    }
                case 40: // down
                    {
                        _default.preventDefault();
                        vm.data.currentElementCount++;
                        if (vm.data.currentElementCount === (vm.data.query || []).length) {
                            vm.data.currentElementCount = data.originalElemCount;
                        }
                        $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        break;
                    }
                case 13:
                    { //enter
                        _default.preventDefault();
                        if (vm.data.currentElementCount >= 0) {
                            vm.fun.changeText(vm.data.query[vm.data.currentElementCount], vm.data.currentElementCount);
                            $scope.$root && $scope.$root.$$phase || $scope.$apply();
                        }
                        return false;
                    }
            }
        }
    }
})();