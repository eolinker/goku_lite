(function() {
    'use strict';
    angular.module('eolinker')
        .config(['$stateProvider', 'RouteHelpersProvider', function($stateProvider, helper) {
            $stateProvider
                .state('gepdit.inside', {
                    url: '/inside?projectHashKey',
                    template: '<gepdit-inside></gepdit-inside>'
                });
        }])
        .component('gepditInside', {
            templateUrl: 'app/ui/content/gpedit/inside/index.html',
            controller: indexController
        })

    indexController.$inject = ['$scope','$state','Authority_CommonService'];

    function indexController($scope,$state,Authority_CommonService) {

        var vm = this;
        vm.data = {
            info: {
                strategyID:$state.params.strategyID,
                shrinkObject: {}
            }
        }
        vm.service={
            authority:Authority_CommonService
        }
        vm.$onInit=function(){
            $scope.$emit('$Home_ShrinkSidebar', { shrink: false });
        }
    }
})();
