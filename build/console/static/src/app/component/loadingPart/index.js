(function () {
    'use strict';

    /**
     * loading加载组件
     * @param {function} fun 绑定方法
     * @param {object} interaction 交互参数
     */
    angular.module('eolinker')
        .component('loadingPartCommonComponent', {
            templateUrl: 'app/component/loadingPart/index.html',
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
            broadcast: null
        },fun={};
        vm.data = {
            isEnd: true
        }
        fun.dataProcessing = function (arg) {
            vm.data.isEnd = false;
            var template = {
                promise: vm.fun({
                    arg: arg
                })
            }
            if (template.promise) {
                template.promise.finally(function () {
                    vm.data.isEnd = true;
                })
            } else {
                vm.data.isEnd = true;
            }
        }
        fun.$Destory = function() {
            data.broadcast();
        }
        fun.$LoadingInit = function (_default, arg) {
            fun.dataProcessing(arg);
        }
        vm.$onInit = function(arg) {
            vm.interaction=vm.interaction||{ request: {}, response: {} };
            if (!vm.interaction.request.delay) {
                fun.dataProcessing(arg);
            }
            data.broadcast = $scope.$on('$Init_LoadingPartCommonComponent', fun.$LoadingInit);
            $scope.$on('$destroy', fun.$Destory);
        }

    }
})();