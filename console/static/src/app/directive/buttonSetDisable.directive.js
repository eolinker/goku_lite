(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 提交表单禁用button按钮指令js
     */
    angular.module('eolinker.directive')

    .directive('buttonSetDisableDirective', [function() {
        return {
            restrict: 'AE',
            scope: {
                buttonSetDisableDirective: '&' //绑定设置回调函数
            },
            link: function($scope, elem, attrs, ctrl) {
                var data = {
                    fun: {
                        init: null, //初始化功能函数
                        btnFun: null //按钮相关功能函数
                    }
                }
                data.fun.btnFun = function(event) {
                    event.stopPropagation();
                    var template = {
                        promise: $scope.buttonSetDisableDirective()
                    }
                    elem.prop('disabled', true);
                    if (template.promise) {
                        template.promise.finally(function() {
                            elem.prop('disabled', false);
                        })
                    } else {
                        elem.prop('disabled', false);
                        $scope.$root && $scope.$root.$$phase || $scope.$apply();
                    }
                }
                data.fun.init = (function() {
                    elem.bind(attrs.buttonFunction || 'click', data.fun.btnFun);
                })()
            }
        };
    }]);
})();
