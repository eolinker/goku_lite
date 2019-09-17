(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 检测确认密码以及原密码是否相同指令js
     */
    angular.module('eolinker.directive')

    .directive('passwordConfirmDirective', [function() {
        return {
            restrict: 'A',
            require: "ngModel",
            link: function(scope, elem, attrs, ngModel) {
                var data = {
                    info:{
                        origin:elem.inheritedData("$formController")[attrs.passwordConfirmDirective]
                    },
                    fun: {
                        init: null, //初始化功能函数
                        origin: null, //原始密码监听函数
                        current: null //当前密码监听函数
                    }
                }
                data.fun.current = function(_default) {
                    ngModel.$setValidity("passwordConfirmDirective", _default === data.info.origin.$viewValue);
                    return _default;
                }
                data.fun.origin = function(_default) {
                    ngModel.$setValidity("passwordConfirmDirective", _default === ngModel.$viewValue);
                    return _default;
                }
                data.fun.init = (function() {
                    ngModel.$parsers.push(data.fun.current);
                    data.info.origin.$parsers.push(data.fun.origin);
                })()

            }
        };
    }]);
})();
