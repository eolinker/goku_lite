(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 上传文件重置指令js
     */
    angular.module('eolinker.directive')

    .directive('fileResetDirective', ['$compile', function($compile) {
        return {
            restrict: 'A',
            link: function($scope, elem, attrs, ctrl) {
                var data = {
                    fun: {
                        init: null, //初始化功能函数
                        change: null //file按钮内容更改触发功能函数
                    }
                }
                data.fun.change = function(_default) {
                    elem[0].parentNode.replaceChild($compile(elem[0].outerHTML)($scope)[0], elem[0])
                    $scope.$root && $scope.$root.$$phase || $scope.$apply();
                }
                data.fun.init = (function() {
                    elem.bind(attrs.buttonFunction || 'click', data.fun.change);
                })()
            }
        };
    }]);
})();
