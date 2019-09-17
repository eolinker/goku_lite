(function () {
    'use strict';

    /**
     * loading加载组件
     * @param {function} fun 绑定方法
     * @param {object} interaction 交互参数
     */
    angular.module('eolinker')
        .component('loadingCommonComponent', {
            template: '<div class="loading-content" ng-hide="$ctrl.data.info.isEnd">' +
                '<div class="loading">' +
                '    <ul>' +
                '        <li>' +
                '            <div><span class="iconfont  icon-jiazai_shuaxin"></span></div>正在{{$ctrl.data.tips||"载入"}}' +
                '        </li>' +
                '    </ul>' +
                '</div>' +
                '</div>',
            controller: indexController,
            bindings: {
                fun: '&',
                interaction: '<'
            }
        })

    indexController.$inject = ['$scope', '$state'];

    function indexController($scope, $state) {
        var vm = this;
        var data = {
            fun: {
                $Destory: null, //资源回收
                dataProcessing: null, //数据处理功能函数
                $LoadingInit: null //监听LoadingInit广播功能函数
            },
            info: {
                broadcast: null
            }
        }
        vm.data = {
            info: {
                isEnd: true
            }
        }
        data.fun.dataProcessing = function (arg) {
            vm.data.info.isEnd = false;
            var template = {
                promise: vm.fun({
                    arg: arg
                })
            }
            if (template.promise) {
                template.promise.finally(function () {
                    vm.data.info.isEnd = true;
                    vm.data.tips=null;
                })
            } else {
                vm.data.info.isEnd = true;
                vm.data.tips=null;
            }
        }
        data.fun.$LoadingInit = function (_default, arg) {
            arg=arg;
            vm.data.tips=arg?arg.tips:"载入";
            data.fun.dataProcessing(arg);
        }
        data.fun.$Destory = function () {
            data.info.broadcast();
        }
        vm.$onInit = function (arg) {
            vm.interaction = vm.interaction || {
                request: {},
                response: {}
            };
            if (!vm.interaction.request.delay) {
                data.fun.dataProcessing(arg);
            }
            data.info.broadcast = $scope.$on('$Init_LoadingCommonComponent', data.fun.$LoadingInit);
            $scope.$on('$destroy', data.fun.$Destory);
        }

    }
})();