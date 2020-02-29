(function () {
    'use strict';
    /*
     * author：广州银云信息科技有限公司
     * 导出文件指令js
     */
    angular.module('eolinker.directive')

        .directive('dumpDirective', ['PATH_INFO', function (PATH_INFO) {
            return {
                restrict: 'AE',
                transclude: true,
                template: '<a class="eo-export {{interaction.request.class}}" ng-click="data.fun.dump()"><p>{{interaction.request.text}}</p></a><loading-common-component fun="dumpDirective(arg)" interaction="{request:{delay:true}}"></loading-common-component>',
                scope: {
                    interaction: '<',
                    dumpDirective: '&'
                },
                link: function ($scope, elem, attrs, ctrl) {
                    $scope.data = {
                        info: {
                            elem: document.getElementById('dump-directive_js')
                        },
                        fun: {
                            dump: null, //导出功能函数
                        }
                    }
                    var data = {
                        info: {
                            broadcast: null
                        },
                        fun: {
                            init: null, //初始化功能函数
                            $DumpDirective_Click: null, //导出监听请求返回功能函数
                            $Destory: null, //资源回收
                        }
                    }
                    $scope.data.fun.dump = function () {
                        $scope.$broadcast('$Init_LoadingCommonComponent', {
                            arg: {
                                switch: $scope.interaction.request.switch
                            }
                        });
                    }
                    data.fun.$DumpDirective_Click = function (_default, arg) {
                        $scope.data.info.elem.href = PATH_INFO.HOST + 'export/' + arg.response.fileName;
                        // $scope.data.info.elem.download = arg.fileName || arg.response.fileName;
                        $scope.data.info.elem.click();

                    }
                    data.fun.$Destory = function () {
                        data.info.broadcast();
                    }
                    data.fun.init = (function () {
                        $scope.interaction = $scope.interaction || {
                            request: {}
                        };
                        data.info.broadcast = $scope.$on('$DumpDirective_Click_' + ($scope.interaction.request.broadcast || $scope.interaction.request.switch || ''), data.fun.$DumpDirective_Click);
                        $scope.$on('$destroy', data.fun.$Destory);
                    })()
                }
            };
        }]);
})();