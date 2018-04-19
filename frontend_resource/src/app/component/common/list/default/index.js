(function() {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 默认列表组件
     * @extend {object} authorityObject 权限类{operate}
     * @extend {object} funObject 第一部分功能集类{showVar,btnGroupList{edit:{fun,key,class,showable,icon,tips},sort:{default,cancel,confirm:{fun,key,showable,class,icon,tips}}}}
     * @extend {object} mainObject 主类{baseInfo:{colspan,warning},item:{default,fun}}
     */
    angular.module('goku')
        .component('listDefaultCommonComponent', {
            templateUrl: 'app/component/common/list/default/index.html',
            controller: indexController,
            bindings: {
                authorityObject: '<',
                funObject: '<',
                mainObject: '<',
                list: '<'
            }
        })

    indexController.$inject = ['$scope', '$compile'];

    function indexController($scope, $compile) {
        var vm = this;
        vm.data = {
            info: {
                html: ''
            },
            fun: {
                more: null,
                filter: null
            }
        }

        /**
         * 初始化单项表格
         */
        vm.$onInit = function() {
            var template = {
                loopHtml: '',
                html: '<td ng-if="$ctrl.mainObject.item.fun&&(!$ctrl.mainObject.baseInfo.operate||($ctrl.mainObject.baseInfo.operate&&$ctrl.authorityObject.operate))">' +
                    '<div ng-if="' + (vm.mainObject.item.fun ? vm.mainObject.item.fun.power : true) + '">' +
                    '<button class="home-function-edit" ng-repeat="funItem in $ctrl.mainObject.item.fun.array" ng-if="funItem.show==-1||item[funItem.showType]==funItem.show" ng-click="$ctrl.data.fun.common(funItem,{item:item,$index:$outerIndex,$event:$event})" ng-disabled="' + vm.mainObject.baseInfo.disabled + '&&funItem.status!==\'allowed\'"><span class="iconfont icon-{{funItem.icon}}" ng-if="funItem.icon"></span>{{funItem.key}}</button>' +
                    '</div>' +
                    '</td>'
            }
            angular.forEach(vm.mainObject.item.default, function(val, key) {
                template.loopHtml = template.loopHtml + '<td class="' + (val.contentClass || '') + (val.switch?('" ng-switch="item.' + val.switch+'" '):'" ') + (val.title ? ('title="' + val.title + '"') : '') + '>' + val.html + '</td>';
            })
            vm.data.info.html = '<tr ng-style="$ctrl.mainObject.baseInfo.style" ng-class="{\'disabled-tr\':' + (vm.mainObject.baseInfo.disabled || false) + '}" ng-repeat=\'($outerIndex,item) in $ctrl.list\' ng-click="$ctrl.mainObject.baseFun.click({item:item})" >' + template.loopHtml + template.html + '</tr>';
        }

        /**
         * @description 统筹绑定调用页面列表功能单击函数
         * @param {extend} obejct 方式值
         * @param {object} arg 共用体变量，后根据传值函数回调方法
         */
        vm.data.fun.common = function(extend, arg) {
            arg.$event.stopPropagation();
            var template = {
                params: angular.copy(arg)
            }

            for (var key in extend.params) {
                if (extend.params[key] == null) {
                    template.params[key] = arg[key];
                } else {
                    template.params[key] = extend.params[key];
                }
            }
            extend.fun(template.params);
        }
    }
})();
