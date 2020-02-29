(function () {
    'use strict';

    /**
     * 进度条选项组件
     * @param {object} mainObject 配置
     */
    angular.module('eolinker')
        .component('progressBarCommonComponent', {
            templateUrl: 'app/component/progressBar/index.html',
            controller: indexController,
            bindings: {
                mainObject: '<',
                modelKey:'=',
                list:'<'
            }
        })

    indexController.$inject = ['$scope'];

    function indexController($scope) {
        let vm=this;

    }
})();