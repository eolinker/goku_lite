(function() {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 上传文件重置指令js
     */
    angular.module('goku.directive')

        .directive('innerHtmlCommonDirective', ['$compile', function($compile) {
            return {
                restrict: 'AE',
                scope:{
                    html:'@'
                },
                link: function($scope, elem, attrs, ctrl) {
                    var data = {
                        fun: {
                            init: null //初始话检测功能函数
                        }
                    }
                    $scope.$watch('html',function(){
                        if(!$scope.html)return;
                        elem.append($compile($scope.html)($scope.$parent));
                    })
                }
            };
        }]);
})();