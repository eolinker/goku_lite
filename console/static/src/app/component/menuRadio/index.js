(function () {
    'use strict';
    angular.module('eolinker')
        .component('menuRadioCommonComponent', {
            templateUrl: 'app/component/menuRadio/index.html',
            controller: indexController,
            bindings: {
                output: '=',
                list: '<',
                modelKey: '@',
                bindFun: '&',
                cancelBindFun:'@',
                disabled:'<'
            }
        })

    indexController.$inject = [];

    function indexController() {
        var vm = this;
        vm.fun = {};
        vm.fun.clickMenu = function (inputValue) {
            if (!vm.cancelBindFun) {
                vm.bindFun({
                    arg: inputValue
                });
            } else {
                vm.output[vm.modelKey] = inputValue;
            }
        }
    }
})();