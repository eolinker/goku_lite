(function() {
    'use strict';

    /**
     * 监听dom结构变化
     * @param {string} bindId 绑定监听id
     * @param {function} watchDomCommonDirective 绑定方法
     */
    angular.module('eolinker.directive')
        .directive('watchDomCommonDirective', [function() {
            return {
                restrict: 'A',
                scope: {
                    bindId: '@',
                    watchDomCommonDirective: '&'
                },
                link: function($scope,elem,attrs,ngModel) {
                    var data = {
                        elem: document.getElementById($scope.bindId),
                        fun: {
                            init: null
                        }
                    }

                    /**
                     * @description dom节点变化监听函数
                     */
                    data.fun.DOMSubtreeModified = function() {
                        $scope.watchDomCommonDirective({input:data.elem.innerText});
                    }

                    /**
                     * 自启动初始化功能函数
                     */
                    data.fun.init = (function() {
                        angular.element(data.elem).bind('DOMSubtreeModified', data.fun.DOMSubtreeModified);
                    })()
                }
            };
        }]);

})();
