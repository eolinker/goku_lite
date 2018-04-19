(function() {
    'use strict';
    /**
     * @author 广州银云信息科技有限公司
     * @description 工具栏组件
     * @extend {object} authorityObject 权限类{operate}
     * @extend {object} activeObject 活动/聚焦标志
     * @extend {object} showObject 显示标志
     * @extends {array} list 列表
     * @extend {object} otherObject 不可预期辅助类
     */
    angular.module('goku')
        .component('menuCommonComponent', {
            templateUrl: 'app/component/common/menu/index.html',
            controller: indexController,
            bindings: {
                authorityObject: '<',
                otherObject: '<',
                activeObject: '<',
                showObject: '<',
                list: '<'
            }
        })

    indexController.$inject = ['$scope', '$compile'];

    function indexController($scope, $compile) {
        var vm = this;
        vm.data = {
            fun: {
                common: null
            }
        }



        /**
         * @description 统筹绑定调用页面列表功能单击函数
         * @param {extend} obejct 方式值
         * @param {object} arg 共用体变量，后根据传值函数回调方法
         */
        vm.data.fun.common = function(extend, arg) {
            if(!extend)return;
            var template = {
                params: arg
            }
            switch (typeof(extend.params)) {
                case 'string':
                    {
                        return eval('extend.default(' + extend.params + ')');
                        break;
                    }
                default:
                    {
                        for (var key in extend.params) {
                            if (extend.params[key] == null) {
                                template.params[key] = arg[key];
                            } else {
                                template.params[key] = extend.params[key];
                            }
                        }
                        return extend.default(template.params);
                        break;
                    }
            }

        }
    }
})();
